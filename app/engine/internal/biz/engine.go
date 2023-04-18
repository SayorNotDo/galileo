package biz

import "github.com/go-kratos/kratos/v2/log"

type EngineRepo interface {
}

type EngineUseCase struct {
	Repo EngineRepo
	log  *log.Helper
}

func NewEngineUseCase(repo EngineRepo, logger log.Logger) *EngineUseCase {
	helper := log.NewHelper(log.With(logger, "module", "core.useCase"))
	return &EngineUseCase{Repo: repo, log: helper}
}
