package data

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	managementV1 "galileo/api/management/v1"
	"galileo/app/engine/internal/conf"
	"galileo/app/engine/pkg/constant"
	"galileo/ent"
	"github.com/RichardKnop/machinery/v1"
	machineryConf "github.com/RichardKnop/machinery/v1/config"
	"github.com/docker/docker/client"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	grpcx "google.golang.org/grpc"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewEntDB,
	NewEngineRepo,
	NewDockerRepo,
	NewRedis,
	NewDockerClient,
	NewRegistrar,
	NewMachinery,
	NewDiscovery,
	NewTaskServiceClient,
)

var (
	RedisCli redis.Cmdable
)

// Data .
type Data struct {
	entDB     *ent.Client
	log       *log.Helper
	redisCli  redis.Cmdable
	ManageCli managementV1.ManagementClient
	dockerCli *client.Client
	cron      *cron.Cron
	machinery *machinery.Server
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, redisCli redis.Cmdable, ManageCli managementV1.ManagementClient, dockerCli *client.Client) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "engine.DataService"))
	cronJob := cron.New(cron.WithSeconds())
	cronJob.Start()
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return &Data{log: l, redisCli: redisCli, ManageCli: ManageCli, cron: cronJob, dockerCli: dockerCli}, cleanup, nil
}

func NewEntDB(c *conf.Data) (*ent.Client, error) {
	drv, err := sql.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	sqlDB := drv.DB()
	sqlDB.SetMaxOpenConns(150)
	sqlDB.SetConnMaxLifetime(25 * time.Second)
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		log.Context(ctx).Info(i)
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(attribute.String("sql", fmt.Sprintln(i...))),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	cli := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		return nil, err
	}
	if err := cli.Schema.Create(context.Background()); err != nil {
		return nil, err
	}
	return cli, nil
}

func NewTaskServiceClient(sr *conf.Service, rr registry.Discovery) managementV1.ManagementClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Management.Endpoint),
		grpc.WithDiscovery(rr),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
			metadata.Client(),
		),
		grpc.WithTimeout(2*time.Second),
		grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)
	if err != nil {
		panic(err)
	}
	c := managementV1.NewManagementClient(conn)
	return c
}

func NewRedis(conf *conf.Data, logger log.Logger) redis.Cmdable {
	logs := log.NewHelper(log.With(logger, "module", "engineService/data/redis"))
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	err := rdb.Ping(timeout).Err()
	defer cancelFunc()
	if err != nil {
		logs.Fatalf("redis connect error: %v", err)
	}
	RedisCli = rdb
	return rdb
}

func NewMachinery(conf *conf.Data, logger log.Logger) *machinery.Server {
	logs := log.NewHelper(log.With(logger, "module", "engineService/data/machinery"))

	cnf := &machineryConf.Config{
		Broker:        conf.Redis.Addr,
		DefaultQueue:  constant.DefaultQueue,
		ResultBackend: conf.Redis.Addr,
		Lock:          conf.Redis.Addr,
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		logs.Fatalf("redis connect error: %v", err)
	}

	worker := server.NewWorker("scheduler", 1)

	errorsChan := make(chan error, 1)

	worker.LaunchAsync(errorsChan)

	return server
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme

	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewDockerClient(conf *conf.Service, logger log.Logger) *client.Client {
	logs := log.NewHelper(log.With(logger, "module", "data.docker"))
	/* 远程连接：生产环境部署docker-hub */
	/* 获取连接助手对象 */
	//helper, err := connhelper.GetConnectionHelper(conf.Docker.Endpoint)
	//logs.Info(helper)
	//if err != nil {
	//	logs.Fatalf("docker get connect helper error: %v", err)
	//	return nil
	//}
	//httpClient := &http.Client{
	//	Transport: &http.Transport{
	//		DialContext: helper.Dialer,
	//	},
	//}
	//cli, err := client.NewClientWithOpts(
	//	client.WithHTTPClient(httpClient),
	//	client.WithHost(helper.Host),
	//	client.WithDialContext(helper.Dialer),
	//)
	/* 本地连接: 仅适用在开发环境 */
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		logs.Fatalf("docker connect error: %v", err)
		return nil
	}
	return cli
}
