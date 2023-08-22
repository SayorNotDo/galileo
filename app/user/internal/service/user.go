package service

import (
	"context"
	v1 "galileo/api/user/v1"
	"galileo/app/user/internal/biz"
	"galileo/app/user/internal/pkg/util"
	"galileo/pkg/ctxdata"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

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

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	ret, err := s.uc.Create(ctx, &biz.User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}
	userInfoRep := v1.CreateUserReply{
		Id:        ret.Id,
		Username:  ret.Username,
		CreatedAt: timestamppb.New(ret.CreatedAt),
	}
	return &userInfoRep, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Update(ctx, &biz.User{Id: req.Id, Nickname: req.Nickname, Avatar: req.Avatar})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	ok, err := s.uc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUserReply{
		Deleted: ok,
	}, nil
}

func (s *UserService) SoftDeleteUser(ctx context.Context, req *v1.SoftDeleteRequest) (*v1.SoftDeleteReply, error) {
	md, _ := metadata.FromServerContext(ctx)
	uidStr := md.Get("x-md-local-uid")
	deletedAt := time.Now()
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	ok, err := s.uc.SoftDeleteById(ctx, uint32(uid), req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.SoftDeleteReply{
		DeletedAt: timestamppb.New(deletedAt),
		Deleted:   ok,
		Status:    ok,
	}, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, req *v1.UsernameRequest) (*v1.UserInfoReply, error) {
	if user, err := s.uc.GetByUsername(ctx, req.Username); err != nil {
		return nil, err
	} else {
		return &v1.UserInfoReply{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
			Nickname: user.Nickname,
			Phone:    user.Phone,
			Status:   user.Status,
			Password: user.Password,
			Role:     user.Role,
		}, nil
	}
}

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.UserInfoReply, error) {
	user, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoReply{
		Id:       user.Id,
		Username: user.Username,
		Status:   user.Status,
		Nickname: user.Nickname,
		Email:    user.Email,
		Phone:    user.Phone,
	}, nil
}

func (s *UserService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	users, total, err := s.uc.List(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	return &v1.ListUserReply{
		Total: total,
		Data:  users,
	}, nil
}

func (s *UserService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *UserService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.VerifyPasswordReply, error) {
	ok, err := s.uc.VerifyPassword(req.Password, req.HashedPassword)
	if err != nil {
		return nil, err
	}
	return &v1.VerifyPasswordReply{
		Success: ok,
	}, nil
}

func (s *UserService) SetToken(ctx context.Context, req *v1.SetTokenRequest) (*v1.SetTokenReply, error) {
	ok, err := s.uc.SetToken(ctx, req.Username, req.Token)
	if err != nil {
		return nil, err
	}
	return &v1.SetTokenReply{
		Success: ok,
	}, nil
}

func (s *UserService) EmptyToken(ctx context.Context, req *emptypb.Empty) (*v1.EmptyTokenReply, error) {
	md, _ := metadata.FromServerContext(ctx)
	uidStr := md.Get("x-md-local-uid")
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	ok, err := s.uc.EmptyToken(ctx, uint32(uid))
	if err != nil {
		return nil, err
	}
	return &v1.EmptyTokenReply{
		IsEmpty: ok,
	}, nil
}

func (s *UserService) ResetPassword(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *UserService) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromServerContext(ctx)
	uidStr := md.Get(ctxdata.UserIdKey)
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	u, err := s.uc.Get(ctx, uint32(uid))
	if err != nil {
		return nil, err
	}
	if _, err = s.uc.UpdatePassword(ctx, u, req.NewPassword); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) GetUserGroupList(ctx context.Context, empty *empty.Empty) (*v1.UserGroupListReply, error) {
	md, _ := metadata.FromServerContext(ctx)
	uidStr := md.Get(ctxdata.UserIdKey)
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	res, err := s.uc.GetUserGroupList(ctx, uint32(uid))
	if err != nil {
		return nil, err
	}
	/* 初始化返回的分组列表: UserGroup */
	var userGroupList = make([]*v1.UserGroup, 0)
	lo.ForEach(res, func(item *biz.UserGroup, _ int) {
		userGroupList = append(userGroupList, &v1.UserGroup{
			GroupMemberId: item.GroupMemberId,
			Role:          int32(item.Role),
			Group: &v1.GroupInfo{
				Id:          item.GroupInfo.Id,
				Name:        item.GroupInfo.Name,
				Avatar:      item.GroupInfo.Avatar,
				Description: item.GroupInfo.Description,
				CreatedAt:   timestamppb.New(item.GroupInfo.CreatedAt),
				CreatedBy:   item.GroupInfo.CreatedBy,
				UpdatedAt:   timestamppb.New(item.GroupInfo.UpdatedAt),
				UpdatedBy:   item.GroupInfo.UpdatedBy,
				DeletedAt:   timestamppb.New(item.GroupInfo.DeletedAt),
				DeletedBy:   item.GroupInfo.DeletedBy,
				Headcount:   item.GroupInfo.Headcount,
			},
		})
	})
	return &v1.UserGroupListReply{
		Total:     int32(len(res)),
		GroupList: userGroupList,
	}, nil
}
