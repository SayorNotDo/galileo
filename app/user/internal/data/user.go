package data

import (
	"context"
	"galileo/app/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	return user, nil
}

func (repo *userRepo) Get(ctx context.Context, id uint32) (*biz.User, error) {
	var user *biz.User
	repo.data.gormDB.Where("id = ?", id).First(&user)
	repo.log.WithContext(ctx).Info("gormDB: GetUser, id: ", id)
	return &biz.User{
		ID:        user.ID,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Status:    user.Status,
		Phone:     user.Phone,
		UpdateAt:  user.UpdateAt,
		CreatedAt: user.CreatedAt,
	}, nil
}
