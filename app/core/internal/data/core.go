package data

import (
	"context"
	v1 "galileo/api/core/v1"
	userService "galileo/api/user/v1"
	"galileo/app/core/internal/biz"
	"galileo/app/core/internal/pkg/middleware/auth"
	"galileo/pkg/encryption"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"strings"
	"time"
)

type coreRepo struct {
	data *Data
	log  *log.Helper
}

func NewCoreRepo(data *Data, logger log.Logger) biz.CoreRepo {
	return &coreRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "core.Repo")),
	}
}

func (r *coreRepo) UserByUsername(c context.Context, username string) (*biz.User, error) {
	user, err := r.data.uc.GetUserByUsername(c, &userService.UsernameRequest{
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	log.Debugf("UserByUsername %v", user)
	return &biz.User{
		Id:       user.Id,
		Phone:    user.Phone,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}, nil
}

func (r *coreRepo) UserById(c context.Context, id uint32) (*biz.User, error) {
	user, err := r.data.uc.GetUser(c, &userService.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
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
		Id:        createUser.Id,
		Username:  createUser.Username,
		CreatedAt: time.Unix(createUser.CreatedAt, 0),
	}, nil
}

func (r *coreRepo) UpdatePassword(ctx context.Context, password string) (bool, error) {
	if _, err := r.data.uc.UpdatePassword(ctx, &userService.UpdatePasswordRequest{NewPassword: password}); err != nil {
		return false, err
	}
	return true, nil
}

func (r *coreRepo) VerifyPassword(c context.Context, password, encryptedPassword string) (bool, error) {
	if ret, err := r.data.uc.VerifyPassword(c, &userService.VerifyPasswordRequest{Password: password, HashedPassword: encryptedPassword}); err != nil {
		return false, err
	} else {
		return ret.Success, nil
	}
}

func (r *coreRepo) ListUser(c context.Context, pageNum, pageSize int32) ([]*v1.UserDetail, int32, error) {
	rsp, err := r.data.uc.ListUser(c, &userService.ListUserRequest{PageNum: pageNum, PageSize: pageSize})
	if err != nil {
		return nil, 0, err
	}
	rv := make([]*v1.UserDetail, 0)
	for _, u := range rsp.Data {
		rv = append(rv, &v1.UserDetail{
			Id:          u.Id,
			Username:    u.Username,
			Nickname:    u.Nickname,
			ChineseName: u.ChineseName,
			Phone:       u.Phone,
			Email:       u.Email,
			Role:        u.Role,
		})
	}
	total := rsp.Total
	return rv, total, nil
}

func (r *coreRepo) SoftDeleteUser(c context.Context, uid uint32) (bool, error) {
	rsp, err := r.data.uc.SoftDeleteUser(c, &userService.SoftDeleteRequest{Id: uid})
	if err != nil {
		return false, err
	}
	return rsp.Deleted, nil
}

func (r *coreRepo) UpdateUserInfo(c context.Context, user *biz.User) (bool, error) {
	return false, nil
}

func (r *coreRepo) SetToken(ctx context.Context, token string) (string, error) {
	key := encryption.EncodeMD5(token)
	r.data.redisCli.Set(ctx, "token:"+key, token, auth.TokenExpiration)
	return key, nil
}

func (r *coreRepo) DestroyToken(ctx context.Context) error {
	tr, _ := transport.FromServerContext(ctx)
	jwtToken := strings.SplitN(tr.RequestHeader().Get("Authorization"), " ", 2)[1]
	key := encryption.EncodeMD5(jwtToken)
	r.data.redisCli.Del(ctx, "token:"+key)
	return nil
}
