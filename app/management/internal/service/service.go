package service

import (
	v1 "galileo/api/management/v1"
	"galileo/app/management/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is scheduler providers.
var ProviderSet = wire.NewSet(
	NewManagementService,
)

type ManagementService struct {
	v1.UnimplementedManagementServer

	ac     *biz.ApiUseCase
	pc     *biz.ProjectUseCase
	tc     *biz.TestcaseUseCase
	sc     *biz.TaskUseCase
	uc     *biz.ManagementUseCase
	logger *log.Helper
}

func NewManagementService(
	ac *biz.ApiUseCase,
	pc *biz.ProjectUseCase,
	tc *biz.TestcaseUseCase,
	uc *biz.ManagementUseCase,
	sc *biz.TaskUseCase,
	logger log.Logger,
) *ManagementService {
	return &ManagementService{
		ac:     ac,
		pc:     pc,
		tc:     tc,
		sc:     sc,
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "management.Service")),
	}
}
