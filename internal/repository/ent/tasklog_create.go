// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"form-survey-cs-service/internal/repository/ent/tasklog"
	"form-survey-cs-service/internal/repository/ent/taskrecord"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskLogCreate is the builder for creating a TaskLog entity.
type TaskLogCreate struct {
	config
	mutation *TaskLogMutation
	hooks    []Hook
}

// SetMessage sets the "message" field.
func (tlc *TaskLogCreate) SetMessage(s string) *TaskLogCreate {
	tlc.mutation.SetMessage(s)
	return tlc
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (tlc *TaskLogCreate) SetNillableMessage(s *string) *TaskLogCreate {
	if s != nil {
		tlc.SetMessage(*s)
	}
	return tlc
}

// SetID sets the "id" field.
func (tlc *TaskLogCreate) SetID(i int) *TaskLogCreate {
	tlc.mutation.SetID(i)
	return tlc
}

// AddTaskRecordIDs adds the "task_records" edge to the TaskRecord entity by IDs.
func (tlc *TaskLogCreate) AddTaskRecordIDs(ids ...int) *TaskLogCreate {
	tlc.mutation.AddTaskRecordIDs(ids...)
	return tlc
}

// AddTaskRecords adds the "task_records" edges to the TaskRecord entity.
func (tlc *TaskLogCreate) AddTaskRecords(t ...*TaskRecord) *TaskLogCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tlc.AddTaskRecordIDs(ids...)
}

// Mutation returns the TaskLogMutation object of the builder.
func (tlc *TaskLogCreate) Mutation() *TaskLogMutation {
	return tlc.mutation
}

// Save creates the TaskLog in the database.
func (tlc *TaskLogCreate) Save(ctx context.Context) (*TaskLog, error) {
	tlc.defaults()
	return withHooks(ctx, tlc.sqlSave, tlc.mutation, tlc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tlc *TaskLogCreate) SaveX(ctx context.Context) *TaskLog {
	v, err := tlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tlc *TaskLogCreate) Exec(ctx context.Context) error {
	_, err := tlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlc *TaskLogCreate) ExecX(ctx context.Context) {
	if err := tlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tlc *TaskLogCreate) defaults() {
	if _, ok := tlc.mutation.Message(); !ok {
		v := tasklog.DefaultMessage
		tlc.mutation.SetMessage(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tlc *TaskLogCreate) check() error {
	if _, ok := tlc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "TaskLog.message"`)}
	}
	return nil
}

func (tlc *TaskLogCreate) sqlSave(ctx context.Context) (*TaskLog, error) {
	if err := tlc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	tlc.mutation.id = &_node.ID
	tlc.mutation.done = true
	return _node, nil
}

func (tlc *TaskLogCreate) createSpec() (*TaskLog, *sqlgraph.CreateSpec) {
	var (
		_node = &TaskLog{config: tlc.config}
		_spec = sqlgraph.NewCreateSpec(tasklog.Table, sqlgraph.NewFieldSpec(tasklog.FieldID, field.TypeInt))
	)
	if id, ok := tlc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tlc.mutation.Message(); ok {
		_spec.SetField(tasklog.FieldMessage, field.TypeString, value)
		_node.Message = value
	}
	if nodes := tlc.mutation.TaskRecordsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TaskLogCreateBulk is the builder for creating many TaskLog entities in bulk.
type TaskLogCreateBulk struct {
	config
	err      error
	builders []*TaskLogCreate
}

// Save creates the TaskLog entities in the database.
func (tlcb *TaskLogCreateBulk) Save(ctx context.Context) ([]*TaskLog, error) {
	if tlcb.err != nil {
		return nil, tlcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tlcb.builders))
	nodes := make([]*TaskLog, len(tlcb.builders))
	mutators := make([]Mutator, len(tlcb.builders))
	for i := range tlcb.builders {
		func(i int, root context.Context) {
			builder := tlcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskLogMutation)
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
					_, err = mutators[i+1].Mutate(root, tlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tlcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tlcb *TaskLogCreateBulk) SaveX(ctx context.Context) []*TaskLog {
	v, err := tlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tlcb *TaskLogCreateBulk) Exec(ctx context.Context) error {
	_, err := tlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tlcb *TaskLogCreateBulk) ExecX(ctx context.Context) {
	if err := tlcb.Exec(ctx); err != nil {
		panic(err)
	}
}
