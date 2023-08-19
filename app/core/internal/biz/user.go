package biz

import (
	"context"
	"galileo/ent"
)

type UserRepo interface {
	GetUserProjectList(ctx context.Context, uid uint32) ([]*ent.Project, error)
}

func (e *UserUseCase) GetUserProjectList(ctx context.Context, uid uint32) ([]*ent.Project, error) {
	return e.repo.GetUserProjectList(ctx, uid)
}
