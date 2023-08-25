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
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"
)

type UserRepo interface {
	GetUserInfo(context.Context) (*User, error)
	CreateUser(ctx context.Context, u *User) (*User, error)
	UpdatePassword(ctx context.Context, password string) (bool, error)
	SoftDeleteUser(ctx context.Context, uid uint32) (bool, error)
	UserByUsername(ctx context.Context, username string) (*User, error)
	SetToken(ctx context.Context, token string) (string, error)
	DestroyToken(ctx context.Context) error
	ListUser(ctx context.Context, pageNum, pageSize int32) ([]*v1.UserDetail, int32, error)
	GetUserProjectList(ctx context.Context) ([]*v1.ProjectInfo, error)
	GetUserGroup(ctx context.Context, groupId int32) (*Group, error)
	CreateUserGroup(ctx context.Context, group *Group) (*Group, error)
	UpdateUserGroup(ctx context.Context, group *Group) error
	GetUserGroupList(ctx context.Context) ([]*Group, error)
	VerifyPassword(ctx context.Context, password, encryptedPassword string) (bool, error)
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

func NewUser(phone, username, password, email string) (User, error) {
	if len(phone) <= 0 {
		return User{}, SetErrByReason(ReasonParamsError)
	}
	if len(username) <= 0 {
		return User{}, SetErrByReason(ReasonParamsError)
	}
	if len(password) <= 0 {
		return User{}, SetErrByReason(ReasonParamsError)
	}
	if len(email) <= 0 {
		return User{}, SetErrByReason(ReasonParamsError)
	}
	return User{
		Phone:    phone,
		Nickname: username,
		Password: password,
		Email:    email,
	}, nil
}

func newUserClaim(u *User) *auth.CustomClaims {
	return &auth.CustomClaims{
		ID:          u.Id,
		Username:    u.Username,
		AuthorityId: int(u.Role),
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),
			ExpiresAt: jwt2.NewNumericDate(time.Now().AddDate(0, 1, 0)),
			IssuedAt:  jwt2.NewNumericDate(time.Now()),
			Issuer:    "Galileo",
		},
	}
}

func (u *UserUseCase) CreateUser(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	newUser, err := NewUser(req.Phone, req.Username, req.Password, req.Email)
	if err != nil {
		return nil, err
	}
	createUser, err := u.repo.CreateUser(ctx, &newUser)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Id:        createUser.Id,
		Username:  createUser.Username,
		CreatedAt: timestamppb.New(createUser.CreatedAt),
	}, nil
}

func (u *UserUseCase) UpdateUserInfo(ctx context.Context, req *v1.UserInfoRequest) (*v1.UserInfoReply, error) {
	userClaim, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, SetErrByReason(ReasonUnknownError)
	}
	uid := fmt.Sprintf("%v", userClaim.(jwt2.MapClaims)["ID"])
	ctx = metadata.AppendToClientContext(ctx, "x-md-local-uid", uid)
	return nil, nil
}

func (u *UserUseCase) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) (*emptypb.Empty, error) {
	userClaim, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, SetErrByReason(ReasonUnknownError)
	}
	uid := fmt.Sprintf("%v", userClaim.(jwt2.MapClaims)["ID"])
	ctx = metadata.AppendToClientContext(ctx, "x-md-local-uid", uid)
	if ok, _ := u.repo.UpdatePassword(ctx, req.Password); !ok {
		return nil, SetErrByReason(ReasonUnknownError)
	}
	err := u.repo.DestroyToken(ctx)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (u *UserUseCase) DeleteUser(ctx context.Context, deleteId uint32) (*v1.DeleteReply, error) {
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
	_, err := u.repo.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	ok, err = u.repo.SoftDeleteUser(ctx, deleteId)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteReply{
		Data: nil,
	}, nil
}

// Login TODO: 修复逻辑，禁止密码的传递
func (u *UserUseCase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	if len(req.Username) <= 0 {
		return nil, SetErrByReason(ReasonParamsError)
	}
	if len(req.Password) <= 0 {
		return nil, SetErrByReason(ReasonParamsError)
	}
	user, err := u.repo.UserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if _, passErr := u.repo.VerifyPassword(ctx, req.Password, user.Password); passErr != nil {
		return nil, passErr
	}
	claims := newUserClaim(user)
	token, err := auth.CreateToken(*claims, u.signingKey)
	if err != nil {
		return nil, SetErrByReason(ReasonParamsError)
	}
	encryptionToken, _ := u.repo.SetToken(ctx, token)
	return &v1.LoginReply{
		Type:      "Bearer",
		Token:     encryptionToken,
		ExpiresAt: timestamppb.New(time.Unix(claims.ExpiresAt.Unix(), 0)),
	}, nil
}

func (u *UserUseCase) Logout(ctx context.Context) (*emptypb.Empty, error) {
	if _, err := u.repo.GetUserInfo(ctx); err != nil {
		return nil, err
	}
	_ = u.repo.DestroyToken(ctx)
	return nil, nil
}

func (u *UserUseCase) GetUserInfo(ctx context.Context) (*v1.UserDetailReply, error) {
	user, err := u.repo.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UserDetailReply{
		Id:          user.Id,
		Nickname:    user.Nickname,
		ChineseName: user.ChineseName,
		Phone:       user.Phone,
		Email:       user.Email,
		Role:        user.Role,
	}, nil
}

func (u *UserUseCase) ListUser(ctx context.Context, pageNum, pageSize int32) (*v1.ListUserReply, error) {
	user, total, err := u.repo.ListUser(ctx, pageNum, pageSize)
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
