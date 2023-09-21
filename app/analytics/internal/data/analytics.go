package data

import (
	"galileo/app/analytics/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type analyticsRepo struct {
	data *Data
	log  *log.Helper
}

func NewAnalyticsRepo(data *Data, logger log.Logger) biz.AnalyticsRepo {
	return &analyticsRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "analytics.Repo")),
	}
}
