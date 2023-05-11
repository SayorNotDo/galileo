// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"galileo/ent/apihistory"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApiHistoryCreate is the builder for creating a ApiHistory entity.
type ApiHistoryCreate struct {
	config
	mutation *ApiHistoryMutation
	hooks    []Hook
}

// SetVersion sets the "version" field.
func (ahc *ApiHistoryCreate) SetVersion(i int64) *ApiHistoryCreate {
	ahc.mutation.SetVersion(i)
	return ahc
}

// SetQueryParams sets the "query_params" field.
func (ahc *ApiHistoryCreate) SetQueryParams(s string) *ApiHistoryCreate {
	ahc.mutation.SetQueryParams(s)
	return ahc
}

// SetCreatedAt sets the "created_at" field.
func (ahc *ApiHistoryCreate) SetCreatedAt(t time.Time) *ApiHistoryCreate {
	ahc.mutation.SetCreatedAt(t)
	return ahc
}

// SetCreatedBy sets the "created_by" field.
func (ahc *ApiHistoryCreate) SetCreatedBy(u uint32) *ApiHistoryCreate {
	ahc.mutation.SetCreatedBy(u)
	return ahc
}

// SetDescription sets the "description" field.
func (ahc *ApiHistoryCreate) SetDescription(s string) *ApiHistoryCreate {
	ahc.mutation.SetDescription(s)
	return ahc
}

// Mutation returns the ApiHistoryMutation object of the builder.
func (ahc *ApiHistoryCreate) Mutation() *ApiHistoryMutation {
	return ahc.mutation
}

// Save creates the ApiHistory in the database.
func (ahc *ApiHistoryCreate) Save(ctx context.Context) (*ApiHistory, error) {
	return withHooks[*ApiHistory, ApiHistoryMutation](ctx, ahc.sqlSave, ahc.mutation, ahc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ahc *ApiHistoryCreate) SaveX(ctx context.Context) *ApiHistory {
	v, err := ahc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ahc *ApiHistoryCreate) Exec(ctx context.Context) error {
	_, err := ahc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ahc *ApiHistoryCreate) ExecX(ctx context.Context) {
	if err := ahc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ahc *ApiHistoryCreate) check() error {
	if _, ok := ahc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "ApiHistory.version"`)}
	}
	if _, ok := ahc.mutation.QueryParams(); !ok {
		return &ValidationError{Name: "query_params", err: errors.New(`ent: missing required field "ApiHistory.query_params"`)}
	}
	if _, ok := ahc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ApiHistory.created_at"`)}
	}
	if _, ok := ahc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "ApiHistory.created_by"`)}
	}
	if _, ok := ahc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "ApiHistory.description"`)}
	}
	return nil
}

func (ahc *ApiHistoryCreate) sqlSave(ctx context.Context) (*ApiHistory, error) {
	if err := ahc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ahc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ahc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ahc.mutation.id = &_node.ID
	ahc.mutation.done = true
	return _node, nil
}

func (ahc *ApiHistoryCreate) createSpec() (*ApiHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &ApiHistory{config: ahc.config}
		_spec = sqlgraph.NewCreateSpec(apihistory.Table, sqlgraph.NewFieldSpec(apihistory.FieldID, field.TypeInt))
	)
	if value, ok := ahc.mutation.Version(); ok {
		_spec.SetField(apihistory.FieldVersion, field.TypeInt64, value)
		_node.Version = value
	}
	if value, ok := ahc.mutation.QueryParams(); ok {
		_spec.SetField(apihistory.FieldQueryParams, field.TypeString, value)
		_node.QueryParams = value
	}
	if value, ok := ahc.mutation.CreatedAt(); ok {
		_spec.SetField(apihistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ahc.mutation.CreatedBy(); ok {
		_spec.SetField(apihistory.FieldCreatedBy, field.TypeUint32, value)
		_node.CreatedBy = value
	}
	if value, ok := ahc.mutation.Description(); ok {
		_spec.SetField(apihistory.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	return _node, _spec
}

// ApiHistoryCreateBulk is the builder for creating many ApiHistory entities in bulk.
type ApiHistoryCreateBulk struct {
	config
	builders []*ApiHistoryCreate
}

// Save creates the ApiHistory entities in the database.
func (ahcb *ApiHistoryCreateBulk) Save(ctx context.Context) ([]*ApiHistory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ahcb.builders))
	nodes := make([]*ApiHistory, len(ahcb.builders))
	mutators := make([]Mutator, len(ahcb.builders))
	for i := range ahcb.builders {
		func(i int, root context.Context) {
			builder := ahcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApiHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ahcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ahcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ahcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ahcb *ApiHistoryCreateBulk) SaveX(ctx context.Context) []*ApiHistory {
	v, err := ahcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ahcb *ApiHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ahcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ahcb *ApiHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := ahcb.Exec(ctx); err != nil {
		panic(err)
	}
}
