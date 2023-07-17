// Code generated by ent, DO NOT EDIT.

package task

import (
	"galileo/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldName, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v uint32) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreatedBy, v))
}

// Assignee applies equality check predicate on the "assignee" field. It's identical to AssigneeEQ.
func Assignee(v uint32) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldAssignee, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int8) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldType, v))
}

// Config applies equality check predicate on the "config" field. It's identical to ConfigEQ.
func Config(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldConfig, v))
}

// Rank applies equality check predicate on the "rank" field. It's identical to RankEQ.
func Rank(v int8) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldRank, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldStatus, v))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldStartTime, v))
}

// CompletedAt applies equality check predicate on the "completed_at" field. It's identical to CompletedAtEQ.
func CompletedAt(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCompletedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldUpdatedAt, v))
}

// StatusUpdatedAt applies equality check predicate on the "status_updated_at" field. It's identical to StatusUpdatedAtEQ.
func StatusUpdatedAt(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldStatusUpdatedAt, v))
}

// Deadline applies equality check predicate on the "deadline" field. It's identical to DeadlineEQ.
func Deadline(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDeadline, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v uint32) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDeletedBy, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDescription, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldName, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v uint32) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v uint32) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...uint32) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...uint32) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v uint32) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v uint32) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v uint32) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v uint32) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCreatedBy, v))
}

// AssigneeEQ applies the EQ predicate on the "assignee" field.
func AssigneeEQ(v uint32) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldAssignee, v))
}

// AssigneeNEQ applies the NEQ predicate on the "assignee" field.
func AssigneeNEQ(v uint32) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldAssignee, v))
}

// AssigneeIn applies the In predicate on the "assignee" field.
func AssigneeIn(vs ...uint32) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldAssignee, vs...))
}

// AssigneeNotIn applies the NotIn predicate on the "assignee" field.
func AssigneeNotIn(vs ...uint32) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldAssignee, vs...))
}

// AssigneeGT applies the GT predicate on the "assignee" field.
func AssigneeGT(v uint32) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldAssignee, v))
}

// AssigneeGTE applies the GTE predicate on the "assignee" field.
func AssigneeGTE(v uint32) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldAssignee, v))
}

// AssigneeLT applies the LT predicate on the "assignee" field.
func AssigneeLT(v uint32) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldAssignee, v))
}

// AssigneeLTE applies the LTE predicate on the "assignee" field.
func AssigneeLTE(v uint32) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldAssignee, v))
}

// AssigneeIsNil applies the IsNil predicate on the "assignee" field.
func AssigneeIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldAssignee))
}

// AssigneeNotNil applies the NotNil predicate on the "assignee" field.
func AssigneeNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldAssignee))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int8) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int8) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int8) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int8) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v int8) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v int8) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v int8) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v int8) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldType, v))
}

// ConfigEQ applies the EQ predicate on the "config" field.
func ConfigEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldConfig, v))
}

// ConfigNEQ applies the NEQ predicate on the "config" field.
func ConfigNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldConfig, v))
}

// ConfigIn applies the In predicate on the "config" field.
func ConfigIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldConfig, vs...))
}

// ConfigNotIn applies the NotIn predicate on the "config" field.
func ConfigNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldConfig, vs...))
}

// ConfigGT applies the GT predicate on the "config" field.
func ConfigGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldConfig, v))
}

// ConfigGTE applies the GTE predicate on the "config" field.
func ConfigGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldConfig, v))
}

// ConfigLT applies the LT predicate on the "config" field.
func ConfigLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldConfig, v))
}

// ConfigLTE applies the LTE predicate on the "config" field.
func ConfigLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldConfig, v))
}

// ConfigContains applies the Contains predicate on the "config" field.
func ConfigContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldConfig, v))
}

// ConfigHasPrefix applies the HasPrefix predicate on the "config" field.
func ConfigHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldConfig, v))
}

// ConfigHasSuffix applies the HasSuffix predicate on the "config" field.
func ConfigHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldConfig, v))
}

// ConfigIsNil applies the IsNil predicate on the "config" field.
func ConfigIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldConfig))
}

// ConfigNotNil applies the NotNil predicate on the "config" field.
func ConfigNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldConfig))
}

// ConfigEqualFold applies the EqualFold predicate on the "config" field.
func ConfigEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldConfig, v))
}

// ConfigContainsFold applies the ContainsFold predicate on the "config" field.
func ConfigContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldConfig, v))
}

// RankEQ applies the EQ predicate on the "rank" field.
func RankEQ(v int8) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldRank, v))
}

// RankNEQ applies the NEQ predicate on the "rank" field.
func RankNEQ(v int8) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldRank, v))
}

// RankIn applies the In predicate on the "rank" field.
func RankIn(vs ...int8) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldRank, vs...))
}

// RankNotIn applies the NotIn predicate on the "rank" field.
func RankNotIn(vs ...int8) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldRank, vs...))
}

// RankGT applies the GT predicate on the "rank" field.
func RankGT(v int8) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldRank, v))
}

// RankGTE applies the GTE predicate on the "rank" field.
func RankGTE(v int8) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldRank, v))
}

// RankLT applies the LT predicate on the "rank" field.
func RankLT(v int8) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldRank, v))
}

// RankLTE applies the LTE predicate on the "rank" field.
func RankLTE(v int8) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldRank, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldStatus, v))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldStartTime, v))
}

// StartTimeIsNil applies the IsNil predicate on the "start_time" field.
func StartTimeIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldStartTime))
}

// StartTimeNotNil applies the NotNil predicate on the "start_time" field.
func StartTimeNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldStartTime))
}

// CompletedAtEQ applies the EQ predicate on the "completed_at" field.
func CompletedAtEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCompletedAt, v))
}

// CompletedAtNEQ applies the NEQ predicate on the "completed_at" field.
func CompletedAtNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCompletedAt, v))
}

// CompletedAtIn applies the In predicate on the "completed_at" field.
func CompletedAtIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCompletedAt, vs...))
}

// CompletedAtNotIn applies the NotIn predicate on the "completed_at" field.
func CompletedAtNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCompletedAt, vs...))
}

// CompletedAtGT applies the GT predicate on the "completed_at" field.
func CompletedAtGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCompletedAt, v))
}

// CompletedAtGTE applies the GTE predicate on the "completed_at" field.
func CompletedAtGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCompletedAt, v))
}

// CompletedAtLT applies the LT predicate on the "completed_at" field.
func CompletedAtLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCompletedAt, v))
}

// CompletedAtLTE applies the LTE predicate on the "completed_at" field.
func CompletedAtLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCompletedAt, v))
}

// CompletedAtIsNil applies the IsNil predicate on the "completed_at" field.
func CompletedAtIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldCompletedAt))
}

// CompletedAtNotNil applies the NotNil predicate on the "completed_at" field.
func CompletedAtNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldCompletedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldUpdatedAt))
}

// StatusUpdatedAtEQ applies the EQ predicate on the "status_updated_at" field.
func StatusUpdatedAtEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldStatusUpdatedAt, v))
}

// StatusUpdatedAtNEQ applies the NEQ predicate on the "status_updated_at" field.
func StatusUpdatedAtNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldStatusUpdatedAt, v))
}

// StatusUpdatedAtIn applies the In predicate on the "status_updated_at" field.
func StatusUpdatedAtIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldStatusUpdatedAt, vs...))
}

// StatusUpdatedAtNotIn applies the NotIn predicate on the "status_updated_at" field.
func StatusUpdatedAtNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldStatusUpdatedAt, vs...))
}

// StatusUpdatedAtGT applies the GT predicate on the "status_updated_at" field.
func StatusUpdatedAtGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldStatusUpdatedAt, v))
}

// StatusUpdatedAtGTE applies the GTE predicate on the "status_updated_at" field.
func StatusUpdatedAtGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldStatusUpdatedAt, v))
}

// StatusUpdatedAtLT applies the LT predicate on the "status_updated_at" field.
func StatusUpdatedAtLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldStatusUpdatedAt, v))
}

// StatusUpdatedAtLTE applies the LTE predicate on the "status_updated_at" field.
func StatusUpdatedAtLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldStatusUpdatedAt, v))
}

// StatusUpdatedAtIsNil applies the IsNil predicate on the "status_updated_at" field.
func StatusUpdatedAtIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldStatusUpdatedAt))
}

// StatusUpdatedAtNotNil applies the NotNil predicate on the "status_updated_at" field.
func StatusUpdatedAtNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldStatusUpdatedAt))
}

// DeadlineEQ applies the EQ predicate on the "deadline" field.
func DeadlineEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDeadline, v))
}

// DeadlineNEQ applies the NEQ predicate on the "deadline" field.
func DeadlineNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldDeadline, v))
}

// DeadlineIn applies the In predicate on the "deadline" field.
func DeadlineIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldDeadline, vs...))
}

// DeadlineNotIn applies the NotIn predicate on the "deadline" field.
func DeadlineNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldDeadline, vs...))
}

// DeadlineGT applies the GT predicate on the "deadline" field.
func DeadlineGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldDeadline, v))
}

// DeadlineGTE applies the GTE predicate on the "deadline" field.
func DeadlineGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldDeadline, v))
}

// DeadlineLT applies the LT predicate on the "deadline" field.
func DeadlineLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldDeadline, v))
}

// DeadlineLTE applies the LTE predicate on the "deadline" field.
func DeadlineLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldDeadline, v))
}

// DeadlineIsNil applies the IsNil predicate on the "deadline" field.
func DeadlineIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldDeadline))
}

// DeadlineNotNil applies the NotNil predicate on the "deadline" field.
func DeadlineNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldDeadline))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v uint32) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v uint32) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...uint32) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...uint32) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v uint32) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v uint32) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v uint32) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v uint32) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldDeletedBy))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldDescription, v))
}

// HasTestcaseSuite applies the HasEdge predicate on the "testcase_suite" edge.
func HasTestcaseSuite() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestcaseSuiteTable, TestcaseSuiteColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestcaseSuiteWith applies the HasEdge predicate on the "testcase_suite" edge with a given conditions (other predicates).
func HasTestcaseSuiteWith(preds ...predicate.TestcaseSuite) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestcaseSuiteInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestcaseSuiteTable, TestcaseSuiteColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
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
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		p(s.Not())
	})
}
