// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"galileo/ent/predicate"
	"galileo/ent/task"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TaskUpdate) SetName(s string) *TaskUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetAssignee sets the "assignee" field.
func (tu *TaskUpdate) SetAssignee(u uint32) *TaskUpdate {
	tu.mutation.ResetAssignee()
	tu.mutation.SetAssignee(u)
	return tu
}

// SetNillableAssignee sets the "assignee" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableAssignee(u *uint32) *TaskUpdate {
	if u != nil {
		tu.SetAssignee(*u)
	}
	return tu
}

// AddAssignee adds u to the "assignee" field.
func (tu *TaskUpdate) AddAssignee(u int32) *TaskUpdate {
	tu.mutation.AddAssignee(u)
	return tu
}

// ClearAssignee clears the value of the "assignee" field.
func (tu *TaskUpdate) ClearAssignee() *TaskUpdate {
	tu.mutation.ClearAssignee()
	return tu
}

// SetType sets the "type" field.
func (tu *TaskUpdate) SetType(i int8) *TaskUpdate {
	tu.mutation.ResetType()
	tu.mutation.SetType(i)
	return tu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableType(i *int8) *TaskUpdate {
	if i != nil {
		tu.SetType(*i)
	}
	return tu
}

// AddType adds i to the "type" field.
func (tu *TaskUpdate) AddType(i int8) *TaskUpdate {
	tu.mutation.AddType(i)
	return tu
}

// SetFrequency sets the "frequency" field.
func (tu *TaskUpdate) SetFrequency(i int8) *TaskUpdate {
	tu.mutation.ResetFrequency()
	tu.mutation.SetFrequency(i)
	return tu
}

// SetNillableFrequency sets the "frequency" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableFrequency(i *int8) *TaskUpdate {
	if i != nil {
		tu.SetFrequency(*i)
	}
	return tu
}

// AddFrequency adds i to the "frequency" field.
func (tu *TaskUpdate) AddFrequency(i int8) *TaskUpdate {
	tu.mutation.AddFrequency(i)
	return tu
}

// ClearFrequency clears the value of the "frequency" field.
func (tu *TaskUpdate) ClearFrequency() *TaskUpdate {
	tu.mutation.ClearFrequency()
	return tu
}

// SetScheduleTime sets the "schedule_time" field.
func (tu *TaskUpdate) SetScheduleTime(t time.Time) *TaskUpdate {
	tu.mutation.SetScheduleTime(t)
	return tu
}

// SetNillableScheduleTime sets the "schedule_time" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableScheduleTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetScheduleTime(*t)
	}
	return tu
}

// ClearScheduleTime clears the value of the "schedule_time" field.
func (tu *TaskUpdate) ClearScheduleTime() *TaskUpdate {
	tu.mutation.ClearScheduleTime()
	return tu
}

// SetRank sets the "rank" field.
func (tu *TaskUpdate) SetRank(i int8) *TaskUpdate {
	tu.mutation.ResetRank()
	tu.mutation.SetRank(i)
	return tu
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableRank(i *int8) *TaskUpdate {
	if i != nil {
		tu.SetRank(*i)
	}
	return tu
}

// AddRank adds i to the "rank" field.
func (tu *TaskUpdate) AddRank(i int8) *TaskUpdate {
	tu.mutation.AddRank(i)
	return tu
}

// SetStatus sets the "status" field.
func (tu *TaskUpdate) SetStatus(i int8) *TaskUpdate {
	tu.mutation.ResetStatus()
	tu.mutation.SetStatus(i)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableStatus(i *int8) *TaskUpdate {
	if i != nil {
		tu.SetStatus(*i)
	}
	return tu
}

// AddStatus adds i to the "status" field.
func (tu *TaskUpdate) AddStatus(i int8) *TaskUpdate {
	tu.mutation.AddStatus(i)
	return tu
}

// SetStartTime sets the "start_time" field.
func (tu *TaskUpdate) SetStartTime(t time.Time) *TaskUpdate {
	tu.mutation.SetStartTime(t)
	return tu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableStartTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetStartTime(*t)
	}
	return tu
}

// ClearStartTime clears the value of the "start_time" field.
func (tu *TaskUpdate) ClearStartTime() *TaskUpdate {
	tu.mutation.ClearStartTime()
	return tu
}

// SetCompletedAt sets the "completed_at" field.
func (tu *TaskUpdate) SetCompletedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetCompletedAt(t)
	return tu
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCompletedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetCompletedAt(*t)
	}
	return tu
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (tu *TaskUpdate) ClearCompletedAt() *TaskUpdate {
	tu.mutation.ClearCompletedAt()
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TaskUpdate) SetUpdatedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tu *TaskUpdate) ClearUpdatedAt() *TaskUpdate {
	tu.mutation.ClearUpdatedAt()
	return tu
}

// SetUpdatedBy sets the "updated_by" field.
func (tu *TaskUpdate) SetUpdatedBy(u uint32) *TaskUpdate {
	tu.mutation.ResetUpdatedBy()
	tu.mutation.SetUpdatedBy(u)
	return tu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableUpdatedBy(u *uint32) *TaskUpdate {
	if u != nil {
		tu.SetUpdatedBy(*u)
	}
	return tu
}

// AddUpdatedBy adds u to the "updated_by" field.
func (tu *TaskUpdate) AddUpdatedBy(u int32) *TaskUpdate {
	tu.mutation.AddUpdatedBy(u)
	return tu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (tu *TaskUpdate) ClearUpdatedBy() *TaskUpdate {
	tu.mutation.ClearUpdatedBy()
	return tu
}

// SetStatusUpdatedAt sets the "status_updated_at" field.
func (tu *TaskUpdate) SetStatusUpdatedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetStatusUpdatedAt(t)
	return tu
}

// SetNillableStatusUpdatedAt sets the "status_updated_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableStatusUpdatedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetStatusUpdatedAt(*t)
	}
	return tu
}

// ClearStatusUpdatedAt clears the value of the "status_updated_at" field.
func (tu *TaskUpdate) ClearStatusUpdatedAt() *TaskUpdate {
	tu.mutation.ClearStatusUpdatedAt()
	return tu
}

// SetDeadline sets the "deadline" field.
func (tu *TaskUpdate) SetDeadline(t time.Time) *TaskUpdate {
	tu.mutation.SetDeadline(t)
	return tu
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDeadline(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetDeadline(*t)
	}
	return tu
}

// ClearDeadline clears the value of the "deadline" field.
func (tu *TaskUpdate) ClearDeadline() *TaskUpdate {
	tu.mutation.ClearDeadline()
	return tu
}

// SetDeletedAt sets the "deleted_at" field.
func (tu *TaskUpdate) SetDeletedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetDeletedAt(t)
	return tu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDeletedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetDeletedAt(*t)
	}
	return tu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tu *TaskUpdate) ClearDeletedAt() *TaskUpdate {
	tu.mutation.ClearDeletedAt()
	return tu
}

// SetDeletedBy sets the "deleted_by" field.
func (tu *TaskUpdate) SetDeletedBy(u uint32) *TaskUpdate {
	tu.mutation.ResetDeletedBy()
	tu.mutation.SetDeletedBy(u)
	return tu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDeletedBy(u *uint32) *TaskUpdate {
	if u != nil {
		tu.SetDeletedBy(*u)
	}
	return tu
}

// AddDeletedBy adds u to the "deleted_by" field.
func (tu *TaskUpdate) AddDeletedBy(u int32) *TaskUpdate {
	tu.mutation.AddDeletedBy(u)
	return tu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (tu *TaskUpdate) ClearDeletedBy() *TaskUpdate {
	tu.mutation.ClearDeletedBy()
	return tu
}

// SetDescription sets the "description" field.
func (tu *TaskUpdate) SetDescription(s string) *TaskUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDescription(s *string) *TaskUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// ClearDescription clears the value of the "description" field.
func (tu *TaskUpdate) ClearDescription() *TaskUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// SetTestplanID sets the "testplan_id" field.
func (tu *TaskUpdate) SetTestplanID(i int64) *TaskUpdate {
	tu.mutation.ResetTestplanID()
	tu.mutation.SetTestplanID(i)
	return tu
}

// SetNillableTestplanID sets the "testplan_id" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableTestplanID(i *int64) *TaskUpdate {
	if i != nil {
		tu.SetTestplanID(*i)
	}
	return tu
}

// AddTestplanID adds i to the "testplan_id" field.
func (tu *TaskUpdate) AddTestplanID(i int64) *TaskUpdate {
	tu.mutation.AddTestplanID(i)
	return tu
}

// ClearTestplanID clears the value of the "testplan_id" field.
func (tu *TaskUpdate) ClearTestplanID() *TaskUpdate {
	tu.mutation.ClearTestplanID()
	return tu
}

// SetTestcaseSuite sets the "testcase_suite" field.
func (tu *TaskUpdate) SetTestcaseSuite(i []int64) *TaskUpdate {
	tu.mutation.SetTestcaseSuite(i)
	return tu
}

// AppendTestcaseSuite appends i to the "testcase_suite" field.
func (tu *TaskUpdate) AppendTestcaseSuite(i []int64) *TaskUpdate {
	tu.mutation.AppendTestcaseSuite(i)
	return tu
}

// ClearTestcaseSuite clears the value of the "testcase_suite" field.
func (tu *TaskUpdate) ClearTestcaseSuite() *TaskUpdate {
	tu.mutation.ClearTestcaseSuite()
	return tu
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks[int, TaskMutation](ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TaskUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok && !tu.mutation.UpdatedAtCleared() {
		v := task.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := task.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Task.name": %w`, err)}
		}
	}
	return nil
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt64))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(task.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Assignee(); ok {
		_spec.SetField(task.FieldAssignee, field.TypeUint32, value)
	}
	if value, ok := tu.mutation.AddedAssignee(); ok {
		_spec.AddField(task.FieldAssignee, field.TypeUint32, value)
	}
	if tu.mutation.AssigneeCleared() {
		_spec.ClearField(task.FieldAssignee, field.TypeUint32)
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.SetField(task.FieldType, field.TypeInt8, value)
	}
	if value, ok := tu.mutation.AddedType(); ok {
		_spec.AddField(task.FieldType, field.TypeInt8, value)
	}
	if value, ok := tu.mutation.Frequency(); ok {
		_spec.SetField(task.FieldFrequency, field.TypeInt8, value)
	}
	if value, ok := tu.mutation.AddedFrequency(); ok {
		_spec.AddField(task.FieldFrequency, field.TypeInt8, value)
	}
	if tu.mutation.FrequencyCleared() {
		_spec.ClearField(task.FieldFrequency, field.TypeInt8)
	}
	if value, ok := tu.mutation.ScheduleTime(); ok {
		_spec.SetField(task.FieldScheduleTime, field.TypeTime, value)
	}
	if tu.mutation.ScheduleTimeCleared() {
		_spec.ClearField(task.FieldScheduleTime, field.TypeTime)
	}
	if value, ok := tu.mutation.Rank(); ok {
		_spec.SetField(task.FieldRank, field.TypeInt8, value)
	}
	if value, ok := tu.mutation.AddedRank(); ok {
		_spec.AddField(task.FieldRank, field.TypeInt8, value)
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := tu.mutation.AddedStatus(); ok {
		_spec.AddField(task.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := tu.mutation.StartTime(); ok {
		_spec.SetField(task.FieldStartTime, field.TypeTime, value)
	}
	if tu.mutation.StartTimeCleared() {
		_spec.ClearField(task.FieldStartTime, field.TypeTime)
	}
	if value, ok := tu.mutation.CompletedAt(); ok {
		_spec.SetField(task.FieldCompletedAt, field.TypeTime, value)
	}
	if tu.mutation.CompletedAtCleared() {
		_spec.ClearField(task.FieldCompletedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(task.FieldUpdatedAt, field.TypeTime, value)
	}
	if tu.mutation.UpdatedAtCleared() {
		_spec.ClearField(task.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.UpdatedBy(); ok {
		_spec.SetField(task.FieldUpdatedBy, field.TypeUint32, value)
	}
	if value, ok := tu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(task.FieldUpdatedBy, field.TypeUint32, value)
	}
	if tu.mutation.UpdatedByCleared() {
		_spec.ClearField(task.FieldUpdatedBy, field.TypeUint32)
	}
	if value, ok := tu.mutation.StatusUpdatedAt(); ok {
		_spec.SetField(task.FieldStatusUpdatedAt, field.TypeTime, value)
	}
	if tu.mutation.StatusUpdatedAtCleared() {
		_spec.ClearField(task.FieldStatusUpdatedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.Deadline(); ok {
		_spec.SetField(task.FieldDeadline, field.TypeTime, value)
	}
	if tu.mutation.DeadlineCleared() {
		_spec.ClearField(task.FieldDeadline, field.TypeTime)
	}
	if value, ok := tu.mutation.DeletedAt(); ok {
		_spec.SetField(task.FieldDeletedAt, field.TypeTime, value)
	}
	if tu.mutation.DeletedAtCleared() {
		_spec.ClearField(task.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.DeletedBy(); ok {
		_spec.SetField(task.FieldDeletedBy, field.TypeUint32, value)
	}
	if value, ok := tu.mutation.AddedDeletedBy(); ok {
		_spec.AddField(task.FieldDeletedBy, field.TypeUint32, value)
	}
	if tu.mutation.DeletedByCleared() {
		_spec.ClearField(task.FieldDeletedBy, field.TypeUint32)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(task.FieldDescription, field.TypeString, value)
	}
	if tu.mutation.DescriptionCleared() {
		_spec.ClearField(task.FieldDescription, field.TypeString)
	}
	if value, ok := tu.mutation.TestplanID(); ok {
		_spec.SetField(task.FieldTestplanID, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedTestplanID(); ok {
		_spec.AddField(task.FieldTestplanID, field.TypeInt64, value)
	}
	if tu.mutation.TestplanIDCleared() {
		_spec.ClearField(task.FieldTestplanID, field.TypeInt64)
	}
	if value, ok := tu.mutation.TestcaseSuite(); ok {
		_spec.SetField(task.FieldTestcaseSuite, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.AppendedTestcaseSuite(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, task.FieldTestcaseSuite, value)
		})
	}
	if tu.mutation.TestcaseSuiteCleared() {
		_spec.ClearField(task.FieldTestcaseSuite, field.TypeJSON)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetName sets the "name" field.
func (tuo *TaskUpdateOne) SetName(s string) *TaskUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetAssignee sets the "assignee" field.
func (tuo *TaskUpdateOne) SetAssignee(u uint32) *TaskUpdateOne {
	tuo.mutation.ResetAssignee()
	tuo.mutation.SetAssignee(u)
	return tuo
}

// SetNillableAssignee sets the "assignee" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableAssignee(u *uint32) *TaskUpdateOne {
	if u != nil {
		tuo.SetAssignee(*u)
	}
	return tuo
}

// AddAssignee adds u to the "assignee" field.
func (tuo *TaskUpdateOne) AddAssignee(u int32) *TaskUpdateOne {
	tuo.mutation.AddAssignee(u)
	return tuo
}

// ClearAssignee clears the value of the "assignee" field.
func (tuo *TaskUpdateOne) ClearAssignee() *TaskUpdateOne {
	tuo.mutation.ClearAssignee()
	return tuo
}

// SetType sets the "type" field.
func (tuo *TaskUpdateOne) SetType(i int8) *TaskUpdateOne {
	tuo.mutation.ResetType()
	tuo.mutation.SetType(i)
	return tuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableType(i *int8) *TaskUpdateOne {
	if i != nil {
		tuo.SetType(*i)
	}
	return tuo
}

// AddType adds i to the "type" field.
func (tuo *TaskUpdateOne) AddType(i int8) *TaskUpdateOne {
	tuo.mutation.AddType(i)
	return tuo
}

// SetFrequency sets the "frequency" field.
func (tuo *TaskUpdateOne) SetFrequency(i int8) *TaskUpdateOne {
	tuo.mutation.ResetFrequency()
	tuo.mutation.SetFrequency(i)
	return tuo
}

// SetNillableFrequency sets the "frequency" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableFrequency(i *int8) *TaskUpdateOne {
	if i != nil {
		tuo.SetFrequency(*i)
	}
	return tuo
}

// AddFrequency adds i to the "frequency" field.
func (tuo *TaskUpdateOne) AddFrequency(i int8) *TaskUpdateOne {
	tuo.mutation.AddFrequency(i)
	return tuo
}

// ClearFrequency clears the value of the "frequency" field.
func (tuo *TaskUpdateOne) ClearFrequency() *TaskUpdateOne {
	tuo.mutation.ClearFrequency()
	return tuo
}

// SetScheduleTime sets the "schedule_time" field.
func (tuo *TaskUpdateOne) SetScheduleTime(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetScheduleTime(t)
	return tuo
}

// SetNillableScheduleTime sets the "schedule_time" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableScheduleTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetScheduleTime(*t)
	}
	return tuo
}

// ClearScheduleTime clears the value of the "schedule_time" field.
func (tuo *TaskUpdateOne) ClearScheduleTime() *TaskUpdateOne {
	tuo.mutation.ClearScheduleTime()
	return tuo
}

// SetRank sets the "rank" field.
func (tuo *TaskUpdateOne) SetRank(i int8) *TaskUpdateOne {
	tuo.mutation.ResetRank()
	tuo.mutation.SetRank(i)
	return tuo
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableRank(i *int8) *TaskUpdateOne {
	if i != nil {
		tuo.SetRank(*i)
	}
	return tuo
}

// AddRank adds i to the "rank" field.
func (tuo *TaskUpdateOne) AddRank(i int8) *TaskUpdateOne {
	tuo.mutation.AddRank(i)
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TaskUpdateOne) SetStatus(i int8) *TaskUpdateOne {
	tuo.mutation.ResetStatus()
	tuo.mutation.SetStatus(i)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableStatus(i *int8) *TaskUpdateOne {
	if i != nil {
		tuo.SetStatus(*i)
	}
	return tuo
}

// AddStatus adds i to the "status" field.
func (tuo *TaskUpdateOne) AddStatus(i int8) *TaskUpdateOne {
	tuo.mutation.AddStatus(i)
	return tuo
}

// SetStartTime sets the "start_time" field.
func (tuo *TaskUpdateOne) SetStartTime(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetStartTime(t)
	return tuo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableStartTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetStartTime(*t)
	}
	return tuo
}

// ClearStartTime clears the value of the "start_time" field.
func (tuo *TaskUpdateOne) ClearStartTime() *TaskUpdateOne {
	tuo.mutation.ClearStartTime()
	return tuo
}

// SetCompletedAt sets the "completed_at" field.
func (tuo *TaskUpdateOne) SetCompletedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetCompletedAt(t)
	return tuo
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCompletedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetCompletedAt(*t)
	}
	return tuo
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (tuo *TaskUpdateOne) ClearCompletedAt() *TaskUpdateOne {
	tuo.mutation.ClearCompletedAt()
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TaskUpdateOne) SetUpdatedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tuo *TaskUpdateOne) ClearUpdatedAt() *TaskUpdateOne {
	tuo.mutation.ClearUpdatedAt()
	return tuo
}

// SetUpdatedBy sets the "updated_by" field.
func (tuo *TaskUpdateOne) SetUpdatedBy(u uint32) *TaskUpdateOne {
	tuo.mutation.ResetUpdatedBy()
	tuo.mutation.SetUpdatedBy(u)
	return tuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableUpdatedBy(u *uint32) *TaskUpdateOne {
	if u != nil {
		tuo.SetUpdatedBy(*u)
	}
	return tuo
}

// AddUpdatedBy adds u to the "updated_by" field.
func (tuo *TaskUpdateOne) AddUpdatedBy(u int32) *TaskUpdateOne {
	tuo.mutation.AddUpdatedBy(u)
	return tuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (tuo *TaskUpdateOne) ClearUpdatedBy() *TaskUpdateOne {
	tuo.mutation.ClearUpdatedBy()
	return tuo
}

// SetStatusUpdatedAt sets the "status_updated_at" field.
func (tuo *TaskUpdateOne) SetStatusUpdatedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetStatusUpdatedAt(t)
	return tuo
}

// SetNillableStatusUpdatedAt sets the "status_updated_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableStatusUpdatedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetStatusUpdatedAt(*t)
	}
	return tuo
}

// ClearStatusUpdatedAt clears the value of the "status_updated_at" field.
func (tuo *TaskUpdateOne) ClearStatusUpdatedAt() *TaskUpdateOne {
	tuo.mutation.ClearStatusUpdatedAt()
	return tuo
}

// SetDeadline sets the "deadline" field.
func (tuo *TaskUpdateOne) SetDeadline(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetDeadline(t)
	return tuo
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDeadline(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetDeadline(*t)
	}
	return tuo
}

// ClearDeadline clears the value of the "deadline" field.
func (tuo *TaskUpdateOne) ClearDeadline() *TaskUpdateOne {
	tuo.mutation.ClearDeadline()
	return tuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tuo *TaskUpdateOne) SetDeletedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetDeletedAt(t)
	return tuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDeletedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetDeletedAt(*t)
	}
	return tuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tuo *TaskUpdateOne) ClearDeletedAt() *TaskUpdateOne {
	tuo.mutation.ClearDeletedAt()
	return tuo
}

// SetDeletedBy sets the "deleted_by" field.
func (tuo *TaskUpdateOne) SetDeletedBy(u uint32) *TaskUpdateOne {
	tuo.mutation.ResetDeletedBy()
	tuo.mutation.SetDeletedBy(u)
	return tuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDeletedBy(u *uint32) *TaskUpdateOne {
	if u != nil {
		tuo.SetDeletedBy(*u)
	}
	return tuo
}

// AddDeletedBy adds u to the "deleted_by" field.
func (tuo *TaskUpdateOne) AddDeletedBy(u int32) *TaskUpdateOne {
	tuo.mutation.AddDeletedBy(u)
	return tuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (tuo *TaskUpdateOne) ClearDeletedBy() *TaskUpdateOne {
	tuo.mutation.ClearDeletedBy()
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TaskUpdateOne) SetDescription(s string) *TaskUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDescription(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// ClearDescription clears the value of the "description" field.
func (tuo *TaskUpdateOne) ClearDescription() *TaskUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// SetTestplanID sets the "testplan_id" field.
func (tuo *TaskUpdateOne) SetTestplanID(i int64) *TaskUpdateOne {
	tuo.mutation.ResetTestplanID()
	tuo.mutation.SetTestplanID(i)
	return tuo
}

// SetNillableTestplanID sets the "testplan_id" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableTestplanID(i *int64) *TaskUpdateOne {
	if i != nil {
		tuo.SetTestplanID(*i)
	}
	return tuo
}

// AddTestplanID adds i to the "testplan_id" field.
func (tuo *TaskUpdateOne) AddTestplanID(i int64) *TaskUpdateOne {
	tuo.mutation.AddTestplanID(i)
	return tuo
}

// ClearTestplanID clears the value of the "testplan_id" field.
func (tuo *TaskUpdateOne) ClearTestplanID() *TaskUpdateOne {
	tuo.mutation.ClearTestplanID()
	return tuo
}

// SetTestcaseSuite sets the "testcase_suite" field.
func (tuo *TaskUpdateOne) SetTestcaseSuite(i []int64) *TaskUpdateOne {
	tuo.mutation.SetTestcaseSuite(i)
	return tuo
}

// AppendTestcaseSuite appends i to the "testcase_suite" field.
func (tuo *TaskUpdateOne) AppendTestcaseSuite(i []int64) *TaskUpdateOne {
	tuo.mutation.AppendTestcaseSuite(i)
	return tuo
}

// ClearTestcaseSuite clears the value of the "testcase_suite" field.
func (tuo *TaskUpdateOne) ClearTestcaseSuite() *TaskUpdateOne {
	tuo.mutation.ClearTestcaseSuite()
	return tuo
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tuo *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	tuo.defaults()
	return withHooks[*Task, TaskMutation](ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TaskUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok && !tuo.mutation.UpdatedAtCleared() {
		v := task.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := task.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Task.name": %w`, err)}
		}
	}
	return nil
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt64))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(task.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Assignee(); ok {
		_spec.SetField(task.FieldAssignee, field.TypeUint32, value)
	}
	if value, ok := tuo.mutation.AddedAssignee(); ok {
		_spec.AddField(task.FieldAssignee, field.TypeUint32, value)
	}
	if tuo.mutation.AssigneeCleared() {
		_spec.ClearField(task.FieldAssignee, field.TypeUint32)
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.SetField(task.FieldType, field.TypeInt8, value)
	}
	if value, ok := tuo.mutation.AddedType(); ok {
		_spec.AddField(task.FieldType, field.TypeInt8, value)
	}
	if value, ok := tuo.mutation.Frequency(); ok {
		_spec.SetField(task.FieldFrequency, field.TypeInt8, value)
	}
	if value, ok := tuo.mutation.AddedFrequency(); ok {
		_spec.AddField(task.FieldFrequency, field.TypeInt8, value)
	}
	if tuo.mutation.FrequencyCleared() {
		_spec.ClearField(task.FieldFrequency, field.TypeInt8)
	}
	if value, ok := tuo.mutation.ScheduleTime(); ok {
		_spec.SetField(task.FieldScheduleTime, field.TypeTime, value)
	}
	if tuo.mutation.ScheduleTimeCleared() {
		_spec.ClearField(task.FieldScheduleTime, field.TypeTime)
	}
	if value, ok := tuo.mutation.Rank(); ok {
		_spec.SetField(task.FieldRank, field.TypeInt8, value)
	}
	if value, ok := tuo.mutation.AddedRank(); ok {
		_spec.AddField(task.FieldRank, field.TypeInt8, value)
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := tuo.mutation.AddedStatus(); ok {
		_spec.AddField(task.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := tuo.mutation.StartTime(); ok {
		_spec.SetField(task.FieldStartTime, field.TypeTime, value)
	}
	if tuo.mutation.StartTimeCleared() {
		_spec.ClearField(task.FieldStartTime, field.TypeTime)
	}
	if value, ok := tuo.mutation.CompletedAt(); ok {
		_spec.SetField(task.FieldCompletedAt, field.TypeTime, value)
	}
	if tuo.mutation.CompletedAtCleared() {
		_spec.ClearField(task.FieldCompletedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(task.FieldUpdatedAt, field.TypeTime, value)
	}
	if tuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(task.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.UpdatedBy(); ok {
		_spec.SetField(task.FieldUpdatedBy, field.TypeUint32, value)
	}
	if value, ok := tuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(task.FieldUpdatedBy, field.TypeUint32, value)
	}
	if tuo.mutation.UpdatedByCleared() {
		_spec.ClearField(task.FieldUpdatedBy, field.TypeUint32)
	}
	if value, ok := tuo.mutation.StatusUpdatedAt(); ok {
		_spec.SetField(task.FieldStatusUpdatedAt, field.TypeTime, value)
	}
	if tuo.mutation.StatusUpdatedAtCleared() {
		_spec.ClearField(task.FieldStatusUpdatedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.Deadline(); ok {
		_spec.SetField(task.FieldDeadline, field.TypeTime, value)
	}
	if tuo.mutation.DeadlineCleared() {
		_spec.ClearField(task.FieldDeadline, field.TypeTime)
	}
	if value, ok := tuo.mutation.DeletedAt(); ok {
		_spec.SetField(task.FieldDeletedAt, field.TypeTime, value)
	}
	if tuo.mutation.DeletedAtCleared() {
		_spec.ClearField(task.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.DeletedBy(); ok {
		_spec.SetField(task.FieldDeletedBy, field.TypeUint32, value)
	}
	if value, ok := tuo.mutation.AddedDeletedBy(); ok {
		_spec.AddField(task.FieldDeletedBy, field.TypeUint32, value)
	}
	if tuo.mutation.DeletedByCleared() {
		_spec.ClearField(task.FieldDeletedBy, field.TypeUint32)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(task.FieldDescription, field.TypeString, value)
	}
	if tuo.mutation.DescriptionCleared() {
		_spec.ClearField(task.FieldDescription, field.TypeString)
	}
	if value, ok := tuo.mutation.TestplanID(); ok {
		_spec.SetField(task.FieldTestplanID, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedTestplanID(); ok {
		_spec.AddField(task.FieldTestplanID, field.TypeInt64, value)
	}
	if tuo.mutation.TestplanIDCleared() {
		_spec.ClearField(task.FieldTestplanID, field.TypeInt64)
	}
	if value, ok := tuo.mutation.TestcaseSuite(); ok {
		_spec.SetField(task.FieldTestcaseSuite, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.AppendedTestcaseSuite(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, task.FieldTestcaseSuite, value)
		})
	}
	if tuo.mutation.TestcaseSuiteCleared() {
		_spec.ClearField(task.FieldTestcaseSuite, field.TypeJSON)
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
