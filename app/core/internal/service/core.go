package service

import (
	"context"
	v1 "galileo/api/core/v1"
)

func (c *CoreService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	return c.uc.CreateUser(ctx, req)
}
