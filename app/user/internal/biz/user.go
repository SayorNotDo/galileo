package biz

import (
	"context"
	"galileo/app/user/internal/pkg/util"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	GetUserInfo(context.Context, uint32) (*User, error)
	CreateUser(context.Context, *User) (*User, error)
	ListUser(ctx context.Context, pageToken string, pageSize int32) ([]*User, int32, string, error)
	UpdateUser(context.Context, *User) (bool, error)
	ValidateUser(ctx context.Context, username, password string) (*User, error)
	UpdatePassword(context.Context, *User) (bool, error)
	SoftDeleteById(context.Context, uint32, uint32) (bool, error)
	SetToken(context.Context, string, string) (bool, error)
	EmptyToken(context.Context, string) (bool, error)
	GetUserGroupList(context.Context, uint32) ([]*Group, error)
	GetUserGroup(ctx context.Context, groupId int64) (*Group, error)
	UpdateUserGroup(context.Context, *Group) error
	CreateUserGroup(context.Context, *Group) (*Group, error)
}

type UserUseCase struct {
	repo       UserRepo
	log        *log.Helper
	signingKey string
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	helper := log.NewHelper(log.With(logger, "module", "UseCase.user"))
	return &UserUseCase{repo: repo, log: helper}
}

func (uc *UserUseCase) EmptyToken(ctx context.Context, uid uint32) (bool, error) {
	u, err := uc.repo.GetUserInfo(ctx, uid)
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

func (uc *UserUseCase) CreateUser(ctx context.Context, user *User) (*User, error) {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUseCase) GetUserInfo(ctx context.Context, id uint32) (*User, error) {
	return uc.repo.GetUserInfo(ctx, id)
}

func (uc *UserUseCase) ListUser(ctx context.Context, pageToken string, pageSize int32) ([]*User, int32, string, error) {
	return uc.repo.ListUser(ctx, pageToken, pageSize)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, user *User) (bool, error) {
	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserUseCase) ValidateUser(ctx context.Context, username, password string) (*User, error) {
	return uc.repo.ValidateUser(ctx, username, password)
}

func (uc *UserUseCase) UpdatePassword(ctx context.Context, u *User, newPassword string) (bool, error) {
	encryptedPassword, err := util.HashPassword(newPassword)
	if err != nil {
		return false, err
	}
	u.Password = encryptedPassword
	return uc.repo.UpdatePassword(ctx, u)
}

func (uc *UserUseCase) GetUserGroupList(ctx context.Context, uid uint32) ([]*Group, error) {
	return uc.repo.GetUserGroupList(ctx, uid)
}

func (uc *UserUseCase) GetUserGroup(ctx context.Context, groupId int64) (*Group, error) {
	return uc.repo.GetUserGroup(ctx, groupId)
}

func (uc *UserUseCase) UpdateUserGroup(ctx context.Context, group *Group) error {
	return uc.repo.UpdateUserGroup(ctx, group)
}

func (uc *UserUseCase) CreateUserGroup(ctx context.Context, group *Group) (*Group, error) {
	return uc.repo.CreateUserGroup(ctx, group)
}
