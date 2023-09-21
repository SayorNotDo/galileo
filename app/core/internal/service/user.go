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

func (c *CoreService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	return c.uc.DeleteUser(ctx, req.ID)
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
			ID:          item.ID,
			Name:        item.Name,
			Avatar:      item.Avatar,
			Description: item.Description,
			CreatedAt:   timestamppb.New(item.CreatedAt),
			CreatedBy:   item.CreatedBy,
			Headcount:   item.Headcount,
		})
	})
	return &v1.ListUserGroupReply{
		Total:     int32(len(ret)),
		GroupList: userGroupList,
	}, nil
}

func convertToGroupInfo(group *biz.Group) *v1.GroupInfo {
	var groupMemberList []*v1.GroupMember
	lo.ForEach(group.GroupMemberList, func(member *biz.GroupMember, _ int) {
		groupMemberList = append(groupMemberList, &v1.GroupMember{
			Uid:       member.Uid,
			Username:  member.Username,
			Role:      uint32(member.Role),
			CreatedBy: member.CreatedBy,
			CreatedAt: timestamppb.New(member.CreatedAt),
		})
	})
	return &v1.GroupInfo{
		ID:              group.ID,
		Name:            group.Name,
		Avatar:          group.Avatar,
		Description:     group.Description,
		CreatedBy:       group.CreatedBy,
		CreatedAt:       timestamppb.New(group.CreatedAt),
		Headcount:       group.Headcount,
		GroupMemberList: groupMemberList,
	}
}

func (c *CoreService) UserGroups(ctx context.Context, req *v1.GroupInfoRequest) (*v1.GroupInfo, error) {
	switch ctxdata.MethodFromContext(ctx) {
	case http.MethodPost:
		if len(req.Name) <= 0 {
			return nil, errResponse.SetCustomizeErrMsg(ReasonParamsError, "group's name can not be empty")
		}
		ret, err := c.uc.CreateUserGroup(ctx, &biz.Group{
			Name:        req.Name,
			Avatar:      req.Avatar,
			Description: req.Description,
		})
		if err != nil {
			return nil, err
		}
		return convertToGroupInfo(ret), nil
	case http.MethodPut:
		err := c.uc.UpdateUserGroup(ctx, &biz.Group{
			ID:          req.ID,
			Name:        req.Name,
			Avatar:      req.Avatar,
			Description: req.Description,
		})
		if err != nil {
			return nil, err
		}
		return nil, nil
	case http.MethodGet:
		ret, err := c.uc.GetUserGroup(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return convertToGroupInfo(ret), nil
	default:
		return nil, errResponse.SetErrByReason(errResponse.ReasonUnknownError)
	}
}
