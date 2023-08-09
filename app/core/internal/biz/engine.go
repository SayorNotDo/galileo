package biz

import (
	"context"
	engineV1 "galileo/api/engine/v1"
)

type EngineRepo interface {
	UploadEngineFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error)
	InspectContainer(ctx context.Context, id string) (*engineV1.Container, error)
}

func (e *EngineUseCase) UploadEngineFile(ctx context.Context, fileName string, filetype string, content []byte) (string, error) {
	return e.repo.UploadEngineFile(ctx, fileName, filetype, content)
}

func (e *EngineUseCase) InspectContainer(ctx context.Context, id string) (*engineV1.Container, error) {
	return e.repo.InspectContainer(ctx, id)
}
