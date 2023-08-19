package service

import (
	"context"
	v1 "galileo/api/core/v1"
	"galileo/pkg/ctxdata"
	"github.com/golang/protobuf/ptypes/empty"
)

func (c *CoreService) GetUserProjectList(ctx context.Context, empty *empty.Empty) (*v1.UserProjectReply, error) {
	userId := ctxdata.GetUserId(ctx)
	_, err := c.uu.GetUserProjectList(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &v1.UserProjectReply{}, nil
}

func (c *CoreService) GetUserLatestActivity(ctx context.Context, empty *empty.Empty) (*v1.UserLatestActivityReply, error) {
	return nil, nil
}

func (c *CoreService) GetUserGroupList(ctx context.Context, empty *empty.Empty) (*v1.UserGroupListReply, error) {
	return nil, nil
}
