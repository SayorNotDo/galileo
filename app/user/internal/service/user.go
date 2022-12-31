package service

import (
	"context"
	"galileo/app/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "galileo/api/user"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc     *biz.UserCase
	logger *log.Helper
}

func NewUserService(uc *biz.UserCase, logger log.Logger) *UserService {
	return &UserService{
		uc:     uc,
		logger: log.NewHelper(logger),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Id:       user.ID,
		Status:   user.Status,
		Username: user.Username,
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}

// SayHello implements helloworld.GreeterServer.
func (s *UserService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	g, err := s.uc.CreateUser(ctx, &biz.User{Nickname: in.Name})
	if err != nil {
		return nil, err
	}
	return &pb.HelloReply{Message: "Hello " + g.Nickname}, nil
}
