package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Api struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	Type        int8      `json:"type"`
	Status      int8      `json:"status"`
	Body        []byte    `json:"body"`
	QueryParams []byte    `json:"query_params"`
	Response    []byte    `json:"response"`
	Module      string    `json:"module"`
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
