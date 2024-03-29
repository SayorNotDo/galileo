// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"galileo/ent/predicate"
	"galileo/ent/testplan"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestPlanDelete is the builder for deleting a TestPlan entity.
type TestPlanDelete struct {
	config
	hooks    []Hook
	mutation *TestPlanMutation
}

// Where appends a list predicates to the TestPlanDelete builder.
func (tpd *TestPlanDelete) Where(ps ...predicate.TestPlan) *TestPlanDelete {
	tpd.mutation.Where(ps...)
	return tpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tpd *TestPlanDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, TestPlanMutation](ctx, tpd.sqlExec, tpd.mutation, tpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tpd *TestPlanDelete) ExecX(ctx context.Context) int {
	n, err := tpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tpd *TestPlanDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(testplan.Table, sqlgraph.NewFieldSpec(testplan.FieldID, field.TypeInt64))
	if ps := tpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tpd.mutation.done = true
	return affected, err
}

// TestPlanDeleteOne is the builder for deleting a single TestPlan entity.
type TestPlanDeleteOne struct {
	tpd *TestPlanDelete
}

// Where appends a list predicates to the TestPlanDelete builder.
func (tpdo *TestPlanDeleteOne) Where(ps ...predicate.TestPlan) *TestPlanDeleteOne {
	tpdo.tpd.mutation.Where(ps...)
	return tpdo
}

// Exec executes the deletion query.
func (tpdo *TestPlanDeleteOne) Exec(ctx context.Context) error {
	n, err := tpdo.tpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{testplan.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tpdo *TestPlanDeleteOne) ExecX(ctx context.Context) {
	if err := tpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
