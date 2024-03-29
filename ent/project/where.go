// Code generated by ent, DO NOT EDIT.

package project

import (
	"galileo/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldName, v))
}

// Identifier applies equality check predicate on the "identifier" field. It's identical to IdentifierEQ.
func Identifier(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldIdentifier, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v uint32) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v uint32) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v uint32) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDeletedBy, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldStatus, v))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldStartTime, v))
}

// Deadline applies equality check predicate on the "deadline" field. It's identical to DeadlineEQ.
func Deadline(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDeadline, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDescription, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldRemark, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldName, v))
}

// IdentifierEQ applies the EQ predicate on the "identifier" field.
func IdentifierEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldIdentifier, v))
}

// IdentifierNEQ applies the NEQ predicate on the "identifier" field.
func IdentifierNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldIdentifier, v))
}

// IdentifierIn applies the In predicate on the "identifier" field.
func IdentifierIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldIdentifier, vs...))
}

// IdentifierNotIn applies the NotIn predicate on the "identifier" field.
func IdentifierNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldIdentifier, vs...))
}

// IdentifierGT applies the GT predicate on the "identifier" field.
func IdentifierGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldIdentifier, v))
}

// IdentifierGTE applies the GTE predicate on the "identifier" field.
func IdentifierGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldIdentifier, v))
}

// IdentifierLT applies the LT predicate on the "identifier" field.
func IdentifierLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldIdentifier, v))
}

// IdentifierLTE applies the LTE predicate on the "identifier" field.
func IdentifierLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldIdentifier, v))
}

// IdentifierContains applies the Contains predicate on the "identifier" field.
func IdentifierContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldIdentifier, v))
}

// IdentifierHasPrefix applies the HasPrefix predicate on the "identifier" field.
func IdentifierHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldIdentifier, v))
}

// IdentifierHasSuffix applies the HasSuffix predicate on the "identifier" field.
func IdentifierHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldIdentifier, v))
}

// IdentifierEqualFold applies the EqualFold predicate on the "identifier" field.
func IdentifierEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldIdentifier, v))
}

// IdentifierContainsFold applies the ContainsFold predicate on the "identifier" field.
func IdentifierContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldIdentifier, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v uint32) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v uint32) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...uint32) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...uint32) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v uint32) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v uint32) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v uint32) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v uint32) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldUpdatedAt))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v uint32) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v uint32) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...uint32) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...uint32) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v uint32) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v uint32) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v uint32) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v uint32) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldUpdatedBy))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v uint32) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v uint32) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...uint32) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...uint32) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v uint32) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v uint32) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v uint32) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v uint32) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldDeletedBy))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldStatus, v))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldStartTime, v))
}

// StartTimeIsNil applies the IsNil predicate on the "start_time" field.
func StartTimeIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldStartTime))
}

// StartTimeNotNil applies the NotNil predicate on the "start_time" field.
func StartTimeNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldStartTime))
}

// DeadlineEQ applies the EQ predicate on the "deadline" field.
func DeadlineEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDeadline, v))
}

// DeadlineNEQ applies the NEQ predicate on the "deadline" field.
func DeadlineNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldDeadline, v))
}

// DeadlineIn applies the In predicate on the "deadline" field.
func DeadlineIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldDeadline, vs...))
}

// DeadlineNotIn applies the NotIn predicate on the "deadline" field.
func DeadlineNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldDeadline, vs...))
}

// DeadlineGT applies the GT predicate on the "deadline" field.
func DeadlineGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldDeadline, v))
}

// DeadlineGTE applies the GTE predicate on the "deadline" field.
func DeadlineGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldDeadline, v))
}

// DeadlineLT applies the LT predicate on the "deadline" field.
func DeadlineLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldDeadline, v))
}

// DeadlineLTE applies the LTE predicate on the "deadline" field.
func DeadlineLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldDeadline, v))
}

// DeadlineIsNil applies the IsNil predicate on the "deadline" field.
func DeadlineIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldDeadline))
}

// DeadlineNotNil applies the NotNil predicate on the "deadline" field.
func DeadlineNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldDeadline))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldDescription, v))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldRemark, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
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
func Not(p predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		p(s.Not())
	})
}
