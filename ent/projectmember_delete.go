// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"galileo/ent/predicate"
	"galileo/ent/projectmember"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectMemberDelete is the builder for deleting a ProjectMember entity.
type ProjectMemberDelete struct {
	config
	hooks    []Hook
	mutation *ProjectMemberMutation
}

// Where appends a list predicates to the ProjectMemberDelete builder.
func (pmd *ProjectMemberDelete) Where(ps ...predicate.ProjectMember) *ProjectMemberDelete {
	pmd.mutation.Where(ps...)
	return pmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pmd *ProjectMemberDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ProjectMemberMutation](ctx, pmd.sqlExec, pmd.mutation, pmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pmd *ProjectMemberDelete) ExecX(ctx context.Context) int {
	n, err := pmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pmd *ProjectMemberDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(projectmember.Table, sqlgraph.NewFieldSpec(projectmember.FieldID, field.TypeInt))
	if ps := pmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pmd.mutation.done = true
	return affected, err
}

// ProjectMemberDeleteOne is the builder for deleting a single ProjectMember entity.
type ProjectMemberDeleteOne struct {
	pmd *ProjectMemberDelete
}

// Where appends a list predicates to the ProjectMemberDelete builder.
func (pmdo *ProjectMemberDeleteOne) Where(ps ...predicate.ProjectMember) *ProjectMemberDeleteOne {
	pmdo.pmd.mutation.Where(ps...)
	return pmdo
}

// Exec executes the deletion query.
func (pmdo *ProjectMemberDeleteOne) Exec(ctx context.Context) error {
	n, err := pmdo.pmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{projectmember.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pmdo *ProjectMemberDeleteOne) ExecX(ctx context.Context) {
	if err := pmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
