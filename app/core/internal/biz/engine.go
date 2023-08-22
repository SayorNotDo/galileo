package biz

import (
	"context"
	engineV1 "galileo/api/engine/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type EngineRepo interface {
	UploadEngineFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error)
	InspectContainer(ctx context.Context, id string) (*engineV1.Container, error)
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

func (e *EngineUseCase) InspectContainer(ctx context.Context, id string) (*engineV1.Container, error) {
	return e.repo.InspectContainer(ctx, id)
}
