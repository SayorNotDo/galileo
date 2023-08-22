package biz

import (
	"github.com/google/wire"
	"time"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUseCase)

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
	UpdateAt    time.Time `json:"update_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	DeletedBy   uint32    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserGroup struct {
	GroupMemberId int32 `json:"group_member_id"`
	Role          uint8 `json:"role"`
	GroupInfo     Group `json:"group_info"`
}

type Group struct {
	Id          int32     `json:"id"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	Description string    `json:"description"`
	Headcount   int32     `json:"headcount"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   uint32    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   uint32    `json:"updated_by"`
	DeletedAt   time.Time `json:"deleted_at"`
	DeletedBy   uint32    `json:"deleted_by"`
}
