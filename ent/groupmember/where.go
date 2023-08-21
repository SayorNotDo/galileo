// Code generated by ent, DO NOT EDIT.

package groupmember

import (
	"galileo/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLTE(FieldID, id))
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldGroupID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldUserID, v))
}

// Role applies equality check predicate on the "role" field. It's identical to RoleEQ.
func Role(v uint8) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldRole, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldCreatedBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldDeletedBy, v))
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldGroupID, v))
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNEQ(FieldGroupID, v))
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldIn(FieldGroupID, vs...))
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNotIn(FieldGroupID, vs...))
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGT(FieldGroupID, v))
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGTE(FieldGroupID, v))
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLT(FieldGroupID, v))
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLTE(FieldGroupID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLTE(FieldUserID, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v uint8) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v uint8) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...uint8) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...uint8) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNotIn(FieldRole, vs...))
}

// RoleGT applies the GT predicate on the "role" field.
func RoleGT(v uint8) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGT(FieldRole, v))
}

// RoleGTE applies the GTE predicate on the "role" field.
func RoleGTE(v uint8) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGTE(FieldRole, v))
}

// RoleLT applies the LT predicate on the "role" field.
func RoleLT(v uint8) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLT(FieldRole, v))
}

// RoleLTE applies the LTE predicate on the "role" field.
func RoleLTE(v uint8) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLTE(FieldRole, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLTE(FieldCreatedBy, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.GroupMember {
	return predicate.GroupMember(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v uint32) predicate.GroupMember {
	return predicate.GroupMember(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.GroupMember {
	return predicate.GroupMember(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.GroupMember {
	return predicate.GroupMember(sql.FieldNotNull(FieldDeletedBy))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GroupMember) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GroupMember) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
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
func Not(p predicate.GroupMember) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		p(s.Not())
	})
}
