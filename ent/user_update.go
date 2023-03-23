// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"galileo/ent/predicate"
	"galileo/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetActive sets the "active" field.
func (uu *UserUpdate) SetActive(b bool) *UserUpdate {
	uu.mutation.SetActive(b)
	return uu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (uu *UserUpdate) SetNillableActive(b *bool) *UserUpdate {
	if b != nil {
		uu.SetActive(*b)
	}
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetChineseName sets the "chineseName" field.
func (uu *UserUpdate) SetChineseName(s string) *UserUpdate {
	uu.mutation.SetChineseName(s)
	return uu
}

// SetNillableChineseName sets the "chineseName" field if the given value is not nil.
func (uu *UserUpdate) SetNillableChineseName(s *string) *UserUpdate {
	if s != nil {
		uu.SetChineseName(*s)
	}
	return uu
}

// ClearChineseName clears the value of the "chineseName" field.
func (uu *UserUpdate) ClearChineseName() *UserUpdate {
	uu.mutation.ClearChineseName()
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetPhone sets the "phone" field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatar(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatar(*s)
	}
	return uu
}

// ClearAvatar clears the value of the "avatar" field.
func (uu *UserUpdate) ClearAvatar() *UserUpdate {
	uu.mutation.ClearAvatar()
	return uu
}

// SetRole sets the "role" field.
func (uu *UserUpdate) SetRole(u uint8) *UserUpdate {
	uu.mutation.ResetRole()
	uu.mutation.SetRole(u)
	return uu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRole(u *uint8) *UserUpdate {
	if u != nil {
		uu.SetRole(*u)
	}
	return uu
}

// AddRole adds u to the "role" field.
func (uu *UserUpdate) AddRole(u int8) *UserUpdate {
	uu.mutation.AddRole(u)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetDeletedAt sets the "deleted_at" field.
func (uu *UserUpdate) SetDeletedAt(t time.Time) *UserUpdate {
	uu.mutation.SetDeletedAt(t)
	return uu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDeletedAt(*t)
	}
	return uu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uu *UserUpdate) ClearDeletedAt() *UserUpdate {
	uu.mutation.ClearDeletedAt()
	return uu
}

// SetDeletedBy sets the "deleted_by" field.
func (uu *UserUpdate) SetDeletedBy(u uint32) *UserUpdate {
	uu.mutation.ResetDeletedBy()
	uu.mutation.SetDeletedBy(u)
	return uu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedBy(u *uint32) *UserUpdate {
	if u != nil {
		uu.SetDeletedBy(*u)
	}
	return uu
}

// AddDeletedBy adds u to the "deleted_by" field.
func (uu *UserUpdate) AddDeletedBy(u int32) *UserUpdate {
	uu.mutation.AddDeletedBy(u)
	return uu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (uu *UserUpdate) ClearDeletedBy() *UserUpdate {
	uu.mutation.ClearDeletedBy()
	return uu
}

// SetIsDeleted sets the "is_deleted" field.
func (uu *UserUpdate) SetIsDeleted(b bool) *UserUpdate {
	uu.mutation.SetIsDeleted(b)
	return uu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsDeleted(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsDeleted(*b)
	}
	return uu
}

// ClearIsDeleted clears the value of the "is_deleted" field.
func (uu *UserUpdate) ClearIsDeleted() *UserUpdate {
	uu.mutation.ClearIsDeleted()
	return uu
}

// SetUUID sets the "uuid" field.
func (uu *UserUpdate) SetUUID(u uuid.UUID) *UserUpdate {
	uu.mutation.SetUUID(u)
	return uu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUUID(u *uuid.UUID) *UserUpdate {
	if u != nil {
		uu.SetUUID(*u)
	}
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Active(); ok {
		_spec.SetField(user.FieldActive, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.ChineseName(); ok {
		_spec.SetField(user.FieldChineseName, field.TypeString, value)
	}
	if uu.mutation.ChineseNameCleared() {
		_spec.ClearField(user.FieldChineseName, field.TypeString)
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uu.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeUint8, value)
	}
	if value, ok := uu.mutation.AddedRole(); ok {
		_spec.AddField(user.FieldRole, field.TypeUint8, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uu.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uu.mutation.DeletedBy(); ok {
		_spec.SetField(user.FieldDeletedBy, field.TypeUint32, value)
	}
	if value, ok := uu.mutation.AddedDeletedBy(); ok {
		_spec.AddField(user.FieldDeletedBy, field.TypeUint32, value)
	}
	if uu.mutation.DeletedByCleared() {
		_spec.ClearField(user.FieldDeletedBy, field.TypeUint32)
	}
	if value, ok := uu.mutation.IsDeleted(); ok {
		_spec.SetField(user.FieldIsDeleted, field.TypeBool, value)
	}
	if uu.mutation.IsDeletedCleared() {
		_spec.ClearField(user.FieldIsDeleted, field.TypeBool)
	}
	if value, ok := uu.mutation.UUID(); ok {
		_spec.SetField(user.FieldUUID, field.TypeUUID, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetActive sets the "active" field.
func (uuo *UserUpdateOne) SetActive(b bool) *UserUpdateOne {
	uuo.mutation.SetActive(b)
	return uuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableActive(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetActive(*b)
	}
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetChineseName sets the "chineseName" field.
func (uuo *UserUpdateOne) SetChineseName(s string) *UserUpdateOne {
	uuo.mutation.SetChineseName(s)
	return uuo
}

// SetNillableChineseName sets the "chineseName" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableChineseName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetChineseName(*s)
	}
	return uuo
}

// ClearChineseName clears the value of the "chineseName" field.
func (uuo *UserUpdateOne) ClearChineseName() *UserUpdateOne {
	uuo.mutation.ClearChineseName()
	return uuo
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetPhone sets the "phone" field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatar(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatar(*s)
	}
	return uuo
}

// ClearAvatar clears the value of the "avatar" field.
func (uuo *UserUpdateOne) ClearAvatar() *UserUpdateOne {
	uuo.mutation.ClearAvatar()
	return uuo
}

// SetRole sets the "role" field.
func (uuo *UserUpdateOne) SetRole(u uint8) *UserUpdateOne {
	uuo.mutation.ResetRole()
	uuo.mutation.SetRole(u)
	return uuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRole(u *uint8) *UserUpdateOne {
	if u != nil {
		uuo.SetRole(*u)
	}
	return uuo
}

// AddRole adds u to the "role" field.
func (uuo *UserUpdateOne) AddRole(u int8) *UserUpdateOne {
	uuo.mutation.AddRole(u)
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetDeletedAt sets the "deleted_at" field.
func (uuo *UserUpdateOne) SetDeletedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDeletedAt(t)
	return uuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDeletedAt(*t)
	}
	return uuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uuo *UserUpdateOne) ClearDeletedAt() *UserUpdateOne {
	uuo.mutation.ClearDeletedAt()
	return uuo
}

// SetDeletedBy sets the "deleted_by" field.
func (uuo *UserUpdateOne) SetDeletedBy(u uint32) *UserUpdateOne {
	uuo.mutation.ResetDeletedBy()
	uuo.mutation.SetDeletedBy(u)
	return uuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedBy(u *uint32) *UserUpdateOne {
	if u != nil {
		uuo.SetDeletedBy(*u)
	}
	return uuo
}

// AddDeletedBy adds u to the "deleted_by" field.
func (uuo *UserUpdateOne) AddDeletedBy(u int32) *UserUpdateOne {
	uuo.mutation.AddDeletedBy(u)
	return uuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (uuo *UserUpdateOne) ClearDeletedBy() *UserUpdateOne {
	uuo.mutation.ClearDeletedBy()
	return uuo
}

// SetIsDeleted sets the "is_deleted" field.
func (uuo *UserUpdateOne) SetIsDeleted(b bool) *UserUpdateOne {
	uuo.mutation.SetIsDeleted(b)
	return uuo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsDeleted(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsDeleted(*b)
	}
	return uuo
}

// ClearIsDeleted clears the value of the "is_deleted" field.
func (uuo *UserUpdateOne) ClearIsDeleted() *UserUpdateOne {
	uuo.mutation.ClearIsDeleted()
	return uuo
}

// SetUUID sets the "uuid" field.
func (uuo *UserUpdateOne) SetUUID(u uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetUUID(u)
	return uuo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUUID(u *uuid.UUID) *UserUpdateOne {
	if u != nil {
		uuo.SetUUID(*u)
	}
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Active(); ok {
		_spec.SetField(user.FieldActive, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.ChineseName(); ok {
		_spec.SetField(user.FieldChineseName, field.TypeString, value)
	}
	if uuo.mutation.ChineseNameCleared() {
		_spec.ClearField(user.FieldChineseName, field.TypeString)
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uuo.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeUint8, value)
	}
	if value, ok := uuo.mutation.AddedRole(); ok {
		_spec.AddField(user.FieldRole, field.TypeUint8, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uuo.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.DeletedBy(); ok {
		_spec.SetField(user.FieldDeletedBy, field.TypeUint32, value)
	}
	if value, ok := uuo.mutation.AddedDeletedBy(); ok {
		_spec.AddField(user.FieldDeletedBy, field.TypeUint32, value)
	}
	if uuo.mutation.DeletedByCleared() {
		_spec.ClearField(user.FieldDeletedBy, field.TypeUint32)
	}
	if value, ok := uuo.mutation.IsDeleted(); ok {
		_spec.SetField(user.FieldIsDeleted, field.TypeBool, value)
	}
	if uuo.mutation.IsDeletedCleared() {
		_spec.ClearField(user.FieldIsDeleted, field.TypeBool)
	}
	if value, ok := uuo.mutation.UUID(); ok {
		_spec.SetField(user.FieldUUID, field.TypeUUID, value)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
