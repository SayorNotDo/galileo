package service

import (
	"context"
	"encoding/json"
	v1 "galileo/api/core/v1"
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

func (c *CoreService) DataReportTrack(ctx context.Context, req *v1.DataReportTrackRequest) (*emptypb.Empty, error) {
	originData := req.Data
	var data map[string]interface{}
	/* 解析接口上报的数据 */
	if err := json.Unmarshal(originData, &data); err != nil {
		return nil, err
	}
	/* 调用业务函数 */
	if err := c.uc.DataReportTrack(ctx, data); err != nil {
		return nil, err
	}
	return nil, nil
}
