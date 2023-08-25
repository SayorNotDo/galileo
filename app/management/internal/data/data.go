package data

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	engineV1 "galileo/api/engine/v1"
	fileV1 "galileo/api/file/v1"
	"galileo/app/management/internal/conf"
	"galileo/ent"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	grpcx "google.golang.org/grpc"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewEntDB,
	NewRedis,
	NewProjectRepo,
	NewTaskRepo,
	NewTestCaseRepo,
	NewApiRepo,
	NewRegistrar,
	NewDiscovery,
	NewManagementRepo,
	NewFileServiceClient,
	NewEngineServiceClient,
)

var RedisCli redis.Cmdable

// Data .
type Data struct {
	engineCli engineV1.EngineClient
	fileCli   fileV1.FileClient
	entDB     *ent.Client
	log       *log.Helper
	redisCli  *redis.Client
}

// NewData .
func NewData(c *conf.Data, db *ent.Client, logger log.Logger, redisCli *redis.Client, fc fileV1.FileClient, eg engineV1.EngineClient) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "management.DataService"))
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{entDB: db, log: l, redisCli: redisCli, fileCli: fc, engineCli: eg}, cleanup, nil
}

func NewEngineServiceClient(sr *conf.Service, rr registry.Discovery) engineV1.EngineClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Engine.Endpoint),
		grpc.WithDiscovery(rr),
		grpc.WithMiddleware(
			tracing.Client(), // 链路追踪
			recovery.Recovery(),
			metadata.Client(),
		),
		grpc.WithTimeout(2*time.Second),
		grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)
	if err != nil {
		panic(err)
	}
	c := engineV1.NewEngineClient(conn)
	return c
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

func NewEntDB(c *conf.Data) (*ent.Client, error) {
	drv, err := sql.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	sqlDB := drv.DB()
	sqlDB.SetMaxIdleConns(50)
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
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}
	return client, nil
}

func NewRedis(conf *conf.Data, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "management.dataService.redis"))
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	// rdb.AddHook()
	/* 开启tracing instrumentation */
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		l.Fatalf(err.Error())
	}
	/* 开启 metrics instrumentation */
	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		l.Fatalf(err.Error())
	}
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	err := rdb.Ping(timeout).Err()
	defer cancelFunc()
	if err != nil {
		l.Fatalf("redis connect error: %v", err)
	}
	RedisCli = rdb
	return rdb
}

func NewFileServiceClient(sr *conf.Service, rr registry.Discovery) fileV1.FileClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.File.Endpoint),
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
	c := fileV1.NewFileClient(conn)
	return c
}

func rollback(tx *ent.Tx, err error) error {
	if rErr := tx.Rollback(); rErr != nil {
		err = fmt.Errorf("%w: %v", err, rErr)
	}
	return err
}
