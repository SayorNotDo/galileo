package biz

import (
	"context"
	v1 "galileo/api/user/v1"
	"galileo/app/user/internal/pkg/util"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type User struct {
	Id          uint32    `json:"id"`
	Username    string    `json:"name"`
	ChineseName string    `json:"chinese_name"`
	Nickname    string    `json:"nickname"`
	Password    string    `json:"password"`
	Avatar      string    `json:"avatar"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Status      bool      `json:"status"`
	Role        int32     `json:"role"`
	DeletedAt   time.Time `json:"deleted_at"`
	DeletedBy   uint32    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	GetById(context.Context, uint32) (*User, error)
	GetByUsername(context.Context, string) (*User, error)
	Create(context.Context, *User) (*User, error)
	List(ctx context.Context, pageNum, pageSize int32) ([]*v1.UserInfoReply, int32, error)
	Update(context.Context, *User) (bool, error)
	UpdatePassword(context.Context, *User) (bool, error)
	DeleteById(context.Context, uint32) (bool, error)
	SoftDeleteById(context.Context, uint32, uint32) (bool, error)
	SetToken(context.Context, string, string) (bool, error)
	EmptyToken(context.Context, string) (bool, error)
}

type UserUseCase struct {
	repo       UserRepo
	log        *log.Helper
	signingKey string
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	helper := log.NewHelper(log.With(logger, "module", "useCase/user"))
	return &UserUseCase{repo: repo, log: helper}
}

func (uc *UserUseCase) EmptyToken(ctx context.Context, uid uint32) (bool, error) {
	u, err := uc.repo.GetById(ctx, uid)
	if err != nil {
		return false, err
	}
	return uc.repo.EmptyToken(ctx, u.Username)
}

func (uc *UserUseCase) SetToken(ctx context.Context, username string, token string) (bool, error) {
	return uc.repo.SetToken(ctx, username, token)
}

func (uc *UserUseCase) SoftDeleteById(ctx context.Context, uid, deleteId uint32) (bool, error) {
	return uc.repo.SoftDeleteById(ctx, uid, deleteId)
}

func (uc *UserUseCase) Create(ctx context.Context, user *User) (*User, error) {
	return uc.repo.Create(ctx, user)
}

func (uc *UserUseCase) Get(ctx context.Context, id uint32) (*User, error) {
	return uc.repo.GetById(ctx, id)
}

func (uc *UserUseCase) List(ctx context.Context, pageNum, pageSize int32) ([]*v1.UserInfoReply, int32, error) {
	return uc.repo.List(ctx, pageNum, pageSize)
}

func (uc *UserUseCase) Update(ctx context.Context, user *User) (bool, error) {
	return uc.repo.Update(ctx, user)
}

func (uc *UserUseCase) Delete(ctx context.Context, id uint32) (bool, error) {
	return uc.repo.DeleteById(ctx, id)
}

func (uc *UserUseCase) GetByUsername(ctx context.Context, username string) (*User, error) {
	return uc.repo.GetByUsername(ctx, username)
}

func (uc *UserUseCase) VerifyPassword(password, hashedPassword string) (bool, error) {
	if ok := util.ComparePassword(password, hashedPassword); !ok {
		return false, SetErrByReason(ReasonUserPasswordError)
	}
	return true, nil
}

func (uc *UserUseCase) UpdatePassword(ctx context.Context, u *User, newPassword string) (bool, error) {
	encryptedPassword, err := util.HashPassword(newPassword)
	if err != nil {
		return false, err
	}
	u.Password = encryptedPassword
	return uc.repo.UpdatePassword(ctx, u)
}
