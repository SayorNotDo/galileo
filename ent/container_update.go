// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"galileo/ent/container"
	"galileo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// ContainerUpdate is the builder for updating Container entities.
type ContainerUpdate struct {
	config
	hooks    []Hook
	mutation *ContainerMutation
}

// Where appends a list predicates to the ContainerUpdate builder.
func (cu *ContainerUpdate) Where(ps ...predicate.Container) *ContainerUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetHostname sets the "hostname" field.
func (cu *ContainerUpdate) SetHostname(s string) *ContainerUpdate {
	cu.mutation.SetHostname(s)
	return cu
}

// SetDomainname sets the "domainname" field.
func (cu *ContainerUpdate) SetDomainname(s string) *ContainerUpdate {
	cu.mutation.SetDomainname(s)
	return cu
}

// SetUser sets the "user" field.
func (cu *ContainerUpdate) SetUser(s string) *ContainerUpdate {
	cu.mutation.SetUser(s)
	return cu
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (cu *ContainerUpdate) SetNillableUser(s *string) *ContainerUpdate {
	if s != nil {
		cu.SetUser(*s)
	}
	return cu
}

// ClearUser clears the value of the "user" field.
func (cu *ContainerUpdate) ClearUser() *ContainerUpdate {
	cu.mutation.ClearUser()
	return cu
}

// SetEnv sets the "env" field.
func (cu *ContainerUpdate) SetEnv(s []string) *ContainerUpdate {
	cu.mutation.SetEnv(s)
	return cu
}

// AppendEnv appends s to the "env" field.
func (cu *ContainerUpdate) AppendEnv(s []string) *ContainerUpdate {
	cu.mutation.AppendEnv(s)
	return cu
}

// ClearEnv clears the value of the "env" field.
func (cu *ContainerUpdate) ClearEnv() *ContainerUpdate {
	cu.mutation.ClearEnv()
	return cu
}

// SetCmd sets the "cmd" field.
func (cu *ContainerUpdate) SetCmd(s []string) *ContainerUpdate {
	cu.mutation.SetCmd(s)
	return cu
}

// AppendCmd appends s to the "cmd" field.
func (cu *ContainerUpdate) AppendCmd(s []string) *ContainerUpdate {
	cu.mutation.AppendCmd(s)
	return cu
}

// ClearCmd clears the value of the "cmd" field.
func (cu *ContainerUpdate) ClearCmd() *ContainerUpdate {
	cu.mutation.ClearCmd()
	return cu
}

// SetImage sets the "image" field.
func (cu *ContainerUpdate) SetImage(s string) *ContainerUpdate {
	cu.mutation.SetImage(s)
	return cu
}

// SetLabels sets the "labels" field.
func (cu *ContainerUpdate) SetLabels(s []string) *ContainerUpdate {
	cu.mutation.SetLabels(s)
	return cu
}

// AppendLabels appends s to the "labels" field.
func (cu *ContainerUpdate) AppendLabels(s []string) *ContainerUpdate {
	cu.mutation.AppendLabels(s)
	return cu
}

// ClearLabels clears the value of the "labels" field.
func (cu *ContainerUpdate) ClearLabels() *ContainerUpdate {
	cu.mutation.ClearLabels()
	return cu
}

// SetVolumes sets the "volumes" field.
func (cu *ContainerUpdate) SetVolumes(s []string) *ContainerUpdate {
	cu.mutation.SetVolumes(s)
	return cu
}

// AppendVolumes appends s to the "volumes" field.
func (cu *ContainerUpdate) AppendVolumes(s []string) *ContainerUpdate {
	cu.mutation.AppendVolumes(s)
	return cu
}

// ClearVolumes clears the value of the "volumes" field.
func (cu *ContainerUpdate) ClearVolumes() *ContainerUpdate {
	cu.mutation.ClearVolumes()
	return cu
}

// SetWorkingDir sets the "working_dir" field.
func (cu *ContainerUpdate) SetWorkingDir(s string) *ContainerUpdate {
	cu.mutation.SetWorkingDir(s)
	return cu
}

// SetNillableWorkingDir sets the "working_dir" field if the given value is not nil.
func (cu *ContainerUpdate) SetNillableWorkingDir(s *string) *ContainerUpdate {
	if s != nil {
		cu.SetWorkingDir(*s)
	}
	return cu
}

// ClearWorkingDir clears the value of the "working_dir" field.
func (cu *ContainerUpdate) ClearWorkingDir() *ContainerUpdate {
	cu.mutation.ClearWorkingDir()
	return cu
}

// SetEntrypoint sets the "entrypoint" field.
func (cu *ContainerUpdate) SetEntrypoint(s []string) *ContainerUpdate {
	cu.mutation.SetEntrypoint(s)
	return cu
}

// AppendEntrypoint appends s to the "entrypoint" field.
func (cu *ContainerUpdate) AppendEntrypoint(s []string) *ContainerUpdate {
	cu.mutation.AppendEntrypoint(s)
	return cu
}

// ClearEntrypoint clears the value of the "entrypoint" field.
func (cu *ContainerUpdate) ClearEntrypoint() *ContainerUpdate {
	cu.mutation.ClearEntrypoint()
	return cu
}

// SetMACAddress sets the "mac_address" field.
func (cu *ContainerUpdate) SetMACAddress(s string) *ContainerUpdate {
	cu.mutation.SetMACAddress(s)
	return cu
}

// SetExposePorts sets the "expose_ports" field.
func (cu *ContainerUpdate) SetExposePorts(s []string) *ContainerUpdate {
	cu.mutation.SetExposePorts(s)
	return cu
}

// AppendExposePorts appends s to the "expose_ports" field.
func (cu *ContainerUpdate) AppendExposePorts(s []string) *ContainerUpdate {
	cu.mutation.AppendExposePorts(s)
	return cu
}

// ClearExposePorts clears the value of the "expose_ports" field.
func (cu *ContainerUpdate) ClearExposePorts() *ContainerUpdate {
	cu.mutation.ClearExposePorts()
	return cu
}

// SetComposeFileURL sets the "compose_file_url" field.
func (cu *ContainerUpdate) SetComposeFileURL(s string) *ContainerUpdate {
	cu.mutation.SetComposeFileURL(s)
	return cu
}

// SetNillableComposeFileURL sets the "compose_file_url" field if the given value is not nil.
func (cu *ContainerUpdate) SetNillableComposeFileURL(s *string) *ContainerUpdate {
	if s != nil {
		cu.SetComposeFileURL(*s)
	}
	return cu
}

// ClearComposeFileURL clears the value of the "compose_file_url" field.
func (cu *ContainerUpdate) ClearComposeFileURL() *ContainerUpdate {
	cu.mutation.ClearComposeFileURL()
	return cu
}

// SetDockerfileURL sets the "dockerfile_url" field.
func (cu *ContainerUpdate) SetDockerfileURL(s string) *ContainerUpdate {
	cu.mutation.SetDockerfileURL(s)
	return cu
}

// SetNillableDockerfileURL sets the "dockerfile_url" field if the given value is not nil.
func (cu *ContainerUpdate) SetNillableDockerfileURL(s *string) *ContainerUpdate {
	if s != nil {
		cu.SetDockerfileURL(*s)
	}
	return cu
}

// ClearDockerfileURL clears the value of the "dockerfile_url" field.
func (cu *ContainerUpdate) ClearDockerfileURL() *ContainerUpdate {
	cu.mutation.ClearDockerfileURL()
	return cu
}

// Mutation returns the ContainerMutation object of the builder.
func (cu *ContainerUpdate) Mutation() *ContainerMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ContainerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ContainerMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ContainerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ContainerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ContainerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ContainerUpdate) check() error {
	if v, ok := cu.mutation.Hostname(); ok {
		if err := container.HostnameValidator(v); err != nil {
			return &ValidationError{Name: "hostname", err: fmt.Errorf(`ent: validator failed for field "Container.hostname": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Domainname(); ok {
		if err := container.DomainnameValidator(v); err != nil {
			return &ValidationError{Name: "domainname", err: fmt.Errorf(`ent: validator failed for field "Container.domainname": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Image(); ok {
		if err := container.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "Container.image": %w`, err)}
		}
	}
	if v, ok := cu.mutation.MACAddress(); ok {
		if err := container.MACAddressValidator(v); err != nil {
			return &ValidationError{Name: "mac_address", err: fmt.Errorf(`ent: validator failed for field "Container.mac_address": %w`, err)}
		}
	}
	return nil
}

func (cu *ContainerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(container.Table, container.Columns, sqlgraph.NewFieldSpec(container.FieldID, field.TypeString))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Hostname(); ok {
		_spec.SetField(container.FieldHostname, field.TypeString, value)
	}
	if value, ok := cu.mutation.Domainname(); ok {
		_spec.SetField(container.FieldDomainname, field.TypeString, value)
	}
	if value, ok := cu.mutation.User(); ok {
		_spec.SetField(container.FieldUser, field.TypeString, value)
	}
	if cu.mutation.UserCleared() {
		_spec.ClearField(container.FieldUser, field.TypeString)
	}
	if value, ok := cu.mutation.Env(); ok {
		_spec.SetField(container.FieldEnv, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedEnv(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, container.FieldEnv, value)
		})
	}
	if cu.mutation.EnvCleared() {
		_spec.ClearField(container.FieldEnv, field.TypeJSON)
	}
	if value, ok := cu.mutation.Cmd(); ok {
		_spec.SetField(container.FieldCmd, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedCmd(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, container.FieldCmd, value)
		})
	}
	if cu.mutation.CmdCleared() {
		_spec.ClearField(container.FieldCmd, field.TypeJSON)
	}
	if value, ok := cu.mutation.Image(); ok {
		_spec.SetField(container.FieldImage, field.TypeString, value)
	}
	if value, ok := cu.mutation.Labels(); ok {
		_spec.SetField(container.FieldLabels, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedLabels(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, container.FieldLabels, value)
		})
	}
	if cu.mutation.LabelsCleared() {
		_spec.ClearField(container.FieldLabels, field.TypeJSON)
	}
	if value, ok := cu.mutation.Volumes(); ok {
		_spec.SetField(container.FieldVolumes, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedVolumes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, container.FieldVolumes, value)
		})
	}
	if cu.mutation.VolumesCleared() {
		_spec.ClearField(container.FieldVolumes, field.TypeJSON)
	}
	if value, ok := cu.mutation.WorkingDir(); ok {
		_spec.SetField(container.FieldWorkingDir, field.TypeString, value)
	}
	if cu.mutation.WorkingDirCleared() {
		_spec.ClearField(container.FieldWorkingDir, field.TypeString)
	}
	if value, ok := cu.mutation.Entrypoint(); ok {
		_spec.SetField(container.FieldEntrypoint, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedEntrypoint(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, container.FieldEntrypoint, value)
		})
	}
	if cu.mutation.EntrypointCleared() {
		_spec.ClearField(container.FieldEntrypoint, field.TypeJSON)
	}
	if value, ok := cu.mutation.MACAddress(); ok {
		_spec.SetField(container.FieldMACAddress, field.TypeString, value)
	}
	if value, ok := cu.mutation.ExposePorts(); ok {
		_spec.SetField(container.FieldExposePorts, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedExposePorts(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, container.FieldExposePorts, value)
		})
	}
	if cu.mutation.ExposePortsCleared() {
		_spec.ClearField(container.FieldExposePorts, field.TypeJSON)
	}
	if value, ok := cu.mutation.ComposeFileURL(); ok {
		_spec.SetField(container.FieldComposeFileURL, field.TypeString, value)
	}
	if cu.mutation.ComposeFileURLCleared() {
		_spec.ClearField(container.FieldComposeFileURL, field.TypeString)
	}
	if value, ok := cu.mutation.DockerfileURL(); ok {
		_spec.SetField(container.FieldDockerfileURL, field.TypeString, value)
	}
	if cu.mutation.DockerfileURLCleared() {
		_spec.ClearField(container.FieldDockerfileURL, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{container.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ContainerUpdateOne is the builder for updating a single Container entity.
type ContainerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContainerMutation
}

// SetHostname sets the "hostname" field.
func (cuo *ContainerUpdateOne) SetHostname(s string) *ContainerUpdateOne {
	cuo.mutation.SetHostname(s)
	return cuo
}

// SetDomainname sets the "domainname" field.
func (cuo *ContainerUpdateOne) SetDomainname(s string) *ContainerUpdateOne {
	cuo.mutation.SetDomainname(s)
	return cuo
}

// SetUser sets the "user" field.
func (cuo *ContainerUpdateOne) SetUser(s string) *ContainerUpdateOne {
	cuo.mutation.SetUser(s)
	return cuo
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (cuo *ContainerUpdateOne) SetNillableUser(s *string) *ContainerUpdateOne {
	if s != nil {
		cuo.SetUser(*s)
	}
	return cuo
}

// ClearUser clears the value of the "user" field.
func (cuo *ContainerUpdateOne) ClearUser() *ContainerUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// SetEnv sets the "env" field.
func (cuo *ContainerUpdateOne) SetEnv(s []string) *ContainerUpdateOne {
	cuo.mutation.SetEnv(s)
	return cuo
}

// AppendEnv appends s to the "env" field.
func (cuo *ContainerUpdateOne) AppendEnv(s []string) *ContainerUpdateOne {
	cuo.mutation.AppendEnv(s)
	return cuo
}

// ClearEnv clears the value of the "env" field.
func (cuo *ContainerUpdateOne) ClearEnv() *ContainerUpdateOne {
	cuo.mutation.ClearEnv()
	return cuo
}

// SetCmd sets the "cmd" field.
func (cuo *ContainerUpdateOne) SetCmd(s []string) *ContainerUpdateOne {
	cuo.mutation.SetCmd(s)
	return cuo
}

// AppendCmd appends s to the "cmd" field.
func (cuo *ContainerUpdateOne) AppendCmd(s []string) *ContainerUpdateOne {
	cuo.mutation.AppendCmd(s)
	return cuo
}

// ClearCmd clears the value of the "cmd" field.
func (cuo *ContainerUpdateOne) ClearCmd() *ContainerUpdateOne {
	cuo.mutation.ClearCmd()
	return cuo
}

// SetImage sets the "image" field.
func (cuo *ContainerUpdateOne) SetImage(s string) *ContainerUpdateOne {
	cuo.mutation.SetImage(s)
	return cuo
}

// SetLabels sets the "labels" field.
func (cuo *ContainerUpdateOne) SetLabels(s []string) *ContainerUpdateOne {
	cuo.mutation.SetLabels(s)
	return cuo
}

// AppendLabels appends s to the "labels" field.
func (cuo *ContainerUpdateOne) AppendLabels(s []string) *ContainerUpdateOne {
	cuo.mutation.AppendLabels(s)
	return cuo
}

// ClearLabels clears the value of the "labels" field.
func (cuo *ContainerUpdateOne) ClearLabels() *ContainerUpdateOne {
	cuo.mutation.ClearLabels()
	return cuo
}

// SetVolumes sets the "volumes" field.
func (cuo *ContainerUpdateOne) SetVolumes(s []string) *ContainerUpdateOne {
	cuo.mutation.SetVolumes(s)
	return cuo
}

// AppendVolumes appends s to the "volumes" field.
func (cuo *ContainerUpdateOne) AppendVolumes(s []string) *ContainerUpdateOne {
	cuo.mutation.AppendVolumes(s)
	return cuo
}

// ClearVolumes clears the value of the "volumes" field.
func (cuo *ContainerUpdateOne) ClearVolumes() *ContainerUpdateOne {
	cuo.mutation.ClearVolumes()
	return cuo
}

// SetWorkingDir sets the "working_dir" field.
func (cuo *ContainerUpdateOne) SetWorkingDir(s string) *ContainerUpdateOne {
	cuo.mutation.SetWorkingDir(s)
	return cuo
}

// SetNillableWorkingDir sets the "working_dir" field if the given value is not nil.
func (cuo *ContainerUpdateOne) SetNillableWorkingDir(s *string) *ContainerUpdateOne {
	if s != nil {
		cuo.SetWorkingDir(*s)
	}
	return cuo
}

// ClearWorkingDir clears the value of the "working_dir" field.
func (cuo *ContainerUpdateOne) ClearWorkingDir() *ContainerUpdateOne {
	cuo.mutation.ClearWorkingDir()
	return cuo
}

// SetEntrypoint sets the "entrypoint" field.
func (cuo *ContainerUpdateOne) SetEntrypoint(s []string) *ContainerUpdateOne {
	cuo.mutation.SetEntrypoint(s)
	return cuo
}

// AppendEntrypoint appends s to the "entrypoint" field.
func (cuo *ContainerUpdateOne) AppendEntrypoint(s []string) *ContainerUpdateOne {
	cuo.mutation.AppendEntrypoint(s)
	return cuo
}

// ClearEntrypoint clears the value of the "entrypoint" field.
func (cuo *ContainerUpdateOne) ClearEntrypoint() *ContainerUpdateOne {
	cuo.mutation.ClearEntrypoint()
	return cuo
}

// SetMACAddress sets the "mac_address" field.
func (cuo *ContainerUpdateOne) SetMACAddress(s string) *ContainerUpdateOne {
	cuo.mutation.SetMACAddress(s)
	return cuo
}

// SetExposePorts sets the "expose_ports" field.
func (cuo *ContainerUpdateOne) SetExposePorts(s []string) *ContainerUpdateOne {
	cuo.mutation.SetExposePorts(s)
	return cuo
}

// AppendExposePorts appends s to the "expose_ports" field.
func (cuo *ContainerUpdateOne) AppendExposePorts(s []string) *ContainerUpdateOne {
	cuo.mutation.AppendExposePorts(s)
	return cuo
}

// ClearExposePorts clears the value of the "expose_ports" field.
func (cuo *ContainerUpdateOne) ClearExposePorts() *ContainerUpdateOne {
	cuo.mutation.ClearExposePorts()
	return cuo
}

// SetComposeFileURL sets the "compose_file_url" field.
func (cuo *ContainerUpdateOne) SetComposeFileURL(s string) *ContainerUpdateOne {
	cuo.mutation.SetComposeFileURL(s)
	return cuo
}

// SetNillableComposeFileURL sets the "compose_file_url" field if the given value is not nil.
func (cuo *ContainerUpdateOne) SetNillableComposeFileURL(s *string) *ContainerUpdateOne {
	if s != nil {
		cuo.SetComposeFileURL(*s)
	}
	return cuo
}

// ClearComposeFileURL clears the value of the "compose_file_url" field.
func (cuo *ContainerUpdateOne) ClearComposeFileURL() *ContainerUpdateOne {
	cuo.mutation.ClearComposeFileURL()
	return cuo
}

// SetDockerfileURL sets the "dockerfile_url" field.
func (cuo *ContainerUpdateOne) SetDockerfileURL(s string) *ContainerUpdateOne {
	cuo.mutation.SetDockerfileURL(s)
	return cuo
}

// SetNillableDockerfileURL sets the "dockerfile_url" field if the given value is not nil.
func (cuo *ContainerUpdateOne) SetNillableDockerfileURL(s *string) *ContainerUpdateOne {
	if s != nil {
		cuo.SetDockerfileURL(*s)
	}
	return cuo
}

// ClearDockerfileURL clears the value of the "dockerfile_url" field.
func (cuo *ContainerUpdateOne) ClearDockerfileURL() *ContainerUpdateOne {
	cuo.mutation.ClearDockerfileURL()
	return cuo
}

// Mutation returns the ContainerMutation object of the builder.
func (cuo *ContainerUpdateOne) Mutation() *ContainerMutation {
	return cuo.mutation
}

// Where appends a list predicates to the ContainerUpdate builder.
func (cuo *ContainerUpdateOne) Where(ps ...predicate.Container) *ContainerUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ContainerUpdateOne) Select(field string, fields ...string) *ContainerUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Container entity.
func (cuo *ContainerUpdateOne) Save(ctx context.Context) (*Container, error) {
	return withHooks[*Container, ContainerMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ContainerUpdateOne) SaveX(ctx context.Context) *Container {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ContainerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ContainerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ContainerUpdateOne) check() error {
	if v, ok := cuo.mutation.Hostname(); ok {
		if err := container.HostnameValidator(v); err != nil {
			return &ValidationError{Name: "hostname", err: fmt.Errorf(`ent: validator failed for field "Container.hostname": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Domainname(); ok {
		if err := container.DomainnameValidator(v); err != nil {
			return &ValidationError{Name: "domainname", err: fmt.Errorf(`ent: validator failed for field "Container.domainname": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Image(); ok {
		if err := container.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "Container.image": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.MACAddress(); ok {
		if err := container.MACAddressValidator(v); err != nil {
			return &ValidationError{Name: "mac_address", err: fmt.Errorf(`ent: validator failed for field "Container.mac_address": %w`, err)}
		}
	}
	return nil
}

func (cuo *ContainerUpdateOne) sqlSave(ctx context.Context) (_node *Container, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(container.Table, container.Columns, sqlgraph.NewFieldSpec(container.FieldID, field.TypeString))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Container.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, container.FieldID)
		for _, f := range fields {
			if !container.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != container.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Hostname(); ok {
		_spec.SetField(container.FieldHostname, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Domainname(); ok {
		_spec.SetField(container.FieldDomainname, field.TypeString, value)
	}
	if value, ok := cuo.mutation.User(); ok {
		_spec.SetField(container.FieldUser, field.TypeString, value)
	}
	if cuo.mutation.UserCleared() {
		_spec.ClearField(container.FieldUser, field.TypeString)
	}
	if value, ok := cuo.mutation.Env(); ok {
		_spec.SetField(container.FieldEnv, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedEnv(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, container.FieldEnv, value)
		})
	}
	if cuo.mutation.EnvCleared() {
		_spec.ClearField(container.FieldEnv, field.TypeJSON)
	}
	if value, ok := cuo.mutation.Cmd(); ok {
		_spec.SetField(container.FieldCmd, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedCmd(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, container.FieldCmd, value)
		})
	}
	if cuo.mutation.CmdCleared() {
		_spec.ClearField(container.FieldCmd, field.TypeJSON)
	}
	if value, ok := cuo.mutation.Image(); ok {
		_spec.SetField(container.FieldImage, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Labels(); ok {
		_spec.SetField(container.FieldLabels, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedLabels(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, container.FieldLabels, value)
		})
	}
	if cuo.mutation.LabelsCleared() {
		_spec.ClearField(container.FieldLabels, field.TypeJSON)
	}
	if value, ok := cuo.mutation.Volumes(); ok {
		_spec.SetField(container.FieldVolumes, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedVolumes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, container.FieldVolumes, value)
		})
	}
	if cuo.mutation.VolumesCleared() {
		_spec.ClearField(container.FieldVolumes, field.TypeJSON)
	}
	if value, ok := cuo.mutation.WorkingDir(); ok {
		_spec.SetField(container.FieldWorkingDir, field.TypeString, value)
	}
	if cuo.mutation.WorkingDirCleared() {
		_spec.ClearField(container.FieldWorkingDir, field.TypeString)
	}
	if value, ok := cuo.mutation.Entrypoint(); ok {
		_spec.SetField(container.FieldEntrypoint, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedEntrypoint(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, container.FieldEntrypoint, value)
		})
	}
	if cuo.mutation.EntrypointCleared() {
		_spec.ClearField(container.FieldEntrypoint, field.TypeJSON)
	}
	if value, ok := cuo.mutation.MACAddress(); ok {
		_spec.SetField(container.FieldMACAddress, field.TypeString, value)
	}
	if value, ok := cuo.mutation.ExposePorts(); ok {
		_spec.SetField(container.FieldExposePorts, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedExposePorts(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, container.FieldExposePorts, value)
		})
	}
	if cuo.mutation.ExposePortsCleared() {
		_spec.ClearField(container.FieldExposePorts, field.TypeJSON)
	}
	if value, ok := cuo.mutation.ComposeFileURL(); ok {
		_spec.SetField(container.FieldComposeFileURL, field.TypeString, value)
	}
	if cuo.mutation.ComposeFileURLCleared() {
		_spec.ClearField(container.FieldComposeFileURL, field.TypeString)
	}
	if value, ok := cuo.mutation.DockerfileURL(); ok {
		_spec.SetField(container.FieldDockerfileURL, field.TypeString, value)
	}
	if cuo.mutation.DockerfileURLCleared() {
		_spec.ClearField(container.FieldDockerfileURL, field.TypeString)
	}
	_node = &Container{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{container.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}