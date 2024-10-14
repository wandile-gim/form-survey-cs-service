// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"form-survey-cs-service/internal/repository/ent/predicate"
	"form-survey-cs-service/internal/repository/ent/task"
	"form-survey-cs-service/internal/repository/ent/tasklog"
	"form-survey-cs-service/internal/repository/ent/taskrecord"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskRecordQuery is the builder for querying TaskRecord entities.
type TaskRecordQuery struct {
	config
	ctx          *QueryContext
	order        []taskrecord.OrderOption
	inters       []Interceptor
	predicates   []predicate.TaskRecord
	withTask     *TaskQuery
	withTaskLogs *TaskLogQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TaskRecordQuery builder.
func (trq *TaskRecordQuery) Where(ps ...predicate.TaskRecord) *TaskRecordQuery {
	trq.predicates = append(trq.predicates, ps...)
	return trq
}

// Limit the number of records to be returned by this query.
func (trq *TaskRecordQuery) Limit(limit int) *TaskRecordQuery {
	trq.ctx.Limit = &limit
	return trq
}

// Offset to start from.
func (trq *TaskRecordQuery) Offset(offset int) *TaskRecordQuery {
	trq.ctx.Offset = &offset
	return trq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (trq *TaskRecordQuery) Unique(unique bool) *TaskRecordQuery {
	trq.ctx.Unique = &unique
	return trq
}

// Order specifies how the records should be ordered.
func (trq *TaskRecordQuery) Order(o ...taskrecord.OrderOption) *TaskRecordQuery {
	trq.order = append(trq.order, o...)
	return trq
}

// QueryTask chains the current query on the "task" edge.
func (trq *TaskRecordQuery) QueryTask() *TaskQuery {
	query := (&TaskClient{config: trq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := trq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := trq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(taskrecord.Table, taskrecord.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, taskrecord.TaskTable, taskrecord.TaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(trq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTaskLogs chains the current query on the "task_logs" edge.
func (trq *TaskRecordQuery) QueryTaskLogs() *TaskLogQuery {
	query := (&TaskLogClient{config: trq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := trq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := trq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(taskrecord.Table, taskrecord.FieldID, selector),
			sqlgraph.To(tasklog.Table, tasklog.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, taskrecord.TaskLogsTable, taskrecord.TaskLogsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(trq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TaskRecord entity from the query.
// Returns a *NotFoundError when no TaskRecord was found.
func (trq *TaskRecordQuery) First(ctx context.Context) (*TaskRecord, error) {
	nodes, err := trq.Limit(1).All(setContextOp(ctx, trq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{taskrecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (trq *TaskRecordQuery) FirstX(ctx context.Context) *TaskRecord {
	node, err := trq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TaskRecord ID from the query.
// Returns a *NotFoundError when no TaskRecord ID was found.
func (trq *TaskRecordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = trq.Limit(1).IDs(setContextOp(ctx, trq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{taskrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (trq *TaskRecordQuery) FirstIDX(ctx context.Context) int {
	id, err := trq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TaskRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TaskRecord entity is found.
// Returns a *NotFoundError when no TaskRecord entities are found.
func (trq *TaskRecordQuery) Only(ctx context.Context) (*TaskRecord, error) {
	nodes, err := trq.Limit(2).All(setContextOp(ctx, trq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{taskrecord.Label}
	default:
		return nil, &NotSingularError{taskrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (trq *TaskRecordQuery) OnlyX(ctx context.Context) *TaskRecord {
	node, err := trq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TaskRecord ID in the query.
// Returns a *NotSingularError when more than one TaskRecord ID is found.
// Returns a *NotFoundError when no entities are found.
func (trq *TaskRecordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = trq.Limit(2).IDs(setContextOp(ctx, trq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{taskrecord.Label}
	default:
		err = &NotSingularError{taskrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (trq *TaskRecordQuery) OnlyIDX(ctx context.Context) int {
	id, err := trq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TaskRecords.
func (trq *TaskRecordQuery) All(ctx context.Context) ([]*TaskRecord, error) {
	ctx = setContextOp(ctx, trq.ctx, ent.OpQueryAll)
	if err := trq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TaskRecord, *TaskRecordQuery]()
	return withInterceptors[[]*TaskRecord](ctx, trq, qr, trq.inters)
}

// AllX is like All, but panics if an error occurs.
func (trq *TaskRecordQuery) AllX(ctx context.Context) []*TaskRecord {
	nodes, err := trq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TaskRecord IDs.
func (trq *TaskRecordQuery) IDs(ctx context.Context) (ids []int, err error) {
	if trq.ctx.Unique == nil && trq.path != nil {
		trq.Unique(true)
	}
	ctx = setContextOp(ctx, trq.ctx, ent.OpQueryIDs)
	if err = trq.Select(taskrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (trq *TaskRecordQuery) IDsX(ctx context.Context) []int {
	ids, err := trq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (trq *TaskRecordQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, trq.ctx, ent.OpQueryCount)
	if err := trq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, trq, querierCount[*TaskRecordQuery](), trq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (trq *TaskRecordQuery) CountX(ctx context.Context) int {
	count, err := trq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (trq *TaskRecordQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, trq.ctx, ent.OpQueryExist)
	switch _, err := trq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (trq *TaskRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := trq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TaskRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (trq *TaskRecordQuery) Clone() *TaskRecordQuery {
	if trq == nil {
		return nil
	}
	return &TaskRecordQuery{
		config:       trq.config,
		ctx:          trq.ctx.Clone(),
		order:        append([]taskrecord.OrderOption{}, trq.order...),
		inters:       append([]Interceptor{}, trq.inters...),
		predicates:   append([]predicate.TaskRecord{}, trq.predicates...),
		withTask:     trq.withTask.Clone(),
		withTaskLogs: trq.withTaskLogs.Clone(),
		// clone intermediate query.
		sql:  trq.sql.Clone(),
		path: trq.path,
	}
}

// WithTask tells the query-builder to eager-load the nodes that are connected to
// the "task" edge. The optional arguments are used to configure the query builder of the edge.
func (trq *TaskRecordQuery) WithTask(opts ...func(*TaskQuery)) *TaskRecordQuery {
	query := (&TaskClient{config: trq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	trq.withTask = query
	return trq
}

// WithTaskLogs tells the query-builder to eager-load the nodes that are connected to
// the "task_logs" edge. The optional arguments are used to configure the query builder of the edge.
func (trq *TaskRecordQuery) WithTaskLogs(opts ...func(*TaskLogQuery)) *TaskRecordQuery {
	query := (&TaskLogClient{config: trq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	trq.withTaskLogs = query
	return trq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		State string `json:"state,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TaskRecord.Query().
//		GroupBy(taskrecord.FieldState).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (trq *TaskRecordQuery) GroupBy(field string, fields ...string) *TaskRecordGroupBy {
	trq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TaskRecordGroupBy{build: trq}
	grbuild.flds = &trq.ctx.Fields
	grbuild.label = taskrecord.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		State string `json:"state,omitempty"`
//	}
//
//	client.TaskRecord.Query().
//		Select(taskrecord.FieldState).
//		Scan(ctx, &v)
func (trq *TaskRecordQuery) Select(fields ...string) *TaskRecordSelect {
	trq.ctx.Fields = append(trq.ctx.Fields, fields...)
	sbuild := &TaskRecordSelect{TaskRecordQuery: trq}
	sbuild.label = taskrecord.Label
	sbuild.flds, sbuild.scan = &trq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TaskRecordSelect configured with the given aggregations.
func (trq *TaskRecordQuery) Aggregate(fns ...AggregateFunc) *TaskRecordSelect {
	return trq.Select().Aggregate(fns...)
}

func (trq *TaskRecordQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range trq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, trq); err != nil {
				return err
			}
		}
	}
	for _, f := range trq.ctx.Fields {
		if !taskrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if trq.path != nil {
		prev, err := trq.path(ctx)
		if err != nil {
			return err
		}
		trq.sql = prev
	}
	return nil
}

func (trq *TaskRecordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TaskRecord, error) {
	var (
		nodes       = []*TaskRecord{}
		withFKs     = trq.withFKs
		_spec       = trq.querySpec()
		loadedTypes = [2]bool{
			trq.withTask != nil,
			trq.withTaskLogs != nil,
		}
	)
	if trq.withTask != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, taskrecord.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TaskRecord).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TaskRecord{config: trq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, trq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := trq.withTask; query != nil {
		if err := trq.loadTask(ctx, query, nodes, nil,
			func(n *TaskRecord, e *Task) { n.Edges.Task = e }); err != nil {
			return nil, err
		}
	}
	if query := trq.withTaskLogs; query != nil {
		if err := trq.loadTaskLogs(ctx, query, nodes,
			func(n *TaskRecord) { n.Edges.TaskLogs = []*TaskLog{} },
			func(n *TaskRecord, e *TaskLog) { n.Edges.TaskLogs = append(n.Edges.TaskLogs, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (trq *TaskRecordQuery) loadTask(ctx context.Context, query *TaskQuery, nodes []*TaskRecord, init func(*TaskRecord), assign func(*TaskRecord, *Task)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TaskRecord)
	for i := range nodes {
		if nodes[i].task_task_record == nil {
			continue
		}
		fk := *nodes[i].task_task_record
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(task.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "task_task_record" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (trq *TaskRecordQuery) loadTaskLogs(ctx context.Context, query *TaskLogQuery, nodes []*TaskRecord, init func(*TaskRecord), assign func(*TaskRecord, *TaskLog)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*TaskRecord)
	nids := make(map[int]map[*TaskRecord]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(taskrecord.TaskLogsTable)
		s.Join(joinT).On(s.C(tasklog.FieldID), joinT.C(taskrecord.TaskLogsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(taskrecord.TaskLogsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(taskrecord.TaskLogsPrimaryKey[0]))
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
					nids[inValue] = map[*TaskRecord]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*TaskLog](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "task_logs" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (trq *TaskRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := trq.querySpec()
	_spec.Node.Columns = trq.ctx.Fields
	if len(trq.ctx.Fields) > 0 {
		_spec.Unique = trq.ctx.Unique != nil && *trq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, trq.driver, _spec)
}

func (trq *TaskRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(taskrecord.Table, taskrecord.Columns, sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt))
	_spec.From = trq.sql
	if unique := trq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if trq.path != nil {
		_spec.Unique = true
	}
	if fields := trq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taskrecord.FieldID)
		for i := range fields {
			if fields[i] != taskrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := trq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := trq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := trq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := trq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (trq *TaskRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(trq.driver.Dialect())
	t1 := builder.Table(taskrecord.Table)
	columns := trq.ctx.Fields
	if len(columns) == 0 {
		columns = taskrecord.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if trq.sql != nil {
		selector = trq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if trq.ctx.Unique != nil && *trq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range trq.predicates {
		p(selector)
	}
	for _, p := range trq.order {
		p(selector)
	}
	if offset := trq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := trq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TaskRecordGroupBy is the group-by builder for TaskRecord entities.
type TaskRecordGroupBy struct {
	selector
	build *TaskRecordQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (trgb *TaskRecordGroupBy) Aggregate(fns ...AggregateFunc) *TaskRecordGroupBy {
	trgb.fns = append(trgb.fns, fns...)
	return trgb
}

// Scan applies the selector query and scans the result into the given value.
func (trgb *TaskRecordGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, trgb.build.ctx, ent.OpQueryGroupBy)
	if err := trgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TaskRecordQuery, *TaskRecordGroupBy](ctx, trgb.build, trgb, trgb.build.inters, v)
}

func (trgb *TaskRecordGroupBy) sqlScan(ctx context.Context, root *TaskRecordQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(trgb.fns))
	for _, fn := range trgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*trgb.flds)+len(trgb.fns))
		for _, f := range *trgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*trgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := trgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TaskRecordSelect is the builder for selecting fields of TaskRecord entities.
type TaskRecordSelect struct {
	*TaskRecordQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (trs *TaskRecordSelect) Aggregate(fns ...AggregateFunc) *TaskRecordSelect {
	trs.fns = append(trs.fns, fns...)
	return trs
}

// Scan applies the selector query and scans the result into the given value.
func (trs *TaskRecordSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, trs.ctx, ent.OpQuerySelect)
	if err := trs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TaskRecordQuery, *TaskRecordSelect](ctx, trs.TaskRecordQuery, trs, trs.inters, v)
}

func (trs *TaskRecordSelect) sqlScan(ctx context.Context, root *TaskRecordQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(trs.fns))
	for _, fn := range trs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*trs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := trs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
