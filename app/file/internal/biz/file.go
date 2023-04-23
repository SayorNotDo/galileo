package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// File is a Greeter model.
type File struct {
	Hello string
}

type OssStsToken struct {
	AccessKey     string
	AccessSecret  string
	Expiration    string
	SecurityToken string
	EndPoint      string
	BucketName    string
	Region        string
	Url           string
}

// FileRepo is a Greater repo.
type FileRepo interface {
	GetOssStsToken(ctx context.Context) (*OssStsToken, error)
	UploadFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error)
}

// FileUseCase is a File useCase.
type FileUseCase struct {
	repo FileRepo
	log  *log.Helper
}

// NewFileUseCase new a Greeter usecase.
func NewFileUseCase(repo FileRepo, logger log.Logger) *FileUseCase {
	return &FileUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c *FileUseCase) UploadFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error) {
	return c.repo.UploadFile(ctx, fileName, fileType, content)
}

func (c *FileUseCase) GetOssStsToken(ctx context.Context) (*OssStsToken, error) {
	return c.repo.GetOssStsToken(ctx)
}
