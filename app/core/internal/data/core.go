package data

import (
	"context"
	userService "galileo/api/user/v1"
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
		log:  log.NewHelper(log.With(logger, "module", "repo/core")),
	}
}

func (r *coreRepo) UserByUsername(c context.Context, username string) (*biz.User, error) {
	user, err := r.data.uc.GetUserByUsername(c, &userService.UsernameRequest{
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:       user.Id,
		Phone:    user.Phone,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (r *coreRepo) CreateUser(c context.Context, user *biz.User) (*biz.User, error) {
	createUser, err := r.data.uc.CreateUser(c, &userService.CreateUserRequest{
		Username: user.Nickname,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:       createUser.Data.Id,
		Phone:    createUser.Data.Phone,
		Nickname: createUser.Data.Nickname,
		Email:    createUser.Data.Email,
	}, nil
}

func (r *coreRepo) CheckPassword(c context.Context, password, encryptedPassword string) (bool, error) {
	if ret, err := r.data.uc.CheckPassword(c, &userService.CheckPasswordRequest{Password: password, HashedPassword: encryptedPassword}); err != nil {
		return false, err
	} else {
		return ret.Success, nil
	}
}
