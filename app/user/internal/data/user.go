package data

import (
	"context"
	"fmt"
	v1 "galileo/api/user/v1"
	"galileo/app/user/internal/biz"
	"galileo/ent"
	"galileo/ent/group"
	"galileo/ent/groupmember"
	"galileo/ent/user"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"time"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	return user, nil
}

func (repo *userRepo) GetById(ctx context.Context, id uint32) (*biz.User, error) {
	u, err := repo.data.entDB.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       u.ID,
		Username: u.Username,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Email:    u.Email,
		Status:   u.Active,
		Phone:    u.Phone,
		Password: u.Password,
	}, nil
}

func (repo *userRepo) GetByUsername(ctx context.Context, username string) (*biz.User, error) {
	u, err := repo.data.entDB.User.Query().Where(user.Username(username)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       u.ID,
		Username: u.Username,
		Nickname: u.Nickname,
		Email:    u.Email,
		Password: u.Password,
		Phone:    u.Phone,
		Status:   u.Active,
		Role:     int32(u.Role),
	}, nil
}

func (repo *userRepo) DeleteById(ctx context.Context, id uint32) (bool, error) {
	_, err := repo.data.entDB.User.Delete().Where(user.ID(id)).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *userRepo) SoftDeleteById(ctx context.Context, id, deleteId uint32) (bool, error) {
	_, err := repo.data.entDB.User.
		UpdateOneID(deleteId).
		SetDeletedAt(time.Now()).
		SetIsDeleted(true).
		SetDeletedBy(id).
		Save(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *userRepo) List(ctx context.Context, pageNum, pageSize int32) ([]*v1.UserInfoReply, int32, error) {
	if pageNum == 0 {
		pageNum = 1
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 10:
		pageSize = 10
	}
	offset := (pageNum - 1) * pageSize
	userList, err := repo.data.entDB.User.Query().Offset(int(offset)).Limit(int(pageSize)).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	total := len(userList)
	rv := make([]*v1.UserInfoReply, 0)
	for _, u := range userList {
		rv = append(rv, &v1.UserInfoReply{
			Id:          u.ID,
			Username:    u.Username,
			Nickname:    u.Nickname,
			ChineseName: u.ChineseName,
			Email:       u.Email,
			Phone:       u.Phone,
			Role:        int32(u.Role),
			Status:      u.Active,
		})
	}
	return rv, int32(total), nil
}

func (repo *userRepo) Create(ctx context.Context, u *biz.User) (*biz.User, error) {
	createUser, err := repo.data.entDB.User.Create().
		SetUsername(u.Username).
		SetEmail(u.Email).
		SetPhone(u.Phone).
		SetNickname(u.Username).
		SetPassword(u.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:        createUser.ID,
		Username:  createUser.Username,
		CreatedAt: createUser.CreatedAt,
	}, nil
}

func (repo *userRepo) Update(ctx context.Context, u *biz.User) (bool, error) {
	err := repo.data.entDB.User.
		UpdateOneID(u.Id).
		SetAvatar(u.Avatar).
		SetNickname(u.Nickname).
		Exec(ctx)
	switch {
	case ent.IsNotFound(err):
		return false, errors.NotFound("User Not Found", err.Error())
	case err != nil:
		return false, err
	}
	return true, nil
}

func (repo *userRepo) UpdatePassword(ctx context.Context, u *biz.User) (bool, error) {
	err := repo.data.entDB.User.
		UpdateOneID(u.Id).
		SetPassword(u.Password).
		Exec(ctx)
	switch {
	case ent.IsNotFound(err):
		return false, errors.NotFound("User Not Found", err.Error())
	case err != nil:
		return false, err
	}
	return true, nil
}

func (repo *userRepo) SetToken(ctx context.Context, username string, token string) (bool, error) {
	conn := repo.data.redisDB.Conn()
	defer func(conn *redis.Conn) {
		err := conn.Close()
		if err != nil {
			log.Error("redis client closed connection error")
		}
	}(conn)
	res := conn.Set(ctx, fmt.Sprintf("%s:%s", username, "token"), token, time.Hour*24*30)
	if res.Err() != nil {
		return false, errors.InternalServer("InternalServer Error", res.Err().Error())
	}
	return true, nil
}

func (repo *userRepo) EmptyToken(ctx context.Context, username string) (bool, error) {
	conn := repo.data.redisDB.Conn()
	defer func(conn *redis.Conn) {
		err := conn.Close()
		if err != nil {
			log.Error("redis client closed connection")
		}
	}(conn)
	if res := conn.Del(ctx, fmt.Sprintf("%s:token", username)); res.Err() != nil {
		return false, errors.InternalServer("InternalServer Error", res.Err().Error())
	}
	return true, nil
}

func (repo *userRepo) GetUserGroupList(ctx context.Context, uid uint32) ([]*biz.UserGroup, error) {
	/* 初始化返回值 */
	var groupList = make([]*biz.UserGroup, 0)
	var err error
	/* 关系表查询对应uid的记录 */
	ret, err := repo.data.entDB.GroupMember.Query().
		Where(groupmember.UserID(uid)).All(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, SetCustomizeErrInfoByReason(ReasonRecordNotFound)
	case err != nil:
		return nil, err
	}
	lo.ForEach(ret, func(item *ent.GroupMember, _ int) {
		var queryGroup *ent.Group
		/* 基于获取的记录查询对应的GroupInfo */
		queryGroup, err = repo.data.entDB.Group.Query().Where(group.ID(item.GroupID)).Only(ctx)
		if err != nil {
			return
		}
		groupList = append(groupList, &biz.UserGroup{
			GroupMemberId: item.ID,
			Role:          item.Role,
			GroupInfo: biz.Group{
				Id:          queryGroup.ID,
				Name:        queryGroup.Name,
				Avatar:      queryGroup.Avatar,
				Description: queryGroup.Description,
				CreatedBy:   queryGroup.CreatedBy,
				CreatedAt:   queryGroup.CreatedAt,
				UpdatedBy:   queryGroup.UpdatedBy,
				UpdatedAt:   queryGroup.UpdatedAt,
				DeletedBy:   queryGroup.DeletedBy,
				DeletedAt:   queryGroup.DeletedAt,
				Headcount:   queryGroup.Headcount,
			},
		})
	})
	if err != nil {
		return nil, err
	}
	return groupList, nil
}
