package biz

import (
	"context"
	"fmt"
	v1 "galileo/api/core/v1"
	"galileo/app/core/internal/conf"
	"galileo/app/core/internal/pkg/constant"
	"galileo/app/core/internal/pkg/middleware/auth"
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

type CoreRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	VerifyPassword(ctx context.Context, password, encryptedPassword string) (bool, error)
	UserByUsername(ctx context.Context, username string) (*User, error)
	UserById(context.Context, uint32) (*User, error)
	ListUser(ctx context.Context, pageNum, pageSize int32) ([]*v1.UserDetail, int32, error)
	SoftDeleteUser(ctx context.Context, uid uint32) (bool, error)
	SetToken(ctx context.Context, token string) (string, error)
	DestroyToken(ctx context.Context) error
	UpdatePassword(ctx context.Context, password string) (bool, error)
	DataReportTrack(ctx context.Context, data []map[string]interface{}, claims *auth.ReportClaims) error
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

func NewCoreUseCase(repo CoreRepo, logger log.Logger, conf *conf.Auth) *CoreUseCase {
	helper := log.NewHelper(log.With(logger, "module", "core.useCase"))
	return &CoreUseCase{repo: repo, log: helper, signingKey: conf.JwtKey}
}
func (c *CoreUseCase) CreateUser(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	newUser, err := NewUser(req.Phone, req.Username, req.Password, req.Email)
	if err != nil {
		return nil, err
	}
	createUser, err := c.repo.CreateUser(ctx, &newUser)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Id:        createUser.Id,
		Username:  createUser.Username,
		CreatedAt: timestamppb.New(createUser.CreatedAt),
	}, nil
}

func (c *CoreUseCase) Logout(ctx context.Context) (*emptypb.Empty, error) {
	claims, _ := jwt.FromContext(ctx)
	id := uint32(claims.(jwt2.MapClaims)["ID"].(float64))
	if _, err := c.repo.UserById(ctx, id); err != nil {
		return nil, err
	}
	_ = c.repo.DestroyToken(ctx)
	return nil, nil
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

func newReportClaim(machine string) *auth.ReportClaims {
	return &auth.ReportClaims{
		Machine: machine,
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),
			ExpiresAt: jwt2.NewNumericDate(time.Now().AddDate(1, 0, 0)),
			IssuedAt:  jwt2.NewNumericDate(time.Now()),
			Issuer:    "Galileo",
		},
	}
}

func (c *CoreUseCase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	if len(req.Username) <= 0 {
		return nil, SetErrByReason(ReasonParamsError)
	}
	if len(req.Password) <= 0 {
		return nil, SetErrByReason(ReasonParamsError)
	}
	user, err := c.repo.UserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if _, passErr := c.repo.VerifyPassword(ctx, req.Password, user.Password); passErr != nil {
		return nil, passErr
	}
	claims := newUserClaim(user)
	token, err := auth.CreateToken(*claims, c.signingKey)
	if err != nil {
		return nil, SetErrByReason(ReasonParamsError)
	}
	encryptionToken, _ := c.repo.SetToken(ctx, token)
	return &v1.LoginReply{
		Type:      "Bearer",
		Token:     encryptionToken,
		ExpiresAt: timestamppb.New(time.Unix(claims.ExpiresAt.Unix(), 0)),
	}, nil
}

func (c *CoreUseCase) UserDetail(ctx context.Context, empty *emptypb.Empty) (*v1.UserDetailReply, error) {
	userClaim, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, SetErrByReason(ReasonUnknownError)
	}
	uid := uint32(userClaim.(jwt2.MapClaims)["ID"].(float64))
	user, err := c.repo.UserById(ctx, uid)
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

func (c *CoreUseCase) ListUser(ctx context.Context, pageNum, pageSize int32) (*v1.ListUserReply, error) {
	user, total, err := c.repo.ListUser(ctx, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return &v1.ListUserReply{
		Total:    total,
		UserList: user,
	}, nil
}

func (c *CoreUseCase) DeleteUser(ctx context.Context, deleteId uint32) (*v1.DeleteReply, error) {
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
	ctx = metadata.AppendToClientContext(ctx, "x-md-local-uid", uid)
	_, err := c.repo.UserById(ctx, deleteId)
	if err != nil {
		return nil, err
	}
	ok, err = c.repo.SoftDeleteUser(ctx, deleteId)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteReply{
		Data: nil,
	}, nil
}

func (c *CoreUseCase) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) (*emptypb.Empty, error) {
	userClaim, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, SetErrByReason(ReasonUnknownError)
	}
	uid := fmt.Sprintf("%v", userClaim.(jwt2.MapClaims)["ID"])
	ctx = metadata.AppendToClientContext(ctx, "x-md-local-uid", uid)
	if ok, _ := c.repo.UpdatePassword(ctx, req.Password); !ok {
		return nil, SetErrByReason(ReasonUnknownError)
	}
	err := c.repo.DestroyToken(ctx)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (c *CoreUseCase) UpdateUserInfo(ctx context.Context, req *v1.UserInfoUpdateRequest) (*v1.UserInfoUpdateReply, error) {
	userClaim, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, SetErrByReason(ReasonUnknownError)
	}
	uid := fmt.Sprintf("%v", userClaim.(jwt2.MapClaims)["ID"])
	ctx = metadata.AppendToClientContext(ctx, "x-md-local-uid", uid)
	return nil, nil
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

func (c *CoreUseCase) DataReportTrack(ctx context.Context, data []map[string]interface{}, claims *auth.ReportClaims) error {
	return c.repo.DataReportTrack(ctx, data, claims)
}

func (c *CoreUseCase) ExecuteToken(ctx context.Context, machine string) (string, error) {
	if len(machine) <= 0 {
		return "", SetErrByReason(ReasonParamsError)
	}
	claims := newReportClaim(machine)
	token, err := auth.GenerateExecuteToken(*claims, constant.ReportKey)
	if err != nil {
		return "", SetErrByReason(ReasonSystemError)
	}
	return token, nil
}
