// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"form-survey-cs-service/internal/repository/ent/predicate"
	"form-survey-cs-service/internal/repository/ent/tasklog"
	"form-survey-cs-service/internal/repository/ent/taskrecord"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskLogQuery is the builder for querying TaskLog entities.
type TaskLogQuery struct {
	config
	ctx             *QueryContext
	order           []tasklog.OrderOption
	inters          []Interceptor
	predicates      []predicate.TaskLog
	withTaskRecords *TaskRecordQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TaskLogQuery builder.
func (tlq *TaskLogQuery) Where(ps ...predicate.TaskLog) *TaskLogQuery {
	tlq.predicates = append(tlq.predicates, ps...)
	return tlq
}

// Limit the number of records to be returned by this query.
func (tlq *TaskLogQuery) Limit(limit int) *TaskLogQuery {
	tlq.ctx.Limit = &limit
	return tlq
}

// Offset to start from.
func (tlq *TaskLogQuery) Offset(offset int) *TaskLogQuery {
	tlq.ctx.Offset = &offset
	return tlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tlq *TaskLogQuery) Unique(unique bool) *TaskLogQuery {
	tlq.ctx.Unique = &unique
	return tlq
}

// Order specifies how the records should be ordered.
func (tlq *TaskLogQuery) Order(o ...tasklog.OrderOption) *TaskLogQuery {
	tlq.order = append(tlq.order, o...)
	return tlq
}

// QueryTaskRecords chains the current query on the "task_records" edge.
func (tlq *TaskLogQuery) QueryTaskRecords() *TaskRecordQuery {
	query := (&TaskRecordClient{config: tlq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tasklog.Table, tasklog.FieldID, selector),
			sqlgraph.To(taskrecord.Table, taskrecord.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tasklog.TaskRecordsTable, tasklog.TaskRecordsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TaskLog entity from the query.
// Returns a *NotFoundError when no TaskLog was found.
func (tlq *TaskLogQuery) First(ctx context.Context) (*TaskLog, error) {
	nodes, err := tlq.Limit(1).All(setContextOp(ctx, tlq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tasklog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tlq *TaskLogQuery) FirstX(ctx context.Context) *TaskLog {
	node, err := tlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TaskLog ID from the query.
// Returns a *NotFoundError when no TaskLog ID was found.
func (tlq *TaskLogQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tlq.Limit(1).IDs(setContextOp(ctx, tlq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tasklog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tlq *TaskLogQuery) FirstIDX(ctx context.Context) int {
	id, err := tlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TaskLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TaskLog entity is found.
// Returns a *NotFoundError when no TaskLog entities are found.
func (tlq *TaskLogQuery) Only(ctx context.Context) (*TaskLog, error) {
	nodes, err := tlq.Limit(2).All(setContextOp(ctx, tlq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tasklog.Label}
	default:
		return nil, &NotSingularError{tasklog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tlq *TaskLogQuery) OnlyX(ctx context.Context) *TaskLog {
	node, err := tlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TaskLog ID in the query.
// Returns a *NotSingularError when more than one TaskLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (tlq *TaskLogQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tlq.Limit(2).IDs(setContextOp(ctx, tlq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tasklog.Label}
	default:
		err = &NotSingularError{tasklog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tlq *TaskLogQuery) OnlyIDX(ctx context.Context) int {
	id, err := tlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TaskLogs.
func (tlq *TaskLogQuery) All(ctx context.Context) ([]*TaskLog, error) {
	ctx = setContextOp(ctx, tlq.ctx, ent.OpQueryAll)
	if err := tlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TaskLog, *TaskLogQuery]()
	return withInterceptors[[]*TaskLog](ctx, tlq, qr, tlq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tlq *TaskLogQuery) AllX(ctx context.Context) []*TaskLog {
	nodes, err := tlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TaskLog IDs.
func (tlq *TaskLogQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tlq.ctx.Unique == nil && tlq.path != nil {
		tlq.Unique(true)
	}
	ctx = setContextOp(ctx, tlq.ctx, ent.OpQueryIDs)
	if err = tlq.Select(tasklog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tlq *TaskLogQuery) IDsX(ctx context.Context) []int {
	ids, err := tlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tlq *TaskLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tlq.ctx, ent.OpQueryCount)
	if err := tlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tlq, querierCount[*TaskLogQuery](), tlq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tlq *TaskLogQuery) CountX(ctx context.Context) int {
	count, err := tlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tlq *TaskLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tlq.ctx, ent.OpQueryExist)
	switch _, err := tlq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tlq *TaskLogQuery) ExistX(ctx context.Context) bool {
	exist, err := tlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TaskLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tlq *TaskLogQuery) Clone() *TaskLogQuery {
	if tlq == nil {
		return nil
	}
	return &TaskLogQuery{
		config:          tlq.config,
		ctx:             tlq.ctx.Clone(),
		order:           append([]tasklog.OrderOption{}, tlq.order...),
		inters:          append([]Interceptor{}, tlq.inters...),
		predicates:      append([]predicate.TaskLog{}, tlq.predicates...),
		withTaskRecords: tlq.withTaskRecords.Clone(),
		// clone intermediate query.
		sql:  tlq.sql.Clone(),
		path: tlq.path,
	}
}

// WithTaskRecords tells the query-builder to eager-load the nodes that are connected to
// the "task_records" edge. The optional arguments are used to configure the query builder of the edge.
func (tlq *TaskLogQuery) WithTaskRecords(opts ...func(*TaskRecordQuery)) *TaskLogQuery {
	query := (&TaskRecordClient{config: tlq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tlq.withTaskRecords = query
	return tlq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Message string `json:"message,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TaskLog.Query().
//		GroupBy(tasklog.FieldMessage).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tlq *TaskLogQuery) GroupBy(field string, fields ...string) *TaskLogGroupBy {
	tlq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TaskLogGroupBy{build: tlq}
	grbuild.flds = &tlq.ctx.Fields
	grbuild.label = tasklog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Message string `json:"message,omitempty"`
//	}
//
//	client.TaskLog.Query().
//		Select(tasklog.FieldMessage).
//		Scan(ctx, &v)
func (tlq *TaskLogQuery) Select(fields ...string) *TaskLogSelect {
	tlq.ctx.Fields = append(tlq.ctx.Fields, fields...)
	sbuild := &TaskLogSelect{TaskLogQuery: tlq}
	sbuild.label = tasklog.Label
	sbuild.flds, sbuild.scan = &tlq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TaskLogSelect configured with the given aggregations.
func (tlq *TaskLogQuery) Aggregate(fns ...AggregateFunc) *TaskLogSelect {
	return tlq.Select().Aggregate(fns...)
}

func (tlq *TaskLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tlq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tlq); err != nil {
				return err
			}
		}
	}
	for _, f := range tlq.ctx.Fields {
		if !tasklog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tlq.path != nil {
		prev, err := tlq.path(ctx)
		if err != nil {
			return err
		}
		tlq.sql = prev
	}
	return nil
}

func (tlq *TaskLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TaskLog, error) {
	var (
		nodes       = []*TaskLog{}
		_spec       = tlq.querySpec()
		loadedTypes = [1]bool{
			tlq.withTaskRecords != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TaskLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TaskLog{config: tlq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tlq.withTaskRecords; query != nil {
		if err := tlq.loadTaskRecords(ctx, query, nodes,
			func(n *TaskLog) { n.Edges.TaskRecords = []*TaskRecord{} },
			func(n *TaskLog, e *TaskRecord) { n.Edges.TaskRecords = append(n.Edges.TaskRecords, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tlq *TaskLogQuery) loadTaskRecords(ctx context.Context, query *TaskRecordQuery, nodes []*TaskLog, init func(*TaskLog), assign func(*TaskLog, *TaskRecord)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*TaskLog)
	nids := make(map[int]map[*TaskLog]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(tasklog.TaskRecordsTable)
		s.Join(joinT).On(s.C(taskrecord.FieldID), joinT.C(tasklog.TaskRecordsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(tasklog.TaskRecordsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(tasklog.TaskRecordsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*TaskLog]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*TaskRecord](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "task_records" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (tlq *TaskLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tlq.querySpec()
	_spec.Node.Columns = tlq.ctx.Fields
	if len(tlq.ctx.Fields) > 0 {
		_spec.Unique = tlq.ctx.Unique != nil && *tlq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tlq.driver, _spec)
}

func (tlq *TaskLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tasklog.Table, tasklog.Columns, sqlgraph.NewFieldSpec(tasklog.FieldID, field.TypeInt))
	_spec.From = tlq.sql
	if unique := tlq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tlq.path != nil {
		_spec.Unique = true
	}
	if fields := tlq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tasklog.FieldID)
		for i := range fields {
			if fields[i] != tasklog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tlq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tlq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tlq *TaskLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tlq.driver.Dialect())
	t1 := builder.Table(tasklog.Table)
	columns := tlq.ctx.Fields
	if len(columns) == 0 {
		columns = tasklog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tlq.sql != nil {
		selector = tlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tlq.ctx.Unique != nil && *tlq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tlq.predicates {
		p(selector)
	}
	for _, p := range tlq.order {
		p(selector)
	}
	if offset := tlq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tlq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TaskLogGroupBy is the group-by builder for TaskLog entities.
type TaskLogGroupBy struct {
	selector
	build *TaskLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tlgb *TaskLogGroupBy) Aggregate(fns ...AggregateFunc) *TaskLogGroupBy {
	tlgb.fns = append(tlgb.fns, fns...)
	return tlgb
}

// Scan applies the selector query and scans the result into the given value.
func (tlgb *TaskLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tlgb.build.ctx, ent.OpQueryGroupBy)
	if err := tlgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TaskLogQuery, *TaskLogGroupBy](ctx, tlgb.build, tlgb, tlgb.build.inters, v)
}

func (tlgb *TaskLogGroupBy) sqlScan(ctx context.Context, root *TaskLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tlgb.fns))
	for _, fn := range tlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tlgb.flds)+len(tlgb.fns))
		for _, f := range *tlgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tlgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tlgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TaskLogSelect is the builder for selecting fields of TaskLog entities.
type TaskLogSelect struct {
	*TaskLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tls *TaskLogSelect) Aggregate(fns ...AggregateFunc) *TaskLogSelect {
	tls.fns = append(tls.fns, fns...)
	return tls
}

// Scan applies the selector query and scans the result into the given value.
func (tls *TaskLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tls.ctx, ent.OpQuerySelect)
	if err := tls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TaskLogQuery, *TaskLogSelect](ctx, tls.TaskLogQuery, tls, tls.inters, v)
}

func (tls *TaskLogSelect) sqlScan(ctx context.Context, root *TaskLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tls.fns))
	for _, fn := range tls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
