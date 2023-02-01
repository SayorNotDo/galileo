package biz

import (
	"context"
	"errors"
	v1 "galileo/api/user/v1"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrPasswordInvalid     = errors.New("password is invalid")
	ErrUsernameInvalid     = errors.New("username is invalid")
	ErrPhoneInvalid        = errors.New("phone is invalid")
	ErrUserNotFound        = errors.New("user not found")
	ErrLoginFailed         = errors.New("login failed")
	ErrGenerateTokenFailed = errors.New("generate token failed")
	ErrAuthFailed          = errors.New("authentication failed")
	ErrEmailInvalid        = errors.New("email is invalid")
)

type User struct {
	ID             uint32 `json:"id"`
	Username       string `json:"name"`
	ChineseName    string `json:"chinese_name"`
	Nickname       string `json:"nickname"`
	HashedPassword string `json:"hashed_password"`
	Avatar         string `json:"avatar"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Status         int32  `json:"status"`
}

type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	GetById(context.Context, uint32) (*User, error)
	Create(context.Context, *User) (*User, error)
	List(ctx context.Context, pageNum, pageSize int32) ([]*v1.UserInfoReply, int32, error)
	Update(context.Context, *User) (bool, error)
	DeleteById(context.Context, uint32) (bool, error)
	CheckPassword(ctx context.Context, password string, hashedPassword string) (bool, error)
}

type UserUseCase struct {
	repo       UserRepo
	log        *log.Helper
	signingKey string
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	helper := log.NewHelper(log.With(logger, "module", "useCase/user"))
	return &UserUseCase{repo: repo, log: helper}
}

func (uc *UserUseCase) Create(ctx context.Context, user *User) (*User, error) {
	return uc.repo.Create(ctx, user)
}

func (uc *UserUseCase) Get(ctx context.Context, id uint32) (*User, error) {
	return uc.repo.GetById(ctx, id)
}

func (uc *UserUseCase) List(ctx context.Context, pageNum, pageSize int32) ([]*v1.UserInfoReply, int32, error) {
	return uc.repo.List(ctx, pageNum, pageSize)
}

func (uc *UserUseCase) Update(ctx context.Context, user *User) (bool, error) {
	return uc.repo.Update(ctx, user)
}

func (uc *UserUseCase) Delete(ctx context.Context, id uint32) (bool, error) {
	return uc.repo.DeleteById(ctx, id)
}

func (uc *UserUseCase) CheckPassword(ctx context.Context, password, hashedPassword string) (bool, error) {
	return uc.repo.CheckPassword(ctx, password, hashedPassword)
}
