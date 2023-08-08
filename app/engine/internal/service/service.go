package service

import (
	v1 "galileo/api/engine/v1"
	"galileo/app/engine/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is scheduler providers.
var ProviderSet = wire.NewSet(NewEngineService)

type EngineService struct {
	v1.UnimplementedEngineServer

	dc  *biz.DockerUseCase
	uc  *biz.EngineUseCase
	log *log.Helper
}

func NewEngineService(uc *biz.EngineUseCase, dc *biz.DockerUseCase, logger log.Logger) *EngineService {
	return &EngineService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "engine.Service")),
	}
}
