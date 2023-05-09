package data

import (
	"context"
	"galileo/app/management/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type ApiRepo struct {
	data *Data
	log  *log.Helper
}

func NewApiRepo(data *Data, logger log.Logger) biz.ApiRepo {
	return &ApiRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "management.ApiRepo")),
	}
}

func (repo *ApiRepo) CreateApi(ctx context.Context, api *biz.Api) (*biz.Api, error) {
	createApi, err := repo.data.entDB.Api.Create().
		SetName(api.Name).
		SetURL(api.Url).
		SetType(api.Type).
		SetCreatedBy(api.CreatedBy).
		SetBody(api.Body).
		SetModule(api.Module).
		SetQueryParams(api.QueryParams).
		SetResponse(api.Response).
		SetDescription(api.Description).
		SetStatus(api.Status).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Api{
		ID:        createApi.ID,
		CreatedAt: createApi.CreatedAt,
	}, nil
}

func (repo *ApiRepo) UpdateApi(ctx context.Context, api *biz.Api) (bool, error) {
	return true, nil
}
