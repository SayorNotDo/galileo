package service

import (
	"context"
	v1 "galileo/api/core/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *CoreService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	// add trace
	tr := otel.Tracer("scheduler")
	ctx, span := tr.Start(ctx, "get user info")
	span.SpanContext()
	defer span.End()
	return c.uc.CreateUser(ctx, req)
}

func (c *CoreService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	// add trace
	tr := otel.Tracer("scheduler")
	ctx, span := tr.Start(ctx, "login")
	span.SpanContext()
	defer span.End()
	return c.uc.Login(ctx, req)
}

func (c *CoreService) Logout(ctx context.Context, req *v1.LogoutRequest) (*emptypb.Empty, error) {
	// add trace
	tr := otel.Tracer("scheduler")
	ctx, span := tr.Start(ctx, "logout")
	span.SpanContext()
	defer span.End()
	return c.uc.Logout(ctx)
}

func (c *CoreService) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) (*emptypb.Empty, error) {
	return c.uc.UpdatePassword(ctx, req)
}

func (c *CoreService) UserDetail(ctx context.Context, req *emptypb.Empty) (*v1.UserDetailReply, error) {
	return c.uc.UserDetail(ctx, &emptypb.Empty{})
}

func (c *CoreService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	return c.uc.ListUser(ctx, req.PageNum, req.PageSize)
}

func (c *CoreService) UpdateUserInfo(ctx context.Context, req *v1.UserInfoUpdateRequest) (*v1.UserInfoUpdateReply, error) {
	return c.uc.UpdateUserInfo(ctx, req)
}

func (c *CoreService) DeleteUser(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteReply, error) {
	return c.uc.DeleteUser(ctx, req.Id)
}

func (c *CoreService) ExecuteToken(ctx context.Context, req *v1.ExecuteTokenRequest) (*v1.ExecuteTokenReply, error) {
	token, err := c.cc.ExecuteToken(ctx, req.Machine)
	if err != nil {
		return nil, err
	}
	return &v1.ExecuteTokenReply{
		ExecuteToken: token,
	}, nil
}

func (c *CoreService) GetUserProjectList(ctx context.Context, empty *empty.Empty) (*v1.UserProjectListReply, error) {
	ret, err := c.uc.GetUserProjectList(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UserProjectListReply{
		Total:       int32(len(ret)),
		ProjectList: ret,
	}, nil
}

func (c *CoreService) GetUserLatestActivity(ctx context.Context, empty *empty.Empty) (*v1.UserLatestActivityReply, error) {
	return nil, nil
}

func (c *CoreService) GetUserGroupList(ctx context.Context, empty *empty.Empty) (*v1.UserGroupListReply, error) {
	return nil, nil
}
