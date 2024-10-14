// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"form-survey-cs-service/internal/repository/ent/task"
	"form-survey-cs-service/internal/repository/ent/tasklog"
	"form-survey-cs-service/internal/repository/ent/taskrecord"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskRecordCreate is the builder for creating a TaskRecord entity.
type TaskRecordCreate struct {
	config
	mutation *TaskRecordMutation
	hooks    []Hook
}

// SetState sets the "state" field.
func (trc *TaskRecordCreate) SetState(s string) *TaskRecordCreate {
	trc.mutation.SetState(s)
	return trc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (trc *TaskRecordCreate) SetNillableState(s *string) *TaskRecordCreate {
	if s != nil {
		trc.SetState(*s)
	}
	return trc
}

// SetRetryCount sets the "retry_count" field.
func (trc *TaskRecordCreate) SetRetryCount(i int) *TaskRecordCreate {
	trc.mutation.SetRetryCount(i)
	return trc
}

// SetNillableRetryCount sets the "retry_count" field if the given value is not nil.
func (trc *TaskRecordCreate) SetNillableRetryCount(i *int) *TaskRecordCreate {
	if i != nil {
		trc.SetRetryCount(*i)
	}
	return trc
}

// SetID sets the "id" field.
func (trc *TaskRecordCreate) SetID(i int) *TaskRecordCreate {
	trc.mutation.SetID(i)
	return trc
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (trc *TaskRecordCreate) SetTaskID(id int) *TaskRecordCreate {
	trc.mutation.SetTaskID(id)
	return trc
}

// SetTask sets the "task" edge to the Task entity.
func (trc *TaskRecordCreate) SetTask(t *Task) *TaskRecordCreate {
	return trc.SetTaskID(t.ID)
}

// AddTaskLogIDs adds the "task_logs" edge to the TaskLog entity by IDs.
func (trc *TaskRecordCreate) AddTaskLogIDs(ids ...int) *TaskRecordCreate {
	trc.mutation.AddTaskLogIDs(ids...)
	return trc
}

// AddTaskLogs adds the "task_logs" edges to the TaskLog entity.
func (trc *TaskRecordCreate) AddTaskLogs(t ...*TaskLog) *TaskRecordCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return trc.AddTaskLogIDs(ids...)
}

// Mutation returns the TaskRecordMutation object of the builder.
func (trc *TaskRecordCreate) Mutation() *TaskRecordMutation {
	return trc.mutation
}

// Save creates the TaskRecord in the database.
func (trc *TaskRecordCreate) Save(ctx context.Context) (*TaskRecord, error) {
	trc.defaults()
	return withHooks(ctx, trc.sqlSave, trc.mutation, trc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (trc *TaskRecordCreate) SaveX(ctx context.Context) *TaskRecord {
	v, err := trc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (trc *TaskRecordCreate) Exec(ctx context.Context) error {
	_, err := trc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (trc *TaskRecordCreate) ExecX(ctx context.Context) {
	if err := trc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (trc *TaskRecordCreate) defaults() {
	if _, ok := trc.mutation.State(); !ok {
		v := taskrecord.DefaultState
		trc.mutation.SetState(v)
	}
	if _, ok := trc.mutation.RetryCount(); !ok {
		v := taskrecord.DefaultRetryCount
		trc.mutation.SetRetryCount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (trc *TaskRecordCreate) check() error {
	if _, ok := trc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "TaskRecord.state"`)}
	}
	if _, ok := trc.mutation.RetryCount(); !ok {
		return &ValidationError{Name: "retry_count", err: errors.New(`ent: missing required field "TaskRecord.retry_count"`)}
	}
	if len(trc.mutation.TaskIDs()) == 0 {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required edge "TaskRecord.task"`)}
	}
	return nil
}

func (trc *TaskRecordCreate) sqlSave(ctx context.Context) (*TaskRecord, error) {
	if err := trc.check(); err != nil {
		return nil, err
	}
	_node, _spec := trc.createSpec()
	if err := sqlgraph.CreateNode(ctx, trc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	trc.mutation.id = &_node.ID
	trc.mutation.done = true
	return _node, nil
}

func (trc *TaskRecordCreate) createSpec() (*TaskRecord, *sqlgraph.CreateSpec) {
	var (
		_node = &TaskRecord{config: trc.config}
		_spec = sqlgraph.NewCreateSpec(taskrecord.Table, sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt))
	)
	if id, ok := trc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := trc.mutation.State(); ok {
		_spec.SetField(taskrecord.FieldState, field.TypeString, value)
		_node.State = value
	}
	if value, ok := trc.mutation.RetryCount(); ok {
		_spec.SetField(taskrecord.FieldRetryCount, field.TypeInt, value)
		_node.RetryCount = value
	}
	if nodes := trc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   taskrecord.TaskTable,
			Columns: []string{taskrecord.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.task_task_record = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := trc.mutation.TaskLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   taskrecord.TaskLogsTable,
			Columns: taskrecord.TaskLogsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tasklog.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TaskRecordCreateBulk is the builder for creating many TaskRecord entities in bulk.
type TaskRecordCreateBulk struct {
	config
	err      error
	builders []*TaskRecordCreate
}

// Save creates the TaskRecord entities in the database.
func (trcb *TaskRecordCreateBulk) Save(ctx context.Context) ([]*TaskRecord, error) {
	if trcb.err != nil {
		return nil, trcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(trcb.builders))
	nodes := make([]*TaskRecord, len(trcb.builders))
	mutators := make([]Mutator, len(trcb.builders))
	for i := range trcb.builders {
		func(i int, root context.Context) {
			builder := trcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskRecordMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, trcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, trcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, trcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (trcb *TaskRecordCreateBulk) SaveX(ctx context.Context) []*TaskRecord {
	v, err := trcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (trcb *TaskRecordCreateBulk) Exec(ctx context.Context) error {
	_, err := trcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (trcb *TaskRecordCreateBulk) ExecX(ctx context.Context) {
	if err := trcb.Exec(ctx); err != nil {
		panic(err)
	}
}