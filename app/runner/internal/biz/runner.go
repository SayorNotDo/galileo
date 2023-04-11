package biz

import "github.com/go-kratos/kratos/v2/log"

type Runner struct {
}

type RunnerRepo interface{}

type RunnerUseCase struct {
	repo RunnerRepo
	log  *log.Helper
}

func NewRunnerUseCase(repo RunnerRepo, logger log.Logger) *RunnerUseCase {
	return &RunnerUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "runnerUseCase")),
	}
}
