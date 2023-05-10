package service

import (
	"context"
	v1 "galileo/api/management/api/v1"
	"galileo/app/management/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type ApiService struct {
	v1.UnimplementedApiServer
	uc     *biz.ApiUseCase
	logger *log.Helper
}

func NewApiService(uc *biz.ApiUseCase, logger log.Logger) *ApiService {
	return &ApiService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "management.apiService")),
	}
}

func NewApi(req *v1.CreateApiRequest) (biz.Api, error) {
	return biz.Api{
		Name:        req.Name,
		Url:         req.Url,
		Module:      req.Module,
		Type:        int8(req.Type.Number()),
		Body:        req.Body,
		QueryParams: req.QueryParams,
		Response:    req.Response,
		Description: req.Description,
	}, nil
}

func (s *ApiService) CreateApi(ctx context.Context, req *v1.CreateApiRequest) (*v1.CreateApiReply, error) {
	newApi, err := NewApi(req)
	if err != nil {
		return nil, err
	}
	ret, err := s.uc.CreateApi(ctx, &newApi)
	if err != nil {
		return nil, err
	}
	return &v1.CreateApiReply{
		Id:        ret.ID,
		CreatedAt: ret.CreatedAt.Unix(),
	}, nil
}
