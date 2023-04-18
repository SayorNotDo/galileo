package data

import (
	"galileo/app/engine/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type engineRepo struct {
	data *Data
	log  *log.Helper
}

func NewEngineRepo(data *Data, logger log.Logger) biz.EngineRepo {
	return &engineRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "core.Repo")),
	}
}
