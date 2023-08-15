package service

import (
	"context"
	v1 "galileo/api/management/project/v1"
	"galileo/app/management/internal/biz"
	. "galileo/pkg/errResponse"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ManagementService) CreateProject(ctx context.Context, req *v1.CreateProjectRequest) (*v1.CreateProjectReply, error) {
	userId := ctx.Value("x-md-global-userId")
	createProject, err := biz.NewProject(req.Name, req.Identifier, userId.(uint32), req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}
	ret, err := s.pc.CreateProject(ctx, &createProject)
	if err != nil {
		return nil, err
	}
	return &v1.CreateProjectReply{
		Id: ret.ID,
	}, nil
}

func (s *ManagementService) GetProject(ctx context.Context, req *v1.GetProjectRequest) (*v1.ProjectInfo, error) {
	if req.Id <= 0 {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "project id must be greater than zero")
	}
	ret, err := s.pc.GetProjectById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ProjectInfo{
		Id:          ret.ID,
		Name:        ret.Name,
		Identifier:  ret.Identifier,
		Status:      ret.Status,
		CreatedAt:   timestamppb.New(ret.CreatedAt),
		CreatedBy:   ret.CreatedBy,
		UpdatedAt:   timestamppb.New(ret.UpdatedAt),
		UpdatedBy:   ret.UpdatedBy,
		Description: ret.Description,
		Remark:      ret.Remark,
	}, nil
}

func (s *ManagementService) UpdateProject(ctx context.Context, req *v1.UpdateProjectRequest) (*empty.Empty, error) {
	if req.Id <= 0 {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "project id must be greater than zero")
	}
	ret, err := s.pc.GetProjectById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	updateProject, err := biz.NewProject(req.Name, req.Identifier, ret.CreatedBy, req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}
	updateProject.ID = req.Id
	updateProject.Description = req.Description
	updateProject.Remark = req.Remark
	updateProject.Status = req.Status
	if err := s.pc.UpdateProject(ctx, &updateProject); err != nil {
		return nil, err
	}
	return nil, nil
}

//func (s *ProjectService) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*pb.DeleteProjectReply, errResponse) {
//	return &pb.DeleteProjectReply{}, nil
//}

//func (s *ProjectService) ListProject(ctx context.Context, req *pb.ListProjectRequest) (*pb.ListProjectReply, errResponse) {
//	return &pb.ListProjectReply{}, nil
//}
