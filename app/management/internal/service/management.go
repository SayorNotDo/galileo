package service

import (
	"context"
	v1 "galileo/api/management/v1"
	"galileo/app/management/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
)

type ManagementService struct {
	v1.UnimplementedManagementServer

	uc     *biz.ManagementUseCase
	logger *log.Helper
}

func NewManagementService(uc *biz.ManagementUseCase, logger log.Logger) *ManagementService {
	return &ManagementService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "management.Service")),
	}
}

func (s *ManagementService) BaseInfo(ctx context.Context, empty *empty.Empty) (*v1.BaseInfoReply, error) {
	ret, err := s.uc.BaseInfo(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.BaseInfoReply{
		ApiCount:       ret.ApiCount,
		SceneCaseCount: ret.SceneCaseCount,
		ApiCaseCount:   ret.ApiCaseCount,
		CronJobCount:   ret.CronJobCount,
	}, nil
}
