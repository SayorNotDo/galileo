package data

import (
	"context"
	"errors"
	v1 "galileo/api/user/v1"
	"galileo/app/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	var user *biz.User
	res := repo.data.gormDB.Where("id = ?", id).First(&user)
	if res.RowsAffected == 0 {
		return nil, res.Error
	}
	return &biz.User{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Status:   user.Status,
		Phone:    user.Phone,
		Password: user.Password,
	}, nil
}

func (repo *userRepo) GetByUsername(ctx context.Context, username string) (*biz.User, error) {
	var user *biz.User
	res := repo.data.gormDB.Where("username = ?", username).First(&user)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}
	return user, nil
}

func (repo *userRepo) DeleteById(ctx context.Context, id uint32) (bool, error) {
	var user *biz.User
	ret := repo.data.gormDB.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&user)
	if ret.RowsAffected == 0 {
		return false, ret.Error
	}
	return true, nil
}

func (repo *userRepo) List(ctx context.Context, pageNum, pageSize int32) ([]*v1.UserInfoReply, int32, error) {
	var users []User
	result := repo.data.gormDB.Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	total := int32(result.RowsAffected)
	repo.data.gormDB.Scopes(paginate(pageNum, pageSize)).Find(&users)
	rv := make([]*v1.UserInfoReply, 0)
	for _, u := range users {
		rv = append(rv, &v1.UserInfoReply{
			Id:          u.Id,
			Username:    u.Username,
			Nickname:    u.Nickname,
			ChineseName: u.ChineseName,
			Email:       u.Email,
			Phone:       u.Phone,
			Role:        u.Role,
			Status:      u.Status,
		})
	}
	log.Debugf("total: %v", total)
	log.Debugf("userList: %v", rv)
	log.Debugf("users: %v", users)
	return rv, total, nil
}

func paginate(pageNum, pageSize int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum == 0 {
			pageNum = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (pageNum - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}

func (repo *userRepo) Create(ctx context.Context, u *biz.User) (*biz.User, error) {
	var user User
	res := repo.data.gormDB.Where("phone = ?", u.Phone).Or("username = ?", u.Username).Or("email = ?", u.Email).First(&user)
	if res.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "duplicated field value exists")
	}
	user.Phone = u.Phone
	user.Username = u.Username
	user.Email = u.Email
	user.Password = u.Password
	if result := repo.data.gormDB.Create(&user); result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &biz.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Status:   user.Status,
	}, nil
}

func (repo *userRepo) Update(ctx context.Context, u *biz.User) (bool, error) {
	var user User
	res := repo.data.gormDB.Where("id = ?", u.Id).Find(&user)
	if res.RowsAffected == 0 {
		return false, errors.New("record not found")
	}
	if result := repo.data.gormDB.Updates(&u); result.Error != nil {
		return false, status.Errorf(codes.Internal, result.Error.Error())
	}
	return true, nil
}

func (repo *userRepo) MapUpdate(ctx context.Context, u map[string]interface{}) (bool, error) {
	if res := repo.data.gormDB.Model(User{}).Where("id = ?", u["id"]).Updates(u); res.Error != nil {
		return false, status.Errorf(codes.Internal, res.Error.Error())
	}
	return true, nil
}
