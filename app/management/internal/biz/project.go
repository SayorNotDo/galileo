package biz

import (
	"context"
	v1 "galileo/api/management/v1"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// Project is a Project model.
type Project struct {
	ID          int64     `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Identifier  string    `json:"identifier,omitempty"`
	CreatedBy   uint32    `json:"created_by,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedBy   uint32    `json:"updated_by,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	StartTime   time.Time `json:"start_time,omitempty"`
	Deadline    time.Time `json:"deadline,omitempty"`
	Description string    `json:"description,omitempty"`
	Remark      string    `json:"remark"`
	Status      v1.Status `json:"status,omitempty"`
}

// ProjectRepo is a Project repo.
type ProjectRepo interface {
	CreateProject(context.Context, *Project) (*Project, error)
	GetProjectById(context.Context, int64) (*Project, error)
	UpdateProject(context.Context, *Project) error
	GetUserProjectList(context.Context, uint32) ([]*v1.ProjectInfo, error)
}

// ProjectUseCase is a Project use case.
type ProjectUseCase struct {
	repo ProjectRepo
	log  *log.Helper
}

// NewProjectUseCase new a Project use case.
func NewProjectUseCase(repo ProjectRepo, logger log.Logger) *ProjectUseCase {
	return &ProjectUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "management.projectUseCase")),
	}
}

// CreateProject creates a Project, and returns the new Project.
func (uc *ProjectUseCase) CreateProject(ctx context.Context, project *Project) (*Project, error) {
	return uc.repo.CreateProject(ctx, project)
}

// GetProjectById returns the Project
func (uc *ProjectUseCase) GetProjectById(ctx context.Context, id int64) (*Project, error) {
	return uc.repo.GetProjectById(ctx, id)
}

// UpdateProject updates the Project
func (uc *ProjectUseCase) UpdateProject(ctx context.Context, project *Project) error {
	return uc.repo.UpdateProject(ctx, project)
}

func (uc *ProjectUseCase) GetUserProjectList(ctx context.Context, uid uint32) ([]*v1.ProjectInfo, error) {
	return uc.repo.GetUserProjectList(ctx, uid)
}

func NewProject(name, identifier string, createdBy uint32, startTime *timestamppb.Timestamp, endTime *timestamppb.Timestamp) (Project, error) {
	if len(name) <= 0 {
		return Project{}, SetCustomizeErrMsg(ReasonParamsError, "project name must not be empty")
	} else if len(identifier) <= 0 {
		return Project{}, SetCustomizeErrMsg(ReasonParamsError, "identifier must not be empty")
	} else if createdBy <= 0 {
		return Project{}, SetCustomizeErrMsg(ReasonParamsError, "creator id must be greater than zero")
	}
	return Project{
		Name:       name,
		Identifier: identifier,
		CreatedBy:  createdBy,
		StartTime:  startTime.AsTime(),
		Deadline:   endTime.AsTime(),
	}, nil
}
