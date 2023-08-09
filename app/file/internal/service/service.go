package service

import (
	v1 "galileo/api/file/v1"
	"galileo/app/file/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is scheduler providers.
var ProviderSet = wire.NewSet(NewFileService)

type FileService struct {
	v1.UnimplementedFileServer
	uc     *biz.FileUseCase
	logger *log.Helper
}

func NewFileService(uc *biz.FileUseCase, logger log.Logger) *FileService {
	return &FileService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "file")),
	}
}
