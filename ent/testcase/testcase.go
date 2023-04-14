// Code generated by ent, DO NOT EDIT.

package testcase

import (
	"time"
)

const (
	// Label holds the string label denoting the testcase type in the database.
	Label = "test_case"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdateBy holds the string denoting the update_by field in the database.
	FieldUpdateBy = "update_by"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// Table holds the table name of the testcase in the database.
	Table = "test_cases"
)

// Columns holds all SQL columns for testcase fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdateBy,
	FieldUpdateAt,
	FieldStatus,
	FieldType,
	FieldPriority,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldDescription,
	FieldURL,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "test_cases"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"test_case_suite_testcase",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType int16
	// DefaultPriority holds the default value on creation for the "priority" field.
	DefaultPriority int8
)