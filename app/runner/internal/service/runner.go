package service

import (
	"galileo/app/runner/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type RunnerService struct {
	uc     *biz.RunnerUseCase
	logger *log.Helper
}

func NewRunnerService(uc *biz.RunnerUseCase, logger log.Logger) *RunnerService {
	return &RunnerService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "runnerService")),
	}
}
