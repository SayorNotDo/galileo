package service

import (
	"context"
	v1 "galileo/api/project/v1"
	"galileo/app/project/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type ProjectService struct {
	v1.UnimplementedProjectServer
	uc     *biz.ProjectUseCase
	logger *log.Helper
}

func NewProjectService(uc *biz.ProjectUseCase, logger log.Logger) *ProjectService {
	return &ProjectService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "project")),
	}
}

func (s *ProjectService) CreateProject(ctx context.Context, req *v1.CreateProjectRequest) (*v1.CreateProjectReply, error) {
	_, err := s.uc.CreateProject(ctx, &biz.Project{})
	if err != nil {
		return nil, err
	}
	return &v1.CreateProjectReply{
		Data: &v1.ProjectInfo{
			Name:       req.Name,
			Identifier: req.Identifier,
			Remark:     req.Remark,
		},
	}, nil
}

//func (s *ProjectService) UpdateProject(ctx context.Context, req *pb.UpdateProjectRequest) (*pb.UpdateProjectReply, errors) {
//	return &pb.UpdateProjectReply{}, nil
//}
//func (s *ProjectService) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*pb.DeleteProjectReply, errors) {
//	return &pb.DeleteProjectReply{}, nil
//}
//func (s *ProjectService) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.GetProjectReply, errors) {
//	return &pb.GetProjectReply{}, nil
//}
//func (s *ProjectService) ListProject(ctx context.Context, req *pb.ListProjectRequest) (*pb.ListProjectReply, errors) {
//	return &pb.ListProjectReply{}, nil
//}
