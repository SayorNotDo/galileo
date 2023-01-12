package data

import (
	"galileo/app/user/internal/conf"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	slog "log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGormDB, NewRedis, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	gormDB  *gorm.DB
	redisDB *redis.Client
}

func NewGormDB(c *conf.Data) (*gorm.DB, error) {
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond * 200, // 慢查询 SQL 阈值
			Colorful:                  true,                   // 禁用彩色打印
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Info, // Log lever
		},
	)
	dsn := c.Database.Source
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(150)
	sqlDB.SetConnMaxLifetime(25 * time.Second)
	return db, nil
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{gormDB: db, redisDB: rdb}, cleanup, nil
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
