package biz

import (
	"context"
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
	Get(context.Context, uint32) (*User, error)
	Create(context.Context, *User) (*User, error)
	List(ctx context.Context, pageNum, pageSize int32) ([]*User, int32, error)
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
	return uc.repo.Get(ctx, id)
}

func (uc *UserUseCase) List(ctx context.Context, pageNum, pageSize int32) ([]*User, int32, error) {
	return uc.repo.List(ctx, pageNum, pageSize)
}
