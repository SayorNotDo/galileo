// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"galileo/ent/apicategory"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApiCategoryCreate is the builder for creating a ApiCategory entity.
type ApiCategoryCreate struct {
	config
	mutation *ApiCategoryMutation
	hooks    []Hook
}

// Mutation returns the ApiCategoryMutation object of the builder.
func (acc *ApiCategoryCreate) Mutation() *ApiCategoryMutation {
	return acc.mutation
}

// Save creates the ApiCategory in the database.
func (acc *ApiCategoryCreate) Save(ctx context.Context) (*ApiCategory, error) {
	return withHooks[*ApiCategory, ApiCategoryMutation](ctx, acc.sqlSave, acc.mutation, acc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (acc *ApiCategoryCreate) SaveX(ctx context.Context) *ApiCategory {
	v, err := acc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acc *ApiCategoryCreate) Exec(ctx context.Context) error {
	_, err := acc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acc *ApiCategoryCreate) ExecX(ctx context.Context) {
	if err := acc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acc *ApiCategoryCreate) check() error {
	return nil
}

func (acc *ApiCategoryCreate) sqlSave(ctx context.Context) (*ApiCategory, error) {
	if err := acc.check(); err != nil {
		return nil, err
	}
	_node, _spec := acc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	acc.mutation.id = &_node.ID
	acc.mutation.done = true
	return _node, nil
}

func (acc *ApiCategoryCreate) createSpec() (*ApiCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &ApiCategory{config: acc.config}
		_spec = sqlgraph.NewCreateSpec(apicategory.Table, sqlgraph.NewFieldSpec(apicategory.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// ApiCategoryCreateBulk is the builder for creating many ApiCategory entities in bulk.
type ApiCategoryCreateBulk struct {
	config
	builders []*ApiCategoryCreate
}

// Save creates the ApiCategory entities in the database.
func (accb *ApiCategoryCreateBulk) Save(ctx context.Context) ([]*ApiCategory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(accb.builders))
	nodes := make([]*ApiCategory, len(accb.builders))
	mutators := make([]Mutator, len(accb.builders))
	for i := range accb.builders {
		func(i int, root context.Context) {
			builder := accb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApiCategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, accb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, accb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, accb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (accb *ApiCategoryCreateBulk) SaveX(ctx context.Context) []*ApiCategory {
	v, err := accb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (accb *ApiCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := accb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (accb *ApiCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := accb.Exec(ctx); err != nil {
		panic(err)
	}
}
