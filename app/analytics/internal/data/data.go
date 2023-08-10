package data

import (
	"galileo/app/analytics/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAnalyticsRepo)

// Data .
type Data struct {
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "analytics.DataService"))

	return &Data{log: l}, nil
}
