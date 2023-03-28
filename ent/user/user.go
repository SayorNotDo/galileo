// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldChineseName holds the string denoting the chinesename field in the database.
	FieldChineseName = "chinese_name"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// Table holds the table name of the user in the database.
	Table = "user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldActive,
	FieldUsername,
	FieldChineseName,
	FieldNickname,
	FieldPassword,
	FieldPhone,
	FieldEmail,
	FieldAvatar,
	FieldRole,
	FieldUpdateAt,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldIsDeleted,
	FieldUUID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// DefaultRole holds the default value on creation for the "role" field.
	DefaultRole uint8
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() time.Time
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted bool
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
)
