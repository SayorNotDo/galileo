// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"galileo/ent/metaevent"
	"galileo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MetaEventDelete is the builder for deleting a MetaEvent entity.
type MetaEventDelete struct {
	config
	hooks    []Hook
	mutation *MetaEventMutation
}

// Where appends a list predicates to the MetaEventDelete builder.
func (med *MetaEventDelete) Where(ps ...predicate.MetaEvent) *MetaEventDelete {
	med.mutation.Where(ps...)
	return med
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (med *MetaEventDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, MetaEventMutation](ctx, med.sqlExec, med.mutation, med.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (med *MetaEventDelete) ExecX(ctx context.Context) int {
	n, err := med.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (med *MetaEventDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(metaevent.Table, sqlgraph.NewFieldSpec(metaevent.FieldID, field.TypeInt64))
	if ps := med.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, med.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	med.mutation.done = true
	return affected, err
}

// MetaEventDeleteOne is the builder for deleting a single MetaEvent entity.
type MetaEventDeleteOne struct {
	med *MetaEventDelete
}

// Where appends a list predicates to the MetaEventDelete builder.
func (medo *MetaEventDeleteOne) Where(ps ...predicate.MetaEvent) *MetaEventDeleteOne {
	medo.med.mutation.Where(ps...)
	return medo
}

// Exec executes the deletion query.
func (medo *MetaEventDeleteOne) Exec(ctx context.Context) error {
	n, err := medo.med.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{metaevent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (medo *MetaEventDeleteOne) ExecX(ctx context.Context) {
	if err := medo.Exec(ctx); err != nil {
		panic(err)
	}
}