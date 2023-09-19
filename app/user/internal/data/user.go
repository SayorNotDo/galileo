package data

import (
	"context"
	"fmt"
	"galileo/app/user/internal/biz"
	"galileo/app/user/internal/pkg/util"
	"galileo/ent"
	"galileo/ent/group"
	"galileo/ent/groupmember"
	"galileo/ent/user"
	"galileo/pkg/constant"
	. "galileo/pkg/errResponse"
	"galileo/pkg/pagination"
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

func (repo *userRepo) GetUserInfo(ctx context.Context, id uint32) (*biz.User, error) {
	u, err := repo.data.entDB.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:            u.ID,
		Username:      u.Username,
		ChineseName:   u.ChineseName,
		Avatar:        u.Avatar,
		Email:         u.Email,
		Phone:         u.Phone,
		Active:        u.Active,
		Location:      u.Location,
		LastLoginTime: u.LastLoginTime,
		CreatedAt:     u.CreatedAt,
	}, nil
}

func (repo *userRepo) SoftDeleteById(ctx context.Context, id, deleteId uint32) (bool, error) {
	_, err := repo.data.entDB.User.
		UpdateOneID(deleteId).
		SetDeletedAt(time.Now()).
		SetDeletedBy(id).
		Save(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *userRepo) ListUser(ctx context.Context, pageToken string, pageSize int32) ([]*biz.User, int32, string, error) {
	/* 初始化偏移量 */
	var offset int
	/* 解析pageToken获得分页状态信息结构 */
	if len(pageToken) == 0 {
		offset = 0
	} else {
		state, err := pagination.ParsePageToken(pageToken)
		if err != nil {
			return nil, 0, "", err
		}
		offset = state.Offset
	}
	//offset = (pageNum - 1) * pageSize
	userList, err := repo.data.entDB.User.Query().Offset(offset).Limit(int(pageSize)).All(ctx)
	if err != nil {
		return nil, 0, "", err
	}
	total := len(userList)
	rv := make([]*biz.User, 0)
	for _, u := range userList {
		rv = append(rv, &biz.User{
			Id:            u.ID,
			Username:      u.Username,
			ChineseName:   u.ChineseName,
			Email:         u.Email,
			Phone:         u.Phone,
			Avatar:        u.Avatar,
			Active:        u.Active,
			Location:      u.Location,
			LastLoginTime: u.LastLoginTime,
			CreatedAt:     u.CreatedAt,
		})
	}
	/* 生成序列化的分页状态信息 */
	nextPageToken, err := pagination.GeneratePageToken(pagination.State{
		Offset: offset + int(pageSize),
	})
	if err != nil {
		return nil, 0, "", err
	}
	return rv, int32(total), nextPageToken, nil
}

func (repo *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	createUser, err := repo.data.entDB.User.Create().
		SetUsername(u.Username).
		SetEmail(u.Email).
		SetPhone(u.Phone).
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

func (repo *userRepo) UpdateUser(ctx context.Context, u *biz.User) (bool, error) {
	err := repo.data.entDB.User.
		UpdateOneID(u.Id).
		SetAvatar(u.Avatar).
		SetLocation(u.Location).
		SetChineseName(u.ChineseName).
		Exec(ctx)
	switch {
	case ent.IsNotFound(err):
		return false, errors.NotFound(ReasonRecordNotFound, err.Error())
	case err != nil:
		return false, err
	}
	return true, nil
}

func (repo *userRepo) ValidateUser(ctx context.Context, username, password string) (*biz.User, error) {
	ret, err := repo.data.entDB.User.Query().Where(user.Username(username)).Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, SetErrByReason(ReasonRecordNotFound)
	case err != nil:
		return nil, err
	}
	if ok := util.ComparePassword(password, ret.Password); !ok {
		return nil, SetErrByReason(ReasonUserPasswordError)
	}
	if err := repo.data.entDB.User.UpdateOneID(ret.ID).SetLastLoginTime(time.Now()).Exec(ctx); err != nil {
		return nil, SetErrByReason(ReasonSystemError)
	}
	return &biz.User{
		Id:       ret.ID,
		UUID:     ret.UUID,
		Username: ret.Username,
	}, nil
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

func (repo *userRepo) GetUserGroupList(ctx context.Context, uid uint32) ([]*biz.Group, error) {
	/* 初始化返回值 */
	var groupList = make([]*biz.Group, 0)
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
		groupList = append(groupList, &biz.Group{
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
		})
	})
	if err != nil {
		return nil, err
	}
	return groupList, nil
}

func (repo *userRepo) GetUserGroup(ctx context.Context, groupId int32) (*biz.Group, error) {
	/* 获取指定的Group */
	var err error
	ret, err := repo.data.entDB.Group.Get(ctx, groupId)
	switch {
	case ent.IsNotFound(err):
		return nil, SetErrByReason(ReasonRecordNotFound)
	case err != nil:
		return nil, err
	}
	var groupMemberList []*biz.GroupMember
	res, err := repo.data.entDB.GroupMember.Query().
		Where(
			groupmember.GroupID(groupId),
		).All(ctx)
	lo.ForEach(res, func(item *ent.GroupMember, _ int) {
		var username string
		err = repo.data.entDB.User.Query().
			Where(user.ID(item.UserID)).
			Select(user.FieldUsername).
			Scan(ctx, username)
		groupMemberList = append(groupMemberList, &biz.GroupMember{
			Uid:       item.UserID,
			Username:  username,
			Role:      item.Role,
			CreatedAt: item.CreatedAt,
			CreatedBy: item.CreatedBy,
			DeletedAt: item.DeletedAt,
			DeletedBy: item.DeletedBy,
		})
	})
	if err != nil {
		return nil, err
	}
	return &biz.Group{
		Id:          ret.ID,
		Name:        ret.Name,
		Avatar:      ret.Avatar,
		Description: ret.Description,
		Headcount:   ret.Headcount,
		CreatedAt:   ret.CreatedAt,
		CreatedBy:   ret.CreatedBy,
		UpdatedAt:   ret.UpdatedAt,
		UpdatedBy:   ret.UpdatedBy,
		DeletedAt:   ret.DeletedAt,
		DeletedBy:   ret.DeletedBy,
	}, nil
}

func (repo *userRepo) UpdateUserGroup(ctx context.Context, updateGroup *biz.Group) error {
	/* 查询当前用户分组是否存在 */
	if exist, _ := repo.data.entDB.Group.Query().Where(group.ID(updateGroup.Id)).Exist(ctx); !exist {
		return SetCustomizeErrInfoByReason(ReasonRecordNotFound)
	}
	/* 更新 */
	if err := repo.data.entDB.Group.UpdateOneID(updateGroup.Id).
		SetName(updateGroup.Name).
		SetAvatar(updateGroup.Avatar).
		SetDescription(updateGroup.Description).
		SetUpdatedBy(updateGroup.UpdatedBy).
		Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (repo *userRepo) CreateUserGroup(ctx context.Context, group *biz.Group) (*biz.Group, error) {
	/* 创建用户分组逻辑：
	1.新建分组信息
	2.设置创建人的角色为分组负责人
	*/
	/* 创建事务函数 */
	tx, err := repo.data.entDB.Tx(ctx)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonSystemError, err.Error())
	}
	ret, err := tx.Group.Create().
		SetName(group.Name).
		SetAvatar(group.Avatar).
		SetDescription(group.Description).
		SetCreatedAt(group.CreatedAt).
		SetCreatedBy(group.CreatedBy).
		SetHeadcount(1).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	if err := tx.GroupMember.Create().
		SetGroupID(ret.ID).
		SetUserID(group.CreatedBy).
		SetRole(constant.Leader).
		Exec(ctx); err != nil {
		return nil, rollback(tx, err)
	}
	/* 提交事务，失败则回滚 */
	if err := tx.Commit(); err != nil {
		return nil, rollback(tx, err)
	}
	return &biz.Group{
		Id:          ret.ID,
		Name:        ret.Name,
		Avatar:      ret.Avatar,
		Description: ret.Description,
		Headcount:   ret.Headcount,
		CreatedAt:   ret.CreatedAt,
		CreatedBy:   ret.CreatedBy,
	}, nil
}
