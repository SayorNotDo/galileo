// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"galileo/ent/testplan"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TestPlan is the model entity for the TestPlan schema.
type TestPlan struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy uint32 `json:"created_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy uint32 `json:"updated_by,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// StartTime holds the value of the "start_time" field.
	StartTime time.Time `json:"start_time,omitempty"`
	// Deadline holds the value of the "deadline" field.
	Deadline time.Time `json:"deadline,omitempty"`
	// StatusUpdatedAt holds the value of the "status_updated_at" field.
	StatusUpdatedAt time.Time `json:"status_updated_at,omitempty"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
	// Tasks holds the value of the "tasks" field.
	Tasks []int32 `json:"tasks,omitempty"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID int32 `json:"project_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TestPlan) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case testplan.FieldTasks:
			values[i] = new([]byte)
		case testplan.FieldID, testplan.FieldCreatedBy, testplan.FieldUpdatedBy, testplan.FieldStatus, testplan.FieldProjectID:
			values[i] = new(sql.NullInt64)
		case testplan.FieldName, testplan.FieldDescription:
			values[i] = new(sql.NullString)
		case testplan.FieldCreatedAt, testplan.FieldUpdatedAt, testplan.FieldStartTime, testplan.FieldDeadline, testplan.FieldStatusUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TestPlan", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TestPlan fields.
func (tp *TestPlan) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case testplan.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tp.ID = int32(value.Int64)
		case testplan.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tp.Name = value.String
			}
		case testplan.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tp.CreatedAt = value.Time
			}
		case testplan.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				tp.CreatedBy = uint32(value.Int64)
			}
		case testplan.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tp.UpdatedAt = value.Time
			}
		case testplan.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				tp.UpdatedBy = uint32(value.Int64)
			}
		case testplan.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				tp.Description = value.String
			}
		case testplan.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				tp.StartTime = value.Time
			}
		case testplan.FieldDeadline:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deadline", values[i])
			} else if value.Valid {
				tp.Deadline = value.Time
			}
		case testplan.FieldStatusUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field status_updated_at", values[i])
			} else if value.Valid {
				tp.StatusUpdatedAt = value.Time
			}
		case testplan.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				tp.Status = int8(value.Int64)
			}
		case testplan.FieldTasks:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tasks", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &tp.Tasks); err != nil {
					return fmt.Errorf("unmarshal field tasks: %w", err)
				}
			}
		case testplan.FieldProjectID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value.Valid {
				tp.ProjectID = int32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TestPlan.
// Note that you need to call TestPlan.Unwrap() before calling this method if this TestPlan
// was returned from a transaction, and the transaction was committed or rolled back.
func (tp *TestPlan) Update() *TestPlanUpdateOne {
	return NewTestPlanClient(tp.config).UpdateOne(tp)
}

// Unwrap unwraps the TestPlan entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tp *TestPlan) Unwrap() *TestPlan {
	_tx, ok := tp.config.driver.(*txDriver)
	if !ok {
		panic("ent: TestPlan is not a transactional entity")
	}
	tp.config.driver = _tx.drv
	return tp
}

// String implements the fmt.Stringer.
func (tp *TestPlan) String() string {
	var builder strings.Builder
	builder.WriteString("TestPlan(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tp.ID))
	builder.WriteString("name=")
	builder.WriteString(tp.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(tp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", tp.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(tp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", tp.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(tp.Description)
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(tp.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deadline=")
	builder.WriteString(tp.Deadline.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status_updated_at=")
	builder.WriteString(tp.StatusUpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", tp.Status))
	builder.WriteString(", ")
	builder.WriteString("tasks=")
	builder.WriteString(fmt.Sprintf("%v", tp.Tasks))
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", tp.ProjectID))
	builder.WriteByte(')')
	return builder.String()
}

// TestPlans is a parsable slice of TestPlan.
type TestPlans []*TestPlan
