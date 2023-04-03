package data

import (
	"context"
	"galileo/app/runner/internal/conf"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRunnerRepo, NewRedis)

var RedisCli redis.Cmdable

// Data .
type Data struct {
	log      *log.Helper
	redisCli redis.Cmdable
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		log:      log.NewHelper(log.With(logger, "module", "runnerService/data")),
		redisCli: RedisCli,
	}, cleanup, nil
}

func NewRedis(conf *conf.Data, logger log.Logger) redis.Cmdable {
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
