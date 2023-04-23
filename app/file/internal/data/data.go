package data

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"galileo/app/file/internal/conf"
	"galileo/ent"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewFileRepo, NewEntDB, NewRedis, NewRegistrar, NewDiscovery, NewOssClient)

var RedisCli redis.Cmdable

// Data .
type Data struct {
	// TODO wrapped database client
	entDB    *ent.Client
	log      *log.Helper
	oss      *OssClient
	config   *conf.Data
	redisCli redis.Cmdable
}

type OssClient struct {
	*oss.Client
}

// NewData .
func NewData(c *conf.Data, db *ent.Client, logger log.Logger, redisCli redis.Cmdable, oss *OssClient) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "file.DataService"))
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return &Data{
		entDB:    db,
		log:      l,
		redisCli: redisCli,
		config:   c,
		oss:      oss,
	}, cleanup, nil
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

func NewOssClient(conf *conf.Data) (*OssClient, error) {
	client, err := oss.New(conf.Oss.Endpoint, conf.Oss.AccessKey, conf.Oss.AccessSecret)
	if err != nil {
		return nil, err
	}
	return &OssClient{client}, nil
}
