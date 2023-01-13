package biz

import (
	"context"
	v1 "galileo/api/user/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID             uint32 `json:"id"`
	Username       string `json:"name"`
	ChineseName    string `json:"chinese_name"`
	Nickname       string `json:"nickname"`
	HashedPassword []byte `json:"hashed_password"`
	Avatar         string `json:"avatar"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Status         int32  `json:"status"`
}

type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	GetById(context.Context, uint32) (*User, error)
	Create(context.Context, *User) (*User, error)
	List(ctx context.Context, pageNum, pageSize int32) ([]*v1.UserInfo, int32, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Create(ctx context.Context, user *User) (*User, error) {
	return uc.repo.Create(ctx, user)
}

func (uc *UserUseCase) Get(ctx context.Context, id uint32) (*User, error) {
	return uc.repo.GetById(ctx, id)
}

func (uc *UserUseCase) List(ctx context.Context, pageNum, pageSize int32) ([]*v1.UserInfo, int32, error) {
	return uc.repo.List(ctx, pageNum, pageSize)
}

func (uc *UserUseCase) Update(ctx context.Context, user *User) (bool, error) {
	panic("update user information")
}
