package service

import (
	"context"
	v1 "galileo/api/core/v1"
	"galileo/app/core/internal/biz"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	var userGroupList = make([]*v1.UserGroup, 0)
	ret, err := c.uc.GetUserGroupList(ctx)
	if err != nil {
		return nil, err
	}
	lo.ForEach(ret, func(item *biz.UserGroup, _ int) {
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
		Total:     int32(len(ret)),
		GroupList: userGroupList,
	}, nil
}
