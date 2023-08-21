// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"galileo/ent/testcase"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestcaseCreate is the builder for creating a Testcase entity.
type TestcaseCreate struct {
	config
	mutation *TestcaseMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TestcaseCreate) SetName(s string) *TestcaseCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetCreatedBy sets the "created_by" field.
func (tc *TestcaseCreate) SetCreatedBy(u uint32) *TestcaseCreate {
	tc.mutation.SetCreatedBy(u)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TestcaseCreate) SetCreatedAt(t time.Time) *TestcaseCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TestcaseCreate) SetNillableCreatedAt(t *time.Time) *TestcaseCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedBy sets the "updated_by" field.
func (tc *TestcaseCreate) SetUpdatedBy(u uint32) *TestcaseCreate {
	tc.mutation.SetUpdatedBy(u)
	return tc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tc *TestcaseCreate) SetNillableUpdatedBy(u *uint32) *TestcaseCreate {
	if u != nil {
		tc.SetUpdatedBy(*u)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TestcaseCreate) SetUpdatedAt(t time.Time) *TestcaseCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TestcaseCreate) SetNillableUpdatedAt(t *time.Time) *TestcaseCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetStatus sets the "status" field.
func (tc *TestcaseCreate) SetStatus(i int8) *TestcaseCreate {
	tc.mutation.SetStatus(i)
	return tc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tc *TestcaseCreate) SetNillableStatus(i *int8) *TestcaseCreate {
	if i != nil {
		tc.SetStatus(*i)
	}
	return tc
}

// SetType sets the "type" field.
func (tc *TestcaseCreate) SetType(i int8) *TestcaseCreate {
	tc.mutation.SetType(i)
	return tc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tc *TestcaseCreate) SetNillableType(i *int8) *TestcaseCreate {
	if i != nil {
		tc.SetType(*i)
	}
	return tc
}

// SetPriority sets the "priority" field.
func (tc *TestcaseCreate) SetPriority(i int8) *TestcaseCreate {
	tc.mutation.SetPriority(i)
	return tc
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tc *TestcaseCreate) SetNillablePriority(i *int8) *TestcaseCreate {
	if i != nil {
		tc.SetPriority(*i)
	}
	return tc
}

// SetDeletedAt sets the "deleted_at" field.
func (tc *TestcaseCreate) SetDeletedAt(t time.Time) *TestcaseCreate {
	tc.mutation.SetDeletedAt(t)
	return tc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tc *TestcaseCreate) SetNillableDeletedAt(t *time.Time) *TestcaseCreate {
	if t != nil {
		tc.SetDeletedAt(*t)
	}
	return tc
}

// SetDeletedBy sets the "deleted_by" field.
func (tc *TestcaseCreate) SetDeletedBy(u uint32) *TestcaseCreate {
	tc.mutation.SetDeletedBy(u)
	return tc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (tc *TestcaseCreate) SetNillableDeletedBy(u *uint32) *TestcaseCreate {
	if u != nil {
		tc.SetDeletedBy(*u)
	}
	return tc
}

// SetDescription sets the "description" field.
func (tc *TestcaseCreate) SetDescription(s string) *TestcaseCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TestcaseCreate) SetNillableDescription(s *string) *TestcaseCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetLabel sets the "label" field.
func (tc *TestcaseCreate) SetLabel(s string) *TestcaseCreate {
	tc.mutation.SetLabel(s)
	return tc
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (tc *TestcaseCreate) SetNillableLabel(s *string) *TestcaseCreate {
	if s != nil {
		tc.SetLabel(*s)
	}
	return tc
}

// SetURL sets the "url" field.
func (tc *TestcaseCreate) SetURL(s string) *TestcaseCreate {
	tc.mutation.SetURL(s)
	return tc
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (tc *TestcaseCreate) SetNillableURL(s *string) *TestcaseCreate {
	if s != nil {
		tc.SetURL(*s)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TestcaseCreate) SetID(i int32) *TestcaseCreate {
	tc.mutation.SetID(i)
	return tc
}

// Mutation returns the TestcaseMutation object of the builder.
func (tc *TestcaseCreate) Mutation() *TestcaseMutation {
	return tc.mutation
}

// Save creates the Testcase in the database.
func (tc *TestcaseCreate) Save(ctx context.Context) (*Testcase, error) {
	tc.defaults()
	return withHooks[*Testcase, TestcaseMutation](ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TestcaseCreate) SaveX(ctx context.Context) *Testcase {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TestcaseCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TestcaseCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TestcaseCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := testcase.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.Status(); !ok {
		v := testcase.DefaultStatus
		tc.mutation.SetStatus(v)
	}
	if _, ok := tc.mutation.GetType(); !ok {
		v := testcase.DefaultType
		tc.mutation.SetType(v)
	}
	if _, ok := tc.mutation.Priority(); !ok {
		v := testcase.DefaultPriority
		tc.mutation.SetPriority(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TestcaseCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Testcase.name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := testcase.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Testcase.name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Testcase.created_by"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Testcase.created_at"`)}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Testcase.status"`)}
	}
	if _, ok := tc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Testcase.type"`)}
	}
	if _, ok := tc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "Testcase.priority"`)}
	}
	return nil
}

func (tc *TestcaseCreate) sqlSave(ctx context.Context) (*Testcase, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TestcaseCreate) createSpec() (*Testcase, *sqlgraph.CreateSpec) {
	var (
		_node = &Testcase{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(testcase.Table, sqlgraph.NewFieldSpec(testcase.FieldID, field.TypeInt32))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(testcase.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.CreatedBy(); ok {
		_spec.SetField(testcase.FieldCreatedBy, field.TypeUint32, value)
		_node.CreatedBy = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(testcase.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedBy(); ok {
		_spec.SetField(testcase.FieldUpdatedBy, field.TypeUint32, value)
		_node.UpdatedBy = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(testcase.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(testcase.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.GetType(); ok {
		_spec.SetField(testcase.FieldType, field.TypeInt8, value)
		_node.Type = value
	}
	if value, ok := tc.mutation.Priority(); ok {
		_spec.SetField(testcase.FieldPriority, field.TypeInt8, value)
		_node.Priority = value
	}
	if value, ok := tc.mutation.DeletedAt(); ok {
		_spec.SetField(testcase.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := tc.mutation.DeletedBy(); ok {
		_spec.SetField(testcase.FieldDeletedBy, field.TypeUint32, value)
		_node.DeletedBy = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(testcase.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.Label(); ok {
		_spec.SetField(testcase.FieldLabel, field.TypeString, value)
		_node.Label = value
	}
	if value, ok := tc.mutation.URL(); ok {
		_spec.SetField(testcase.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	return _node, _spec
}

// TestcaseCreateBulk is the builder for creating many Testcase entities in bulk.
type TestcaseCreateBulk struct {
	config
	builders []*TestcaseCreate
}

// Save creates the Testcase entities in the database.
func (tcb *TestcaseCreateBulk) Save(ctx context.Context) ([]*Testcase, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Testcase, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TestcaseMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TestcaseCreateBulk) SaveX(ctx context.Context) []*Testcase {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TestcaseCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TestcaseCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
