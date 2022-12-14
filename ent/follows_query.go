// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bypepe77/api/ent/follows"
	"github.com/bypepe77/api/ent/predicate"
)

// FollowsQuery is the builder for querying Follows entities.
type FollowsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Follows
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FollowsQuery builder.
func (fq *FollowsQuery) Where(ps ...predicate.Follows) *FollowsQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit adds a limit step to the query.
func (fq *FollowsQuery) Limit(limit int) *FollowsQuery {
	fq.limit = &limit
	return fq
}

// Offset adds an offset step to the query.
func (fq *FollowsQuery) Offset(offset int) *FollowsQuery {
	fq.offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FollowsQuery) Unique(unique bool) *FollowsQuery {
	fq.unique = &unique
	return fq
}

// Order adds an order step to the query.
func (fq *FollowsQuery) Order(o ...OrderFunc) *FollowsQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// First returns the first Follows entity from the query.
// Returns a *NotFoundError when no Follows was found.
func (fq *FollowsQuery) First(ctx context.Context) (*Follows, error) {
	nodes, err := fq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{follows.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FollowsQuery) FirstX(ctx context.Context) *Follows {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Follows ID from the query.
// Returns a *NotFoundError when no Follows ID was found.
func (fq *FollowsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{follows.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FollowsQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Follows entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Follows entity is found.
// Returns a *NotFoundError when no Follows entities are found.
func (fq *FollowsQuery) Only(ctx context.Context) (*Follows, error) {
	nodes, err := fq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{follows.Label}
	default:
		return nil, &NotSingularError{follows.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FollowsQuery) OnlyX(ctx context.Context) *Follows {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Follows ID in the query.
// Returns a *NotSingularError when more than one Follows ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FollowsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{follows.Label}
	default:
		err = &NotSingularError{follows.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FollowsQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FollowsSlice.
func (fq *FollowsQuery) All(ctx context.Context) ([]*Follows, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fq *FollowsQuery) AllX(ctx context.Context) []*Follows {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Follows IDs.
func (fq *FollowsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fq.Select(follows.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FollowsQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FollowsQuery) Count(ctx context.Context) (int, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FollowsQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FollowsQuery) Exist(ctx context.Context) (bool, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FollowsQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FollowsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FollowsQuery) Clone() *FollowsQuery {
	if fq == nil {
		return nil
	}
	return &FollowsQuery{
		config:     fq.config,
		limit:      fq.limit,
		offset:     fq.offset,
		order:      append([]OrderFunc{}, fq.order...),
		predicates: append([]predicate.Follows{}, fq.predicates...),
		// clone intermediate query.
		sql:    fq.sql.Clone(),
		path:   fq.path,
		unique: fq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Followedby int `json:"followedby,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Follows.Query().
//		GroupBy(follows.FieldFollowedby).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FollowsQuery) GroupBy(field string, fields ...string) *FollowsGroupBy {
	grbuild := &FollowsGroupBy{config: fq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fq.sqlQuery(ctx), nil
	}
	grbuild.label = follows.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Followedby int `json:"followedby,omitempty"`
//	}
//
//	client.Follows.Query().
//		Select(follows.FieldFollowedby).
//		Scan(ctx, &v)
func (fq *FollowsQuery) Select(fields ...string) *FollowsSelect {
	fq.fields = append(fq.fields, fields...)
	selbuild := &FollowsSelect{FollowsQuery: fq}
	selbuild.label = follows.Label
	selbuild.flds, selbuild.scan = &fq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a FollowsSelect configured with the given aggregations.
func (fq *FollowsQuery) Aggregate(fns ...AggregateFunc) *FollowsSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FollowsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fq.fields {
		if !follows.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FollowsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Follows, error) {
	var (
		nodes = []*Follows{}
		_spec = fq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Follows).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Follows{config: fq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (fq *FollowsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Columns = fq.fields
	if len(fq.fields) > 0 {
		_spec.Unique = fq.unique != nil && *fq.unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FollowsQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (fq *FollowsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   follows.Table,
			Columns: follows.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: follows.FieldID,
			},
		},
		From:   fq.sql,
		Unique: true,
	}
	if unique := fq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, follows.FieldID)
		for i := range fields {
			if fields[i] != follows.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FollowsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(follows.Table)
	columns := fq.fields
	if len(columns) == 0 {
		columns = follows.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.unique != nil && *fq.unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FollowsGroupBy is the group-by builder for Follows entities.
type FollowsGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FollowsGroupBy) Aggregate(fns ...AggregateFunc) *FollowsGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fgb *FollowsGroupBy) Scan(ctx context.Context, v any) error {
	query, err := fgb.path(ctx)
	if err != nil {
		return err
	}
	fgb.sql = query
	return fgb.sqlScan(ctx, v)
}

func (fgb *FollowsGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range fgb.fields {
		if !follows.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fgb *FollowsGroupBy) sqlQuery() *sql.Selector {
	selector := fgb.sql.Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fgb.fields)+len(fgb.fns))
		for _, f := range fgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fgb.fields...)...)
}

// FollowsSelect is the builder for selecting fields of Follows entities.
type FollowsSelect struct {
	*FollowsQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FollowsSelect) Aggregate(fns ...AggregateFunc) *FollowsSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FollowsSelect) Scan(ctx context.Context, v any) error {
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	fs.sql = fs.FollowsQuery.sqlQuery(ctx)
	return fs.sqlScan(ctx, v)
}

func (fs *FollowsSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(fs.sql))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		fs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		fs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := fs.sql.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
