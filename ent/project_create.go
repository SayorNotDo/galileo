// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"galileo/ent/project"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectCreate is the builder for creating a Project entity.
type ProjectCreate struct {
	config
	mutation *ProjectMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *ProjectCreate) SetName(s string) *ProjectCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetIdentifier sets the "identifier" field.
func (pc *ProjectCreate) SetIdentifier(s string) *ProjectCreate {
	pc.mutation.SetIdentifier(s)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProjectCreate) SetCreatedAt(t time.Time) *ProjectCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableCreatedAt(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetCreatedBy sets the "created_by" field.
func (pc *ProjectCreate) SetCreatedBy(u uint32) *ProjectCreate {
	pc.mutation.SetCreatedBy(u)
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProjectCreate) SetUpdatedAt(t time.Time) *ProjectCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetUpdateBy sets the "update_by" field.
func (pc *ProjectCreate) SetUpdateBy(s string) *ProjectCreate {
	pc.mutation.SetUpdateBy(s)
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *ProjectCreate) SetDeletedAt(t time.Time) *ProjectCreate {
	pc.mutation.SetDeletedAt(t)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableDeletedAt(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetDeletedAt(*t)
	}
	return pc
}

// SetDeletedBy sets the "deleted_by" field.
func (pc *ProjectCreate) SetDeletedBy(u uint32) *ProjectCreate {
	pc.mutation.SetDeletedBy(u)
	return pc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableDeletedBy(u *uint32) *ProjectCreate {
	if u != nil {
		pc.SetDeletedBy(*u)
	}
	return pc
}

// SetStatus sets the "status" field.
func (pc *ProjectCreate) SetStatus(i int8) *ProjectCreate {
	pc.mutation.SetStatus(i)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableStatus(i *int8) *ProjectCreate {
	if i != nil {
		pc.SetStatus(*i)
	}
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProjectCreate) SetDescription(s string) *ProjectCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableDescription(s *string) *ProjectCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetRemark sets the "remark" field.
func (pc *ProjectCreate) SetRemark(s string) *ProjectCreate {
	pc.mutation.SetRemark(s)
	return pc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableRemark(s *string) *ProjectCreate {
	if s != nil {
		pc.SetRemark(*s)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProjectCreate) SetID(u uint32) *ProjectCreate {
	pc.mutation.SetID(u)
	return pc
}

// Mutation returns the ProjectMutation object of the builder.
func (pc *ProjectCreate) Mutation() *ProjectMutation {
	return pc.mutation
}

// Save creates the Project in the database.
func (pc *ProjectCreate) Save(ctx context.Context) (*Project, error) {
	pc.defaults()
	return withHooks[*Project, ProjectMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProjectCreate) SaveX(ctx context.Context) *Project {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProjectCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProjectCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProjectCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := project.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.Status(); !ok {
		v := project.DefaultStatus
		pc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProjectCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Project.name"`)}
	}
	if _, ok := pc.mutation.Identifier(); !ok {
		return &ValidationError{Name: "identifier", err: errors.New(`ent: missing required field "Project.identifier"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Project.created_at"`)}
	}
	if _, ok := pc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Project.created_by"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Project.updated_at"`)}
	}
	if _, ok := pc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "update_by", err: errors.New(`ent: missing required field "Project.update_by"`)}
	}
	if _, ok := pc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Project.status"`)}
	}
	return nil
}

func (pc *ProjectCreate) sqlSave(ctx context.Context) (*Project, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProjectCreate) createSpec() (*Project, *sqlgraph.CreateSpec) {
	var (
		_node = &Project{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(project.Table, sqlgraph.NewFieldSpec(project.FieldID, field.TypeUint32))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Identifier(); ok {
		_spec.SetField(project.FieldIdentifier, field.TypeString, value)
		_node.Identifier = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(project.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.CreatedBy(); ok {
		_spec.SetField(project.FieldCreatedBy, field.TypeUint32, value)
		_node.CreatedBy = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(project.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.UpdateBy(); ok {
		_spec.SetField(project.FieldUpdateBy, field.TypeString, value)
		_node.UpdateBy = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.SetField(project.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := pc.mutation.DeletedBy(); ok {
		_spec.SetField(project.FieldDeletedBy, field.TypeUint32, value)
		_node.DeletedBy = &value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(project.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if value, ok := pc.mutation.Remark(); ok {
		_spec.SetField(project.FieldRemark, field.TypeString, value)
		_node.Remark = &value
	}
	return _node, _spec
}

// ProjectCreateBulk is the builder for creating many Project entities in bulk.
type ProjectCreateBulk struct {
	config
	builders []*ProjectCreate
}

// Save creates the Project entities in the database.
func (pcb *ProjectCreateBulk) Save(ctx context.Context) ([]*Project, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Project, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
					nodes[i].ID = uint32(id)
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProjectCreateBulk) SaveX(ctx context.Context) []*Project {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProjectCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProjectCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
