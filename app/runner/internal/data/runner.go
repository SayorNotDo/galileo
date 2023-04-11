package data

import (
	"galileo/app/runner/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type runnerRepo struct {
	data *Data
	log  *log.Helper
}

func NewRunnerRepo(data *Data, logger log.Logger) biz.RunnerRepo {
	return &runnerRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "runnerRepo")),
	}
}
