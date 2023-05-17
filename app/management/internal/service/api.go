package service

import (
	"context"
	v1 "galileo/api/management/api/v1"
	"galileo/app/management/internal/biz"
	. "galileo/pkg/errResponse"
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
	if len(req.ApiInfo.Name) <= 0 {
		return biz.Api{}, SetCustomizeErrMsg(ReasonParamsError, "name is required")
	}
	if len(req.ApiInfo.Url) <= 0 {
		return biz.Api{}, SetCustomizeErrMsg(ReasonParamsError, "url is required")
	}
	if req.ApiInfo.Type <= 0 {
		return biz.Api{}, SetCustomizeErrMsg(ReasonParamsError, "type is required")
	}
	return biz.Api{
		Name:        req.ApiInfo.Name,
		Url:         req.ApiInfo.Url,
		Module:      req.ApiInfo.Module,
		Type:        int8(req.ApiInfo.Type.Number()),
		Body:        req.ApiInfo.Body,
		QueryParams: req.ApiInfo.QueryParams,
		Response:    req.ApiInfo.Response,
		Description: req.ApiInfo.Description,
	}, nil
}

func (s *ApiService) CreateApi(ctx context.Context, req *v1.CreateApiRequest) (*v1.CreateApiReply, error) {
	newApi, err := NewApi(req)
	if err != nil {
		return nil, err
	}
	if ret, _ := s.uc.ApiByName(ctx, newApi.Name); ret != nil {
		return nil, SetCustomizeErrInfoByReason(ReasonRecordExists)
	}
	userId := ctx.Value("x-md-global-userId").(uint32)
	newApi.CreatedBy = userId
	ret, err := s.uc.CreateApi(ctx, &newApi)
	if err != nil {
		return nil, err
	}
	return &v1.CreateApiReply{
		Id:        ret.ID,
		CreatedAt: ret.CreatedAt.Unix(),
	}, nil
}

func (s *ApiService) UpdateApi(ctx context.Context, req *v1.UpdateApiRequest) (*v1.UpdateApiReply, error) {
	return nil, nil
}
