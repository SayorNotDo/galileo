package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type EngineRepo interface {
	UploadEngineFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error)
}

type EngineUseCase struct {
	repo EngineRepo
	log  *log.Helper
}

func NewEngineUseCase(repo EngineRepo, logger log.Logger) *EngineUseCase {
	helper := log.NewHelper(log.With(logger, "module", "core.useCase"))
	return &EngineUseCase{
		repo: repo,
		log:  helper,
	}
}

func (e *EngineUseCase) UploadEngineFile(ctx context.Context, fileName string, filetype string, content []byte) (string, error) {
	return e.repo.UploadEngineFile(ctx, fileName, filetype, content)
}
