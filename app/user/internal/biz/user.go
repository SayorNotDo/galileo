package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type User struct {
	ID        uint32    `json:"id"`
	Username  string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Status    int32     `json:"status"`
	UpdateAt  time.Time `json:"update_at"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Get(context.Context, uint32) (*User, error)
}

type UserCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserCase(repo UserRepo, logger log.Logger) *UserCase {
	return &UserCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserCase) CreateUser(ctx context.Context, user *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", user.ID)
	return uc.repo.Save(ctx, user)
}

func (uc *UserCase) Get(ctx context.Context, id uint32) (*User, error) {
	uc.log.WithContext(ctx).Infof("GetUser: %v", id)
	return uc.repo.Get(ctx, id)
}
