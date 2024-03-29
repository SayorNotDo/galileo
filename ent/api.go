// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"galileo/ent/api"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Api is the model entity for the Api schema.
type Api struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Type holds the value of the "type" field.
	Type int8 `json:"type,omitempty"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
	// Headers holds the value of the "headers" field.
	Headers string `json:"headers,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// Label holds the value of the "label" field.
	Label string `json:"label,omitempty"`
	// QueryParams holds the value of the "query_params" field.
	QueryParams string `json:"query_params,omitempty"`
	// Response holds the value of the "response" field.
	Response string `json:"response,omitempty"`
	// Module holds the value of the "module" field.
	Module string `json:"module,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy uint32 `json:"created_by,omitempty"`
	// IncludeFiles holds the value of the "include_files" field.
	IncludeFiles string `json:"include_files,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// UpdateBy holds the value of the "update_by" field.
	UpdateBy uint32 `json:"update_by,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy uint32 `json:"deleted_by,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Api) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case api.FieldID, api.FieldType, api.FieldStatus, api.FieldCreatedBy, api.FieldUpdateBy, api.FieldDeletedBy:
			values[i] = new(sql.NullInt64)
		case api.FieldName, api.FieldURL, api.FieldHeaders, api.FieldBody, api.FieldLabel, api.FieldQueryParams, api.FieldResponse, api.FieldModule, api.FieldDescription, api.FieldIncludeFiles:
			values[i] = new(sql.NullString)
		case api.FieldCreatedAt, api.FieldUpdateAt, api.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Api", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Api fields.
func (a *Api) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case api.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int32(value.Int64)
		case api.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case api.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				a.URL = value.String
			}
		case api.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = int8(value.Int64)
			}
		case api.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				a.Status = int8(value.Int64)
			}
		case api.FieldHeaders:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field headers", values[i])
			} else if value.Valid {
				a.Headers = value.String
			}
		case api.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				a.Body = value.String
			}
		case api.FieldLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field label", values[i])
			} else if value.Valid {
				a.Label = value.String
			}
		case api.FieldQueryParams:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field query_params", values[i])
			} else if value.Valid {
				a.QueryParams = value.String
			}
		case api.FieldResponse:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field response", values[i])
			} else if value.Valid {
				a.Response = value.String
			}
		case api.FieldModule:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field module", values[i])
			} else if value.Valid {
				a.Module = value.String
			}
		case api.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case api.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case api.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				a.CreatedBy = uint32(value.Int64)
			}
		case api.FieldIncludeFiles:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field include_files", values[i])
			} else if value.Valid {
				a.IncludeFiles = value.String
			}
		case api.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				a.UpdateAt = value.Time
			}
		case api.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_by", values[i])
			} else if value.Valid {
				a.UpdateBy = uint32(value.Int64)
			}
		case api.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = value.Time
			}
		case api.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				a.DeletedBy = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Api.
// Note that you need to call Api.Unwrap() before calling this method if this Api
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Api) Update() *APIUpdateOne {
	return NewAPIClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Api entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Api) Unwrap() *Api {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Api is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Api) String() string {
	var builder strings.Builder
	builder.WriteString("Api(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(a.URL)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", a.Type))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", a.Status))
	builder.WriteString(", ")
	builder.WriteString("headers=")
	builder.WriteString(a.Headers)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(a.Body)
	builder.WriteString(", ")
	builder.WriteString("label=")
	builder.WriteString(a.Label)
	builder.WriteString(", ")
	builder.WriteString("query_params=")
	builder.WriteString(a.QueryParams)
	builder.WriteString(", ")
	builder.WriteString("response=")
	builder.WriteString(a.Response)
	builder.WriteString(", ")
	builder.WriteString("module=")
	builder.WriteString(a.Module)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(a.Description)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", a.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("include_files=")
	builder.WriteString(a.IncludeFiles)
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(a.UpdateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_by=")
	builder.WriteString(fmt.Sprintf("%v", a.UpdateBy))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(a.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(fmt.Sprintf("%v", a.DeletedBy))
	builder.WriteByte(')')
	return builder.String()
}

// Apis is a parsable slice of Api.
type Apis []*Api
