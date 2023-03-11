package service

import (
	"context"
	v1 "galileo/api/core/v1"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *CoreService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	// add trace
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "get user info")
	span.SpanContext()
	defer span.End()
	return c.uc.CreateUser(ctx, req)
}

func (c *CoreService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	// add trace
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "login")
	span.SpanContext()
	defer span.End()
	return c.uc.Login(ctx, req)
}

func (c *CoreService) Logout(ctx context.Context, req *emptypb.Empty) (*v1.LogoutReply, error) {
	// add trace
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "logout")
	span.SpanContext()
	defer span.End()
	return c.uc.Logout(ctx)
}

func (c *CoreService) UserDetail(ctx context.Context, req *emptypb.Empty) (*v1.UserDetailReply, error) {
	return c.uc.UserDetail(ctx, &emptypb.Empty{})
}
