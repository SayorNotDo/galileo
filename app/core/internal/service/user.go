package service

import (
	"context"
	v1 "galileo/api/core/v1"
	"galileo/app/core/internal/biz"
	"galileo/pkg/ctxdata"
	"galileo/pkg/errResponse"
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

func (c *CoreService) Logout(ctx context.Context, empty *empty.Empty) (*emptypb.Empty, error) {
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

func (c *CoreService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	return c.uc.ListUser(ctx, req.PageNum, req.PageSize)
}

func (c *CoreService) UserInfo(ctx context.Context, req *v1.UserInfoRequest) (*v1.UserInfoReply, error) {
	var reply *v1.UserInfoReply
	var err error
	switch ctxdata.MethodFromContext(ctx) {
	case "PUT":
		reply, err = c.uc.UpdateUserInfo(ctx, req)
	case "GET":
		//return c.uc.GetUserInfo(ctx)
	case "":
		return nil, errResponse.SetErrByReason(errResponse.ReasonSystemError)
	}
	if err != nil {
		return nil, err
	}
	return reply, nil
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
	var userGroupList = make([]*v1.GroupInfo, 0)
	ret, err := c.uc.GetUserGroupList(ctx)
	if err != nil {
		return nil, err
	}
	lo.ForEach(ret, func(item *biz.Group, _ int) {
		userGroupList = append(userGroupList, &v1.GroupInfo{
			Id:          item.Id,
			Name:        item.Name,
			Avatar:      item.Avatar,
			Description: item.Description,
			CreatedAt:   timestamppb.New(item.CreatedAt),
			CreatedBy:   item.CreatedBy,
			UpdatedAt:   timestamppb.New(item.UpdatedAt),
			UpdatedBy:   item.UpdatedBy,
			DeletedAt:   timestamppb.New(item.DeletedAt),
			DeletedBy:   item.DeletedBy,
			Headcount:   item.Headcount,
		})
	})
	return &v1.UserGroupListReply{
		Total:     int32(len(ret)),
		GroupList: userGroupList,
	}, nil
}

func (c *CoreService) UserGroup(ctx context.Context, req *v1.GroupInfoRequest) (*v1.GroupInfo, error) {
	var reply *v1.GroupInfo
	switch ctxdata.MethodFromContext(ctx) {
	case "POST":
	case "DELETE":
	case "PUT":
	case "GET":
		_, err := c.uc.GetUserGroup(ctx, req.Id)
		if err != nil {
			return nil, err
		}
		return &v1.GroupInfo{}, nil
	default:
		return nil, errResponse.SetErrByReason(errResponse.ReasonUnknownError)
	}
	return reply, nil
}
