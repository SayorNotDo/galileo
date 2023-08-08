package data

import (
	"context"
	fileV1 "galileo/api/file/v1"
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

func (e engineRepo) UploadEngineFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error) {
	res, err := e.data.fileCli.UploadFile(ctx, &fileV1.UploadFileRequest{
		Name:    fileName,
		Type:    fileType,
		Content: content,
	})
	if err != nil {
		return "", err
	}
	return res.Url, err
}
