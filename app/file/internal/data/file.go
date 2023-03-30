package data

import (
	"context"

	"galileo/app/file/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type fileRepo struct {
	data *Data
	log  *log.Helper
}

// NewFileRepo .
func NewFileRepo(data *Data, logger log.Logger) biz.FileRepo {
	return &fileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *fileRepo) UploadFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error) {

	return "", nil
}
func (r *fileRepo) Save(ctx context.Context, g *biz.File) (*biz.File, error) {
	return g, nil
}

func (r *fileRepo) Update(ctx context.Context, g *biz.File) (*biz.File, error) {
	return g, nil
}

func (r *fileRepo) ListAll(context.Context) ([]*biz.File, error) {
	return nil, nil
}
