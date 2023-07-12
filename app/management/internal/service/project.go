package service

import (
	"context"
	v1 "galileo/api/management/project/v1"
	"galileo/app/management/internal/biz"
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
		logger: log.NewHelper(log.With(logger, "module", "management.projectService")),
	}
}

func (s *ProjectService) CreateProject(ctx context.Context, req *v1.CreateProjectRequest) (*v1.CreateProjectReply, error) {
	ret, err := s.uc.CreateProject(ctx, &biz.Project{
		Name:        req.Name,
		Identifier:  req.Identifier,
		Description: req.Description,
		Remark:      req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateProjectReply{
		Id: ret.ID,
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
