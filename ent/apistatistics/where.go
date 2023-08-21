// Code generated by ent, DO NOT EDIT.

package apistatistics

import (
	"galileo/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLTE(FieldID, id))
}

// CallCount applies equality check predicate on the "call_count" field. It's identical to CallCountEQ.
func CallCount(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldCallCount, v))
}

// SuccessCount applies equality check predicate on the "success_count" field. It's identical to SuccessCountEQ.
func SuccessCount(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldSuccessCount, v))
}

// FailureCount applies equality check predicate on the "failure_count" field. It's identical to FailureCountEQ.
func FailureCount(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldFailureCount, v))
}

// AvgResponseTime applies equality check predicate on the "avg_response_time" field. It's identical to AvgResponseTimeEQ.
func AvgResponseTime(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldAvgResponseTime, v))
}

// MaxResponseTime applies equality check predicate on the "max_response_time" field. It's identical to MaxResponseTimeEQ.
func MaxResponseTime(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldMaxResponseTime, v))
}

// MinResponseTime applies equality check predicate on the "min_response_time" field. It's identical to MinResponseTimeEQ.
func MinResponseTime(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldMinResponseTime, v))
}

// AvgTraffic applies equality check predicate on the "avg_traffic" field. It's identical to AvgTrafficEQ.
func AvgTraffic(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldAvgTraffic, v))
}

// MaxTraffic applies equality check predicate on the "max_traffic" field. It's identical to MaxTrafficEQ.
func MaxTraffic(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldMaxTraffic, v))
}

// MinTraffic applies equality check predicate on the "min_traffic" field. It's identical to MinTrafficEQ.
func MinTraffic(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldMinTraffic, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldDescription, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldUpdateAt, v))
}

// APIID applies equality check predicate on the "api_id" field. It's identical to APIIDEQ.
func APIID(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldAPIID, v))
}

// CallCountEQ applies the EQ predicate on the "call_count" field.
func CallCountEQ(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldCallCount, v))
}

// CallCountNEQ applies the NEQ predicate on the "call_count" field.
func CallCountNEQ(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNEQ(FieldCallCount, v))
}

// CallCountIn applies the In predicate on the "call_count" field.
func CallCountIn(vs ...int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldIn(FieldCallCount, vs...))
}

// CallCountNotIn applies the NotIn predicate on the "call_count" field.
func CallCountNotIn(vs ...int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNotIn(FieldCallCount, vs...))
}

// CallCountGT applies the GT predicate on the "call_count" field.
func CallCountGT(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGT(FieldCallCount, v))
}

// CallCountGTE applies the GTE predicate on the "call_count" field.
func CallCountGTE(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGTE(FieldCallCount, v))
}

// CallCountLT applies the LT predicate on the "call_count" field.
func CallCountLT(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLT(FieldCallCount, v))
}

// CallCountLTE applies the LTE predicate on the "call_count" field.
func CallCountLTE(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLTE(FieldCallCount, v))
}

// SuccessCountEQ applies the EQ predicate on the "success_count" field.
func SuccessCountEQ(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldSuccessCount, v))
}

// SuccessCountNEQ applies the NEQ predicate on the "success_count" field.
func SuccessCountNEQ(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNEQ(FieldSuccessCount, v))
}

// SuccessCountIn applies the In predicate on the "success_count" field.
func SuccessCountIn(vs ...int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldIn(FieldSuccessCount, vs...))
}

// SuccessCountNotIn applies the NotIn predicate on the "success_count" field.
func SuccessCountNotIn(vs ...int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNotIn(FieldSuccessCount, vs...))
}

// SuccessCountGT applies the GT predicate on the "success_count" field.
func SuccessCountGT(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGT(FieldSuccessCount, v))
}

// SuccessCountGTE applies the GTE predicate on the "success_count" field.
func SuccessCountGTE(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGTE(FieldSuccessCount, v))
}

// SuccessCountLT applies the LT predicate on the "success_count" field.
func SuccessCountLT(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLT(FieldSuccessCount, v))
}

// SuccessCountLTE applies the LTE predicate on the "success_count" field.
func SuccessCountLTE(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLTE(FieldSuccessCount, v))
}

// FailureCountEQ applies the EQ predicate on the "failure_count" field.
func FailureCountEQ(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldFailureCount, v))
}

// FailureCountNEQ applies the NEQ predicate on the "failure_count" field.
func FailureCountNEQ(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNEQ(FieldFailureCount, v))
}

// FailureCountIn applies the In predicate on the "failure_count" field.
func FailureCountIn(vs ...int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldIn(FieldFailureCount, vs...))
}

// FailureCountNotIn applies the NotIn predicate on the "failure_count" field.
func FailureCountNotIn(vs ...int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNotIn(FieldFailureCount, vs...))
}

// FailureCountGT applies the GT predicate on the "failure_count" field.
func FailureCountGT(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGT(FieldFailureCount, v))
}

// FailureCountGTE applies the GTE predicate on the "failure_count" field.
func FailureCountGTE(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGTE(FieldFailureCount, v))
}

// FailureCountLT applies the LT predicate on the "failure_count" field.
func FailureCountLT(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLT(FieldFailureCount, v))
}

// FailureCountLTE applies the LTE predicate on the "failure_count" field.
func FailureCountLTE(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLTE(FieldFailureCount, v))
}

// AvgResponseTimeEQ applies the EQ predicate on the "avg_response_time" field.
func AvgResponseTimeEQ(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldAvgResponseTime, v))
}

// AvgResponseTimeNEQ applies the NEQ predicate on the "avg_response_time" field.
func AvgResponseTimeNEQ(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNEQ(FieldAvgResponseTime, v))
}

// AvgResponseTimeIn applies the In predicate on the "avg_response_time" field.
func AvgResponseTimeIn(vs ...float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldIn(FieldAvgResponseTime, vs...))
}

// AvgResponseTimeNotIn applies the NotIn predicate on the "avg_response_time" field.
func AvgResponseTimeNotIn(vs ...float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNotIn(FieldAvgResponseTime, vs...))
}

// AvgResponseTimeGT applies the GT predicate on the "avg_response_time" field.
func AvgResponseTimeGT(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGT(FieldAvgResponseTime, v))
}

// AvgResponseTimeGTE applies the GTE predicate on the "avg_response_time" field.
func AvgResponseTimeGTE(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGTE(FieldAvgResponseTime, v))
}

// AvgResponseTimeLT applies the LT predicate on the "avg_response_time" field.
func AvgResponseTimeLT(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLT(FieldAvgResponseTime, v))
}

// AvgResponseTimeLTE applies the LTE predicate on the "avg_response_time" field.
func AvgResponseTimeLTE(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLTE(FieldAvgResponseTime, v))
}

// MaxResponseTimeEQ applies the EQ predicate on the "max_response_time" field.
func MaxResponseTimeEQ(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldMaxResponseTime, v))
}

// MaxResponseTimeNEQ applies the NEQ predicate on the "max_response_time" field.
func MaxResponseTimeNEQ(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNEQ(FieldMaxResponseTime, v))
}

// MaxResponseTimeIn applies the In predicate on the "max_response_time" field.
func MaxResponseTimeIn(vs ...float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldIn(FieldMaxResponseTime, vs...))
}

// MaxResponseTimeNotIn applies the NotIn predicate on the "max_response_time" field.
func MaxResponseTimeNotIn(vs ...float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNotIn(FieldMaxResponseTime, vs...))
}

// MaxResponseTimeGT applies the GT predicate on the "max_response_time" field.
func MaxResponseTimeGT(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGT(FieldMaxResponseTime, v))
}

// MaxResponseTimeGTE applies the GTE predicate on the "max_response_time" field.
func MaxResponseTimeGTE(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGTE(FieldMaxResponseTime, v))
}

// MaxResponseTimeLT applies the LT predicate on the "max_response_time" field.
func MaxResponseTimeLT(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLT(FieldMaxResponseTime, v))
}

// MaxResponseTimeLTE applies the LTE predicate on the "max_response_time" field.
func MaxResponseTimeLTE(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLTE(FieldMaxResponseTime, v))
}

// MinResponseTimeEQ applies the EQ predicate on the "min_response_time" field.
func MinResponseTimeEQ(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldMinResponseTime, v))
}

// MinResponseTimeNEQ applies the NEQ predicate on the "min_response_time" field.
func MinResponseTimeNEQ(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNEQ(FieldMinResponseTime, v))
}

// MinResponseTimeIn applies the In predicate on the "min_response_time" field.
func MinResponseTimeIn(vs ...float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldIn(FieldMinResponseTime, vs...))
}

// MinResponseTimeNotIn applies the NotIn predicate on the "min_response_time" field.
func MinResponseTimeNotIn(vs ...float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNotIn(FieldMinResponseTime, vs...))
}

// MinResponseTimeGT applies the GT predicate on the "min_response_time" field.
func MinResponseTimeGT(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGT(FieldMinResponseTime, v))
}

// MinResponseTimeGTE applies the GTE predicate on the "min_response_time" field.
func MinResponseTimeGTE(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGTE(FieldMinResponseTime, v))
}

// MinResponseTimeLT applies the LT predicate on the "min_response_time" field.
func MinResponseTimeLT(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLT(FieldMinResponseTime, v))
}

// MinResponseTimeLTE applies the LTE predicate on the "min_response_time" field.
func MinResponseTimeLTE(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLTE(FieldMinResponseTime, v))
}

// AvgTrafficEQ applies the EQ predicate on the "avg_traffic" field.
func AvgTrafficEQ(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldAvgTraffic, v))
}

// AvgTrafficNEQ applies the NEQ predicate on the "avg_traffic" field.
func AvgTrafficNEQ(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNEQ(FieldAvgTraffic, v))
}

// AvgTrafficIn applies the In predicate on the "avg_traffic" field.
func AvgTrafficIn(vs ...float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldIn(FieldAvgTraffic, vs...))
}

// AvgTrafficNotIn applies the NotIn predicate on the "avg_traffic" field.
func AvgTrafficNotIn(vs ...float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNotIn(FieldAvgTraffic, vs...))
}

// AvgTrafficGT applies the GT predicate on the "avg_traffic" field.
func AvgTrafficGT(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGT(FieldAvgTraffic, v))
}

// AvgTrafficGTE applies the GTE predicate on the "avg_traffic" field.
func AvgTrafficGTE(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGTE(FieldAvgTraffic, v))
}

// AvgTrafficLT applies the LT predicate on the "avg_traffic" field.
func AvgTrafficLT(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLT(FieldAvgTraffic, v))
}

// AvgTrafficLTE applies the LTE predicate on the "avg_traffic" field.
func AvgTrafficLTE(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLTE(FieldAvgTraffic, v))
}

// MaxTrafficEQ applies the EQ predicate on the "max_traffic" field.
func MaxTrafficEQ(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldMaxTraffic, v))
}

// MaxTrafficNEQ applies the NEQ predicate on the "max_traffic" field.
func MaxTrafficNEQ(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNEQ(FieldMaxTraffic, v))
}

// MaxTrafficIn applies the In predicate on the "max_traffic" field.
func MaxTrafficIn(vs ...float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldIn(FieldMaxTraffic, vs...))
}

// MaxTrafficNotIn applies the NotIn predicate on the "max_traffic" field.
func MaxTrafficNotIn(vs ...float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNotIn(FieldMaxTraffic, vs...))
}

// MaxTrafficGT applies the GT predicate on the "max_traffic" field.
func MaxTrafficGT(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGT(FieldMaxTraffic, v))
}

// MaxTrafficGTE applies the GTE predicate on the "max_traffic" field.
func MaxTrafficGTE(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGTE(FieldMaxTraffic, v))
}

// MaxTrafficLT applies the LT predicate on the "max_traffic" field.
func MaxTrafficLT(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLT(FieldMaxTraffic, v))
}

// MaxTrafficLTE applies the LTE predicate on the "max_traffic" field.
func MaxTrafficLTE(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLTE(FieldMaxTraffic, v))
}

// MinTrafficEQ applies the EQ predicate on the "min_traffic" field.
func MinTrafficEQ(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldMinTraffic, v))
}

// MinTrafficNEQ applies the NEQ predicate on the "min_traffic" field.
func MinTrafficNEQ(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNEQ(FieldMinTraffic, v))
}

// MinTrafficIn applies the In predicate on the "min_traffic" field.
func MinTrafficIn(vs ...float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldIn(FieldMinTraffic, vs...))
}

// MinTrafficNotIn applies the NotIn predicate on the "min_traffic" field.
func MinTrafficNotIn(vs ...float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNotIn(FieldMinTraffic, vs...))
}

// MinTrafficGT applies the GT predicate on the "min_traffic" field.
func MinTrafficGT(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGT(FieldMinTraffic, v))
}

// MinTrafficGTE applies the GTE predicate on the "min_traffic" field.
func MinTrafficGTE(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGTE(FieldMinTraffic, v))
}

// MinTrafficLT applies the LT predicate on the "min_traffic" field.
func MinTrafficLT(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLT(FieldMinTraffic, v))
}

// MinTrafficLTE applies the LTE predicate on the "min_traffic" field.
func MinTrafficLTE(v float64) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLTE(FieldMinTraffic, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldContainsFold(FieldDescription, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldUpdateAt, v))
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNEQ(FieldUpdateAt, v))
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldIn(FieldUpdateAt, vs...))
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNotIn(FieldUpdateAt, vs...))
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGT(FieldUpdateAt, v))
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGTE(FieldUpdateAt, v))
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLT(FieldUpdateAt, v))
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLTE(FieldUpdateAt, v))
}

// APIIDEQ applies the EQ predicate on the "api_id" field.
func APIIDEQ(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldEQ(FieldAPIID, v))
}

// APIIDNEQ applies the NEQ predicate on the "api_id" field.
func APIIDNEQ(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNEQ(FieldAPIID, v))
}

// APIIDIn applies the In predicate on the "api_id" field.
func APIIDIn(vs ...int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldIn(FieldAPIID, vs...))
}

// APIIDNotIn applies the NotIn predicate on the "api_id" field.
func APIIDNotIn(vs ...int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldNotIn(FieldAPIID, vs...))
}

// APIIDGT applies the GT predicate on the "api_id" field.
func APIIDGT(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGT(FieldAPIID, v))
}

// APIIDGTE applies the GTE predicate on the "api_id" field.
func APIIDGTE(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldGTE(FieldAPIID, v))
}

// APIIDLT applies the LT predicate on the "api_id" field.
func APIIDLT(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLT(FieldAPIID, v))
}

// APIIDLTE applies the LTE predicate on the "api_id" field.
func APIIDLTE(v int32) predicate.ApiStatistics {
	return predicate.ApiStatistics(sql.FieldLTE(FieldAPIID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ApiStatistics) predicate.ApiStatistics {
	return predicate.ApiStatistics(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ApiStatistics) predicate.ApiStatistics {
	return predicate.ApiStatistics(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ApiStatistics) predicate.ApiStatistics {
	return predicate.ApiStatistics(func(s *sql.Selector) {
		p(s.Not())
	})
}
