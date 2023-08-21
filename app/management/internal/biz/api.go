package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Api struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	Type        int8      `json:"type"`
	Status      int8      `json:"status"`
	Headers     string    `json:"headers"`
	Body        string    `json:"body"`
	QueryParams string    `json:"query_params"`
	Response    string    `json:"response"`
	Module      string    `json:"module"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   uint32    `json:"created_by"`
	UpdatedAt   time.Time `json:"update_at"`
	UpdatedBy   uint32    `json:"update_by"`
	DeletedAt   time.Time `json:"deleted_at"`
	DeletedBy   uint32    `json:"deleted_by"`
}

type ApiRepo interface {
	CreateApi(ctx context.Context, api *Api) (*Api, error)
	UpdateApi(ctx context.Context, api *Api) (bool, error)
	ApiByName(ctx context.Context, name string) (*Api, error)
	ApiById(ctx context.Context, id int32) (*Api, error)
	IsApiDuplicated(ctx context.Context, method int8, url string) (bool, error)
	UploadApiFile(ctx context.Context, fileName, fileType string, content []byte) (string, error)
}

type ApiUseCase struct {
	repo ApiRepo
	log  *log.Helper
}

func NewApiUseCase(repo ApiRepo, logger log.Logger) *ApiUseCase {
	return &ApiUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "management.apiUseCase")),
	}
}

func (uc *ApiUseCase) CreateApi(ctx context.Context, api *Api) (*Api, error) {
	return uc.repo.CreateApi(ctx, api)
}

func (uc *ApiUseCase) UpdateApi(ctx context.Context, api *Api) (bool, error) {
	return uc.repo.UpdateApi(ctx, api)
}

func (uc *ApiUseCase) ApiByName(ctx context.Context, name string) (*Api, error) {
	return uc.repo.ApiByName(ctx, name)
}

func (uc *ApiUseCase) ApiById(ctx context.Context, id int32) (*Api, error) {
	return uc.repo.ApiById(ctx, id)
}

func (uc *ApiUseCase) IsApiDuplicated(ctx context.Context, method int8, url string) (bool, error) {
	return uc.repo.IsApiDuplicated(ctx, method, url)
}

func (uc *ApiUseCase) UploadInterfaceFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error) {
	return uc.repo.UploadApiFile(ctx, fileName, fileType, content)
}
