package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Project is a Project model.
type Project struct {
}

// ProjectRepo is a Greater repo.
type ProjectRepo interface {
}

// ProjectUseCase is a Project usecase.
type ProjectUseCase struct {
	repo ProjectRepo
	log  *log.Helper
}

// NewProjectUseCase new a Greeter usecase.
func NewProjectUseCase(repo ProjectRepo, logger log.Logger) *ProjectUseCase {
	return &ProjectUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *ProjectUseCase) CreateGreeter(ctx context.Context, g *Project) (*Project, error) {
	return nil, nil
}
