package biz

import (
	"context"
	"galileo/api/core/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	GetUserProjectList(ctx context.Context) ([]*v1.ProjectInfo, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	helper := log.NewHelper(log.With(logger, "module", "core.useCase"))
	return &UserUseCase{
		repo: repo,
		log:  helper,
	}
}

func (e *UserUseCase) GetUserProjectList(ctx context.Context) ([]*v1.ProjectInfo, error) {
	return e.repo.GetUserProjectList(ctx)
}
