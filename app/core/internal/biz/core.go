package biz

import (
	"context"
	"errors"
	v1 "galileo/api/core/v1"
	"galileo/app/core/internal/conf"
	"galileo/app/core/internal/pkg/middleware/auth"
	"github.com/go-kratos/kratos/v2/log"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	ErrPasswordInvalid = errors.New("password is invalid")
	ErrUsernameInvalid = errors.New("username is invalid")
	ErrPhoneInvalid    = errors.New("phone is invalid")
	ErrEmailInvalid    = errors.New("email is invalid")
)

type CoreRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
}
type User struct {
	ID       uint32
	Phone    string
	Nickname string
	Gender   string
	Email    string
	Role     int
	Password string
	CreateAt time.Time
}

type CoreUseCase struct {
	cRepo      CoreRepo
	log        *log.Helper
	signingKey string
}

func NewCoreUseCase(repo CoreRepo, logger log.Logger, conf *conf.Auth) *CoreUseCase {
	helper := log.NewHelper(log.With(logger, "module", "useCase/user"))
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
	claims := auth.CustomClaims{
		ID:          createUser.ID,
		Nickname:    createUser.Nickname,
		AuthorityId: createUser.Role,
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),
			ExpiresAt: jwt2.NewNumericDate(time.Now().AddDate(0, 0, 30)),
			Issuer:    "Gyl",
		},
	}
	token, err := auth.CreateToken(claims, c.signingKey)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Id:        createUser.ID,
		Phone:     createUser.Phone,
		Username:  createUser.Nickname,
		Token:     token,
		Email:     createUser.Email,
		ExpiresAt: jwt2.NewNumericDate(time.Now().AddDate(0, 0, 30)).Unix(),
	}, nil
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
