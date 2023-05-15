// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"galileo/ent/api"
	"galileo/ent/apistatistics"
	"galileo/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApiStatisticsUpdate is the builder for updating ApiStatistics entities.
type ApiStatisticsUpdate struct {
	config
	hooks    []Hook
	mutation *ApiStatisticsMutation
}

// Where appends a list predicates to the ApiStatisticsUpdate builder.
func (asu *ApiStatisticsUpdate) Where(ps ...predicate.ApiStatistics) *ApiStatisticsUpdate {
	asu.mutation.Where(ps...)
	return asu
}

// SetCallCount sets the "call_count" field.
func (asu *ApiStatisticsUpdate) SetCallCount(i int64) *ApiStatisticsUpdate {
	asu.mutation.ResetCallCount()
	asu.mutation.SetCallCount(i)
	return asu
}

// AddCallCount adds i to the "call_count" field.
func (asu *ApiStatisticsUpdate) AddCallCount(i int64) *ApiStatisticsUpdate {
	asu.mutation.AddCallCount(i)
	return asu
}

// SetSuccessCount sets the "success_count" field.
func (asu *ApiStatisticsUpdate) SetSuccessCount(i int64) *ApiStatisticsUpdate {
	asu.mutation.ResetSuccessCount()
	asu.mutation.SetSuccessCount(i)
	return asu
}

// AddSuccessCount adds i to the "success_count" field.
func (asu *ApiStatisticsUpdate) AddSuccessCount(i int64) *ApiStatisticsUpdate {
	asu.mutation.AddSuccessCount(i)
	return asu
}

// SetFailureCount sets the "failure_count" field.
func (asu *ApiStatisticsUpdate) SetFailureCount(i int64) *ApiStatisticsUpdate {
	asu.mutation.ResetFailureCount()
	asu.mutation.SetFailureCount(i)
	return asu
}

// AddFailureCount adds i to the "failure_count" field.
func (asu *ApiStatisticsUpdate) AddFailureCount(i int64) *ApiStatisticsUpdate {
	asu.mutation.AddFailureCount(i)
	return asu
}

// SetAvgResponseTime sets the "avg_response_time" field.
func (asu *ApiStatisticsUpdate) SetAvgResponseTime(f float64) *ApiStatisticsUpdate {
	asu.mutation.ResetAvgResponseTime()
	asu.mutation.SetAvgResponseTime(f)
	return asu
}

// AddAvgResponseTime adds f to the "avg_response_time" field.
func (asu *ApiStatisticsUpdate) AddAvgResponseTime(f float64) *ApiStatisticsUpdate {
	asu.mutation.AddAvgResponseTime(f)
	return asu
}

// SetMaxResponseTime sets the "max_response_time" field.
func (asu *ApiStatisticsUpdate) SetMaxResponseTime(f float64) *ApiStatisticsUpdate {
	asu.mutation.ResetMaxResponseTime()
	asu.mutation.SetMaxResponseTime(f)
	return asu
}

// AddMaxResponseTime adds f to the "max_response_time" field.
func (asu *ApiStatisticsUpdate) AddMaxResponseTime(f float64) *ApiStatisticsUpdate {
	asu.mutation.AddMaxResponseTime(f)
	return asu
}

// SetMinResponseTime sets the "min_response_time" field.
func (asu *ApiStatisticsUpdate) SetMinResponseTime(f float64) *ApiStatisticsUpdate {
	asu.mutation.ResetMinResponseTime()
	asu.mutation.SetMinResponseTime(f)
	return asu
}

// AddMinResponseTime adds f to the "min_response_time" field.
func (asu *ApiStatisticsUpdate) AddMinResponseTime(f float64) *ApiStatisticsUpdate {
	asu.mutation.AddMinResponseTime(f)
	return asu
}

// SetAvgTraffic sets the "avg_traffic" field.
func (asu *ApiStatisticsUpdate) SetAvgTraffic(f float64) *ApiStatisticsUpdate {
	asu.mutation.ResetAvgTraffic()
	asu.mutation.SetAvgTraffic(f)
	return asu
}

// AddAvgTraffic adds f to the "avg_traffic" field.
func (asu *ApiStatisticsUpdate) AddAvgTraffic(f float64) *ApiStatisticsUpdate {
	asu.mutation.AddAvgTraffic(f)
	return asu
}

// SetMaxTraffic sets the "max_traffic" field.
func (asu *ApiStatisticsUpdate) SetMaxTraffic(f float64) *ApiStatisticsUpdate {
	asu.mutation.ResetMaxTraffic()
	asu.mutation.SetMaxTraffic(f)
	return asu
}

// AddMaxTraffic adds f to the "max_traffic" field.
func (asu *ApiStatisticsUpdate) AddMaxTraffic(f float64) *ApiStatisticsUpdate {
	asu.mutation.AddMaxTraffic(f)
	return asu
}

// SetMinTraffic sets the "min_traffic" field.
func (asu *ApiStatisticsUpdate) SetMinTraffic(f float64) *ApiStatisticsUpdate {
	asu.mutation.ResetMinTraffic()
	asu.mutation.SetMinTraffic(f)
	return asu
}

// AddMinTraffic adds f to the "min_traffic" field.
func (asu *ApiStatisticsUpdate) AddMinTraffic(f float64) *ApiStatisticsUpdate {
	asu.mutation.AddMinTraffic(f)
	return asu
}

// SetDescription sets the "description" field.
func (asu *ApiStatisticsUpdate) SetDescription(s string) *ApiStatisticsUpdate {
	asu.mutation.SetDescription(s)
	return asu
}

// SetCreatedAt sets the "created_at" field.
func (asu *ApiStatisticsUpdate) SetCreatedAt(t time.Time) *ApiStatisticsUpdate {
	asu.mutation.SetCreatedAt(t)
	return asu
}

// SetUpdateAt sets the "update_at" field.
func (asu *ApiStatisticsUpdate) SetUpdateAt(t time.Time) *ApiStatisticsUpdate {
	asu.mutation.SetUpdateAt(t)
	return asu
}

// SetAPIID sets the "api" edge to the Api entity by ID.
func (asu *ApiStatisticsUpdate) SetAPIID(id int64) *ApiStatisticsUpdate {
	asu.mutation.SetAPIID(id)
	return asu
}

// SetAPI sets the "api" edge to the Api entity.
func (asu *ApiStatisticsUpdate) SetAPI(a *Api) *ApiStatisticsUpdate {
	return asu.SetAPIID(a.ID)
}

// Mutation returns the ApiStatisticsMutation object of the builder.
func (asu *ApiStatisticsUpdate) Mutation() *ApiStatisticsMutation {
	return asu.mutation
}

// ClearAPI clears the "api" edge to the Api entity.
func (asu *ApiStatisticsUpdate) ClearAPI() *ApiStatisticsUpdate {
	asu.mutation.ClearAPI()
	return asu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (asu *ApiStatisticsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ApiStatisticsMutation](ctx, asu.sqlSave, asu.mutation, asu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (asu *ApiStatisticsUpdate) SaveX(ctx context.Context) int {
	affected, err := asu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (asu *ApiStatisticsUpdate) Exec(ctx context.Context) error {
	_, err := asu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asu *ApiStatisticsUpdate) ExecX(ctx context.Context) {
	if err := asu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asu *ApiStatisticsUpdate) check() error {
	if _, ok := asu.mutation.APIID(); asu.mutation.APICleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ApiStatistics.api"`)
	}
	return nil
}

func (asu *ApiStatisticsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := asu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(apistatistics.Table, apistatistics.Columns, sqlgraph.NewFieldSpec(apistatistics.FieldID, field.TypeInt64))
	if ps := asu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := asu.mutation.CallCount(); ok {
		_spec.SetField(apistatistics.FieldCallCount, field.TypeInt64, value)
	}
	if value, ok := asu.mutation.AddedCallCount(); ok {
		_spec.AddField(apistatistics.FieldCallCount, field.TypeInt64, value)
	}
	if value, ok := asu.mutation.SuccessCount(); ok {
		_spec.SetField(apistatistics.FieldSuccessCount, field.TypeInt64, value)
	}
	if value, ok := asu.mutation.AddedSuccessCount(); ok {
		_spec.AddField(apistatistics.FieldSuccessCount, field.TypeInt64, value)
	}
	if value, ok := asu.mutation.FailureCount(); ok {
		_spec.SetField(apistatistics.FieldFailureCount, field.TypeInt64, value)
	}
	if value, ok := asu.mutation.AddedFailureCount(); ok {
		_spec.AddField(apistatistics.FieldFailureCount, field.TypeInt64, value)
	}
	if value, ok := asu.mutation.AvgResponseTime(); ok {
		_spec.SetField(apistatistics.FieldAvgResponseTime, field.TypeFloat64, value)
	}
	if value, ok := asu.mutation.AddedAvgResponseTime(); ok {
		_spec.AddField(apistatistics.FieldAvgResponseTime, field.TypeFloat64, value)
	}
	if value, ok := asu.mutation.MaxResponseTime(); ok {
		_spec.SetField(apistatistics.FieldMaxResponseTime, field.TypeFloat64, value)
	}
	if value, ok := asu.mutation.AddedMaxResponseTime(); ok {
		_spec.AddField(apistatistics.FieldMaxResponseTime, field.TypeFloat64, value)
	}
	if value, ok := asu.mutation.MinResponseTime(); ok {
		_spec.SetField(apistatistics.FieldMinResponseTime, field.TypeFloat64, value)
	}
	if value, ok := asu.mutation.AddedMinResponseTime(); ok {
		_spec.AddField(apistatistics.FieldMinResponseTime, field.TypeFloat64, value)
	}
	if value, ok := asu.mutation.AvgTraffic(); ok {
		_spec.SetField(apistatistics.FieldAvgTraffic, field.TypeFloat64, value)
	}
	if value, ok := asu.mutation.AddedAvgTraffic(); ok {
		_spec.AddField(apistatistics.FieldAvgTraffic, field.TypeFloat64, value)
	}
	if value, ok := asu.mutation.MaxTraffic(); ok {
		_spec.SetField(apistatistics.FieldMaxTraffic, field.TypeFloat64, value)
	}
	if value, ok := asu.mutation.AddedMaxTraffic(); ok {
		_spec.AddField(apistatistics.FieldMaxTraffic, field.TypeFloat64, value)
	}
	if value, ok := asu.mutation.MinTraffic(); ok {
		_spec.SetField(apistatistics.FieldMinTraffic, field.TypeFloat64, value)
	}
	if value, ok := asu.mutation.AddedMinTraffic(); ok {
		_spec.AddField(apistatistics.FieldMinTraffic, field.TypeFloat64, value)
	}
	if value, ok := asu.mutation.Description(); ok {
		_spec.SetField(apistatistics.FieldDescription, field.TypeString, value)
	}
	if value, ok := asu.mutation.CreatedAt(); ok {
		_spec.SetField(apistatistics.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := asu.mutation.UpdateAt(); ok {
		_spec.SetField(apistatistics.FieldUpdateAt, field.TypeTime, value)
	}
	if asu.mutation.APICleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   apistatistics.APITable,
			Columns: []string{apistatistics.APIColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(api.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asu.mutation.APIIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   apistatistics.APITable,
			Columns: []string{apistatistics.APIColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(api.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, asu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apistatistics.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	asu.mutation.done = true
	return n, nil
}

// ApiStatisticsUpdateOne is the builder for updating a single ApiStatistics entity.
type ApiStatisticsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApiStatisticsMutation
}

// SetCallCount sets the "call_count" field.
func (asuo *ApiStatisticsUpdateOne) SetCallCount(i int64) *ApiStatisticsUpdateOne {
	asuo.mutation.ResetCallCount()
	asuo.mutation.SetCallCount(i)
	return asuo
}

// AddCallCount adds i to the "call_count" field.
func (asuo *ApiStatisticsUpdateOne) AddCallCount(i int64) *ApiStatisticsUpdateOne {
	asuo.mutation.AddCallCount(i)
	return asuo
}

// SetSuccessCount sets the "success_count" field.
func (asuo *ApiStatisticsUpdateOne) SetSuccessCount(i int64) *ApiStatisticsUpdateOne {
	asuo.mutation.ResetSuccessCount()
	asuo.mutation.SetSuccessCount(i)
	return asuo
}

// AddSuccessCount adds i to the "success_count" field.
func (asuo *ApiStatisticsUpdateOne) AddSuccessCount(i int64) *ApiStatisticsUpdateOne {
	asuo.mutation.AddSuccessCount(i)
	return asuo
}

// SetFailureCount sets the "failure_count" field.
func (asuo *ApiStatisticsUpdateOne) SetFailureCount(i int64) *ApiStatisticsUpdateOne {
	asuo.mutation.ResetFailureCount()
	asuo.mutation.SetFailureCount(i)
	return asuo
}

// AddFailureCount adds i to the "failure_count" field.
func (asuo *ApiStatisticsUpdateOne) AddFailureCount(i int64) *ApiStatisticsUpdateOne {
	asuo.mutation.AddFailureCount(i)
	return asuo
}

// SetAvgResponseTime sets the "avg_response_time" field.
func (asuo *ApiStatisticsUpdateOne) SetAvgResponseTime(f float64) *ApiStatisticsUpdateOne {
	asuo.mutation.ResetAvgResponseTime()
	asuo.mutation.SetAvgResponseTime(f)
	return asuo
}

// AddAvgResponseTime adds f to the "avg_response_time" field.
func (asuo *ApiStatisticsUpdateOne) AddAvgResponseTime(f float64) *ApiStatisticsUpdateOne {
	asuo.mutation.AddAvgResponseTime(f)
	return asuo
}

// SetMaxResponseTime sets the "max_response_time" field.
func (asuo *ApiStatisticsUpdateOne) SetMaxResponseTime(f float64) *ApiStatisticsUpdateOne {
	asuo.mutation.ResetMaxResponseTime()
	asuo.mutation.SetMaxResponseTime(f)
	return asuo
}

// AddMaxResponseTime adds f to the "max_response_time" field.
func (asuo *ApiStatisticsUpdateOne) AddMaxResponseTime(f float64) *ApiStatisticsUpdateOne {
	asuo.mutation.AddMaxResponseTime(f)
	return asuo
}

// SetMinResponseTime sets the "min_response_time" field.
func (asuo *ApiStatisticsUpdateOne) SetMinResponseTime(f float64) *ApiStatisticsUpdateOne {
	asuo.mutation.ResetMinResponseTime()
	asuo.mutation.SetMinResponseTime(f)
	return asuo
}

// AddMinResponseTime adds f to the "min_response_time" field.
func (asuo *ApiStatisticsUpdateOne) AddMinResponseTime(f float64) *ApiStatisticsUpdateOne {
	asuo.mutation.AddMinResponseTime(f)
	return asuo
}

// SetAvgTraffic sets the "avg_traffic" field.
func (asuo *ApiStatisticsUpdateOne) SetAvgTraffic(f float64) *ApiStatisticsUpdateOne {
	asuo.mutation.ResetAvgTraffic()
	asuo.mutation.SetAvgTraffic(f)
	return asuo
}

// AddAvgTraffic adds f to the "avg_traffic" field.
func (asuo *ApiStatisticsUpdateOne) AddAvgTraffic(f float64) *ApiStatisticsUpdateOne {
	asuo.mutation.AddAvgTraffic(f)
	return asuo
}

// SetMaxTraffic sets the "max_traffic" field.
func (asuo *ApiStatisticsUpdateOne) SetMaxTraffic(f float64) *ApiStatisticsUpdateOne {
	asuo.mutation.ResetMaxTraffic()
	asuo.mutation.SetMaxTraffic(f)
	return asuo
}

// AddMaxTraffic adds f to the "max_traffic" field.
func (asuo *ApiStatisticsUpdateOne) AddMaxTraffic(f float64) *ApiStatisticsUpdateOne {
	asuo.mutation.AddMaxTraffic(f)
	return asuo
}

// SetMinTraffic sets the "min_traffic" field.
func (asuo *ApiStatisticsUpdateOne) SetMinTraffic(f float64) *ApiStatisticsUpdateOne {
	asuo.mutation.ResetMinTraffic()
	asuo.mutation.SetMinTraffic(f)
	return asuo
}

// AddMinTraffic adds f to the "min_traffic" field.
func (asuo *ApiStatisticsUpdateOne) AddMinTraffic(f float64) *ApiStatisticsUpdateOne {
	asuo.mutation.AddMinTraffic(f)
	return asuo
}

// SetDescription sets the "description" field.
func (asuo *ApiStatisticsUpdateOne) SetDescription(s string) *ApiStatisticsUpdateOne {
	asuo.mutation.SetDescription(s)
	return asuo
}

// SetCreatedAt sets the "created_at" field.
func (asuo *ApiStatisticsUpdateOne) SetCreatedAt(t time.Time) *ApiStatisticsUpdateOne {
	asuo.mutation.SetCreatedAt(t)
	return asuo
}

// SetUpdateAt sets the "update_at" field.
func (asuo *ApiStatisticsUpdateOne) SetUpdateAt(t time.Time) *ApiStatisticsUpdateOne {
	asuo.mutation.SetUpdateAt(t)
	return asuo
}

// SetAPIID sets the "api" edge to the Api entity by ID.
func (asuo *ApiStatisticsUpdateOne) SetAPIID(id int64) *ApiStatisticsUpdateOne {
	asuo.mutation.SetAPIID(id)
	return asuo
}

// SetAPI sets the "api" edge to the Api entity.
func (asuo *ApiStatisticsUpdateOne) SetAPI(a *Api) *ApiStatisticsUpdateOne {
	return asuo.SetAPIID(a.ID)
}

// Mutation returns the ApiStatisticsMutation object of the builder.
func (asuo *ApiStatisticsUpdateOne) Mutation() *ApiStatisticsMutation {
	return asuo.mutation
}

// ClearAPI clears the "api" edge to the Api entity.
func (asuo *ApiStatisticsUpdateOne) ClearAPI() *ApiStatisticsUpdateOne {
	asuo.mutation.ClearAPI()
	return asuo
}

// Where appends a list predicates to the ApiStatisticsUpdate builder.
func (asuo *ApiStatisticsUpdateOne) Where(ps ...predicate.ApiStatistics) *ApiStatisticsUpdateOne {
	asuo.mutation.Where(ps...)
	return asuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (asuo *ApiStatisticsUpdateOne) Select(field string, fields ...string) *ApiStatisticsUpdateOne {
	asuo.fields = append([]string{field}, fields...)
	return asuo
}

// Save executes the query and returns the updated ApiStatistics entity.
func (asuo *ApiStatisticsUpdateOne) Save(ctx context.Context) (*ApiStatistics, error) {
	return withHooks[*ApiStatistics, ApiStatisticsMutation](ctx, asuo.sqlSave, asuo.mutation, asuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (asuo *ApiStatisticsUpdateOne) SaveX(ctx context.Context) *ApiStatistics {
	node, err := asuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (asuo *ApiStatisticsUpdateOne) Exec(ctx context.Context) error {
	_, err := asuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asuo *ApiStatisticsUpdateOne) ExecX(ctx context.Context) {
	if err := asuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asuo *ApiStatisticsUpdateOne) check() error {
	if _, ok := asuo.mutation.APIID(); asuo.mutation.APICleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ApiStatistics.api"`)
	}
	return nil
}

func (asuo *ApiStatisticsUpdateOne) sqlSave(ctx context.Context) (_node *ApiStatistics, err error) {
	if err := asuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(apistatistics.Table, apistatistics.Columns, sqlgraph.NewFieldSpec(apistatistics.FieldID, field.TypeInt64))
	id, ok := asuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ApiStatistics.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := asuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apistatistics.FieldID)
		for _, f := range fields {
			if !apistatistics.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apistatistics.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := asuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := asuo.mutation.CallCount(); ok {
		_spec.SetField(apistatistics.FieldCallCount, field.TypeInt64, value)
	}
	if value, ok := asuo.mutation.AddedCallCount(); ok {
		_spec.AddField(apistatistics.FieldCallCount, field.TypeInt64, value)
	}
	if value, ok := asuo.mutation.SuccessCount(); ok {
		_spec.SetField(apistatistics.FieldSuccessCount, field.TypeInt64, value)
	}
	if value, ok := asuo.mutation.AddedSuccessCount(); ok {
		_spec.AddField(apistatistics.FieldSuccessCount, field.TypeInt64, value)
	}
	if value, ok := asuo.mutation.FailureCount(); ok {
		_spec.SetField(apistatistics.FieldFailureCount, field.TypeInt64, value)
	}
	if value, ok := asuo.mutation.AddedFailureCount(); ok {
		_spec.AddField(apistatistics.FieldFailureCount, field.TypeInt64, value)
	}
	if value, ok := asuo.mutation.AvgResponseTime(); ok {
		_spec.SetField(apistatistics.FieldAvgResponseTime, field.TypeFloat64, value)
	}
	if value, ok := asuo.mutation.AddedAvgResponseTime(); ok {
		_spec.AddField(apistatistics.FieldAvgResponseTime, field.TypeFloat64, value)
	}
	if value, ok := asuo.mutation.MaxResponseTime(); ok {
		_spec.SetField(apistatistics.FieldMaxResponseTime, field.TypeFloat64, value)
	}
	if value, ok := asuo.mutation.AddedMaxResponseTime(); ok {
		_spec.AddField(apistatistics.FieldMaxResponseTime, field.TypeFloat64, value)
	}
	if value, ok := asuo.mutation.MinResponseTime(); ok {
		_spec.SetField(apistatistics.FieldMinResponseTime, field.TypeFloat64, value)
	}
	if value, ok := asuo.mutation.AddedMinResponseTime(); ok {
		_spec.AddField(apistatistics.FieldMinResponseTime, field.TypeFloat64, value)
	}
	if value, ok := asuo.mutation.AvgTraffic(); ok {
		_spec.SetField(apistatistics.FieldAvgTraffic, field.TypeFloat64, value)
	}
	if value, ok := asuo.mutation.AddedAvgTraffic(); ok {
		_spec.AddField(apistatistics.FieldAvgTraffic, field.TypeFloat64, value)
	}
	if value, ok := asuo.mutation.MaxTraffic(); ok {
		_spec.SetField(apistatistics.FieldMaxTraffic, field.TypeFloat64, value)
	}
	if value, ok := asuo.mutation.AddedMaxTraffic(); ok {
		_spec.AddField(apistatistics.FieldMaxTraffic, field.TypeFloat64, value)
	}
	if value, ok := asuo.mutation.MinTraffic(); ok {
		_spec.SetField(apistatistics.FieldMinTraffic, field.TypeFloat64, value)
	}
	if value, ok := asuo.mutation.AddedMinTraffic(); ok {
		_spec.AddField(apistatistics.FieldMinTraffic, field.TypeFloat64, value)
	}
	if value, ok := asuo.mutation.Description(); ok {
		_spec.SetField(apistatistics.FieldDescription, field.TypeString, value)
	}
	if value, ok := asuo.mutation.CreatedAt(); ok {
		_spec.SetField(apistatistics.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := asuo.mutation.UpdateAt(); ok {
		_spec.SetField(apistatistics.FieldUpdateAt, field.TypeTime, value)
	}
	if asuo.mutation.APICleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   apistatistics.APITable,
			Columns: []string{apistatistics.APIColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(api.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asuo.mutation.APIIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   apistatistics.APITable,
			Columns: []string{apistatistics.APIColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(api.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ApiStatistics{config: asuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, asuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apistatistics.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	asuo.mutation.done = true
	return _node, nil
}