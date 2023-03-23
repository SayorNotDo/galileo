package data

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"galileo/app/user/internal/conf"
	"galileo/ent"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntDB, NewRedis, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	entDB   *ent.Client
	redisDB *redis.Client
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

//func NewGormDB(c *conf.Data) (*gorm.DB, error) {
//	newLogger := logger.New(
//		slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
//		logger.Config{
//			SlowThreshold:             time.Millisecond * 200, // 慢查询 SQL 阈值
//			Colorful:                  true,                   // 禁用彩色打印
//			IgnoreRecordNotFoundError: false,
//			LogLevel:                  logger.Info, // Log lever
//		},
//	)
//	dsn := c.Database.Source
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: newLogger,
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true,
//		},
//	})
//	if err != nil {
//		return nil, err
//	}
//	sqlDB, err := db.DB()
//	if err != nil {
//		return nil, err
//	}
//	sqlDB.SetMaxIdleConns(50)
//	sqlDB.SetMaxOpenConns(150)
//	sqlDB.SetConnMaxLifetime(25 * time.Second)
//	return db, nil
//}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *ent.Client, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{entDB: db, redisDB: rdb}, cleanup, nil
}

// NewRedis .
func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	if err := rdb.Close(); err != nil {
		log.Error(err)
	}
	return rdb
}
