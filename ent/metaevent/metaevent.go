// Code generated by ent, DO NOT EDIT.

package metaevent

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the metaevent type in the database.
	Label = "meta_event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEventName holds the string denoting the event_name field in the database.
	FieldEventName = "event_name"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldEventDesc holds the string denoting the event_desc field in the database.
	FieldEventDesc = "event_desc"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the metaevent in the database.
	Table = "meta_event"
)

// Columns holds all SQL columns for metaevent fields.
var Columns = []string{
	FieldID,
	FieldEventName,
	FieldUUID,
	FieldDisplayName,
	FieldEventDesc,
	FieldRemark,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// EventNameValidator is a validator for the "event_name" field. It is called by the builders before save.
	EventNameValidator func(string) error
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
	// DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	DisplayNameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
