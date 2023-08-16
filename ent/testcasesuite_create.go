// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"galileo/ent/testcasesuite"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestcaseSuiteCreate is the builder for creating a TestcaseSuite entity.
type TestcaseSuiteCreate struct {
	config
	mutation *TestcaseSuiteMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tsc *TestcaseSuiteCreate) SetName(s string) *TestcaseSuiteCreate {
	tsc.mutation.SetName(s)
	return tsc
}

// SetCreatedAt sets the "created_at" field.
func (tsc *TestcaseSuiteCreate) SetCreatedAt(t time.Time) *TestcaseSuiteCreate {
	tsc.mutation.SetCreatedAt(t)
	return tsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tsc *TestcaseSuiteCreate) SetNillableCreatedAt(t *time.Time) *TestcaseSuiteCreate {
	if t != nil {
		tsc.SetCreatedAt(*t)
	}
	return tsc
}

// SetCreatedBy sets the "created_by" field.
func (tsc *TestcaseSuiteCreate) SetCreatedBy(u uint32) *TestcaseSuiteCreate {
	tsc.mutation.SetCreatedBy(u)
	return tsc
}

// SetUpdatedAt sets the "updated_at" field.
func (tsc *TestcaseSuiteCreate) SetUpdatedAt(t time.Time) *TestcaseSuiteCreate {
	tsc.mutation.SetUpdatedAt(t)
	return tsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tsc *TestcaseSuiteCreate) SetNillableUpdatedAt(t *time.Time) *TestcaseSuiteCreate {
	if t != nil {
		tsc.SetUpdatedAt(*t)
	}
	return tsc
}

// SetUpdatedBy sets the "updated_by" field.
func (tsc *TestcaseSuiteCreate) SetUpdatedBy(u uint32) *TestcaseSuiteCreate {
	tsc.mutation.SetUpdatedBy(u)
	return tsc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tsc *TestcaseSuiteCreate) SetNillableUpdatedBy(u *uint32) *TestcaseSuiteCreate {
	if u != nil {
		tsc.SetUpdatedBy(*u)
	}
	return tsc
}

// SetTestcases sets the "testcases" field.
func (tsc *TestcaseSuiteCreate) SetTestcases(i []int64) *TestcaseSuiteCreate {
	tsc.mutation.SetTestcases(i)
	return tsc
}

// SetID sets the "id" field.
func (tsc *TestcaseSuiteCreate) SetID(i int64) *TestcaseSuiteCreate {
	tsc.mutation.SetID(i)
	return tsc
}

// Mutation returns the TestcaseSuiteMutation object of the builder.
func (tsc *TestcaseSuiteCreate) Mutation() *TestcaseSuiteMutation {
	return tsc.mutation
}

// Save creates the TestcaseSuite in the database.
func (tsc *TestcaseSuiteCreate) Save(ctx context.Context) (*TestcaseSuite, error) {
	tsc.defaults()
	return withHooks[*TestcaseSuite, TestcaseSuiteMutation](ctx, tsc.sqlSave, tsc.mutation, tsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tsc *TestcaseSuiteCreate) SaveX(ctx context.Context) *TestcaseSuite {
	v, err := tsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tsc *TestcaseSuiteCreate) Exec(ctx context.Context) error {
	_, err := tsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsc *TestcaseSuiteCreate) ExecX(ctx context.Context) {
	if err := tsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tsc *TestcaseSuiteCreate) defaults() {
	if _, ok := tsc.mutation.CreatedAt(); !ok {
		v := testcasesuite.DefaultCreatedAt()
		tsc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tsc *TestcaseSuiteCreate) check() error {
	if _, ok := tsc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "TestcaseSuite.name"`)}
	}
	if v, ok := tsc.mutation.Name(); ok {
		if err := testcasesuite.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TestcaseSuite.name": %w`, err)}
		}
	}
	if _, ok := tsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TestcaseSuite.created_at"`)}
	}
	if _, ok := tsc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "TestcaseSuite.created_by"`)}
	}
	return nil
}

func (tsc *TestcaseSuiteCreate) sqlSave(ctx context.Context) (*TestcaseSuite, error) {
	if err := tsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	tsc.mutation.id = &_node.ID
	tsc.mutation.done = true
	return _node, nil
}

func (tsc *TestcaseSuiteCreate) createSpec() (*TestcaseSuite, *sqlgraph.CreateSpec) {
	var (
		_node = &TestcaseSuite{config: tsc.config}
		_spec = sqlgraph.NewCreateSpec(testcasesuite.Table, sqlgraph.NewFieldSpec(testcasesuite.FieldID, field.TypeInt64))
	)
	if id, ok := tsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tsc.mutation.Name(); ok {
		_spec.SetField(testcasesuite.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tsc.mutation.CreatedAt(); ok {
		_spec.SetField(testcasesuite.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tsc.mutation.CreatedBy(); ok {
		_spec.SetField(testcasesuite.FieldCreatedBy, field.TypeUint32, value)
		_node.CreatedBy = value
	}
	if value, ok := tsc.mutation.UpdatedAt(); ok {
		_spec.SetField(testcasesuite.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tsc.mutation.UpdatedBy(); ok {
		_spec.SetField(testcasesuite.FieldUpdatedBy, field.TypeUint32, value)
		_node.UpdatedBy = value
	}
	if value, ok := tsc.mutation.Testcases(); ok {
		_spec.SetField(testcasesuite.FieldTestcases, field.TypeJSON, value)
		_node.Testcases = value
	}
	return _node, _spec
}

// TestcaseSuiteCreateBulk is the builder for creating many TestcaseSuite entities in bulk.
type TestcaseSuiteCreateBulk struct {
	config
	builders []*TestcaseSuiteCreate
}

// Save creates the TestcaseSuite entities in the database.
func (tscb *TestcaseSuiteCreateBulk) Save(ctx context.Context) ([]*TestcaseSuite, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tscb.builders))
	nodes := make([]*TestcaseSuite, len(tscb.builders))
	mutators := make([]Mutator, len(tscb.builders))
	for i := range tscb.builders {
		func(i int, root context.Context) {
			builder := tscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TestcaseSuiteMutation)
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
					_, err = mutators[i+1].Mutate(root, tscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tscb *TestcaseSuiteCreateBulk) SaveX(ctx context.Context) []*TestcaseSuite {
	v, err := tscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tscb *TestcaseSuiteCreateBulk) Exec(ctx context.Context) error {
	_, err := tscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tscb *TestcaseSuiteCreateBulk) ExecX(ctx context.Context) {
	if err := tscb.Exec(ctx); err != nil {
		panic(err)
	}
}
