// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"galileo/ent/project"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Project is the model entity for the Project schema.
type Project struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Identifier holds the value of the "identifier" field.
	Identifier string `json:"identifier,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy uint32 `json:"created_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UpdateBy holds the value of the "update_by" field.
	UpdateBy *uint32 `json:"update_by,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy *uint32 `json:"deleted_by,omitempty"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark *string `json:"remark,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Project) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case project.FieldID, project.FieldCreatedBy, project.FieldUpdateBy, project.FieldDeletedBy, project.FieldStatus:
			values[i] = new(sql.NullInt64)
		case project.FieldName, project.FieldIdentifier, project.FieldDescription, project.FieldRemark:
			values[i] = new(sql.NullString)
		case project.FieldCreatedAt, project.FieldUpdatedAt, project.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Project", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Project fields.
func (pr *Project) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = uint32(value.Int64)
		case project.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case project.FieldIdentifier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identifier", values[i])
			} else if value.Valid {
				pr.Identifier = value.String
			}
		case project.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case project.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pr.CreatedBy = uint32(value.Int64)
			}
		case project.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case project.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_by", values[i])
			} else if value.Valid {
				pr.UpdateBy = new(uint32)
				*pr.UpdateBy = uint32(value.Int64)
			}
		case project.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pr.DeletedAt = new(time.Time)
				*pr.DeletedAt = value.Time
			}
		case project.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				pr.DeletedBy = new(uint32)
				*pr.DeletedBy = uint32(value.Int64)
			}
		case project.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pr.Status = int8(value.Int64)
			}
		case project.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = new(string)
				*pr.Description = value.String
			}
		case project.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				pr.Remark = new(string)
				*pr.Remark = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Project.
// Note that you need to call Project.Unwrap() before calling this method if this Project
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Project) Update() *ProjectUpdateOne {
	return NewProjectClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Project entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Project) Unwrap() *Project {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Project is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Project) String() string {
	var builder strings.Builder
	builder.WriteString("Project(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("identifier=")
	builder.WriteString(pr.Identifier)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", pr.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := pr.UpdateBy; v != nil {
		builder.WriteString("update_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := pr.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := pr.DeletedBy; v != nil {
		builder.WriteString("deleted_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pr.Status))
	builder.WriteString(", ")
	if v := pr.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := pr.Remark; v != nil {
		builder.WriteString("remark=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Projects is a parsable slice of Project.
type Projects []*Project
