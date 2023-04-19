package biz

import (
	"context"
	taskV1 "galileo/api/management/task/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type EngineRepo interface {
	UpdateTaskStatus(ctx context.Context, status taskV1.TaskStatus) (bool, error)
}

type EngineUseCase struct {
	Repo EngineRepo
	log  *log.Helper
}

func NewEngineUseCase(repo EngineRepo, logger log.Logger) *EngineUseCase {
	helper := log.NewHelper(log.With(logger, "module", "engine.useCase"))
	return &EngineUseCase{Repo: repo, log: helper}
}
