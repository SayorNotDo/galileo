package data

import (
	"context"
	"galileo/app/runner/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.RunnerDataRepo = (*runnerDataRepo)(nil)

type runnerDataRepo struct {
	data *Data
	log  *log.Helper
}

func (r *runnerDataRepo) InsertRunnerData(ctx context.Context, req *biz.Runner) error {
	return nil
}

func (r *runnerDataRepo) BatchInsertRunnerData(ctx context.Context, req []*biz.Runner) error {
	return nil
}
