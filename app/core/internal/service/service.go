package service

import (
	v1 "galileo/api/core/v1"
	"galileo/app/core/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is scheduler providers.
var ProviderSet = wire.NewSet(NewCoreService)

type CoreService struct {
	v1.UnimplementedCoreServer

	ec  *biz.EngineUseCase
	uc  *biz.CoreUseCase
	uu  *biz.UserUseCase
	log *log.Helper
}

func NewCoreService(uc *biz.CoreUseCase, ec *biz.EngineUseCase, uu *biz.UserUseCase, logger log.Logger) *CoreService {
	return &CoreService{
		uc:  uc,
		ec:  ec,
		log: log.NewHelper(log.With(logger, "module", "core.Service")),
	}
}
