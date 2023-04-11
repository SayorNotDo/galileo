package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Runner struct {
	Id        int64                  `json:"id"`
	Type      string                 `json:"type"`
	Timestamp int64                  `json:"time"`
	Data      map[string]interface{} `json:"data"`
}

type RunnerDataRepo interface {
	InsertRunnerData(ctx context.Context, req *Runner) error
	BatchInsertRunnerData(ctx context.Context, req []*Runner) error
}

type RunnerDataUseCase struct {
	repo RunnerDataRepo
	log  *log.Helper
}

func NewRunnerDataUseCase(repo RunnerDataRepo, logger log.Logger) *RunnerDataUseCase {
	l := log.NewHelper(log.With(logger, "module", "runner-data/useCase/runner-service"))
	return &RunnerDataUseCase{repo: repo, log: l}
}

func (uc *RunnerDataUseCase) InsertRunnerData(ctx context.Context, req *Runner) error {
	return nil
}

func (uc *RunnerDataUseCase) BatchInsertRunnerData(ctx context.Context, req []*Runner) error {
	return nil
}
