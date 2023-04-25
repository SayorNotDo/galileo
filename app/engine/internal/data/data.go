package data

import (
	"context"
	"galileo/app/engine/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewEngineRepo,
	NewRedis,
	NewRegistrar,
	NewDiscovery,
)

var RedisCli redis.Cmdable

// Data .
type Data struct {
	log      *log.Helper
	redisCli redis.Cmdable
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, redisCli redis.Cmdable) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "engine.DataService"))
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return &Data{log: l, redisCli: redisCli}, cleanup, nil
}

//func NewTaskServiceClient(sr *conf.Service, rr registry.Discovery) taskV1.TaskClient {
//	conn, err := grpc.DialInsecure(
//		context.Background(),
//		grpc.WithEndpoint(sr.Task.Endpoint),
//		grpc.WithDiscovery(rr),
//		grpc.WithMiddleware(
//			tracing.Client(),
//			recovery.Recovery(),
//			metadata.Client(),
//		),
//		grpc.WithTimeout(2*time.Second),
//		grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
//	)
//	if err != nil {
//		panic(err)
//	}
//	c := taskV1.NewTaskClient(conn)
//	return c
//}

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
	rdb.AddHook(redisotel.TracingHook{})
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	err := rdb.Ping(timeout).Err()
	defer cancelFunc()
	if err != nil {
		logs.Fatalf("redis connect error: %v", err)
	}
	RedisCli = rdb
	return rdb
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
