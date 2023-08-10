package service

import (
	v1 "galileo/api/analytics/v1"
	"galileo/app/analytics/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAnalyticsService)

type AnalyticsService struct {
	v1.UnimplementedAnalyticsServer

	ac  *biz.AnalyticsUseCase
	log *log.Helper
}

func NewAnalyticsService(ac *biz.AnalyticsUseCase, logger log.Logger) *AnalyticsService {
	return &AnalyticsService{
		ac:  ac,
		log: log.NewHelper(log.With(logger, "module", "analytics.Service")),
	}
}
