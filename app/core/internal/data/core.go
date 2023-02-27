package data

import (
	"context"
	"galileo/app/core/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type coreRepo struct {
	data *Data
	log  *log.Helper
}

func NewCoreRepo(data *Data, logger log.Logger) biz.CoreRepo {
	return &coreRepo{
		data: data,
		log:  log.NewHelper(log.With(log.With(logger, "module", "repo/core"))),
	}
}

func (r *coreRepo) CreateUser(c context.Context, user *biz.User) (*biz.User, error) {
	return nil, nil
}
