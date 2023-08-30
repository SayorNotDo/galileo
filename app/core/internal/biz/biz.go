package biz

import (
	"github.com/google/uuid"
	"github.com/google/wire"
	"time"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewCoreUseCase, NewEngineUseCase, NewUserUseCase)

type User struct {
	Id            uint32    `json:"id"`
	UUID          uuid.UUID `json:"uuid"`
	Username      string    `json:"username"`
	ChineseName   string    `json:"chinese_name"`
	Password      string    `json:"password"`
	Avatar        string    `json:"avatar"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Active        bool      `json:"active"`
	Location      string    `json:"location"`
	LastLoginTime time.Time `json:"last_login_time"`
	CreatedAt     time.Time `json:"created_at"`
	UpdateAt      time.Time `json:"update_at"`
	DeletedAt     time.Time `json:"deleted_at"`
	DeletedBy     uint32    `json:"deleted_by"`
}

type Group struct {
	Id              int32          `json:"id"`
	Name            string         `json:"name"`
	Avatar          string         `json:"avatar"`
	Description     string         `json:"description"`
	Headcount       int32          `json:"headcount"`
	CreatedAt       time.Time      `json:"created_at"`
	CreatedBy       uint32         `json:"created_by"`
	UpdatedAt       time.Time      `json:"updated_at"`
	UpdatedBy       uint32         `json:"updated_by"`
	DeletedAt       time.Time      `json:"deleted_at"`
	DeletedBy       uint32         `json:"deleted_by"`
	GroupMemberList []*GroupMember `json:"group_member_list"`
}

type GroupMember struct {
	Uid       uint32    `json:"uid"`
	Username  string    `json:"username"`
	Role      uint8     `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy uint32    `json:"created_by"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy uint32    `json:"deleted_by"`
}
