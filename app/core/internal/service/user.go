package service

import (
	"context"
	v1 "galileo/api/core/v1"
	"galileo/app/core/internal/biz"
	"galileo/pkg/ctxdata"
	"galileo/pkg/errResponse"
	. "galileo/pkg/errResponse"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
)

func (c *CoreService) Register(ctx context.Context, req *v1.RegisterRequest) (*emptypb.Empty, error) {
	/* add trace */
	tr := otel.Tracer("core.scheduler")
	ctx, span := tr.Start(ctx, "new user registration")
	span.SpanContext()
	defer span.End()
	/* 创建新用户 */
	newUser, err := biz.NewUser(req.Phone, req.Username, req.Email)
	newUser.Password = req.Password
	if err != nil {
		return nil, err
	}
	if _, err := c.uc.CreateUser(ctx, &newUser); err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *CoreService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	// add trace
	tr := otel.Tracer("core.scheduler")
	ctx, span := tr.Start(ctx, "login")
	span.SpanContext()
	defer span.End()
	if len(req.Username) <= 0 || len(req.Password) <= 0 {
		return nil, SetErrByReason(ReasonParamsError)
	}
	token, err := c.uc.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{
		Token: token,
	}, nil
}

func (c *CoreService) Logout(ctx context.Context, empty *empty.Empty) (*emptypb.Empty, error) {
	// add trace
	tr := otel.Tracer("core.scheduler")
	ctx, span := tr.Start(ctx, "logout")
	span.SpanContext()
	defer span.End()
	uid := ctxdata.GetUserId(ctx)
	if _, err := c.uc.GetUserInfo(ctx, uid); err != nil {
		return nil, err
	}
	return c.uc.Logout(ctx)
}

func (c *CoreService) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) (*emptypb.Empty, error) {
	return c.uc.UpdatePassword(ctx, req)
}

func (c *CoreService) ListUsers(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	return c.uc.ListUser(ctx, req.PageToken, req.PageSize)
}

func (c *CoreService) UserInfo(ctx context.Context, req *v1.UserInfoRequest) (*v1.User, error) {
	ret, err := c.uc.GetUserInfo(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.User{
		ID:            ret.ID,
		Username:      ret.Username,
		ChineseName:   ret.ChineseName,
		Email:         ret.Email,
		Phone:         ret.Phone,
		Avatar:        ret.Avatar,
		Active:        ret.Active,
		Location:      ret.Location,
		CreatedAt:     timestamppb.New(ret.CreatedAt),
		LastLoginTime: timestamppb.New(ret.LastLoginTime),
	}, nil
}

func (c *CoreService) CurrentUserInfo(ctx context.Context, req *v1.UserInfoRequest) (*v1.User, error) {
	uid := ctxdata.GetUserId(ctx)
	switch ctxdata.MethodFromContext(ctx) {
	case http.MethodGet:
		ret, err := c.uc.GetUserInfo(ctx, uid)
		if err != nil {
			return nil, err
		}
		return &v1.User{
			ID:            ret.ID,
			Username:      ret.Username,
			ChineseName:   ret.ChineseName,
			Email:         ret.Email,
			Phone:         ret.Phone,
			Avatar:        ret.Avatar,
			Active:        ret.Active,
			Location:      ret.Location,
			CreatedAt:     timestamppb.New(ret.CreatedAt),
			LastLoginTime: timestamppb.New(ret.LastLoginTime),
		}, nil
	case http.MethodPut:
		user, err := biz.NewUser(req.Phone, req.Username, req.Email)
		user.ID = uid
		if err != nil {
			return nil, err
		}
		ret, err := c.uc.UpdateUserInfo(ctx, &user)
		if err != nil {
			return nil, err
		}
		return &v1.User{
			ID:          ret.ID,
			Username:    ret.Username,
			ChineseName: ret.ChineseName,
			Email:       ret.Email,
			Phone:       ret.Phone,
			Avatar:      ret.Avatar,
			Location:    ret.Location,
		}, nil
	default:
		return nil, errResponse.SetErrByReason(errResponse.ReasonUnsupportedMethod)
	}
}

func (c *CoreService) DeleteUser(ctx context.Context, req *v1.DeleteRequest) (*emptypb.Empty, error) {
	return c.uc.DeleteUser(ctx, req.Id)
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

func (c *CoreService) ListUserGroups(ctx context.Context, request *v1.ListUserGroupsRequest) (*v1.ListUserGroupReply, error) {
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
	return &v1.ListUserGroupReply{
		Total:     int32(len(ret)),
		GroupList: userGroupList,
	}, nil
}

func (c *CoreService) UserGroups(ctx context.Context, req *v1.GroupInfoRequest) (*v1.GroupInfo, error) {
	var reply *v1.GroupInfo
	switch ctxdata.MethodFromContext(ctx) {
	case http.MethodPost:
		ret, err := c.uc.CreateUserGroup(ctx, &biz.Group{
			Id:          req.Id,
			Name:        req.Name,
			Avatar:      req.Avatar,
			Description: req.Description,
		})
		if err != nil {
			return nil, err
		}
		return &v1.GroupInfo{Id: ret.Id}, nil
	case http.MethodDelete:
	case http.MethodPut:
		err := c.uc.UpdateUserGroup(ctx, &biz.Group{
			Id:          req.Id,
			Name:        req.Name,
			Avatar:      req.Avatar,
			Description: req.Description,
		})
		if err != nil {
			return nil, err
		}
		return nil, nil
	case http.MethodGet:
		ret, err := c.uc.GetUserGroup(ctx, req.Id)
		if err != nil {
			return nil, err
		}
		var groupMemberList []*v1.GroupMember
		lo.ForEach(ret.GroupMemberList, func(item *biz.GroupMember, _ int) {
			groupMemberList = append(groupMemberList, &v1.GroupMember{
				Uid:       item.Uid,
				Username:  item.Username,
				Role:      uint32(item.Role),
				CreatedAt: timestamppb.New(ret.CreatedAt),
				CreatedBy: ret.CreatedBy,
			})
		})
		return &v1.GroupInfo{
			Id:              ret.Id,
			Name:            ret.Name,
			Avatar:          ret.Avatar,
			Description:     ret.Description,
			CreatedAt:       timestamppb.New(ret.CreatedAt),
			CreatedBy:       ret.CreatedBy,
			UpdatedAt:       timestamppb.New(ret.UpdatedAt),
			UpdatedBy:       ret.UpdatedBy,
			DeletedAt:       timestamppb.New(ret.DeletedAt),
			DeletedBy:       ret.DeletedBy,
			Headcount:       ret.Headcount,
			GroupMemberList: groupMemberList,
		}, nil
	default:
		return nil, errResponse.SetErrByReason(errResponse.ReasonUnknownError)
	}
	return reply, nil
}
