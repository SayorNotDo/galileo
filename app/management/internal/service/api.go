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

func (s *ApiService) CreateApi(ctx context.Context, req *v1.CreateApiRequest) (*v1.CreateApiReply, error) {
	return nil, nil
}
