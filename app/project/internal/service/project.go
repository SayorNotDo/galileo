package service

import (
	pb "galileo/api/project/v1"
	"galileo/app/project/internal/biz"
)

type ProjectService struct {
	pb.UnimplementedProjectServer
	uc *biz.ProjectUseCase
}

func NewProjectService(uc *biz.ProjectUseCase) *ProjectService {
	return &ProjectService{uc: uc}
}

//func (s *ProjectService) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*pb.CreateProjectReply, error) {
//	return &pb.CreateProjectReply{}, nil
//}
//func (s *ProjectService) UpdateProject(ctx context.Context, req *pb.UpdateProjectRequest) (*pb.UpdateProjectReply, error) {
//	return &pb.UpdateProjectReply{}, nil
//}
//func (s *ProjectService) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*pb.DeleteProjectReply, error) {
//	return &pb.DeleteProjectReply{}, nil
//}
//func (s *ProjectService) GetProject(ctx context.Context, req *pb.GetProjectRequest) (*pb.GetProjectReply, error) {
//	return &pb.GetProjectReply{}, nil
//}
//func (s *ProjectService) ListProject(ctx context.Context, req *pb.ListProjectRequest) (*pb.ListProjectReply, error) {
//	return &pb.ListProjectReply{}, nil
//}
