package service

import (
	v1 "galileo/api/core/v1"
	"galileo/app/core/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is scheduler providers.
var ProviderSet = wire.NewSet(NewCoreService, NewEngineService)

type CoreService struct {
	v1.UnimplementedCoreServer

	uc  *biz.CoreUseCase
	log *log.Helper
}

type EngineService struct {
	v1.UnimplementedCoreServer

	uc  *biz.EngineUseCase
	log *log.Helper
}

func NewCoreService(uc *biz.CoreUseCase, logger log.Logger) *CoreService {
	return &CoreService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "core.Service")),
	}
}

func NewEngineService(uc *biz.EngineUseCase, logger log.Logger) *EngineService {
	return &EngineService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "engine.Service")),
	}
}
