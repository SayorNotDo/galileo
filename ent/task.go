// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"galileo/ent/task"
	"net/url"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy uint32 `json:"created_by,omitempty"`
	// Rank holds the value of the "rank" field.
	Rank int8 `json:"rank,omitempty"`
	// Type holds the value of the "type" field.
	Type int16 `json:"type,omitempty"`
	// Status holds the value of the "status" field.
	Status int16 `json:"status,omitempty"`
	// CompleteAt holds the value of the "complete_at" field.
	CompleteAt *time.Time `json:"complete_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// UpdateBy holds the value of the "update_by" field.
	UpdateBy uint32 `json:"update_by,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy *uint32 `json:"deleted_by,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// URL holds the value of the "url" field.
	URL *url.URL `json:"url,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Task) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case task.FieldURL:
			values[i] = new([]byte)
		case task.FieldIsDeleted:
			values[i] = new(sql.NullBool)
		case task.FieldID, task.FieldCreatedBy, task.FieldRank, task.FieldType, task.FieldStatus, task.FieldUpdateBy, task.FieldDeletedBy:
			values[i] = new(sql.NullInt64)
		case task.FieldName, task.FieldDescription:
			values[i] = new(sql.NullString)
		case task.FieldCreatedAt, task.FieldCompleteAt, task.FieldUpdateAt, task.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Task", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Task fields.
func (t *Task) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case task.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int64(value.Int64)
		case task.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case task.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case task.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				t.CreatedBy = uint32(value.Int64)
			}
		case task.FieldRank:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rank", values[i])
			} else if value.Valid {
				t.Rank = int8(value.Int64)
			}
		case task.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = int16(value.Int64)
			}
		case task.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = int16(value.Int64)
			}
		case task.FieldCompleteAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field complete_at", values[i])
			} else if value.Valid {
				t.CompleteAt = new(time.Time)
				*t.CompleteAt = value.Time
			}
		case task.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				t.UpdateAt = value.Time
			}
		case task.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_by", values[i])
			} else if value.Valid {
				t.UpdateBy = uint32(value.Int64)
			}
		case task.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				t.IsDeleted = new(bool)
				*t.IsDeleted = value.Bool
			}
		case task.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				t.DeletedAt = new(time.Time)
				*t.DeletedAt = value.Time
			}
		case task.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				t.DeletedBy = new(uint32)
				*t.DeletedBy = uint32(value.Int64)
			}
		case task.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case task.FieldURL:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.URL); err != nil {
					return fmt.Errorf("unmarshal field url: %w", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Task.
// Note that you need to call Task.Unwrap() before calling this method if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return NewTaskClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Task entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", t.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("rank=")
	builder.WriteString(fmt.Sprintf("%v", t.Rank))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", t.Type))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", ")
	if v := t.CompleteAt; v != nil {
		builder.WriteString("complete_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(t.UpdateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_by=")
	builder.WriteString(fmt.Sprintf("%v", t.UpdateBy))
	builder.WriteString(", ")
	if v := t.IsDeleted; v != nil {
		builder.WriteString("is_deleted=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := t.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := t.DeletedBy; v != nil {
		builder.WriteString("deleted_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(t.Description)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(fmt.Sprintf("%v", t.URL))
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task