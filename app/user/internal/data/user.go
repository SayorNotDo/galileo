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
	Avatar         string    `json:"avatar"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Status         int32     `json:"status"`
	UpdateAt       time.Time `json:"update_at" gorm:"update_time"`
	LastLoginAt    time.Time `json:"last_login_at" gorm:""`
	CreatedAt      time.Time `json:"created_at" gorm:"column:add_time"`
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
	repo.log.WithContext(ctx).Info("gormDB: GetUser, id: ", id)
	return &biz.User{
		ID:       user.ID,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Status:   user.Status,
		Phone:    user.Phone,
	}, nil
}

func (repo *userRepo) Create(ctx context.Context, u *biz.User) (*biz.User, error) {
	var user User
	res := repo.data.gormDB.Where(&biz.User{Phone: u.Phone}).Or(&biz.User{Username: u.Username}).Or(&biz.User{Email: u.Email}).First(&user)
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
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Status:   user.Status,
	}, nil
}
