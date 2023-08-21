// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"galileo/ent/apistatistics"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ApiStatistics is the model entity for the ApiStatistics schema.
type ApiStatistics struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// CallCount holds the value of the "call_count" field.
	CallCount int32 `json:"call_count,omitempty"`
	// SuccessCount holds the value of the "success_count" field.
	SuccessCount int32 `json:"success_count,omitempty"`
	// FailureCount holds the value of the "failure_count" field.
	FailureCount int32 `json:"failure_count,omitempty"`
	// AvgResponseTime holds the value of the "avg_response_time" field.
	AvgResponseTime float64 `json:"avg_response_time,omitempty"`
	// MaxResponseTime holds the value of the "max_response_time" field.
	MaxResponseTime float64 `json:"max_response_time,omitempty"`
	// MinResponseTime holds the value of the "min_response_time" field.
	MinResponseTime float64 `json:"min_response_time,omitempty"`
	// AvgTraffic holds the value of the "avg_traffic" field.
	AvgTraffic float64 `json:"avg_traffic,omitempty"`
	// MaxTraffic holds the value of the "max_traffic" field.
	MaxTraffic float64 `json:"max_traffic,omitempty"`
	// MinTraffic holds the value of the "min_traffic" field.
	MinTraffic float64 `json:"min_traffic,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// APIID holds the value of the "api_id" field.
	APIID int32 `json:"api_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ApiStatistics) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case apistatistics.FieldAvgResponseTime, apistatistics.FieldMaxResponseTime, apistatistics.FieldMinResponseTime, apistatistics.FieldAvgTraffic, apistatistics.FieldMaxTraffic, apistatistics.FieldMinTraffic:
			values[i] = new(sql.NullFloat64)
		case apistatistics.FieldID, apistatistics.FieldCallCount, apistatistics.FieldSuccessCount, apistatistics.FieldFailureCount, apistatistics.FieldAPIID:
			values[i] = new(sql.NullInt64)
		case apistatistics.FieldDescription:
			values[i] = new(sql.NullString)
		case apistatistics.FieldCreatedAt, apistatistics.FieldUpdateAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ApiStatistics", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ApiStatistics fields.
func (as *ApiStatistics) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apistatistics.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			as.ID = int32(value.Int64)
		case apistatistics.FieldCallCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field call_count", values[i])
			} else if value.Valid {
				as.CallCount = int32(value.Int64)
			}
		case apistatistics.FieldSuccessCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field success_count", values[i])
			} else if value.Valid {
				as.SuccessCount = int32(value.Int64)
			}
		case apistatistics.FieldFailureCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field failure_count", values[i])
			} else if value.Valid {
				as.FailureCount = int32(value.Int64)
			}
		case apistatistics.FieldAvgResponseTime:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field avg_response_time", values[i])
			} else if value.Valid {
				as.AvgResponseTime = value.Float64
			}
		case apistatistics.FieldMaxResponseTime:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field max_response_time", values[i])
			} else if value.Valid {
				as.MaxResponseTime = value.Float64
			}
		case apistatistics.FieldMinResponseTime:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field min_response_time", values[i])
			} else if value.Valid {
				as.MinResponseTime = value.Float64
			}
		case apistatistics.FieldAvgTraffic:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field avg_traffic", values[i])
			} else if value.Valid {
				as.AvgTraffic = value.Float64
			}
		case apistatistics.FieldMaxTraffic:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field max_traffic", values[i])
			} else if value.Valid {
				as.MaxTraffic = value.Float64
			}
		case apistatistics.FieldMinTraffic:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field min_traffic", values[i])
			} else if value.Valid {
				as.MinTraffic = value.Float64
			}
		case apistatistics.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				as.Description = value.String
			}
		case apistatistics.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				as.CreatedAt = value.Time
			}
		case apistatistics.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				as.UpdateAt = value.Time
			}
		case apistatistics.FieldAPIID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field api_id", values[i])
			} else if value.Valid {
				as.APIID = int32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ApiStatistics.
// Note that you need to call ApiStatistics.Unwrap() before calling this method if this ApiStatistics
// was returned from a transaction, and the transaction was committed or rolled back.
func (as *ApiStatistics) Update() *ApiStatisticsUpdateOne {
	return NewApiStatisticsClient(as.config).UpdateOne(as)
}

// Unwrap unwraps the ApiStatistics entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (as *ApiStatistics) Unwrap() *ApiStatistics {
	_tx, ok := as.config.driver.(*txDriver)
	if !ok {
		panic("ent: ApiStatistics is not a transactional entity")
	}
	as.config.driver = _tx.drv
	return as
}

// String implements the fmt.Stringer.
func (as *ApiStatistics) String() string {
	var builder strings.Builder
	builder.WriteString("ApiStatistics(")
	builder.WriteString(fmt.Sprintf("id=%v, ", as.ID))
	builder.WriteString("call_count=")
	builder.WriteString(fmt.Sprintf("%v", as.CallCount))
	builder.WriteString(", ")
	builder.WriteString("success_count=")
	builder.WriteString(fmt.Sprintf("%v", as.SuccessCount))
	builder.WriteString(", ")
	builder.WriteString("failure_count=")
	builder.WriteString(fmt.Sprintf("%v", as.FailureCount))
	builder.WriteString(", ")
	builder.WriteString("avg_response_time=")
	builder.WriteString(fmt.Sprintf("%v", as.AvgResponseTime))
	builder.WriteString(", ")
	builder.WriteString("max_response_time=")
	builder.WriteString(fmt.Sprintf("%v", as.MaxResponseTime))
	builder.WriteString(", ")
	builder.WriteString("min_response_time=")
	builder.WriteString(fmt.Sprintf("%v", as.MinResponseTime))
	builder.WriteString(", ")
	builder.WriteString("avg_traffic=")
	builder.WriteString(fmt.Sprintf("%v", as.AvgTraffic))
	builder.WriteString(", ")
	builder.WriteString("max_traffic=")
	builder.WriteString(fmt.Sprintf("%v", as.MaxTraffic))
	builder.WriteString(", ")
	builder.WriteString("min_traffic=")
	builder.WriteString(fmt.Sprintf("%v", as.MinTraffic))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(as.Description)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(as.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(as.UpdateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("api_id=")
	builder.WriteString(fmt.Sprintf("%v", as.APIID))
	builder.WriteByte(')')
	return builder.String()
}

// ApiStatisticsSlice is a parsable slice of ApiStatistics.
type ApiStatisticsSlice []*ApiStatistics
