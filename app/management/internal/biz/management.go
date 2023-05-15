package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type BaseInfo struct {
	ApiCount       int64
	ApiCaseCount   int64
	CronJobCount   int64
	SceneCaseCount int64
}

type ManagementUseCase struct {
	logger *log.Helper
}

func NewManagementUseCase(logger log.Logger) *ManagementUseCase {
	return &ManagementUseCase{
		logger: log.NewHelper(log.With(logger, "module", "management.UseCase")),
	}
}

func (uc *ManagementUseCase) BaseInfo(ctx context.Context) (*BaseInfo, error) {
	return &BaseInfo{}, nil
}
