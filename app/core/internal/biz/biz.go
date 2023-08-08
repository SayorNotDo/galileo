package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewCoreUseCase, NewEngineUseCase)

type EngineUseCase struct {
	repo EngineRepo
	log  *log.Helper
}

type CoreUseCase struct {
	repo       CoreRepo
	log        *log.Helper
	signingKey string
}

func NewEngineUseCase(repo EngineRepo, logger log.Logger) *EngineUseCase {
	helper := log.NewHelper(log.With(logger, "module", "core.useCase"))
	return &EngineUseCase{
		repo: repo,
		log:  helper,
	}
}
