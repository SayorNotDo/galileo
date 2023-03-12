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
		Id:       user.Id,
		Phone:    user.Phone,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (r *coreRepo) UserById(c context.Context, id uint32) (*biz.User, error) {
	user, err := r.data.uc.GetUser(c, &userService.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
	log.Debugf("UserById user: %v", user.Id)
	return &biz.User{
		Id:       user.Id,
		Username: user.Username,
		Phone:    user.Phone,
		Nickname: user.Nickname,
		Email:    user.Email,
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
		Id:       createUser.Data.Id,
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

func (r *coreRepo) ListUser(c context.Context, pageNum, pageSize int32) ([]*biz.User, error) {
	userList, err := r.data.uc.ListUser(c, &userService.ListUserRequest{PageNum: pageNum, PageSize: pageSize})
	if err != nil {
		return nil, err
	}
	log.Debugf("%v", userList)
	return nil, nil
}
