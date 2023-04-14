// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"galileo/ent/predicate"
	"galileo/ent/testcase"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestCaseQuery is the builder for querying TestCase entities.
type TestCaseQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.TestCase
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TestCaseQuery builder.
func (tcq *TestCaseQuery) Where(ps ...predicate.TestCase) *TestCaseQuery {
	tcq.predicates = append(tcq.predicates, ps...)
	return tcq
}

// Limit the number of records to be returned by this query.
func (tcq *TestCaseQuery) Limit(limit int) *TestCaseQuery {
	tcq.ctx.Limit = &limit
	return tcq
}

// Offset to start from.
func (tcq *TestCaseQuery) Offset(offset int) *TestCaseQuery {
	tcq.ctx.Offset = &offset
	return tcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tcq *TestCaseQuery) Unique(unique bool) *TestCaseQuery {
	tcq.ctx.Unique = &unique
	return tcq
}

// Order specifies how the records should be ordered.
func (tcq *TestCaseQuery) Order(o ...OrderFunc) *TestCaseQuery {
	tcq.order = append(tcq.order, o...)
	return tcq
}

// First returns the first TestCase entity from the query.
// Returns a *NotFoundError when no TestCase was found.
func (tcq *TestCaseQuery) First(ctx context.Context) (*TestCase, error) {
	nodes, err := tcq.Limit(1).All(setContextOp(ctx, tcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{testcase.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tcq *TestCaseQuery) FirstX(ctx context.Context) *TestCase {
	node, err := tcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TestCase ID from the query.
// Returns a *NotFoundError when no TestCase ID was found.
func (tcq *TestCaseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tcq.Limit(1).IDs(setContextOp(ctx, tcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{testcase.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tcq *TestCaseQuery) FirstIDX(ctx context.Context) int {
	id, err := tcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TestCase entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TestCase entity is found.
// Returns a *NotFoundError when no TestCase entities are found.
func (tcq *TestCaseQuery) Only(ctx context.Context) (*TestCase, error) {
	nodes, err := tcq.Limit(2).All(setContextOp(ctx, tcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{testcase.Label}
	default:
		return nil, &NotSingularError{testcase.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tcq *TestCaseQuery) OnlyX(ctx context.Context) *TestCase {
	node, err := tcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TestCase ID in the query.
// Returns a *NotSingularError when more than one TestCase ID is found.
// Returns a *NotFoundError when no entities are found.
func (tcq *TestCaseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tcq.Limit(2).IDs(setContextOp(ctx, tcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{testcase.Label}
	default:
		err = &NotSingularError{testcase.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tcq *TestCaseQuery) OnlyIDX(ctx context.Context) int {
	id, err := tcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TestCases.
func (tcq *TestCaseQuery) All(ctx context.Context) ([]*TestCase, error) {
	ctx = setContextOp(ctx, tcq.ctx, "All")
	if err := tcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TestCase, *TestCaseQuery]()
	return withInterceptors[[]*TestCase](ctx, tcq, qr, tcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tcq *TestCaseQuery) AllX(ctx context.Context) []*TestCase {
	nodes, err := tcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TestCase IDs.
func (tcq *TestCaseQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tcq.ctx.Unique == nil && tcq.path != nil {
		tcq.Unique(true)
	}
	ctx = setContextOp(ctx, tcq.ctx, "IDs")
	if err = tcq.Select(testcase.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tcq *TestCaseQuery) IDsX(ctx context.Context) []int {
	ids, err := tcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tcq *TestCaseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tcq.ctx, "Count")
	if err := tcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tcq, querierCount[*TestCaseQuery](), tcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tcq *TestCaseQuery) CountX(ctx context.Context) int {
	count, err := tcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tcq *TestCaseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tcq.ctx, "Exist")
	switch _, err := tcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tcq *TestCaseQuery) ExistX(ctx context.Context) bool {
	exist, err := tcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TestCaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tcq *TestCaseQuery) Clone() *TestCaseQuery {
	if tcq == nil {
		return nil
	}
	return &TestCaseQuery{
		config:     tcq.config,
		ctx:        tcq.ctx.Clone(),
		order:      append([]OrderFunc{}, tcq.order...),
		inters:     append([]Interceptor{}, tcq.inters...),
		predicates: append([]predicate.TestCase{}, tcq.predicates...),
		// clone intermediate query.
		sql:  tcq.sql.Clone(),
		path: tcq.path,
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
//	client.TestCase.Query().
//		GroupBy(testcase.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tcq *TestCaseQuery) GroupBy(field string, fields ...string) *TestCaseGroupBy {
	tcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TestCaseGroupBy{build: tcq}
	grbuild.flds = &tcq.ctx.Fields
	grbuild.label = testcase.Label
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
//	client.TestCase.Query().
//		Select(testcase.FieldName).
//		Scan(ctx, &v)
func (tcq *TestCaseQuery) Select(fields ...string) *TestCaseSelect {
	tcq.ctx.Fields = append(tcq.ctx.Fields, fields...)
	sbuild := &TestCaseSelect{TestCaseQuery: tcq}
	sbuild.label = testcase.Label
	sbuild.flds, sbuild.scan = &tcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TestCaseSelect configured with the given aggregations.
func (tcq *TestCaseQuery) Aggregate(fns ...AggregateFunc) *TestCaseSelect {
	return tcq.Select().Aggregate(fns...)
}

func (tcq *TestCaseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tcq); err != nil {
				return err
			}
		}
	}
	for _, f := range tcq.ctx.Fields {
		if !testcase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tcq.path != nil {
		prev, err := tcq.path(ctx)
		if err != nil {
			return err
		}
		tcq.sql = prev
	}
	return nil
}

func (tcq *TestCaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TestCase, error) {
	var (
		nodes   = []*TestCase{}
		withFKs = tcq.withFKs
		_spec   = tcq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, testcase.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TestCase).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TestCase{config: tcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tcq *TestCaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tcq.querySpec()
	_spec.Node.Columns = tcq.ctx.Fields
	if len(tcq.ctx.Fields) > 0 {
		_spec.Unique = tcq.ctx.Unique != nil && *tcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tcq.driver, _spec)
}

func (tcq *TestCaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(testcase.Table, testcase.Columns, sqlgraph.NewFieldSpec(testcase.FieldID, field.TypeInt))
	_spec.From = tcq.sql
	if unique := tcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tcq.path != nil {
		_spec.Unique = true
	}
	if fields := tcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testcase.FieldID)
		for i := range fields {
			if fields[i] != testcase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tcq *TestCaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tcq.driver.Dialect())
	t1 := builder.Table(testcase.Table)
	columns := tcq.ctx.Fields
	if len(columns) == 0 {
		columns = testcase.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tcq.sql != nil {
		selector = tcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tcq.ctx.Unique != nil && *tcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tcq.predicates {
		p(selector)
	}
	for _, p := range tcq.order {
		p(selector)
	}
	if offset := tcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TestCaseGroupBy is the group-by builder for TestCase entities.
type TestCaseGroupBy struct {
	selector
	build *TestCaseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tcgb *TestCaseGroupBy) Aggregate(fns ...AggregateFunc) *TestCaseGroupBy {
	tcgb.fns = append(tcgb.fns, fns...)
	return tcgb
}

// Scan applies the selector query and scans the result into the given value.
func (tcgb *TestCaseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tcgb.build.ctx, "GroupBy")
	if err := tcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestCaseQuery, *TestCaseGroupBy](ctx, tcgb.build, tcgb, tcgb.build.inters, v)
}

func (tcgb *TestCaseGroupBy) sqlScan(ctx context.Context, root *TestCaseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tcgb.fns))
	for _, fn := range tcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tcgb.flds)+len(tcgb.fns))
		for _, f := range *tcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TestCaseSelect is the builder for selecting fields of TestCase entities.
type TestCaseSelect struct {
	*TestCaseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tcs *TestCaseSelect) Aggregate(fns ...AggregateFunc) *TestCaseSelect {
	tcs.fns = append(tcs.fns, fns...)
	return tcs
}

// Scan applies the selector query and scans the result into the given value.
func (tcs *TestCaseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tcs.ctx, "Select")
	if err := tcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestCaseQuery, *TestCaseSelect](ctx, tcs.TestCaseQuery, tcs, tcs.inters, v)
}

func (tcs *TestCaseSelect) sqlScan(ctx context.Context, root *TestCaseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tcs.fns))
	for _, fn := range tcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}