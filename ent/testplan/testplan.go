// Code generated by ent, DO NOT EDIT.

package testplan

import (
	"time"
)

const (
	// Label holds the string label denoting the testplan type in the database.
	Label = "test_plan"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldDeadline holds the string denoting the deadline field in the database.
	FieldDeadline = "deadline"
	// FieldStatusUpdatedAt holds the string denoting the status_updated_at field in the database.
	FieldStatusUpdatedAt = "status_updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldTasks holds the string denoting the tasks field in the database.
	FieldTasks = "tasks"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// Table holds the table name of the testplan in the database.
	Table = "test_plan"
)

// Columns holds all SQL columns for testplan fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldDescription,
	FieldStartTime,
	FieldDeadline,
	FieldStatusUpdatedAt,
	FieldStatus,
	FieldTasks,
	FieldProjectID,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
	// ProjectIDValidator is a validator for the "project_id" field. It is called by the builders before save.
	ProjectIDValidator func(int64) error
)
