package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewAnalyticsUseCase)

type AnalyticsUseCase struct {
	repo AnalyticsRepo
	log  *log.Helper
}

func NewAnalyticsUseCase(repo AnalyticsRepo, logger log.Logger) *AnalyticsUseCase {
	helper := log.NewHelper(log.With(logger, "module", "analytics.useCase"))

	return &AnalyticsUseCase{
		repo: repo,
		log:  helper,
	}
}
