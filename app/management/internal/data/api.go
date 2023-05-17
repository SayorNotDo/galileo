package data

import (
	"context"
	"galileo/app/management/internal/biz"
	"galileo/ent"
	"galileo/ent/api"
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

func convertApi(api *ent.Api) *biz.Api {
	return &biz.Api{
		ID:          api.ID,
		Name:        api.Name,
		Status:      api.Status,
		Type:        api.Type,
		Headers:     api.Headers,
		QueryParams: api.QueryParams,
		Body:        api.Body,
		Response:    api.Response,
		Module:      api.Module,
		Label:       api.Label,
		Description: api.Description,
		CreatedAt:   api.CreatedAt,
		CreatedBy:   api.CreatedBy,
		UpdatedBy:   api.UpdateBy,
		UpdatedAt:   api.UpdateAt,
		DeletedBy:   api.DeletedBy,
		DeletedAt:   api.DeletedAt,
	}
}

func (repo *ApiRepo) CreateApi(ctx context.Context, api *biz.Api) (*biz.Api, error) {
	createApi, err := repo.data.entDB.Api.Create().
		SetName(api.Name).
		SetURL(api.Url).
		SetType(api.Type).
		SetLabel(api.Label).
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

func (repo *ApiRepo) ApiByName(ctx context.Context, name string) (*biz.Api, error) {
	queryApi, err := repo.data.entDB.Api.Query().Where(api.Name(name)).First(ctx)
	if err != nil {
		return nil, err
	}
	ret := convertApi(queryApi)
	return ret, nil
}
