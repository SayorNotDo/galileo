// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"galileo/ent/apicategory"
	"galileo/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApiCategoryQuery is the builder for querying ApiCategory entities.
type ApiCategoryQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.ApiCategory
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ApiCategoryQuery builder.
func (acq *ApiCategoryQuery) Where(ps ...predicate.ApiCategory) *ApiCategoryQuery {
	acq.predicates = append(acq.predicates, ps...)
	return acq
}

// Limit the number of records to be returned by this query.
func (acq *ApiCategoryQuery) Limit(limit int) *ApiCategoryQuery {
	acq.ctx.Limit = &limit
	return acq
}

// Offset to start from.
func (acq *ApiCategoryQuery) Offset(offset int) *ApiCategoryQuery {
	acq.ctx.Offset = &offset
	return acq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acq *ApiCategoryQuery) Unique(unique bool) *ApiCategoryQuery {
	acq.ctx.Unique = &unique
	return acq
}

// Order specifies how the records should be ordered.
func (acq *ApiCategoryQuery) Order(o ...OrderFunc) *ApiCategoryQuery {
	acq.order = append(acq.order, o...)
	return acq
}

// First returns the first ApiCategory entity from the query.
// Returns a *NotFoundError when no ApiCategory was found.
func (acq *ApiCategoryQuery) First(ctx context.Context) (*ApiCategory, error) {
	nodes, err := acq.Limit(1).All(setContextOp(ctx, acq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{apicategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acq *ApiCategoryQuery) FirstX(ctx context.Context) *ApiCategory {
	node, err := acq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ApiCategory ID from the query.
// Returns a *NotFoundError when no ApiCategory ID was found.
func (acq *ApiCategoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = acq.Limit(1).IDs(setContextOp(ctx, acq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{apicategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acq *ApiCategoryQuery) FirstIDX(ctx context.Context) int {
	id, err := acq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ApiCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ApiCategory entity is found.
// Returns a *NotFoundError when no ApiCategory entities are found.
func (acq *ApiCategoryQuery) Only(ctx context.Context) (*ApiCategory, error) {
	nodes, err := acq.Limit(2).All(setContextOp(ctx, acq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{apicategory.Label}
	default:
		return nil, &NotSingularError{apicategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acq *ApiCategoryQuery) OnlyX(ctx context.Context) *ApiCategory {
	node, err := acq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ApiCategory ID in the query.
// Returns a *NotSingularError when more than one ApiCategory ID is found.
// Returns a *NotFoundError when no entities are found.
func (acq *ApiCategoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = acq.Limit(2).IDs(setContextOp(ctx, acq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{apicategory.Label}
	default:
		err = &NotSingularError{apicategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acq *ApiCategoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := acq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ApiCategories.
func (acq *ApiCategoryQuery) All(ctx context.Context) ([]*ApiCategory, error) {
	ctx = setContextOp(ctx, acq.ctx, "All")
	if err := acq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ApiCategory, *ApiCategoryQuery]()
	return withInterceptors[[]*ApiCategory](ctx, acq, qr, acq.inters)
}

// AllX is like All, but panics if an error occurs.
func (acq *ApiCategoryQuery) AllX(ctx context.Context) []*ApiCategory {
	nodes, err := acq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ApiCategory IDs.
func (acq *ApiCategoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if acq.ctx.Unique == nil && acq.path != nil {
		acq.Unique(true)
	}
	ctx = setContextOp(ctx, acq.ctx, "IDs")
	if err = acq.Select(apicategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acq *ApiCategoryQuery) IDsX(ctx context.Context) []int {
	ids, err := acq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acq *ApiCategoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, acq.ctx, "Count")
	if err := acq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, acq, querierCount[*ApiCategoryQuery](), acq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (acq *ApiCategoryQuery) CountX(ctx context.Context) int {
	count, err := acq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acq *ApiCategoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, acq.ctx, "Exist")
	switch _, err := acq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (acq *ApiCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := acq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ApiCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acq *ApiCategoryQuery) Clone() *ApiCategoryQuery {
	if acq == nil {
		return nil
	}
	return &ApiCategoryQuery{
		config:     acq.config,
		ctx:        acq.ctx.Clone(),
		order:      append([]OrderFunc{}, acq.order...),
		inters:     append([]Interceptor{}, acq.inters...),
		predicates: append([]predicate.ApiCategory{}, acq.predicates...),
		// clone intermediate query.
		sql:  acq.sql.Clone(),
		path: acq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (acq *ApiCategoryQuery) GroupBy(field string, fields ...string) *ApiCategoryGroupBy {
	acq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ApiCategoryGroupBy{build: acq}
	grbuild.flds = &acq.ctx.Fields
	grbuild.label = apicategory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (acq *ApiCategoryQuery) Select(fields ...string) *ApiCategorySelect {
	acq.ctx.Fields = append(acq.ctx.Fields, fields...)
	sbuild := &ApiCategorySelect{ApiCategoryQuery: acq}
	sbuild.label = apicategory.Label
	sbuild.flds, sbuild.scan = &acq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ApiCategorySelect configured with the given aggregations.
func (acq *ApiCategoryQuery) Aggregate(fns ...AggregateFunc) *ApiCategorySelect {
	return acq.Select().Aggregate(fns...)
}

func (acq *ApiCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range acq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, acq); err != nil {
				return err
			}
		}
	}
	for _, f := range acq.ctx.Fields {
		if !apicategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acq.path != nil {
		prev, err := acq.path(ctx)
		if err != nil {
			return err
		}
		acq.sql = prev
	}
	return nil
}

func (acq *ApiCategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ApiCategory, error) {
	var (
		nodes = []*ApiCategory{}
		_spec = acq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ApiCategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ApiCategory{config: acq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, acq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (acq *ApiCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acq.querySpec()
	_spec.Node.Columns = acq.ctx.Fields
	if len(acq.ctx.Fields) > 0 {
		_spec.Unique = acq.ctx.Unique != nil && *acq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, acq.driver, _spec)
}

func (acq *ApiCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(apicategory.Table, apicategory.Columns, sqlgraph.NewFieldSpec(apicategory.FieldID, field.TypeInt))
	_spec.From = acq.sql
	if unique := acq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if acq.path != nil {
		_spec.Unique = true
	}
	if fields := acq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apicategory.FieldID)
		for i := range fields {
			if fields[i] != apicategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := acq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acq *ApiCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acq.driver.Dialect())
	t1 := builder.Table(apicategory.Table)
	columns := acq.ctx.Fields
	if len(columns) == 0 {
		columns = apicategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acq.sql != nil {
		selector = acq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if acq.ctx.Unique != nil && *acq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range acq.predicates {
		p(selector)
	}
	for _, p := range acq.order {
		p(selector)
	}
	if offset := acq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ApiCategoryGroupBy is the group-by builder for ApiCategory entities.
type ApiCategoryGroupBy struct {
	selector
	build *ApiCategoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acgb *ApiCategoryGroupBy) Aggregate(fns ...AggregateFunc) *ApiCategoryGroupBy {
	acgb.fns = append(acgb.fns, fns...)
	return acgb
}

// Scan applies the selector query and scans the result into the given value.
func (acgb *ApiCategoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acgb.build.ctx, "GroupBy")
	if err := acgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ApiCategoryQuery, *ApiCategoryGroupBy](ctx, acgb.build, acgb, acgb.build.inters, v)
}

func (acgb *ApiCategoryGroupBy) sqlScan(ctx context.Context, root *ApiCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(acgb.fns))
	for _, fn := range acgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*acgb.flds)+len(acgb.fns))
		for _, f := range *acgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*acgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ApiCategorySelect is the builder for selecting fields of ApiCategory entities.
type ApiCategorySelect struct {
	*ApiCategoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (acs *ApiCategorySelect) Aggregate(fns ...AggregateFunc) *ApiCategorySelect {
	acs.fns = append(acs.fns, fns...)
	return acs
}

// Scan applies the selector query and scans the result into the given value.
func (acs *ApiCategorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acs.ctx, "Select")
	if err := acs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ApiCategoryQuery, *ApiCategorySelect](ctx, acs.ApiCategoryQuery, acs, acs.inters, v)
}

func (acs *ApiCategorySelect) sqlScan(ctx context.Context, root *ApiCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(acs.fns))
	for _, fn := range acs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*acs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}