package data

import (
	"context"
	"fmt"
	v1 "galileo/api/user/v1"
	"galileo/app/user/internal/biz"
	"galileo/ent"
	"galileo/ent/user"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"time"
)

type User struct {
	Id          uint32    `json:"id" gorm:"primaryKey"`
	Username    string    `json:"name" gorm:"varchar(255)"`
	ChineseName string    `json:"chinese_name" gorm:"varchar(25)"`
	Nickname    string    `json:"nickname" gorm:"varchar(25)"`
	Password    string    `json:"password" gorm:"varchar(255); not null"`
	Role        int32     `json:"role"`
	Avatar      string    `json:"avatar"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Status      bool      `json:"status" gorm:"default: true"`
	UpdateAt    time.Time `json:"update_at" gorm:"autoUpdateTime"`
	LastLoginAt time.Time `json:"last_login_at"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt   time.Time `json:"deleted_at"`
	DeletedBy   uint32    `json:"deleted_by"`
}

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
		Id:       createUser.ID,
		Username: createUser.Username,
		Email:    createUser.Email,
		Phone:    createUser.Phone,
		Status:   createUser.Active,
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
		return false, errors.NotFound("Not Found", err.Error())
	case err != nil:
		return false, errors.InternalServer("InternalServer Error", err.Error())
	}
	return true, nil
}

func (repo *userRepo) SetToken(ctx context.Context, username string, token string) (bool, error) {
	conn := repo.data.redisDB.Conn(ctx)
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
	conn := repo.data.redisDB.Conn(ctx)
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
