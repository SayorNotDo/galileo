package data

import (
	"context"
	"fmt"
	fileService "galileo/api/file/v1"
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

func rollback(tx *ent.Tx, err error) error {
	if rErr := tx.Rollback(); rErr != nil {
		err = fmt.Errorf("%w: %v", err, rErr)
	}
	return err
}

func (repo *ApiRepo) CreateApi(ctx context.Context, api *biz.Api) (*biz.Api, error) {
	tx, err := repo.data.entDB.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	createApi, err := tx.Api.Create().
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
		return nil, rollback(tx, err)
	}
	if createApi.Label != "" {
		// TODO: add label name to tag table
		println("do not implement")
	}
	return &biz.Api{
		ID:        createApi.ID,
		CreatedAt: createApi.CreatedAt,
	}, nil
}

func (repo *ApiRepo) UpdateApi(ctx context.Context, api *biz.Api) (bool, error) {
	err := repo.data.entDB.Api.UpdateOneID(api.ID).
		SetName(api.Name).
		SetURL(api.Url).
		SetType(api.Type).
		SetLabel(api.Label).
		SetBody(api.Body).
		SetModule(api.Module).
		SetQueryParams(api.QueryParams).
		SetResponse(api.Response).
		SetDescription(api.Description).
		SetStatus(api.Status).
		SetUpdateBy(api.UpdatedBy).
		Exec(ctx)
	if err != nil {
		return false, err
	}
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

func (repo *ApiRepo) ApiById(ctx context.Context, id int64) (*biz.Api, error) {
	queryApi, err := repo.data.entDB.Api.Query().Where(api.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	ret := convertApi(queryApi)
	return ret, err
}

func (repo *ApiRepo) IsApiDuplicated(ctx context.Context, method int8, url string) (bool, error) {
	_, err := repo.data.entDB.Api.Query().
		Where(api.Type(method)).
		Where(api.URL(url)).First(ctx)
	if err != nil && err.Error() != "ent: api not found" {
		return false, err
	}
	return true, nil
}

func (repo *ApiRepo) UploadApiFile(ctx context.Context, fileName, fileType string, content []byte) (string, error) {
	res, err := repo.data.fileCli.UploadFile(ctx, &fileService.UploadFileRequest{
		Name:    fileName,
		Type:    fileType,
		Content: content})
	if err != nil {
		return "", err
	}
	return res.Url, nil
}
