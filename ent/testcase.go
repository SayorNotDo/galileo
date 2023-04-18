// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"galileo/ent/testcase"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TestCase is the model entity for the TestCase schema.
type TestCase struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy uint32 `json:"created_by,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdateBy holds the value of the "update_by" field.
	UpdateBy uint32 `json:"update_by,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
	// Type holds the value of the "type" field.
	Type int8 `json:"type,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority int8 `json:"priority,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy uint32 `json:"deleted_by,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// URL holds the value of the "url" field.
	URL                      string `json:"url,omitempty"`
	test_case_suite_testcase *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TestCase) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case testcase.FieldID, testcase.FieldCreatedBy, testcase.FieldUpdateBy, testcase.FieldStatus, testcase.FieldType, testcase.FieldPriority, testcase.FieldDeletedBy:
			values[i] = new(sql.NullInt64)
		case testcase.FieldName, testcase.FieldDescription, testcase.FieldURL:
			values[i] = new(sql.NullString)
		case testcase.FieldCreatedAt, testcase.FieldUpdateAt, testcase.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case testcase.ForeignKeys[0]: // test_case_suite_testcase
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TestCase", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TestCase fields.
func (tc *TestCase) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case testcase.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tc.ID = int64(value.Int64)
		case testcase.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tc.Name = value.String
			}
		case testcase.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				tc.CreatedBy = uint32(value.Int64)
			}
		case testcase.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tc.CreatedAt = value.Time
			}
		case testcase.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_by", values[i])
			} else if value.Valid {
				tc.UpdateBy = uint32(value.Int64)
			}
		case testcase.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				tc.UpdateAt = value.Time
			}
		case testcase.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				tc.Status = int8(value.Int64)
			}
		case testcase.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				tc.Type = int8(value.Int64)
			}
		case testcase.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				tc.Priority = int8(value.Int64)
			}
		case testcase.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				tc.DeletedAt = value.Time
			}
		case testcase.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				tc.DeletedBy = uint32(value.Int64)
			}
		case testcase.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				tc.Description = value.String
			}
		case testcase.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				tc.URL = value.String
			}
		case testcase.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field test_case_suite_testcase", value)
			} else if value.Valid {
				tc.test_case_suite_testcase = new(int)
				*tc.test_case_suite_testcase = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TestCase.
// Note that you need to call TestCase.Unwrap() before calling this method if this TestCase
// was returned from a transaction, and the transaction was committed or rolled back.
func (tc *TestCase) Update() *TestCaseUpdateOne {
	return NewTestCaseClient(tc.config).UpdateOne(tc)
}

// Unwrap unwraps the TestCase entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tc *TestCase) Unwrap() *TestCase {
	_tx, ok := tc.config.driver.(*txDriver)
	if !ok {
		panic("ent: TestCase is not a transactional entity")
	}
	tc.config.driver = _tx.drv
	return tc
}

// String implements the fmt.Stringer.
func (tc *TestCase) String() string {
	var builder strings.Builder
	builder.WriteString("TestCase(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tc.ID))
	builder.WriteString("name=")
	builder.WriteString(tc.Name)
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", tc.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(tc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_by=")
	builder.WriteString(fmt.Sprintf("%v", tc.UpdateBy))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(tc.UpdateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", tc.Status))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", tc.Type))
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", tc.Priority))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(tc.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(fmt.Sprintf("%v", tc.DeletedBy))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(tc.Description)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(tc.URL)
	builder.WriteByte(')')
	return builder.String()
}

// TestCases is a parsable slice of TestCase.
type TestCases []*TestCase
