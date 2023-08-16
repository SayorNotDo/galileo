// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"galileo/ent/testcasesuite"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TestcaseSuite is the model entity for the TestcaseSuite schema.
type TestcaseSuite struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
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
	// Testcases holds the value of the "testcases" field.
	Testcases           []int64 `json:"testcases,omitempty"`
	task_testcase_suite *int64
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TestcaseSuite) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case testcasesuite.FieldTestcases:
			values[i] = new([]byte)
		case testcasesuite.FieldID, testcasesuite.FieldCreatedBy, testcasesuite.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case testcasesuite.FieldName:
			values[i] = new(sql.NullString)
		case testcasesuite.FieldCreatedAt, testcasesuite.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case testcasesuite.ForeignKeys[0]: // task_testcase_suite
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TestcaseSuite", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TestcaseSuite fields.
func (ts *TestcaseSuite) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case testcasesuite.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ts.ID = int64(value.Int64)
		case testcasesuite.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ts.Name = value.String
			}
		case testcasesuite.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ts.CreatedAt = value.Time
			}
		case testcasesuite.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ts.CreatedBy = uint32(value.Int64)
			}
		case testcasesuite.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ts.UpdatedAt = value.Time
			}
		case testcasesuite.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ts.UpdatedBy = uint32(value.Int64)
			}
		case testcasesuite.FieldTestcases:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field testcases", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ts.Testcases); err != nil {
					return fmt.Errorf("unmarshal field testcases: %w", err)
				}
			}
		case testcasesuite.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field task_testcase_suite", value)
			} else if value.Valid {
				ts.task_testcase_suite = new(int64)
				*ts.task_testcase_suite = int64(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TestcaseSuite.
// Note that you need to call TestcaseSuite.Unwrap() before calling this method if this TestcaseSuite
// was returned from a transaction, and the transaction was committed or rolled back.
func (ts *TestcaseSuite) Update() *TestcaseSuiteUpdateOne {
	return NewTestcaseSuiteClient(ts.config).UpdateOne(ts)
}

// Unwrap unwraps the TestcaseSuite entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ts *TestcaseSuite) Unwrap() *TestcaseSuite {
	_tx, ok := ts.config.driver.(*txDriver)
	if !ok {
		panic("ent: TestcaseSuite is not a transactional entity")
	}
	ts.config.driver = _tx.drv
	return ts
}

// String implements the fmt.Stringer.
func (ts *TestcaseSuite) String() string {
	var builder strings.Builder
	builder.WriteString("TestcaseSuite(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ts.ID))
	builder.WriteString("name=")
	builder.WriteString(ts.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ts.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ts.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ts.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ts.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("testcases=")
	builder.WriteString(fmt.Sprintf("%v", ts.Testcases))
	builder.WriteByte(')')
	return builder.String()
}

// TestcaseSuites is a parsable slice of TestcaseSuite.
type TestcaseSuites []*TestcaseSuite
