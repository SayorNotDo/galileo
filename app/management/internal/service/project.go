package service

import (
	"context"
	v12 "galileo/api/management/project/v1"
	"galileo/app/management/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type ProjectService struct {
	v12.UnimplementedProjectServer
	uc     *biz.ProjectUseCase
	logger *log.Helper
}

func NewProjectService(uc *biz.ProjectUseCase, logger log.Logger) *ProjectService {
	return &ProjectService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "management.projectService")),
	}
}

func (s *ProjectService) CreateProject(ctx context.Context, req *v12.CreateProjectRequest) (*v12.CreateProjectReply, error) {
	_, err := s.uc.CreateProject(ctx, &biz.Project{})
	if err != nil {
		return nil, err
	}
	return &v12.CreateProjectReply{
		Data: &v12.ProjectInfo{
			Name:       req.Name,
			Identifier: req.Identifier,
			Remark:     req.Remark,
		},
	}, nil
}

//func (s *ProjectService) UpdateProject(ctx context.Context, req *pb.UpdateProjectRequest) (*pb.UpdateProjectReply, errResponse) {
//	return &pb.UpdateProjectReply{}, nil
//}
//func (s *ProjectService) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*pb.DeleteProjectReply, errResponse) {
//	return &pb.DeleteProjectReply{}, nil
//}
//func (s *ProjectService) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.GetProjectReply, errResponse) {
//	return &pb.GetProjectReply{}, nil
//}
//func (s *ProjectService) ListProject(ctx context.Context, req *pb.ListProjectRequest) (*pb.ListProjectReply, errResponse) {
//	return &pb.ListProjectReply{}, nil
//}