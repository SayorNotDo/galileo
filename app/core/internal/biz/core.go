package biz

import (
	"context"
	v1 "galileo/api/core/v1"
	"galileo/app/core/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type CoreRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
}
type User struct {
	ID       uint32
	Phone    string
	Nickname string
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
	//newUser, err := NewUser(req.Phone, req.Username, req.Password, req.Email)
	panic("")
}
