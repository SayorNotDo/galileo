package data

import (
	"context"
	"galileo/app/analytics/internal/biz"
	"galileo/app/analytics/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAnalyticsRepo, NewDB, NewTransaction)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

type contextTxKey struct{}

func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB); ok {
		return tx
	}
	return d.db
}

func NewTransaction(d *Data) biz.Transaction {
	return d
}

// NewData .
func NewData(db *gorm.DB, c *conf.Data, logger log.Logger) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "analytics.DataService"))

	return &Data{
		db:  db,
		log: l,
	}, nil
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	l := log.NewHelper(log.With(logger, "module", "data.gorm"))
	/* 连接到 Doris 集群 */
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		l.Fatalf("failed opening connection to mysql: %v", err)
	}
	return db
}
