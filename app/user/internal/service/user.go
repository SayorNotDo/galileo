package service

import (
	"context"
	v1 "galileo/api/user/v1"
	"galileo/app/user/internal/biz"
	"galileo/app/user/internal/pkg/util"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc     *biz.UserUseCase
	logger *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:     uc,
		logger: log.NewHelper(logger),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	ret, err := s.uc.Create(ctx, &biz.User{Username: req.Username, Email: req.Email, Phone: req.Phone, HashedPassword: hashedPassword})
	if err != nil {
		return nil, err
	}
	userInfoRep := v1.CreateUserReply{
		Id:    ret.ID,
		Phone: ret.Phone,
	}
	return &userInfoRep, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	return &v1.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	return &v1.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	user, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserReply{
		Code:    200,
		Message: "success",
		Data: &v1.UserInfo{
			Username: user.Username,
			Status:   user.Status,
			Nickname: user.Nickname,
			Email:    user.Email,
			Phone:    user.Phone,
		},
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	users, total, err := s.uc.List(ctx, req.PageNum, req.PageSize)
	println(users)
	println(total)
	println(err)
	return &v1.ListUserReply{}, nil
}

// SayHello implements helloworld.GreeterServer.
func (s *UserService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: "Hello " + in.Name}, nil
}
