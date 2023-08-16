// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"galileo/ent/predicate"
	"galileo/ent/testcasesuite"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// TestcaseSuiteUpdate is the builder for updating TestcaseSuite entities.
type TestcaseSuiteUpdate struct {
	config
	hooks    []Hook
	mutation *TestcaseSuiteMutation
}

// Where appends a list predicates to the TestcaseSuiteUpdate builder.
func (tsu *TestcaseSuiteUpdate) Where(ps ...predicate.TestcaseSuite) *TestcaseSuiteUpdate {
	tsu.mutation.Where(ps...)
	return tsu
}

// SetName sets the "name" field.
func (tsu *TestcaseSuiteUpdate) SetName(s string) *TestcaseSuiteUpdate {
	tsu.mutation.SetName(s)
	return tsu
}

// SetUpdatedAt sets the "updated_at" field.
func (tsu *TestcaseSuiteUpdate) SetUpdatedAt(t time.Time) *TestcaseSuiteUpdate {
	tsu.mutation.SetUpdatedAt(t)
	return tsu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tsu *TestcaseSuiteUpdate) ClearUpdatedAt() *TestcaseSuiteUpdate {
	tsu.mutation.ClearUpdatedAt()
	return tsu
}

// SetUpdatedBy sets the "updated_by" field.
func (tsu *TestcaseSuiteUpdate) SetUpdatedBy(u uint32) *TestcaseSuiteUpdate {
	tsu.mutation.ResetUpdatedBy()
	tsu.mutation.SetUpdatedBy(u)
	return tsu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tsu *TestcaseSuiteUpdate) SetNillableUpdatedBy(u *uint32) *TestcaseSuiteUpdate {
	if u != nil {
		tsu.SetUpdatedBy(*u)
	}
	return tsu
}

// AddUpdatedBy adds u to the "updated_by" field.
func (tsu *TestcaseSuiteUpdate) AddUpdatedBy(u int32) *TestcaseSuiteUpdate {
	tsu.mutation.AddUpdatedBy(u)
	return tsu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (tsu *TestcaseSuiteUpdate) ClearUpdatedBy() *TestcaseSuiteUpdate {
	tsu.mutation.ClearUpdatedBy()
	return tsu
}

// SetTestcases sets the "testcases" field.
func (tsu *TestcaseSuiteUpdate) SetTestcases(i []int64) *TestcaseSuiteUpdate {
	tsu.mutation.SetTestcases(i)
	return tsu
}

// AppendTestcases appends i to the "testcases" field.
func (tsu *TestcaseSuiteUpdate) AppendTestcases(i []int64) *TestcaseSuiteUpdate {
	tsu.mutation.AppendTestcases(i)
	return tsu
}

// ClearTestcases clears the value of the "testcases" field.
func (tsu *TestcaseSuiteUpdate) ClearTestcases() *TestcaseSuiteUpdate {
	tsu.mutation.ClearTestcases()
	return tsu
}

// Mutation returns the TestcaseSuiteMutation object of the builder.
func (tsu *TestcaseSuiteUpdate) Mutation() *TestcaseSuiteMutation {
	return tsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tsu *TestcaseSuiteUpdate) Save(ctx context.Context) (int, error) {
	tsu.defaults()
	return withHooks[int, TestcaseSuiteMutation](ctx, tsu.sqlSave, tsu.mutation, tsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tsu *TestcaseSuiteUpdate) SaveX(ctx context.Context) int {
	affected, err := tsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tsu *TestcaseSuiteUpdate) Exec(ctx context.Context) error {
	_, err := tsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsu *TestcaseSuiteUpdate) ExecX(ctx context.Context) {
	if err := tsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tsu *TestcaseSuiteUpdate) defaults() {
	if _, ok := tsu.mutation.UpdatedAt(); !ok && !tsu.mutation.UpdatedAtCleared() {
		v := testcasesuite.UpdateDefaultUpdatedAt()
		tsu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tsu *TestcaseSuiteUpdate) check() error {
	if v, ok := tsu.mutation.Name(); ok {
		if err := testcasesuite.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TestcaseSuite.name": %w`, err)}
		}
	}
	return nil
}

func (tsu *TestcaseSuiteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(testcasesuite.Table, testcasesuite.Columns, sqlgraph.NewFieldSpec(testcasesuite.FieldID, field.TypeInt64))
	if ps := tsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tsu.mutation.Name(); ok {
		_spec.SetField(testcasesuite.FieldName, field.TypeString, value)
	}
	if value, ok := tsu.mutation.UpdatedAt(); ok {
		_spec.SetField(testcasesuite.FieldUpdatedAt, field.TypeTime, value)
	}
	if tsu.mutation.UpdatedAtCleared() {
		_spec.ClearField(testcasesuite.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := tsu.mutation.UpdatedBy(); ok {
		_spec.SetField(testcasesuite.FieldUpdatedBy, field.TypeUint32, value)
	}
	if value, ok := tsu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(testcasesuite.FieldUpdatedBy, field.TypeUint32, value)
	}
	if tsu.mutation.UpdatedByCleared() {
		_spec.ClearField(testcasesuite.FieldUpdatedBy, field.TypeUint32)
	}
	if value, ok := tsu.mutation.Testcases(); ok {
		_spec.SetField(testcasesuite.FieldTestcases, field.TypeJSON, value)
	}
	if value, ok := tsu.mutation.AppendedTestcases(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, testcasesuite.FieldTestcases, value)
		})
	}
	if tsu.mutation.TestcasesCleared() {
		_spec.ClearField(testcasesuite.FieldTestcases, field.TypeJSON)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{testcasesuite.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tsu.mutation.done = true
	return n, nil
}

// TestcaseSuiteUpdateOne is the builder for updating a single TestcaseSuite entity.
type TestcaseSuiteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TestcaseSuiteMutation
}

// SetName sets the "name" field.
func (tsuo *TestcaseSuiteUpdateOne) SetName(s string) *TestcaseSuiteUpdateOne {
	tsuo.mutation.SetName(s)
	return tsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tsuo *TestcaseSuiteUpdateOne) SetUpdatedAt(t time.Time) *TestcaseSuiteUpdateOne {
	tsuo.mutation.SetUpdatedAt(t)
	return tsuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tsuo *TestcaseSuiteUpdateOne) ClearUpdatedAt() *TestcaseSuiteUpdateOne {
	tsuo.mutation.ClearUpdatedAt()
	return tsuo
}

// SetUpdatedBy sets the "updated_by" field.
func (tsuo *TestcaseSuiteUpdateOne) SetUpdatedBy(u uint32) *TestcaseSuiteUpdateOne {
	tsuo.mutation.ResetUpdatedBy()
	tsuo.mutation.SetUpdatedBy(u)
	return tsuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tsuo *TestcaseSuiteUpdateOne) SetNillableUpdatedBy(u *uint32) *TestcaseSuiteUpdateOne {
	if u != nil {
		tsuo.SetUpdatedBy(*u)
	}
	return tsuo
}

// AddUpdatedBy adds u to the "updated_by" field.
func (tsuo *TestcaseSuiteUpdateOne) AddUpdatedBy(u int32) *TestcaseSuiteUpdateOne {
	tsuo.mutation.AddUpdatedBy(u)
	return tsuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (tsuo *TestcaseSuiteUpdateOne) ClearUpdatedBy() *TestcaseSuiteUpdateOne {
	tsuo.mutation.ClearUpdatedBy()
	return tsuo
}

// SetTestcases sets the "testcases" field.
func (tsuo *TestcaseSuiteUpdateOne) SetTestcases(i []int64) *TestcaseSuiteUpdateOne {
	tsuo.mutation.SetTestcases(i)
	return tsuo
}

// AppendTestcases appends i to the "testcases" field.
func (tsuo *TestcaseSuiteUpdateOne) AppendTestcases(i []int64) *TestcaseSuiteUpdateOne {
	tsuo.mutation.AppendTestcases(i)
	return tsuo
}

// ClearTestcases clears the value of the "testcases" field.
func (tsuo *TestcaseSuiteUpdateOne) ClearTestcases() *TestcaseSuiteUpdateOne {
	tsuo.mutation.ClearTestcases()
	return tsuo
}

// Mutation returns the TestcaseSuiteMutation object of the builder.
func (tsuo *TestcaseSuiteUpdateOne) Mutation() *TestcaseSuiteMutation {
	return tsuo.mutation
}

// Where appends a list predicates to the TestcaseSuiteUpdate builder.
func (tsuo *TestcaseSuiteUpdateOne) Where(ps ...predicate.TestcaseSuite) *TestcaseSuiteUpdateOne {
	tsuo.mutation.Where(ps...)
	return tsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tsuo *TestcaseSuiteUpdateOne) Select(field string, fields ...string) *TestcaseSuiteUpdateOne {
	tsuo.fields = append([]string{field}, fields...)
	return tsuo
}

// Save executes the query and returns the updated TestcaseSuite entity.
func (tsuo *TestcaseSuiteUpdateOne) Save(ctx context.Context) (*TestcaseSuite, error) {
	tsuo.defaults()
	return withHooks[*TestcaseSuite, TestcaseSuiteMutation](ctx, tsuo.sqlSave, tsuo.mutation, tsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tsuo *TestcaseSuiteUpdateOne) SaveX(ctx context.Context) *TestcaseSuite {
	node, err := tsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tsuo *TestcaseSuiteUpdateOne) Exec(ctx context.Context) error {
	_, err := tsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsuo *TestcaseSuiteUpdateOne) ExecX(ctx context.Context) {
	if err := tsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tsuo *TestcaseSuiteUpdateOne) defaults() {
	if _, ok := tsuo.mutation.UpdatedAt(); !ok && !tsuo.mutation.UpdatedAtCleared() {
		v := testcasesuite.UpdateDefaultUpdatedAt()
		tsuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tsuo *TestcaseSuiteUpdateOne) check() error {
	if v, ok := tsuo.mutation.Name(); ok {
		if err := testcasesuite.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TestcaseSuite.name": %w`, err)}
		}
	}
	return nil
}

func (tsuo *TestcaseSuiteUpdateOne) sqlSave(ctx context.Context) (_node *TestcaseSuite, err error) {
	if err := tsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(testcasesuite.Table, testcasesuite.Columns, sqlgraph.NewFieldSpec(testcasesuite.FieldID, field.TypeInt64))
	id, ok := tsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TestcaseSuite.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testcasesuite.FieldID)
		for _, f := range fields {
			if !testcasesuite.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != testcasesuite.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tsuo.mutation.Name(); ok {
		_spec.SetField(testcasesuite.FieldName, field.TypeString, value)
	}
	if value, ok := tsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(testcasesuite.FieldUpdatedAt, field.TypeTime, value)
	}
	if tsuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(testcasesuite.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := tsuo.mutation.UpdatedBy(); ok {
		_spec.SetField(testcasesuite.FieldUpdatedBy, field.TypeUint32, value)
	}
	if value, ok := tsuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(testcasesuite.FieldUpdatedBy, field.TypeUint32, value)
	}
	if tsuo.mutation.UpdatedByCleared() {
		_spec.ClearField(testcasesuite.FieldUpdatedBy, field.TypeUint32)
	}
	if value, ok := tsuo.mutation.Testcases(); ok {
		_spec.SetField(testcasesuite.FieldTestcases, field.TypeJSON, value)
	}
	if value, ok := tsuo.mutation.AppendedTestcases(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, testcasesuite.FieldTestcases, value)
		})
	}
	if tsuo.mutation.TestcasesCleared() {
		_spec.ClearField(testcasesuite.FieldTestcases, field.TypeJSON)
	}
	_node = &TestcaseSuite{config: tsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{testcasesuite.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tsuo.mutation.done = true
	return _node, nil
}
