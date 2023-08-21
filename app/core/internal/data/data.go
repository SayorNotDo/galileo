package data

import (
	"context"
	engineV1 "galileo/api/engine/v1"
	fileV1 "galileo/api/file/v1"
	projectV1 "galileo/api/management/project/v1"
	taskV1 "galileo/api/management/task/v1"
	userV1 "galileo/api/user/v1"
	"galileo/app/core/internal/conf"
	. "galileo/app/core/internal/pkg/constant"
	"github.com/IBM/sarama"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/v9"
	grpcx "google.golang.org/grpc"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewCoreRepo,
	NewEngineRepo,
	NewUserRepo,
	NewProjectServiceClient,
	NewUserServiceClient,
	NewTaskServiceClient,
	NewFileServiceClient,
	NewEngineServiceClient,
	NewRegistrar,
	NewDiscovery,
	NewRedis,
	NewKafkaProducer,
)

var RedisCli *redis.Client

// Data .
type Data struct {
	log           *log.Helper
	uc            userV1.UserClient
	taskCli       taskV1.TaskClient
	projectCli    projectV1.ProjectClient
	fileCli       fileV1.FileClient
	redisCli      *redis.Client
	engineCli     engineV1.EngineClient
	kafkaProducer sarama.SyncProducer
}

// NewData .
func NewData(c *conf.Data, uc userV1.UserClient, logger log.Logger, redisCli *redis.Client, projectCli projectV1.ProjectClient, taskCli taskV1.TaskClient, fileCli fileV1.FileClient, engineCli engineV1.EngineClient, kafkaProducer sarama.SyncProducer) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "core.DataService"))
	return &Data{
		log:           l,
		uc:            uc,
		redisCli:      redisCli,
		projectCli:    projectCli,
		taskCli:       taskCli,
		fileCli:       fileCli,
		engineCli:     engineCli,
		kafkaProducer: kafkaProducer,
	}, nil
}

func NewEngineServiceClient(sr *conf.Service, rr registry.Discovery) engineV1.EngineClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Engine.Endpoint),
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
	return engineV1.NewEngineClient(conn)
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
	return fileV1.NewFileClient(conn)

}

func NewTaskServiceClient(ac *conf.Auth, sr *conf.Service, rr registry.Discovery) taskV1.TaskClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Management.Endpoint),
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
	c := taskV1.NewTaskClient(conn)
	return c
}

func NewProjectServiceClient(ac *conf.Auth, sr *conf.Service, rr registry.Discovery) projectV1.ProjectClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Management.Endpoint),
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
	c := projectV1.NewProjectClient(conn)
	return c
}

func NewUserServiceClient(ac *conf.Auth, sr *conf.Service, rr registry.Discovery) userV1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.User.Endpoint),
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
	c := userV1.NewUserClient(conn)
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

func NewRedis(conf *conf.Data, logger log.Logger) *redis.Client {
	logs := log.NewHelper(log.With(logger, "module", "coreService/data/redis"))
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

func NewKafkaProducer(conf *conf.Data, logger log.Logger) sarama.SyncProducer {
	logs := log.NewHelper(log.With(logger, "module", "coreService/data/kafka"))
	/* 设置Kafka配置 */
	config := sarama.NewConfig()
	config.Producer.Return.Successes = conf.Kafka.ProducerReturnSuccesses
	config.Producer.Return.Errors = conf.Kafka.ProducerReturnErrors
	config.Producer.Timeout = conf.Kafka.ProducerTimeout.AsDuration()

	/* 创建Kafka生产者 */
	producer, err := sarama.NewSyncProducer(KafkaBrokers, config)
	if err != nil {
		logs.Fatalf("kafka dial error: %v", err)
	}
	return producer
}
