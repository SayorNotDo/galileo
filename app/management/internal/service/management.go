package service

import (
	"context"
	v1 "galileo/api/management/v1"
	"galileo/app/management/internal/biz"
	"galileo/pkg/ctxdata"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ManagementService) BaseInformation(ctx context.Context, empty *empty.Empty) (*v1.BaseInfoReply, error) {
	ret, err := s.uc.BaseInformation(ctx)
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

func (s *ManagementService) CreateTestPlan(ctx context.Context, req *v1.CreateTestPlanRequest) (*v1.CreateTestPlanReply, error) {
	userId := ctxdata.GetUserId(ctx)
	createTestPlan, err := biz.NewTestPlan(userId, req.Name, req.Description, req.StartTime, req.Deadline)
	if err != nil {
		return nil, err
	}
	createTestPlan.ProjectId = req.ProjectId
	createTestPlan.Tasks = req.Tasks
	ret, err := s.uc.CreateTestPlan(ctx, &createTestPlan)
	if err != nil {
		return nil, err
	}
	return &v1.CreateTestPlanReply{
		Id: ret.Id,
	}, nil
}

func (s *ManagementService) UpdateTestPlan(ctx context.Context, req *v1.UpdateTestPlanRequest) (*empty.Empty, error) {
	userId := ctxdata.GetUserId(ctx)
	updateTestPlan, err := biz.NewTestPlan(userId, req.Name, req.Description, req.StartTime, req.Deadline)
	updateTestPlan.Tasks = req.Tasks
	if err != nil {
		return nil, err
	}
	if err := s.uc.UpdateTestPlan(ctx, &updateTestPlan); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *ManagementService) GetTestPlan(ctx context.Context, req *v1.GetTestPlanRequest) (*v1.GetTestPlanReply, error) {
	ret, err := s.uc.GetTestPlan(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetTestPlanReply{
		Id:              ret.Id,
		Name:            ret.Name,
		CreatedAt:       timestamppb.New(ret.CreatedAt),
		CreatedBy:       ret.CreatedBy,
		UpdatedAt:       timestamppb.New(ret.UpdatedAt),
		UpdatedBy:       ret.UpdatedBy,
		Description:     ret.Description,
		StartTime:       timestamppb.New(ret.StartTime),
		Deadline:        timestamppb.New(ret.Deadline),
		StatusUpdatedAt: timestamppb.New(ret.StatusUpdatedAt),
		Status:          ret.Status,
		Tasks:           ret.Tasks,
		ProjectId:       ret.ProjectId,
	}, nil
}
