// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"form-survey-cs-service/internal/repository/ent/predicate"
	"form-survey-cs-service/internal/repository/ent/tasklog"
	"form-survey-cs-service/internal/repository/ent/taskrecord"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskLogUpdate is the builder for updating TaskLog entities.
type TaskLogUpdate struct {
	config
	hooks    []Hook
	mutation *TaskLogMutation
}

// Where appends a list predicates to the TaskLogUpdate builder.
func (tlu *TaskLogUpdate) Where(ps ...predicate.TaskLog) *TaskLogUpdate {
	tlu.mutation.Where(ps...)
	return tlu
}

// SetMessage sets the "message" field.
func (tlu *TaskLogUpdate) SetMessage(s string) *TaskLogUpdate {
	tlu.mutation.SetMessage(s)
	return tlu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (tlu *TaskLogUpdate) SetNillableMessage(s *string) *TaskLogUpdate {
	if s != nil {
		tlu.SetMessage(*s)
	}
	return tlu
}

// AddTaskRecordIDs adds the "task_records" edge to the TaskRecord entity by IDs.
func (tlu *TaskLogUpdate) AddTaskRecordIDs(ids ...int) *TaskLogUpdate {
	tlu.mutation.AddTaskRecordIDs(ids...)
	return tlu
}

// AddTaskRecords adds the "task_records" edges to the TaskRecord entity.
func (tlu *TaskLogUpdate) AddTaskRecords(t ...*TaskRecord) *TaskLogUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tlu.AddTaskRecordIDs(ids...)
}

// Mutation returns the TaskLogMutation object of the builder.
func (tlu *TaskLogUpdate) Mutation() *TaskLogMutation {
	return tlu.mutation
}

// ClearTaskRecords clears all "task_records" edges to the TaskRecord entity.
func (tlu *TaskLogUpdate) ClearTaskRecords() *TaskLogUpdate {
	tlu.mutation.ClearTaskRecords()
	return tlu
}

// RemoveTaskRecordIDs removes the "task_records" edge to TaskRecord entities by IDs.
func (tlu *TaskLogUpdate) RemoveTaskRecordIDs(ids ...int) *TaskLogUpdate {
	tlu.mutation.RemoveTaskRecordIDs(ids...)
	return tlu
}

// RemoveTaskRecords removes "task_records" edges to TaskRecord entities.
func (tlu *TaskLogUpdate) RemoveTaskRecords(t ...*TaskRecord) *TaskLogUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tlu.RemoveTaskRecordIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tlu *TaskLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tlu.sqlSave, tlu.mutation, tlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tlu *TaskLogUpdate) SaveX(ctx context.Context) int {
	affected, err := tlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tlu *TaskLogUpdate) Exec(ctx context.Context) error {
	_, err := tlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlu *TaskLogUpdate) ExecX(ctx context.Context) {
	if err := tlu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tlu *TaskLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tasklog.Table, tasklog.Columns, sqlgraph.NewFieldSpec(tasklog.FieldID, field.TypeInt))
	if ps := tlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tlu.mutation.Message(); ok {
		_spec.SetField(tasklog.FieldMessage, field.TypeString, value)
	}
	if tlu.mutation.TaskRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tasklog.TaskRecordsTable,
			Columns: tasklog.TaskRecordsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tlu.mutation.RemovedTaskRecordsIDs(); len(nodes) > 0 && !tlu.mutation.TaskRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tasklog.TaskRecordsTable,
			Columns: tasklog.TaskRecordsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tlu.mutation.TaskRecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tasklog.TaskRecordsTable,
			Columns: tasklog.TaskRecordsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tasklog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tlu.mutation.done = true
	return n, nil
}

// TaskLogUpdateOne is the builder for updating a single TaskLog entity.
type TaskLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskLogMutation
}

// SetMessage sets the "message" field.
func (tluo *TaskLogUpdateOne) SetMessage(s string) *TaskLogUpdateOne {
	tluo.mutation.SetMessage(s)
	return tluo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (tluo *TaskLogUpdateOne) SetNillableMessage(s *string) *TaskLogUpdateOne {
	if s != nil {
		tluo.SetMessage(*s)
	}
	return tluo
}

// AddTaskRecordIDs adds the "task_records" edge to the TaskRecord entity by IDs.
func (tluo *TaskLogUpdateOne) AddTaskRecordIDs(ids ...int) *TaskLogUpdateOne {
	tluo.mutation.AddTaskRecordIDs(ids...)
	return tluo
}

// AddTaskRecords adds the "task_records" edges to the TaskRecord entity.
func (tluo *TaskLogUpdateOne) AddTaskRecords(t ...*TaskRecord) *TaskLogUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tluo.AddTaskRecordIDs(ids...)
}

// Mutation returns the TaskLogMutation object of the builder.
func (tluo *TaskLogUpdateOne) Mutation() *TaskLogMutation {
	return tluo.mutation
}

// ClearTaskRecords clears all "task_records" edges to the TaskRecord entity.
func (tluo *TaskLogUpdateOne) ClearTaskRecords() *TaskLogUpdateOne {
	tluo.mutation.ClearTaskRecords()
	return tluo
}

// RemoveTaskRecordIDs removes the "task_records" edge to TaskRecord entities by IDs.
func (tluo *TaskLogUpdateOne) RemoveTaskRecordIDs(ids ...int) *TaskLogUpdateOne {
	tluo.mutation.RemoveTaskRecordIDs(ids...)
	return tluo
}

// RemoveTaskRecords removes "task_records" edges to TaskRecord entities.
func (tluo *TaskLogUpdateOne) RemoveTaskRecords(t ...*TaskRecord) *TaskLogUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tluo.RemoveTaskRecordIDs(ids...)
}

// Where appends a list predicates to the TaskLogUpdate builder.
func (tluo *TaskLogUpdateOne) Where(ps ...predicate.TaskLog) *TaskLogUpdateOne {
	tluo.mutation.Where(ps...)
	return tluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tluo *TaskLogUpdateOne) Select(field string, fields ...string) *TaskLogUpdateOne {
	tluo.fields = append([]string{field}, fields...)
	return tluo
}

// Save executes the query and returns the updated TaskLog entity.
func (tluo *TaskLogUpdateOne) Save(ctx context.Context) (*TaskLog, error) {
	return withHooks(ctx, tluo.sqlSave, tluo.mutation, tluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tluo *TaskLogUpdateOne) SaveX(ctx context.Context) *TaskLog {
	node, err := tluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tluo *TaskLogUpdateOne) Exec(ctx context.Context) error {
	_, err := tluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tluo *TaskLogUpdateOne) ExecX(ctx context.Context) {
	if err := tluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tluo *TaskLogUpdateOne) sqlSave(ctx context.Context) (_node *TaskLog, err error) {
	_spec := sqlgraph.NewUpdateSpec(tasklog.Table, tasklog.Columns, sqlgraph.NewFieldSpec(tasklog.FieldID, field.TypeInt))
	id, ok := tluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TaskLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tasklog.FieldID)
		for _, f := range fields {
			if !tasklog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tasklog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tluo.mutation.Message(); ok {
		_spec.SetField(tasklog.FieldMessage, field.TypeString, value)
	}
	if tluo.mutation.TaskRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tasklog.TaskRecordsTable,
			Columns: tasklog.TaskRecordsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tluo.mutation.RemovedTaskRecordsIDs(); len(nodes) > 0 && !tluo.mutation.TaskRecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tasklog.TaskRecordsTable,
			Columns: tasklog.TaskRecordsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tluo.mutation.TaskRecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tasklog.TaskRecordsTable,
			Columns: tasklog.TaskRecordsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TaskLog{config: tluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tasklog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tluo.mutation.done = true
	return _node, nil
}