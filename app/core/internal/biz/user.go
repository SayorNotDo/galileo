package biz

import (
	"context"
	"fmt"
	"galileo/api/core/v1"
	"galileo/app/core/internal/conf"
	"galileo/app/core/internal/pkg/middleware/auth"
	"galileo/pkg/ctxdata"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"time"
)

type UserRepo interface {
	GetUserInfo(ctx context.Context, uid uint32) (*User, error)
	CreateUser(ctx context.Context, u *User) (*User, error)
	UpdateUserInfo(ctx context.Context, user *User) (*User, error)
	UpdatePassword(ctx context.Context, password string) (bool, error)
	SoftDeleteUser(ctx context.Context, uid uint32) (bool, error)
	SetToken(ctx context.Context, token string) (string, error)
	DestroyToken(ctx context.Context) error
	ListUser(ctx context.Context, pageToken string, pageSize int32) ([]*v1.User, int32, error)
	GetUserProjectList(ctx context.Context) ([]*v1.ProjectInfo, error)
	GetUserGroup(ctx context.Context, groupId int32) (*Group, error)
	CreateUserGroup(ctx context.Context, group *Group) (*Group, error)
	UpdateUserGroup(ctx context.Context, group *Group) error
	GetUserGroupList(ctx context.Context) ([]*Group, error)
	ValidateUser(ctx context.Context, username, password string) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper

	signingKey string
}

func NewUserUseCase(repo UserRepo, logger log.Logger, conf *conf.Auth) *UserUseCase {
	helper := log.NewHelper(log.With(logger, "module", "user.useCase"))
	return &UserUseCase{
		repo:       repo,
		log:        helper,
		signingKey: conf.JwtKey,
	}
}

func NewUser(phone, username, email string) (User, error) {
	if len(phone) <= 0 {
		return User{}, SetCustomizeErrMsg(ReasonParamsError, "phone's number can not be empty")
	} else if len(username) <= 0 {
		return User{}, SetCustomizeErrMsg(ReasonParamsError, "username can not be empty")
	} else if len(email) <= 0 {
		return User{}, SetCustomizeErrMsg(ReasonParamsError, "email value illegal")
	}
	return User{
		Phone:    phone,
		Username: username,
		Email:    email,
	}, nil
}

func newUserClaim(u *User) *auth.CustomClaims {
	return &auth.CustomClaims{
		ID:          u.ID,
		Username:    u.Username,
		AuthorityId: u.UUID,
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),
			ExpiresAt: jwt2.NewNumericDate(time.Now().AddDate(0, 1, 0)),
			IssuedAt:  jwt2.NewNumericDate(time.Now()),
			Issuer:    ctxdata.Issuer,
		},
	}
}

func (u *UserUseCase) CreateUser(ctx context.Context, user *User) (*User, error) {
	return u.repo.CreateUser(ctx, user)
}

func (u *UserUseCase) UpdateUserInfo(ctx context.Context, user *User) (*User, error) {
	return u.repo.UpdateUserInfo(ctx, user)
}

func (u *UserUseCase) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) (*emptypb.Empty, error) {
	userClaim, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, SetErrByReason(ReasonUnknownError)
	}
	uid := fmt.Sprintf("%v", userClaim.(jwt2.MapClaims)["ID"])
	ctx = metadata.AppendToClientContext(ctx, ctxdata.UserIdKey, uid)
	if ok, _ := u.repo.UpdatePassword(ctx, req.Password); !ok {
		return nil, SetErrByReason(ReasonUnknownError)
	}
	err := u.repo.DestroyToken(ctx)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (u *UserUseCase) DeleteUser(ctx context.Context, deleteId uint32) (*emptypb.Empty, error) {
	userClaim, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, SetErrByReason(ReasonUnknownError)
	}
	role := int(userClaim.(jwt2.MapClaims)["AuthorityId"].(float64))
	uid := fmt.Sprintf("%v", userClaim.(jwt2.MapClaims)["ID"])
	if role == 0 {
		return nil, errors.Forbidden(http.StatusText(403), "Role must be non-zero")
	} else if role < 5 {
		return nil, errors.Forbidden(http.StatusText(403), "Permission denied")
	}
	ctx = metadata.AppendToClientContext(ctx, ctxdata.UserIdKey, uid)
	_, err := u.repo.GetUserInfo(ctx, deleteId)
	if err != nil {
		return nil, err
	}
	ok, err = u.repo.SoftDeleteUser(ctx, deleteId)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Login TODO: 修复逻辑，禁止密码的传递
func (u *UserUseCase) Login(ctx context.Context, req *v1.LoginRequest) (string, error) {
	user, err := u.repo.ValidateUser(ctx, req.Username, req.Password)
	if err != nil {
		return "", err
	}
	claims := newUserClaim(user)
	token, err := auth.CreateToken(*claims, u.signingKey)
	if err != nil {
		return "", SetErrByReason(ReasonParamsError)
	}
	return u.repo.SetToken(ctx, token)
}

func (u *UserUseCase) Logout(ctx context.Context) (*emptypb.Empty, error) {
	if err := u.repo.DestroyToken(ctx); err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *UserUseCase) GetUserInfo(ctx context.Context, uid uint32) (*User, error) {
	user, err := u.repo.GetUserInfo(ctx, uid)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:            user.ID,
		Username:      user.Username,
		ChineseName:   user.ChineseName,
		Phone:         user.Phone,
		Email:         user.Email,
		Avatar:        user.Avatar,
		Active:        user.Active,
		Location:      user.Location,
		CreatedAt:     user.CreatedAt,
		LastLoginTime: user.LastLoginTime,
	}, nil
}

func (u *UserUseCase) ListUser(ctx context.Context, pageToken string, pageSize int32) (*v1.ListUserReply, error) {
	user, total, err := u.repo.ListUser(ctx, pageToken, pageSize)
	if err != nil {
		return nil, err
	}
	return &v1.ListUserReply{
		Total:    total,
		UserList: user,
	}, nil
}

func (u *UserUseCase) GetUserProjectList(ctx context.Context) ([]*v1.ProjectInfo, error) {
	return u.repo.GetUserProjectList(ctx)
}

func (u *UserUseCase) GetUserGroupList(ctx context.Context) ([]*Group, error) {
	return u.repo.GetUserGroupList(ctx)
}

func (u *UserUseCase) GetUserGroup(ctx context.Context, groupId int32) (*Group, error) {
	return u.repo.GetUserGroup(ctx, groupId)
}

func (u *UserUseCase) UpdateUserGroup(ctx context.Context, group *Group) error {
	return u.repo.UpdateUserGroup(ctx, group)
}

func (u *UserUseCase) CreateUserGroup(ctx context.Context, group *Group) (*Group, error) {
	return u.repo.CreateUserGroup(ctx, group)
}
