package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Project is a Project model.
type Project struct {
	ID         int64
	Name       string
	Identifier string
	Status     uint32
}

// ProjectRepo is a Greater repo.
type ProjectRepo interface {
	CreateProject(context.Context, *Project) (*Project, error)
}

// ProjectUseCase is a Project use case.
type ProjectUseCase struct {
	repo ProjectRepo
	log  *log.Helper
}

// NewProjectUseCase new a Greeter use case.
func NewProjectUseCase(repo ProjectRepo, logger log.Logger) *ProjectUseCase {
	helper := log.NewHelper(log.With(logger, "module", "useCase/project"))
	return &ProjectUseCase{repo: repo, log: helper}
}

// CreateProject creates a Project, and returns the new Project.
func (uc *ProjectUseCase) CreateProject(ctx context.Context, project *Project) (*Project, error) {
	return uc.repo.CreateProject(ctx, project)
}
