package data

import (
	"context"
	"galileo/app/core/internal/biz"
	"galileo/ent"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "user.Repo")),
	}
}

func (repo userRepo) GetUserProjectList(ctx context.Context, uid uint32) ([]*ent.Project, error) {
	return nil, nil
}
