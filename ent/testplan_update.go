// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"galileo/ent/predicate"
	"galileo/ent/testplan"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// TestPlanUpdate is the builder for updating TestPlan entities.
type TestPlanUpdate struct {
	config
	hooks    []Hook
	mutation *TestPlanMutation
}

// Where appends a list predicates to the TestPlanUpdate builder.
func (tpu *TestPlanUpdate) Where(ps ...predicate.TestPlan) *TestPlanUpdate {
	tpu.mutation.Where(ps...)
	return tpu
}

// SetName sets the "name" field.
func (tpu *TestPlanUpdate) SetName(s string) *TestPlanUpdate {
	tpu.mutation.SetName(s)
	return tpu
}

// SetUpdatedAt sets the "updated_at" field.
func (tpu *TestPlanUpdate) SetUpdatedAt(t time.Time) *TestPlanUpdate {
	tpu.mutation.SetUpdatedAt(t)
	return tpu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tpu *TestPlanUpdate) ClearUpdatedAt() *TestPlanUpdate {
	tpu.mutation.ClearUpdatedAt()
	return tpu
}

// SetUpdatedBy sets the "updated_by" field.
func (tpu *TestPlanUpdate) SetUpdatedBy(u uint32) *TestPlanUpdate {
	tpu.mutation.ResetUpdatedBy()
	tpu.mutation.SetUpdatedBy(u)
	return tpu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tpu *TestPlanUpdate) SetNillableUpdatedBy(u *uint32) *TestPlanUpdate {
	if u != nil {
		tpu.SetUpdatedBy(*u)
	}
	return tpu
}

// AddUpdatedBy adds u to the "updated_by" field.
func (tpu *TestPlanUpdate) AddUpdatedBy(u int32) *TestPlanUpdate {
	tpu.mutation.AddUpdatedBy(u)
	return tpu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (tpu *TestPlanUpdate) ClearUpdatedBy() *TestPlanUpdate {
	tpu.mutation.ClearUpdatedBy()
	return tpu
}

// SetDescription sets the "description" field.
func (tpu *TestPlanUpdate) SetDescription(s string) *TestPlanUpdate {
	tpu.mutation.SetDescription(s)
	return tpu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tpu *TestPlanUpdate) SetNillableDescription(s *string) *TestPlanUpdate {
	if s != nil {
		tpu.SetDescription(*s)
	}
	return tpu
}

// ClearDescription clears the value of the "description" field.
func (tpu *TestPlanUpdate) ClearDescription() *TestPlanUpdate {
	tpu.mutation.ClearDescription()
	return tpu
}

// SetStartTime sets the "start_time" field.
func (tpu *TestPlanUpdate) SetStartTime(t time.Time) *TestPlanUpdate {
	tpu.mutation.SetStartTime(t)
	return tpu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (tpu *TestPlanUpdate) SetNillableStartTime(t *time.Time) *TestPlanUpdate {
	if t != nil {
		tpu.SetStartTime(*t)
	}
	return tpu
}

// ClearStartTime clears the value of the "start_time" field.
func (tpu *TestPlanUpdate) ClearStartTime() *TestPlanUpdate {
	tpu.mutation.ClearStartTime()
	return tpu
}

// SetDeadline sets the "deadline" field.
func (tpu *TestPlanUpdate) SetDeadline(t time.Time) *TestPlanUpdate {
	tpu.mutation.SetDeadline(t)
	return tpu
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (tpu *TestPlanUpdate) SetNillableDeadline(t *time.Time) *TestPlanUpdate {
	if t != nil {
		tpu.SetDeadline(*t)
	}
	return tpu
}

// ClearDeadline clears the value of the "deadline" field.
func (tpu *TestPlanUpdate) ClearDeadline() *TestPlanUpdate {
	tpu.mutation.ClearDeadline()
	return tpu
}

// SetStatusUpdatedAt sets the "status_updated_at" field.
func (tpu *TestPlanUpdate) SetStatusUpdatedAt(t time.Time) *TestPlanUpdate {
	tpu.mutation.SetStatusUpdatedAt(t)
	return tpu
}

// SetNillableStatusUpdatedAt sets the "status_updated_at" field if the given value is not nil.
func (tpu *TestPlanUpdate) SetNillableStatusUpdatedAt(t *time.Time) *TestPlanUpdate {
	if t != nil {
		tpu.SetStatusUpdatedAt(*t)
	}
	return tpu
}

// ClearStatusUpdatedAt clears the value of the "status_updated_at" field.
func (tpu *TestPlanUpdate) ClearStatusUpdatedAt() *TestPlanUpdate {
	tpu.mutation.ClearStatusUpdatedAt()
	return tpu
}

// SetStatus sets the "status" field.
func (tpu *TestPlanUpdate) SetStatus(i int8) *TestPlanUpdate {
	tpu.mutation.ResetStatus()
	tpu.mutation.SetStatus(i)
	return tpu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tpu *TestPlanUpdate) SetNillableStatus(i *int8) *TestPlanUpdate {
	if i != nil {
		tpu.SetStatus(*i)
	}
	return tpu
}

// AddStatus adds i to the "status" field.
func (tpu *TestPlanUpdate) AddStatus(i int8) *TestPlanUpdate {
	tpu.mutation.AddStatus(i)
	return tpu
}

// SetTasks sets the "tasks" field.
func (tpu *TestPlanUpdate) SetTasks(i []int32) *TestPlanUpdate {
	tpu.mutation.SetTasks(i)
	return tpu
}

// AppendTasks appends i to the "tasks" field.
func (tpu *TestPlanUpdate) AppendTasks(i []int32) *TestPlanUpdate {
	tpu.mutation.AppendTasks(i)
	return tpu
}

// ClearTasks clears the value of the "tasks" field.
func (tpu *TestPlanUpdate) ClearTasks() *TestPlanUpdate {
	tpu.mutation.ClearTasks()
	return tpu
}

// SetProjectID sets the "project_id" field.
func (tpu *TestPlanUpdate) SetProjectID(i int32) *TestPlanUpdate {
	tpu.mutation.ResetProjectID()
	tpu.mutation.SetProjectID(i)
	return tpu
}

// AddProjectID adds i to the "project_id" field.
func (tpu *TestPlanUpdate) AddProjectID(i int32) *TestPlanUpdate {
	tpu.mutation.AddProjectID(i)
	return tpu
}

// Mutation returns the TestPlanMutation object of the builder.
func (tpu *TestPlanUpdate) Mutation() *TestPlanMutation {
	return tpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tpu *TestPlanUpdate) Save(ctx context.Context) (int, error) {
	tpu.defaults()
	return withHooks[int, TestPlanMutation](ctx, tpu.sqlSave, tpu.mutation, tpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tpu *TestPlanUpdate) SaveX(ctx context.Context) int {
	affected, err := tpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tpu *TestPlanUpdate) Exec(ctx context.Context) error {
	_, err := tpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpu *TestPlanUpdate) ExecX(ctx context.Context) {
	if err := tpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tpu *TestPlanUpdate) defaults() {
	if _, ok := tpu.mutation.UpdatedAt(); !ok && !tpu.mutation.UpdatedAtCleared() {
		v := testplan.UpdateDefaultUpdatedAt()
		tpu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tpu *TestPlanUpdate) check() error {
	if v, ok := tpu.mutation.Name(); ok {
		if err := testplan.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TestPlan.name": %w`, err)}
		}
	}
	if v, ok := tpu.mutation.ProjectID(); ok {
		if err := testplan.ProjectIDValidator(v); err != nil {
			return &ValidationError{Name: "project_id", err: fmt.Errorf(`ent: validator failed for field "TestPlan.project_id": %w`, err)}
		}
	}
	return nil
}

func (tpu *TestPlanUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(testplan.Table, testplan.Columns, sqlgraph.NewFieldSpec(testplan.FieldID, field.TypeInt32))
	if ps := tpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tpu.mutation.Name(); ok {
		_spec.SetField(testplan.FieldName, field.TypeString, value)
	}
	if value, ok := tpu.mutation.UpdatedAt(); ok {
		_spec.SetField(testplan.FieldUpdatedAt, field.TypeTime, value)
	}
	if tpu.mutation.UpdatedAtCleared() {
		_spec.ClearField(testplan.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := tpu.mutation.UpdatedBy(); ok {
		_spec.SetField(testplan.FieldUpdatedBy, field.TypeUint32, value)
	}
	if value, ok := tpu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(testplan.FieldUpdatedBy, field.TypeUint32, value)
	}
	if tpu.mutation.UpdatedByCleared() {
		_spec.ClearField(testplan.FieldUpdatedBy, field.TypeUint32)
	}
	if value, ok := tpu.mutation.Description(); ok {
		_spec.SetField(testplan.FieldDescription, field.TypeString, value)
	}
	if tpu.mutation.DescriptionCleared() {
		_spec.ClearField(testplan.FieldDescription, field.TypeString)
	}
	if value, ok := tpu.mutation.StartTime(); ok {
		_spec.SetField(testplan.FieldStartTime, field.TypeTime, value)
	}
	if tpu.mutation.StartTimeCleared() {
		_spec.ClearField(testplan.FieldStartTime, field.TypeTime)
	}
	if value, ok := tpu.mutation.Deadline(); ok {
		_spec.SetField(testplan.FieldDeadline, field.TypeTime, value)
	}
	if tpu.mutation.DeadlineCleared() {
		_spec.ClearField(testplan.FieldDeadline, field.TypeTime)
	}
	if value, ok := tpu.mutation.StatusUpdatedAt(); ok {
		_spec.SetField(testplan.FieldStatusUpdatedAt, field.TypeTime, value)
	}
	if tpu.mutation.StatusUpdatedAtCleared() {
		_spec.ClearField(testplan.FieldStatusUpdatedAt, field.TypeTime)
	}
	if value, ok := tpu.mutation.Status(); ok {
		_spec.SetField(testplan.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := tpu.mutation.AddedStatus(); ok {
		_spec.AddField(testplan.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := tpu.mutation.Tasks(); ok {
		_spec.SetField(testplan.FieldTasks, field.TypeJSON, value)
	}
	if value, ok := tpu.mutation.AppendedTasks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, testplan.FieldTasks, value)
		})
	}
	if tpu.mutation.TasksCleared() {
		_spec.ClearField(testplan.FieldTasks, field.TypeJSON)
	}
	if value, ok := tpu.mutation.ProjectID(); ok {
		_spec.SetField(testplan.FieldProjectID, field.TypeInt32, value)
	}
	if value, ok := tpu.mutation.AddedProjectID(); ok {
		_spec.AddField(testplan.FieldProjectID, field.TypeInt32, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{testplan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tpu.mutation.done = true
	return n, nil
}

// TestPlanUpdateOne is the builder for updating a single TestPlan entity.
type TestPlanUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TestPlanMutation
}

// SetName sets the "name" field.
func (tpuo *TestPlanUpdateOne) SetName(s string) *TestPlanUpdateOne {
	tpuo.mutation.SetName(s)
	return tpuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tpuo *TestPlanUpdateOne) SetUpdatedAt(t time.Time) *TestPlanUpdateOne {
	tpuo.mutation.SetUpdatedAt(t)
	return tpuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tpuo *TestPlanUpdateOne) ClearUpdatedAt() *TestPlanUpdateOne {
	tpuo.mutation.ClearUpdatedAt()
	return tpuo
}

// SetUpdatedBy sets the "updated_by" field.
func (tpuo *TestPlanUpdateOne) SetUpdatedBy(u uint32) *TestPlanUpdateOne {
	tpuo.mutation.ResetUpdatedBy()
	tpuo.mutation.SetUpdatedBy(u)
	return tpuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tpuo *TestPlanUpdateOne) SetNillableUpdatedBy(u *uint32) *TestPlanUpdateOne {
	if u != nil {
		tpuo.SetUpdatedBy(*u)
	}
	return tpuo
}

// AddUpdatedBy adds u to the "updated_by" field.
func (tpuo *TestPlanUpdateOne) AddUpdatedBy(u int32) *TestPlanUpdateOne {
	tpuo.mutation.AddUpdatedBy(u)
	return tpuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (tpuo *TestPlanUpdateOne) ClearUpdatedBy() *TestPlanUpdateOne {
	tpuo.mutation.ClearUpdatedBy()
	return tpuo
}

// SetDescription sets the "description" field.
func (tpuo *TestPlanUpdateOne) SetDescription(s string) *TestPlanUpdateOne {
	tpuo.mutation.SetDescription(s)
	return tpuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tpuo *TestPlanUpdateOne) SetNillableDescription(s *string) *TestPlanUpdateOne {
	if s != nil {
		tpuo.SetDescription(*s)
	}
	return tpuo
}

// ClearDescription clears the value of the "description" field.
func (tpuo *TestPlanUpdateOne) ClearDescription() *TestPlanUpdateOne {
	tpuo.mutation.ClearDescription()
	return tpuo
}

// SetStartTime sets the "start_time" field.
func (tpuo *TestPlanUpdateOne) SetStartTime(t time.Time) *TestPlanUpdateOne {
	tpuo.mutation.SetStartTime(t)
	return tpuo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (tpuo *TestPlanUpdateOne) SetNillableStartTime(t *time.Time) *TestPlanUpdateOne {
	if t != nil {
		tpuo.SetStartTime(*t)
	}
	return tpuo
}

// ClearStartTime clears the value of the "start_time" field.
func (tpuo *TestPlanUpdateOne) ClearStartTime() *TestPlanUpdateOne {
	tpuo.mutation.ClearStartTime()
	return tpuo
}

// SetDeadline sets the "deadline" field.
func (tpuo *TestPlanUpdateOne) SetDeadline(t time.Time) *TestPlanUpdateOne {
	tpuo.mutation.SetDeadline(t)
	return tpuo
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (tpuo *TestPlanUpdateOne) SetNillableDeadline(t *time.Time) *TestPlanUpdateOne {
	if t != nil {
		tpuo.SetDeadline(*t)
	}
	return tpuo
}

// ClearDeadline clears the value of the "deadline" field.
func (tpuo *TestPlanUpdateOne) ClearDeadline() *TestPlanUpdateOne {
	tpuo.mutation.ClearDeadline()
	return tpuo
}

// SetStatusUpdatedAt sets the "status_updated_at" field.
func (tpuo *TestPlanUpdateOne) SetStatusUpdatedAt(t time.Time) *TestPlanUpdateOne {
	tpuo.mutation.SetStatusUpdatedAt(t)
	return tpuo
}

// SetNillableStatusUpdatedAt sets the "status_updated_at" field if the given value is not nil.
func (tpuo *TestPlanUpdateOne) SetNillableStatusUpdatedAt(t *time.Time) *TestPlanUpdateOne {
	if t != nil {
		tpuo.SetStatusUpdatedAt(*t)
	}
	return tpuo
}

// ClearStatusUpdatedAt clears the value of the "status_updated_at" field.
func (tpuo *TestPlanUpdateOne) ClearStatusUpdatedAt() *TestPlanUpdateOne {
	tpuo.mutation.ClearStatusUpdatedAt()
	return tpuo
}

// SetStatus sets the "status" field.
func (tpuo *TestPlanUpdateOne) SetStatus(i int8) *TestPlanUpdateOne {
	tpuo.mutation.ResetStatus()
	tpuo.mutation.SetStatus(i)
	return tpuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tpuo *TestPlanUpdateOne) SetNillableStatus(i *int8) *TestPlanUpdateOne {
	if i != nil {
		tpuo.SetStatus(*i)
	}
	return tpuo
}

// AddStatus adds i to the "status" field.
func (tpuo *TestPlanUpdateOne) AddStatus(i int8) *TestPlanUpdateOne {
	tpuo.mutation.AddStatus(i)
	return tpuo
}

// SetTasks sets the "tasks" field.
func (tpuo *TestPlanUpdateOne) SetTasks(i []int32) *TestPlanUpdateOne {
	tpuo.mutation.SetTasks(i)
	return tpuo
}

// AppendTasks appends i to the "tasks" field.
func (tpuo *TestPlanUpdateOne) AppendTasks(i []int32) *TestPlanUpdateOne {
	tpuo.mutation.AppendTasks(i)
	return tpuo
}

// ClearTasks clears the value of the "tasks" field.
func (tpuo *TestPlanUpdateOne) ClearTasks() *TestPlanUpdateOne {
	tpuo.mutation.ClearTasks()
	return tpuo
}

// SetProjectID sets the "project_id" field.
func (tpuo *TestPlanUpdateOne) SetProjectID(i int32) *TestPlanUpdateOne {
	tpuo.mutation.ResetProjectID()
	tpuo.mutation.SetProjectID(i)
	return tpuo
}

// AddProjectID adds i to the "project_id" field.
func (tpuo *TestPlanUpdateOne) AddProjectID(i int32) *TestPlanUpdateOne {
	tpuo.mutation.AddProjectID(i)
	return tpuo
}

// Mutation returns the TestPlanMutation object of the builder.
func (tpuo *TestPlanUpdateOne) Mutation() *TestPlanMutation {
	return tpuo.mutation
}

// Where appends a list predicates to the TestPlanUpdate builder.
func (tpuo *TestPlanUpdateOne) Where(ps ...predicate.TestPlan) *TestPlanUpdateOne {
	tpuo.mutation.Where(ps...)
	return tpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tpuo *TestPlanUpdateOne) Select(field string, fields ...string) *TestPlanUpdateOne {
	tpuo.fields = append([]string{field}, fields...)
	return tpuo
}

// Save executes the query and returns the updated TestPlan entity.
func (tpuo *TestPlanUpdateOne) Save(ctx context.Context) (*TestPlan, error) {
	tpuo.defaults()
	return withHooks[*TestPlan, TestPlanMutation](ctx, tpuo.sqlSave, tpuo.mutation, tpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tpuo *TestPlanUpdateOne) SaveX(ctx context.Context) *TestPlan {
	node, err := tpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tpuo *TestPlanUpdateOne) Exec(ctx context.Context) error {
	_, err := tpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpuo *TestPlanUpdateOne) ExecX(ctx context.Context) {
	if err := tpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tpuo *TestPlanUpdateOne) defaults() {
	if _, ok := tpuo.mutation.UpdatedAt(); !ok && !tpuo.mutation.UpdatedAtCleared() {
		v := testplan.UpdateDefaultUpdatedAt()
		tpuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tpuo *TestPlanUpdateOne) check() error {
	if v, ok := tpuo.mutation.Name(); ok {
		if err := testplan.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TestPlan.name": %w`, err)}
		}
	}
	if v, ok := tpuo.mutation.ProjectID(); ok {
		if err := testplan.ProjectIDValidator(v); err != nil {
			return &ValidationError{Name: "project_id", err: fmt.Errorf(`ent: validator failed for field "TestPlan.project_id": %w`, err)}
		}
	}
	return nil
}

func (tpuo *TestPlanUpdateOne) sqlSave(ctx context.Context) (_node *TestPlan, err error) {
	if err := tpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(testplan.Table, testplan.Columns, sqlgraph.NewFieldSpec(testplan.FieldID, field.TypeInt32))
	id, ok := tpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TestPlan.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testplan.FieldID)
		for _, f := range fields {
			if !testplan.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != testplan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tpuo.mutation.Name(); ok {
		_spec.SetField(testplan.FieldName, field.TypeString, value)
	}
	if value, ok := tpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(testplan.FieldUpdatedAt, field.TypeTime, value)
	}
	if tpuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(testplan.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := tpuo.mutation.UpdatedBy(); ok {
		_spec.SetField(testplan.FieldUpdatedBy, field.TypeUint32, value)
	}
	if value, ok := tpuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(testplan.FieldUpdatedBy, field.TypeUint32, value)
	}
	if tpuo.mutation.UpdatedByCleared() {
		_spec.ClearField(testplan.FieldUpdatedBy, field.TypeUint32)
	}
	if value, ok := tpuo.mutation.Description(); ok {
		_spec.SetField(testplan.FieldDescription, field.TypeString, value)
	}
	if tpuo.mutation.DescriptionCleared() {
		_spec.ClearField(testplan.FieldDescription, field.TypeString)
	}
	if value, ok := tpuo.mutation.StartTime(); ok {
		_spec.SetField(testplan.FieldStartTime, field.TypeTime, value)
	}
	if tpuo.mutation.StartTimeCleared() {
		_spec.ClearField(testplan.FieldStartTime, field.TypeTime)
	}
	if value, ok := tpuo.mutation.Deadline(); ok {
		_spec.SetField(testplan.FieldDeadline, field.TypeTime, value)
	}
	if tpuo.mutation.DeadlineCleared() {
		_spec.ClearField(testplan.FieldDeadline, field.TypeTime)
	}
	if value, ok := tpuo.mutation.StatusUpdatedAt(); ok {
		_spec.SetField(testplan.FieldStatusUpdatedAt, field.TypeTime, value)
	}
	if tpuo.mutation.StatusUpdatedAtCleared() {
		_spec.ClearField(testplan.FieldStatusUpdatedAt, field.TypeTime)
	}
	if value, ok := tpuo.mutation.Status(); ok {
		_spec.SetField(testplan.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := tpuo.mutation.AddedStatus(); ok {
		_spec.AddField(testplan.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := tpuo.mutation.Tasks(); ok {
		_spec.SetField(testplan.FieldTasks, field.TypeJSON, value)
	}
	if value, ok := tpuo.mutation.AppendedTasks(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, testplan.FieldTasks, value)
		})
	}
	if tpuo.mutation.TasksCleared() {
		_spec.ClearField(testplan.FieldTasks, field.TypeJSON)
	}
	if value, ok := tpuo.mutation.ProjectID(); ok {
		_spec.SetField(testplan.FieldProjectID, field.TypeInt32, value)
	}
	if value, ok := tpuo.mutation.AddedProjectID(); ok {
		_spec.AddField(testplan.FieldProjectID, field.TypeInt32, value)
	}
	_node = &TestPlan{config: tpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{testplan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tpuo.mutation.done = true
	return _node, nil
}