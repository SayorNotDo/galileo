// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"galileo/ent/predicate"
	"galileo/ent/testplan"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestPlanQuery is the builder for querying TestPlan entities.
type TestPlanQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.TestPlan
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TestPlanQuery builder.
func (tpq *TestPlanQuery) Where(ps ...predicate.TestPlan) *TestPlanQuery {
	tpq.predicates = append(tpq.predicates, ps...)
	return tpq
}

// Limit the number of records to be returned by this query.
func (tpq *TestPlanQuery) Limit(limit int) *TestPlanQuery {
	tpq.ctx.Limit = &limit
	return tpq
}

// Offset to start from.
func (tpq *TestPlanQuery) Offset(offset int) *TestPlanQuery {
	tpq.ctx.Offset = &offset
	return tpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tpq *TestPlanQuery) Unique(unique bool) *TestPlanQuery {
	tpq.ctx.Unique = &unique
	return tpq
}

// Order specifies how the records should be ordered.
func (tpq *TestPlanQuery) Order(o ...OrderFunc) *TestPlanQuery {
	tpq.order = append(tpq.order, o...)
	return tpq
}

// First returns the first TestPlan entity from the query.
// Returns a *NotFoundError when no TestPlan was found.
func (tpq *TestPlanQuery) First(ctx context.Context) (*TestPlan, error) {
	nodes, err := tpq.Limit(1).All(setContextOp(ctx, tpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{testplan.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tpq *TestPlanQuery) FirstX(ctx context.Context) *TestPlan {
	node, err := tpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TestPlan ID from the query.
// Returns a *NotFoundError when no TestPlan ID was found.
func (tpq *TestPlanQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = tpq.Limit(1).IDs(setContextOp(ctx, tpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{testplan.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tpq *TestPlanQuery) FirstIDX(ctx context.Context) int64 {
	id, err := tpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TestPlan entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TestPlan entity is found.
// Returns a *NotFoundError when no TestPlan entities are found.
func (tpq *TestPlanQuery) Only(ctx context.Context) (*TestPlan, error) {
	nodes, err := tpq.Limit(2).All(setContextOp(ctx, tpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{testplan.Label}
	default:
		return nil, &NotSingularError{testplan.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tpq *TestPlanQuery) OnlyX(ctx context.Context) *TestPlan {
	node, err := tpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TestPlan ID in the query.
// Returns a *NotSingularError when more than one TestPlan ID is found.
// Returns a *NotFoundError when no entities are found.
func (tpq *TestPlanQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = tpq.Limit(2).IDs(setContextOp(ctx, tpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{testplan.Label}
	default:
		err = &NotSingularError{testplan.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tpq *TestPlanQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := tpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TestPlans.
func (tpq *TestPlanQuery) All(ctx context.Context) ([]*TestPlan, error) {
	ctx = setContextOp(ctx, tpq.ctx, "All")
	if err := tpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TestPlan, *TestPlanQuery]()
	return withInterceptors[[]*TestPlan](ctx, tpq, qr, tpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tpq *TestPlanQuery) AllX(ctx context.Context) []*TestPlan {
	nodes, err := tpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TestPlan IDs.
func (tpq *TestPlanQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if tpq.ctx.Unique == nil && tpq.path != nil {
		tpq.Unique(true)
	}
	ctx = setContextOp(ctx, tpq.ctx, "IDs")
	if err = tpq.Select(testplan.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tpq *TestPlanQuery) IDsX(ctx context.Context) []int64 {
	ids, err := tpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tpq *TestPlanQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tpq.ctx, "Count")
	if err := tpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tpq, querierCount[*TestPlanQuery](), tpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tpq *TestPlanQuery) CountX(ctx context.Context) int {
	count, err := tpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tpq *TestPlanQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tpq.ctx, "Exist")
	switch _, err := tpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tpq *TestPlanQuery) ExistX(ctx context.Context) bool {
	exist, err := tpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TestPlanQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tpq *TestPlanQuery) Clone() *TestPlanQuery {
	if tpq == nil {
		return nil
	}
	return &TestPlanQuery{
		config:     tpq.config,
		ctx:        tpq.ctx.Clone(),
		order:      append([]OrderFunc{}, tpq.order...),
		inters:     append([]Interceptor{}, tpq.inters...),
		predicates: append([]predicate.TestPlan{}, tpq.predicates...),
		// clone intermediate query.
		sql:  tpq.sql.Clone(),
		path: tpq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TestPlan.Query().
//		GroupBy(testplan.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tpq *TestPlanQuery) GroupBy(field string, fields ...string) *TestPlanGroupBy {
	tpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TestPlanGroupBy{build: tpq}
	grbuild.flds = &tpq.ctx.Fields
	grbuild.label = testplan.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.TestPlan.Query().
//		Select(testplan.FieldName).
//		Scan(ctx, &v)
func (tpq *TestPlanQuery) Select(fields ...string) *TestPlanSelect {
	tpq.ctx.Fields = append(tpq.ctx.Fields, fields...)
	sbuild := &TestPlanSelect{TestPlanQuery: tpq}
	sbuild.label = testplan.Label
	sbuild.flds, sbuild.scan = &tpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TestPlanSelect configured with the given aggregations.
func (tpq *TestPlanQuery) Aggregate(fns ...AggregateFunc) *TestPlanSelect {
	return tpq.Select().Aggregate(fns...)
}

func (tpq *TestPlanQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tpq); err != nil {
				return err
			}
		}
	}
	for _, f := range tpq.ctx.Fields {
		if !testplan.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tpq.path != nil {
		prev, err := tpq.path(ctx)
		if err != nil {
			return err
		}
		tpq.sql = prev
	}
	return nil
}

func (tpq *TestPlanQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TestPlan, error) {
	var (
		nodes = []*TestPlan{}
		_spec = tpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TestPlan).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TestPlan{config: tpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tpq *TestPlanQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tpq.querySpec()
	_spec.Node.Columns = tpq.ctx.Fields
	if len(tpq.ctx.Fields) > 0 {
		_spec.Unique = tpq.ctx.Unique != nil && *tpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tpq.driver, _spec)
}

func (tpq *TestPlanQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(testplan.Table, testplan.Columns, sqlgraph.NewFieldSpec(testplan.FieldID, field.TypeInt64))
	_spec.From = tpq.sql
	if unique := tpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tpq.path != nil {
		_spec.Unique = true
	}
	if fields := tpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testplan.FieldID)
		for i := range fields {
			if fields[i] != testplan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tpq *TestPlanQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tpq.driver.Dialect())
	t1 := builder.Table(testplan.Table)
	columns := tpq.ctx.Fields
	if len(columns) == 0 {
		columns = testplan.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tpq.sql != nil {
		selector = tpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tpq.ctx.Unique != nil && *tpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tpq.predicates {
		p(selector)
	}
	for _, p := range tpq.order {
		p(selector)
	}
	if offset := tpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TestPlanGroupBy is the group-by builder for TestPlan entities.
type TestPlanGroupBy struct {
	selector
	build *TestPlanQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tpgb *TestPlanGroupBy) Aggregate(fns ...AggregateFunc) *TestPlanGroupBy {
	tpgb.fns = append(tpgb.fns, fns...)
	return tpgb
}

// Scan applies the selector query and scans the result into the given value.
func (tpgb *TestPlanGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tpgb.build.ctx, "GroupBy")
	if err := tpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestPlanQuery, *TestPlanGroupBy](ctx, tpgb.build, tpgb, tpgb.build.inters, v)
}

func (tpgb *TestPlanGroupBy) sqlScan(ctx context.Context, root *TestPlanQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tpgb.fns))
	for _, fn := range tpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tpgb.flds)+len(tpgb.fns))
		for _, f := range *tpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TestPlanSelect is the builder for selecting fields of TestPlan entities.
type TestPlanSelect struct {
	*TestPlanQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tps *TestPlanSelect) Aggregate(fns ...AggregateFunc) *TestPlanSelect {
	tps.fns = append(tps.fns, fns...)
	return tps
}

// Scan applies the selector query and scans the result into the given value.
func (tps *TestPlanSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tps.ctx, "Select")
	if err := tps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestPlanQuery, *TestPlanSelect](ctx, tps.TestPlanQuery, tps, tps.inters, v)
}

func (tps *TestPlanSelect) sqlScan(ctx context.Context, root *TestPlanQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tps.fns))
	for _, fn := range tps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
