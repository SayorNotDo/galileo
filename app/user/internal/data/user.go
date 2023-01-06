package data

import (
	"context"
	"galileo/app/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type User struct {
	ID             uint32    `json:"id" gorm:"primaryKey"`
	Username       string    `json:"name" gorm:"varchar(255)"`
	ChineseName    string    `json:"chinese_name" gorm:"varchar(25)"`
	Nickname       string    `json:"nickname" gorm:"varchar(25)"`
	HashedPassword []byte    `json:"hashed_password" gorm:"varchar(255); not null"`
	Role           string    `json:"role"`
	Avatar         string    `json:"avatar"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Status         int32     `json:"status" gorm:"default: 1"`
	UpdateAt       time.Time `json:"update_at" gorm:"autoUpdateTime"`
	LastLoginAt    time.Time `json:"last_login_at" gorm:""`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
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

func (repo *userRepo) Get(ctx context.Context, id uint32) (*biz.User, error) {
	var user *biz.User
	repo.data.gormDB.Where("id = ?", id).First(&user)
	return &biz.User{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Status:   user.Status,
		Phone:    user.Phone,
	}, nil
}

func (repo *userRepo) Create(ctx context.Context, u *biz.User) (*biz.User, error) {
	var user User
	res := repo.data.gormDB.Where("phone = ?", u.Phone).Or("username = ?", u.Username).Or("email = ?", u.Email).First(&user)
	if res.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "User already exists")
	}
	user.Phone = u.Phone
	user.Username = u.Username
	user.Email = u.Email
	user.HashedPassword = u.HashedPassword
	if result := repo.data.gormDB.Create(&user); result.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return &biz.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Status:   user.Status,
	}, nil
}
