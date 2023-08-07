package data

import (
	"context"
	"galileo/app/core/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type engineRepo struct {
	data *Data
	log  *log.Helper
}

func NewEngineRepo(data *Data, logger log.Logger) biz.EngineRepo {
	return &engineRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "engine.Repo")),
	}
}

func (e engineRepo) UploadEngineFile(ctx context.Context, fileName string, fileType string, content []byte) error {
	//TODO implement me
	panic("implement me")
}
