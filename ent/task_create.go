// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"galileo/ent/task"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TaskCreate) SetName(s string) *TaskCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TaskCreate) SetCreatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreatedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetCreatedBy sets the "created_by" field.
func (tc *TaskCreate) SetCreatedBy(u uint32) *TaskCreate {
	tc.mutation.SetCreatedBy(u)
	return tc
}

// SetRank sets the "rank" field.
func (tc *TaskCreate) SetRank(i int8) *TaskCreate {
	tc.mutation.SetRank(i)
	return tc
}

// SetType sets the "type" field.
func (tc *TaskCreate) SetType(i int16) *TaskCreate {
	tc.mutation.SetType(i)
	return tc
}

// SetStatus sets the "status" field.
func (tc *TaskCreate) SetStatus(i int16) *TaskCreate {
	tc.mutation.SetStatus(i)
	return tc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tc *TaskCreate) SetNillableStatus(i *int16) *TaskCreate {
	if i != nil {
		tc.SetStatus(*i)
	}
	return tc
}

// SetCompleteAt sets the "complete_at" field.
func (tc *TaskCreate) SetCompleteAt(t time.Time) *TaskCreate {
	tc.mutation.SetCompleteAt(t)
	return tc
}

// SetNillableCompleteAt sets the "complete_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCompleteAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCompleteAt(*t)
	}
	return tc
}

// SetUpdateAt sets the "update_at" field.
func (tc *TaskCreate) SetUpdateAt(t time.Time) *TaskCreate {
	tc.mutation.SetUpdateAt(t)
	return tc
}

// SetIsDeleted sets the "is_deleted" field.
func (tc *TaskCreate) SetIsDeleted(b bool) *TaskCreate {
	tc.mutation.SetIsDeleted(b)
	return tc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (tc *TaskCreate) SetNillableIsDeleted(b *bool) *TaskCreate {
	if b != nil {
		tc.SetIsDeleted(*b)
	}
	return tc
}

// SetDeletedAt sets the "deleted_at" field.
func (tc *TaskCreate) SetDeletedAt(t time.Time) *TaskCreate {
	tc.mutation.SetDeletedAt(t)
	return tc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableDeletedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetDeletedAt(*t)
	}
	return tc
}

// SetDeletedBy sets the "deleted_by" field.
func (tc *TaskCreate) SetDeletedBy(u uint32) *TaskCreate {
	tc.mutation.SetDeletedBy(u)
	return tc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (tc *TaskCreate) SetNillableDeletedBy(u *uint32) *TaskCreate {
	if u != nil {
		tc.SetDeletedBy(*u)
	}
	return tc
}

// SetDescription sets the "description" field.
func (tc *TaskCreate) SetDescription(s string) *TaskCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TaskCreate) SetNillableDescription(s *string) *TaskCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetURL sets the "url" field.
func (tc *TaskCreate) SetURL(s string) *TaskCreate {
	tc.mutation.SetURL(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TaskCreate) SetID(i int64) *TaskCreate {
	tc.mutation.SetID(i)
	return tc
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	tc.defaults()
	return withHooks[*Task, TaskMutation](ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TaskCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := task.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.Status(); !ok {
		v := task.DefaultStatus
		tc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Task.name"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Task.created_at"`)}
	}
	if _, ok := tc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Task.created_by"`)}
	}
	if _, ok := tc.mutation.Rank(); !ok {
		return &ValidationError{Name: "rank", err: errors.New(`ent: missing required field "Task.rank"`)}
	}
	if _, ok := tc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Task.type"`)}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Task.status"`)}
	}
	if _, ok := tc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "Task.update_at"`)}
	}
	if _, ok := tc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Task.url"`)}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
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
		_node.ID = int64(id)
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(task.Table, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt64))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(task.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(task.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.CreatedBy(); ok {
		_spec.SetField(task.FieldCreatedBy, field.TypeUint32, value)
		_node.CreatedBy = value
	}
	if value, ok := tc.mutation.Rank(); ok {
		_spec.SetField(task.FieldRank, field.TypeInt8, value)
		_node.Rank = value
	}
	if value, ok := tc.mutation.GetType(); ok {
		_spec.SetField(task.FieldType, field.TypeInt16, value)
		_node.Type = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeInt16, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.CompleteAt(); ok {
		_spec.SetField(task.FieldCompleteAt, field.TypeTime, value)
		_node.CompleteAt = &value
	}
	if value, ok := tc.mutation.UpdateAt(); ok {
		_spec.SetField(task.FieldUpdateAt, field.TypeTime, value)
		_node.UpdateAt = value
	}
	if value, ok := tc.mutation.IsDeleted(); ok {
		_spec.SetField(task.FieldIsDeleted, field.TypeBool, value)
		_node.IsDeleted = &value
	}
	if value, ok := tc.mutation.DeletedAt(); ok {
		_spec.SetField(task.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := tc.mutation.DeletedBy(); ok {
		_spec.SetField(task.FieldDeletedBy, field.TypeUint32, value)
		_node.DeletedBy = &value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(task.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.URL(); ok {
		_spec.SetField(task.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	return _node, _spec
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	builders []*TaskCreate
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
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
					nodes[i].ID = int64(id)
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
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
