package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
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
	Description string    `json:"description,omitempty"`
	Remark      string    `json:"remark"`
	Status      int8      `json:"status,omitempty"`
}

// ProjectRepo is a Greater repo.
type ProjectRepo interface {
	CreateProject(context.Context, *Project) (*Project, error)
	GetProjectById(context.Context, int64) (*Project, error)
}

// ProjectUseCase is a Project use case.
type ProjectUseCase struct {
	repo ProjectRepo
	log  *log.Helper
}

// NewProjectUseCase new a Greeter use case.
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
