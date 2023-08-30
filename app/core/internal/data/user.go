package data

import (
	"context"
	v1 "galileo/api/core/v1"
	userService "galileo/api/user/v1"
	"galileo/app/core/internal/biz"
	. "galileo/app/core/internal/pkg/middleware/auth"
	"galileo/pkg/encryption"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"strings"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "user.Repo")),
	}
}

func (u *userRepo) GetUserProjectList(ctx context.Context) ([]*v1.ProjectInfo, error) {
	/* 初始化返回值 */
	var projectList = make([]*v1.ProjectInfo, 0)
	res, err := u.data.managementCli.GetUserProjectList(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}
	for _, v := range res.ProjectList {
		projectList = append(projectList, &v1.ProjectInfo{
			Id:          v.Id,
			Name:        v.Name,
			Identifier:  v.Identifier,
			CreatedAt:   v.CreatedAt,
			CreatedBy:   v.CreatedBy,
			UpdatedBy:   v.UpdatedBy,
			UpdatedAt:   v.UpdatedAt,
			DeletedAt:   v.DeletedAt,
			DeletedBy:   v.DeletedBy,
			StartTime:   v.StartTime,
			Deadline:    v.Deadline,
			Description: v.Description,
			Remark:      v.Remark,
		})
	}
	return projectList, nil
}

func (u *userRepo) GetUserGroupList(ctx context.Context) ([]*biz.Group, error) {
	/* 初始化返回值 */
	var groupList = make([]*biz.Group, 0)
	res, err := u.data.uc.GetUserGroupList(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}
	lo.ForEach(res.GroupList, func(item *userService.GroupInfo, index int) {
		groupList = append(groupList, &biz.Group{
			Id:          item.Id,
			Name:        item.Name,
			Avatar:      item.Avatar,
			Description: item.Description,
			Headcount:   item.Headcount,
			CreatedAt:   item.CreatedAt.AsTime(),
			CreatedBy:   item.CreatedBy,
			UpdatedAt:   item.UpdatedAt.AsTime(),
			UpdatedBy:   item.UpdatedBy,
			DeletedAt:   item.DeletedAt.AsTime(),
			DeletedBy:   item.DeletedBy,
		})
	})
	return groupList, nil
}

func (u *userRepo) ValidateUser(ctx context.Context, username, password string) (*biz.User, error) {
	res, err := u.data.uc.ValidateUser(ctx, &userService.ValidateUserRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	uuidObj, err := uuid.Parse(res.UUID)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       res.Id,
		UUID:     uuidObj,
		Username: res.Username,
	}, nil
}

func (u *userRepo) ListUser(c context.Context, PageToken string, pageSize int32) ([]*v1.UserDetail, int32, error) {
	rsp, err := u.data.uc.ListUser(c, &userService.ListUserRequest{PageToken: PageToken, PageSize: pageSize})
	if err != nil {
		return nil, 0, err
	}
	rv := make([]*v1.UserDetail, 0)
	for _, u := range rsp.Data {
		rv = append(rv, &v1.UserDetail{
			Id:          u.Id,
			ChineseName: u.ChineseName,
			Phone:       u.Phone,
			Email:       u.Email,
		})
	}
	total := rsp.Total
	return rv, total, nil
}

func (u *userRepo) GetUserInfo(ctx context.Context, uid uint32) (*biz.User, error) {
	user, err := u.data.uc.GetUserInfo(ctx, &userService.GetUserInfoRequest{Id: uid})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:            user.Id,
		Username:      user.Username,
		ChineseName:   user.ChineseName,
		Phone:         user.Phone,
		Email:         user.Email,
		Avatar:        user.Avatar,
		Active:        user.Active,
		Location:      user.Location,
		CreatedAt:     user.CreatedAt.AsTime(),
		LastLoginTime: user.LastLoginTime.AsTime(),
	}, nil
}

func (u *userRepo) CreateUser(c context.Context, user *biz.User) (*biz.User, error) {
	createUser, err := u.data.uc.CreateUser(c, &userService.CreateUserRequest{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:        createUser.Id,
		Username:  createUser.Username,
		CreatedAt: createUser.CreatedAt.AsTime(),
	}, nil
}

func (u *userRepo) SoftDeleteUser(c context.Context, uid uint32) (bool, error) {
	rsp, err := u.data.uc.SoftDeleteUser(c, &userService.SoftDeleteRequest{Id: uid})
	if err != nil {
		return false, err
	}
	return rsp.Deleted, nil
}

func (r *coreRepo) UpdateUserInfo(c context.Context, user *biz.User) (*biz.User, error) {
	return nil, nil
}

func (u *userRepo) SetToken(ctx context.Context, token string) (string, error) {
	key := encryption.EncodeMD5(token)
	_, err := u.data.redisCli.Set(ctx, "token:"+key, token, TokenExpiration).Result()
	if err != nil {
		return "", err
	}
	return key, nil
}

func (u *userRepo) DestroyToken(ctx context.Context) error {
	tr, _ := transport.FromServerContext(ctx)
	jwtToken := strings.SplitN(tr.RequestHeader().Get("Authorization"), " ", 2)[1]
	key := encryption.EncodeMD5(jwtToken)
	u.data.redisCli.Del(ctx, "token:"+key)
	return nil
}

func (u *userRepo) GetUserPassword(ctx context.Context) (string, error) {
	return "", nil
}

func (u *userRepo) UpdatePassword(ctx context.Context, password string) (bool, error) {
	if _, err := u.data.uc.UpdatePassword(ctx, &userService.UpdatePasswordRequest{
		NewPassword: password,
	}); err != nil {
		return false, err
	}
	return true, nil
}

func (u *userRepo) GetUserGroup(ctx context.Context, groupId int32) (*biz.Group, error) {
	var groupMemberList []*biz.GroupMember
	res, err := u.data.uc.GetUserGroup(ctx, &userService.UserGroupRequest{Id: groupId})
	if err != nil {
		return nil, err
	}
	lo.ForEach(res.GroupMemberList, func(item *userService.GroupMember, _ int) {
		groupMemberList = append(groupMemberList, &biz.GroupMember{
			Uid:       item.Uid,
			Username:  item.Username,
			Role:      uint8(item.Role),
			CreatedBy: item.CreatedBy,
			CreatedAt: item.CreatedAt.AsTime(),
		})
	})
	return &biz.Group{
		Id:              res.Id,
		Name:            res.Name,
		Avatar:          res.Avatar,
		Description:     res.Description,
		Headcount:       res.Headcount,
		CreatedAt:       res.CreatedAt.AsTime(),
		CreatedBy:       res.CreatedBy,
		UpdatedAt:       res.UpdatedAt.AsTime(),
		UpdatedBy:       res.UpdatedBy,
		DeletedAt:       res.DeletedAt.AsTime(),
		DeletedBy:       res.DeletedBy,
		GroupMemberList: groupMemberList,
	}, nil
}

func (u *userRepo) UpdateUserGroup(ctx context.Context, group *biz.Group) error {
	return nil
}

func (u *userRepo) CreateUserGroup(ctx context.Context, group *biz.Group) (*biz.Group, error) {
	res, err := u.data.uc.CreateUserGroup(ctx, &userService.UserGroupRequest{
		Id:          group.Id,
		Name:        group.Name,
		Avatar:      group.Avatar,
		Description: group.Description,
	})
	if err != nil {
		return nil, err
	}
	return &biz.Group{
		Id:          res.Id,
		Name:        res.Name,
		Avatar:      res.Avatar,
		Description: res.Description,
		CreatedBy:   res.CreatedBy,
		CreatedAt:   res.CreatedAt.AsTime(),
		Headcount:   res.Headcount,
	}, nil
}
