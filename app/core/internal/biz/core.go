package biz

import (
	"context"
	"fmt"
	v1 "galileo/api/core/v1"
	"galileo/app/core/internal/conf"
	"galileo/app/core/internal/pkg/middleware/auth"
	. "galileo/pkg/errors"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"time"
)

type CoreRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error)
	UserByUsername(ctx context.Context, username string) (*User, error)
	UserById(context.Context, uint32) (*User, error)
	ListUser(ctx context.Context, pageNum, pageSize int32) ([]*v1.UserDetail, int32, error)
	SoftDeleteUser(ctx context.Context, uid uint32) (bool, error)
	SetToken(ctx context.Context, username string, token string) (bool, error)
	EmptyToken(ctx context.Context) (bool, error)
}
type User struct {
	Id          uint32
	Phone       string
	ChineseName string
	Username    string
	Nickname    string
	Gender      string
	Email       string
	Role        int32
	Password    string
	Status      bool
	CreatedAt   time.Time
	DeletedAt   time.Time
	DeletedBy   uint32
}

type CoreUseCase struct {
	cRepo      CoreRepo
	log        *log.Helper
	signingKey string
}

func NewCoreUseCase(repo CoreRepo, logger log.Logger, conf *conf.Auth) *CoreUseCase {
	helper := log.NewHelper(log.With(logger, "module", "useCase/core"))
	return &CoreUseCase{cRepo: repo, log: helper, signingKey: conf.JwtKey}
}
func (c *CoreUseCase) CreateUser(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	newUser, err := NewUser(req.Phone, req.Username, req.Password, req.Email)
	if err != nil {
		return nil, err
	}
	createUser, err := c.cRepo.CreateUser(ctx, &newUser)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Id:        createUser.Id,
		Username:  createUser.Username,
		CreatedAt: createUser.CreatedAt.Unix(),
	}, nil
}

func (c *CoreUseCase) Logout(ctx context.Context) (*v1.LogoutReply, error) {
	userClaim, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, ErrUnauthenticated
	}
	uid := fmt.Sprintf("%v", userClaim.(jwt2.MapClaims)["ID"])
	ctx = metadata.AppendToClientContext(ctx, "x-md-local-uid", uid)
	if ok, err := c.cRepo.EmptyToken(ctx); !ok {
		return nil, errors.InternalServer("Internal Server Error", err.Error())
	}
	return &v1.LogoutReply{
		Success: true,
	}, nil
}

func userClaim(u *User) *auth.CustomClaims {
	return &auth.CustomClaims{
		ID:          u.Id,
		Username:    u.Username,
		AuthorityId: int(u.Role),
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),
			ExpiresAt: jwt2.NewNumericDate(time.Now().AddDate(0, 0, 1)),
			Issuer:    "Gyl",
		},
	}
}

func (c *CoreUseCase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	if len(req.Username) <= 0 {
		return nil, ErrUsernameInvalid
	}
	if len(req.Password) <= 0 {
		return nil, ErrPasswordInvalid
	}
	user, err := c.cRepo.UserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if _, passErr := c.cRepo.CheckPassword(ctx, req.Password, user.Password); passErr != nil {
		return nil, passErr
	}
	claims := userClaim(user)
	token, err := auth.CreateToken(*claims, c.signingKey)
	if err != nil {
		return nil, ErrInternalServer
	}
	return &v1.LoginReply{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data: &v1.TokenInfo{
			Type:      "Bearer",
			Token:     token,
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
		},
	}, nil
}

func (c *CoreUseCase) UserDetail(ctx context.Context, empty *emptypb.Empty) (*v1.UserDetailReply, error) {
	userClaim, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, ErrInternalServer
	}
	uid := uint32(userClaim.(jwt2.MapClaims)["ID"].(float64))
	user, err := c.cRepo.UserById(ctx, uid)
	if err != nil {
		return nil, err
	}
	return &v1.UserDetailReply{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data: &v1.UserDetail{
			Id:          user.Id,
			Username:    user.Username,
			Nickname:    user.Nickname,
			ChineseName: user.ChineseName,
			Phone:       user.Phone,
			Email:       user.Email,
			Role:        user.Role,
		},
	}, nil
}

func (c *CoreUseCase) ListUser(ctx context.Context, pageNum, pageSize int32) (*v1.ListUserReply, error) {
	user, total, err := c.cRepo.ListUser(ctx, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return &v1.ListUserReply{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data: &v1.ListUserReply_ListUser{
			Total:    total,
			UserList: user,
		},
	}, nil
}

func (c *CoreUseCase) DeleteUser(ctx context.Context, deleteId uint32) (*v1.DeleteReply, error) {
	userClaim, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, ErrInternalServer
	}
	role := int(userClaim.(jwt2.MapClaims)["AuthorityId"].(float64))
	uid := fmt.Sprintf("%v", userClaim.(jwt2.MapClaims)["ID"])
	if role == 0 {
		return nil, errors.Forbidden(http.StatusText(403), "Role must be non-zero")
	} else if role < 5 {
		return nil, errors.Forbidden(http.StatusText(403), "Permission denied")
	}
	ctx = metadata.AppendToClientContext(ctx, "x-md-local-uid", uid)
	_, err := c.cRepo.UserById(ctx, deleteId)
	if err != nil {
		return nil, err
	}
	ok, err = c.cRepo.SoftDeleteUser(ctx, deleteId)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteReply{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	}, nil
}

func (c *CoreUseCase) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) (*v1.UpdatePasswordReply, error) {
	userClaim, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, ErrInternalServer
	}
	uid := fmt.Sprintf("%v", userClaim.(jwt2.MapClaims)["ID"])
	ctx = metadata.AppendToClientContext(ctx, "x-md-local-uid", uid)
	return nil, nil
}

func (c *CoreUseCase) UpdateUserInfo(ctx context.Context, req *v1.UserInfoUpdateRequest) (*v1.UserInfoUpdateReply, error) {
	userClaim, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, ErrInternalServer
	}
	uid := fmt.Sprintf("%v", userClaim.(jwt2.MapClaims)["ID"])
	ctx = metadata.AppendToClientContext(ctx, "x-md-local-uid", uid)
	return nil, nil
}

func NewUser(phone, username, password, email string) (User, error) {
	if len(phone) <= 0 {
		return User{}, ErrPhoneInvalid
	}
	if len(username) <= 0 {
		return User{}, ErrUsernameInvalid
	}
	if len(password) <= 0 {
		return User{}, ErrPasswordInvalid
	}
	if len(email) <= 0 {
		return User{}, ErrEmailInvalid
	}
	return User{
		Phone:    phone,
		Nickname: username,
		Password: password,
		Email:    email,
	}, nil
}
