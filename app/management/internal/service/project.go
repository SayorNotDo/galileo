package service

import (
	"context"
	v1 "galileo/api/management/project/v1"
	"galileo/app/management/internal/biz"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	if req.Name == "" {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "project name must not be empty")
	} else if req.CreatedBy <= 0 {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "creator id must be greater than zero")
	} else if req.Identifier == "" {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "identifier must not be empty")
	}
	ret, err := s.uc.CreateProject(ctx, &biz.Project{
		Name:        req.Name,
		Identifier:  req.Identifier,
		CreatedBy:   req.CreatedBy,
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

func (s *ProjectService) GetProject(ctx context.Context, req *v1.GetProjectRequest) (*v1.ProjectInfo, error) {
	if req.Id <= 0 {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "project id must be greater than zero")
	}
	ret, err := s.uc.GetProjectById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ProjectInfo{
		Id:          ret.ID,
		Name:        ret.Name,
		Identifier:  ret.Identifier,
		Status:      int32(ret.Status),
		CreatedAt:   timestamppb.New(ret.CreatedAt),
		CreatedBy:   ret.CreatedBy,
		UpdatedAt:   timestamppb.New(ret.UpdatedAt),
		UpdatedBy:   ret.UpdatedBy,
		Description: ret.Description,
		Remark:      ret.Remark,
	}, nil
}

//func (s *ProjectService) UpdateProject(ctx context.Context, req *pb.UpdateProjectRequest) (*pb.UpdateProjectReply, errResponse) {
//	return &pb.UpdateProjectReply{}, nil
//}
//func (s *ProjectService) DeleteProject(ctx context.Context, req *pb.DeleteProjectRequest) (*pb.DeleteProjectReply, errResponse) {
//	return &pb.DeleteProjectReply{}, nil
//}

//func (s *ProjectService) ListProject(ctx context.Context, req *pb.ListProjectRequest) (*pb.ListProjectReply, errResponse) {
//	return &pb.ListProjectReply{}, nil
//}
