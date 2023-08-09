package service

import (
	"context"
	v1 "galileo/api/management/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

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
