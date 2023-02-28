package biz

import (
	"context"
	"errors"
	v1 "galileo/api/core/v1"
	"galileo/app/core/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

var (
	ErrPasswordInvalid = errors.New("password is invalid")
	ErrUsernameInvalid = errors.New("username is invalid")
	ErrPhoneInvalid    = errors.New("phone is invalid")
)

type CoreRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
}
type User struct {
	ID       uint32
	Phone    string
	Nickname string
	Gender   string
	Role     string
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
	println("debug:", createUser)
	return &v1.RegisterReply{}, nil
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
	return User{
		Phone:    phone,
		Nickname: username,
		Password: password,
	}, nil
}
