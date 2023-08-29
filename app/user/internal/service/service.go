package service

import (
	v1 "galileo/api/user/v1"
	"galileo/app/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is scheduler providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	v1.UnimplementedUserServer
	uc     *biz.UserUseCase
	logger *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "user.Service")),
	}
}
