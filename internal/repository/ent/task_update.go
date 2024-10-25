// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"form-survey-cs-service/internal/repository/ent/predicate"
	"form-survey-cs-service/internal/repository/ent/task"
	"form-survey-cs-service/internal/repository/ent/taskrecord"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetType sets the "type" field.
func (tu *TaskUpdate) SetType(s string) *TaskUpdate {
	tu.mutation.SetType(s)
	return tu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableType(s *string) *TaskUpdate {
	if s != nil {
		tu.SetType(*s)
	}
	return tu
}

// ClearType clears the value of the "type" field.
func (tu *TaskUpdate) ClearType() *TaskUpdate {
	tu.mutation.ClearType()
	return tu
}

// SetRowNum sets the "row_num" field.
func (tu *TaskUpdate) SetRowNum(i int) *TaskUpdate {
	tu.mutation.ResetRowNum()
	tu.mutation.SetRowNum(i)
	return tu
}

// SetNillableRowNum sets the "row_num" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableRowNum(i *int) *TaskUpdate {
	if i != nil {
		tu.SetRowNum(*i)
	}
	return tu
}

// AddRowNum adds i to the "row_num" field.
func (tu *TaskUpdate) AddRowNum(i int) *TaskUpdate {
	tu.mutation.AddRowNum(i)
	return tu
}

// SetName sets the "name" field.
func (tu *TaskUpdate) SetName(s string) *TaskUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableName(s *string) *TaskUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// SetPhone sets the "phone" field.
func (tu *TaskUpdate) SetPhone(s string) *TaskUpdate {
	tu.mutation.SetPhone(s)
	return tu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (tu *TaskUpdate) SetNillablePhone(s *string) *TaskUpdate {
	if s != nil {
		tu.SetPhone(*s)
	}
	return tu
}

// SetPayAmount sets the "pay_amount" field.
func (tu *TaskUpdate) SetPayAmount(f float64) *TaskUpdate {
	tu.mutation.ResetPayAmount()
	tu.mutation.SetPayAmount(f)
	return tu
}

// SetNillablePayAmount sets the "pay_amount" field if the given value is not nil.
func (tu *TaskUpdate) SetNillablePayAmount(f *float64) *TaskUpdate {
	if f != nil {
		tu.SetPayAmount(*f)
	}
	return tu
}

// AddPayAmount adds f to the "pay_amount" field.
func (tu *TaskUpdate) AddPayAmount(f float64) *TaskUpdate {
	tu.mutation.AddPayAmount(f)
	return tu
}

// ClearPayAmount clears the value of the "pay_amount" field.
func (tu *TaskUpdate) ClearPayAmount() *TaskUpdate {
	tu.mutation.ClearPayAmount()
	return tu
}

// SetPaidAt sets the "paid_at" field.
func (tu *TaskUpdate) SetPaidAt(s string) *TaskUpdate {
	tu.mutation.SetPaidAt(s)
	return tu
}

// SetNillablePaidAt sets the "paid_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillablePaidAt(s *string) *TaskUpdate {
	if s != nil {
		tu.SetPaidAt(*s)
	}
	return tu
}

// ClearPaidAt clears the value of the "paid_at" field.
func (tu *TaskUpdate) ClearPaidAt() *TaskUpdate {
	tu.mutation.ClearPaidAt()
	return tu
}

// SetGroup sets the "group" field.
func (tu *TaskUpdate) SetGroup(s string) *TaskUpdate {
	tu.mutation.SetGroup(s)
	return tu
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableGroup(s *string) *TaskUpdate {
	if s != nil {
		tu.SetGroup(*s)
	}
	return tu
}

// SetCorps sets the "corps" field.
func (tu *TaskUpdate) SetCorps(s string) *TaskUpdate {
	tu.mutation.SetCorps(s)
	return tu
}

// SetNillableCorps sets the "corps" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCorps(s *string) *TaskUpdate {
	if s != nil {
		tu.SetCorps(*s)
	}
	return tu
}

// ClearCorps clears the value of the "corps" field.
func (tu *TaskUpdate) ClearCorps() *TaskUpdate {
	tu.mutation.ClearCorps()
	return tu
}

// SetFood sets the "food" field.
func (tu *TaskUpdate) SetFood(s string) *TaskUpdate {
	tu.mutation.SetFood(s)
	return tu
}

// SetNillableFood sets the "food" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableFood(s *string) *TaskUpdate {
	if s != nil {
		tu.SetFood(*s)
	}
	return tu
}

// ClearFood clears the value of the "food" field.
func (tu *TaskUpdate) ClearFood() *TaskUpdate {
	tu.mutation.ClearFood()
	return tu
}

// SetGender sets the "gender" field.
func (tu *TaskUpdate) SetGender(s string) *TaskUpdate {
	tu.mutation.SetGender(s)
	return tu
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableGender(s *string) *TaskUpdate {
	if s != nil {
		tu.SetGender(*s)
	}
	return tu
}

// ClearGender clears the value of the "gender" field.
func (tu *TaskUpdate) ClearGender() *TaskUpdate {
	tu.mutation.ClearGender()
	return tu
}

// SetGeneration sets the "generation" field.
func (tu *TaskUpdate) SetGeneration(s string) *TaskUpdate {
	tu.mutation.SetGeneration(s)
	return tu
}

// SetNillableGeneration sets the "generation" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableGeneration(s *string) *TaskUpdate {
	if s != nil {
		tu.SetGeneration(*s)
	}
	return tu
}

// ClearGeneration clears the value of the "generation" field.
func (tu *TaskUpdate) ClearGeneration() *TaskUpdate {
	tu.mutation.ClearGeneration()
	return tu
}

// SetRegion sets the "region" field.
func (tu *TaskUpdate) SetRegion(s string) *TaskUpdate {
	tu.mutation.SetRegion(s)
	return tu
}

// SetNillableRegion sets the "region" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableRegion(s *string) *TaskUpdate {
	if s != nil {
		tu.SetRegion(*s)
	}
	return tu
}

// SetRegisteredAt sets the "registered_at" field.
func (tu *TaskUpdate) SetRegisteredAt(t time.Time) *TaskUpdate {
	tu.mutation.SetRegisteredAt(t)
	return tu
}

// SetNillableRegisteredAt sets the "registered_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableRegisteredAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetRegisteredAt(*t)
	}
	return tu
}

// ClearRegisteredAt clears the value of the "registered_at" field.
func (tu *TaskUpdate) ClearRegisteredAt() *TaskUpdate {
	tu.mutation.ClearRegisteredAt()
	return tu
}

// SetTaskRecordID sets the "task_record" edge to the TaskRecord entity by ID.
func (tu *TaskUpdate) SetTaskRecordID(id int) *TaskUpdate {
	tu.mutation.SetTaskRecordID(id)
	return tu
}

// SetNillableTaskRecordID sets the "task_record" edge to the TaskRecord entity by ID if the given value is not nil.
func (tu *TaskUpdate) SetNillableTaskRecordID(id *int) *TaskUpdate {
	if id != nil {
		tu = tu.SetTaskRecordID(*id)
	}
	return tu
}

// SetTaskRecord sets the "task_record" edge to the TaskRecord entity.
func (tu *TaskUpdate) SetTaskRecord(t *TaskRecord) *TaskUpdate {
	return tu.SetTaskRecordID(t.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// ClearTaskRecord clears the "task_record" edge to the TaskRecord entity.
func (tu *TaskUpdate) ClearTaskRecord() *TaskUpdate {
	tu.mutation.ClearTaskRecord()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.SetField(task.FieldType, field.TypeString, value)
	}
	if tu.mutation.TypeCleared() {
		_spec.ClearField(task.FieldType, field.TypeString)
	}
	if value, ok := tu.mutation.RowNum(); ok {
		_spec.SetField(task.FieldRowNum, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedRowNum(); ok {
		_spec.AddField(task.FieldRowNum, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(task.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Phone(); ok {
		_spec.SetField(task.FieldPhone, field.TypeString, value)
	}
	if value, ok := tu.mutation.PayAmount(); ok {
		_spec.SetField(task.FieldPayAmount, field.TypeFloat64, value)
	}
	if value, ok := tu.mutation.AddedPayAmount(); ok {
		_spec.AddField(task.FieldPayAmount, field.TypeFloat64, value)
	}
	if tu.mutation.PayAmountCleared() {
		_spec.ClearField(task.FieldPayAmount, field.TypeFloat64)
	}
	if value, ok := tu.mutation.PaidAt(); ok {
		_spec.SetField(task.FieldPaidAt, field.TypeString, value)
	}
	if tu.mutation.PaidAtCleared() {
		_spec.ClearField(task.FieldPaidAt, field.TypeString)
	}
	if value, ok := tu.mutation.Group(); ok {
		_spec.SetField(task.FieldGroup, field.TypeString, value)
	}
	if value, ok := tu.mutation.Corps(); ok {
		_spec.SetField(task.FieldCorps, field.TypeString, value)
	}
	if tu.mutation.CorpsCleared() {
		_spec.ClearField(task.FieldCorps, field.TypeString)
	}
	if value, ok := tu.mutation.Food(); ok {
		_spec.SetField(task.FieldFood, field.TypeString, value)
	}
	if tu.mutation.FoodCleared() {
		_spec.ClearField(task.FieldFood, field.TypeString)
	}
	if value, ok := tu.mutation.Gender(); ok {
		_spec.SetField(task.FieldGender, field.TypeString, value)
	}
	if tu.mutation.GenderCleared() {
		_spec.ClearField(task.FieldGender, field.TypeString)
	}
	if value, ok := tu.mutation.Generation(); ok {
		_spec.SetField(task.FieldGeneration, field.TypeString, value)
	}
	if tu.mutation.GenerationCleared() {
		_spec.ClearField(task.FieldGeneration, field.TypeString)
	}
	if value, ok := tu.mutation.Region(); ok {
		_spec.SetField(task.FieldRegion, field.TypeString, value)
	}
	if value, ok := tu.mutation.RegisteredAt(); ok {
		_spec.SetField(task.FieldRegisteredAt, field.TypeTime, value)
	}
	if tu.mutation.RegisteredAtCleared() {
		_spec.ClearField(task.FieldRegisteredAt, field.TypeTime)
	}
	if tu.mutation.TaskRecordCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   task.TaskRecordTable,
			Columns: []string{task.TaskRecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TaskRecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   task.TaskRecordTable,
			Columns: []string{task.TaskRecordColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetType sets the "type" field.
func (tuo *TaskUpdateOne) SetType(s string) *TaskUpdateOne {
	tuo.mutation.SetType(s)
	return tuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableType(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetType(*s)
	}
	return tuo
}

// ClearType clears the value of the "type" field.
func (tuo *TaskUpdateOne) ClearType() *TaskUpdateOne {
	tuo.mutation.ClearType()
	return tuo
}

// SetRowNum sets the "row_num" field.
func (tuo *TaskUpdateOne) SetRowNum(i int) *TaskUpdateOne {
	tuo.mutation.ResetRowNum()
	tuo.mutation.SetRowNum(i)
	return tuo
}

// SetNillableRowNum sets the "row_num" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableRowNum(i *int) *TaskUpdateOne {
	if i != nil {
		tuo.SetRowNum(*i)
	}
	return tuo
}

// AddRowNum adds i to the "row_num" field.
func (tuo *TaskUpdateOne) AddRowNum(i int) *TaskUpdateOne {
	tuo.mutation.AddRowNum(i)
	return tuo
}

// SetName sets the "name" field.
func (tuo *TaskUpdateOne) SetName(s string) *TaskUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableName(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// SetPhone sets the "phone" field.
func (tuo *TaskUpdateOne) SetPhone(s string) *TaskUpdateOne {
	tuo.mutation.SetPhone(s)
	return tuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillablePhone(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetPhone(*s)
	}
	return tuo
}

// SetPayAmount sets the "pay_amount" field.
func (tuo *TaskUpdateOne) SetPayAmount(f float64) *TaskUpdateOne {
	tuo.mutation.ResetPayAmount()
	tuo.mutation.SetPayAmount(f)
	return tuo
}

// SetNillablePayAmount sets the "pay_amount" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillablePayAmount(f *float64) *TaskUpdateOne {
	if f != nil {
		tuo.SetPayAmount(*f)
	}
	return tuo
}

// AddPayAmount adds f to the "pay_amount" field.
func (tuo *TaskUpdateOne) AddPayAmount(f float64) *TaskUpdateOne {
	tuo.mutation.AddPayAmount(f)
	return tuo
}

// ClearPayAmount clears the value of the "pay_amount" field.
func (tuo *TaskUpdateOne) ClearPayAmount() *TaskUpdateOne {
	tuo.mutation.ClearPayAmount()
	return tuo
}

// SetPaidAt sets the "paid_at" field.
func (tuo *TaskUpdateOne) SetPaidAt(s string) *TaskUpdateOne {
	tuo.mutation.SetPaidAt(s)
	return tuo
}

// SetNillablePaidAt sets the "paid_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillablePaidAt(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetPaidAt(*s)
	}
	return tuo
}

// ClearPaidAt clears the value of the "paid_at" field.
func (tuo *TaskUpdateOne) ClearPaidAt() *TaskUpdateOne {
	tuo.mutation.ClearPaidAt()
	return tuo
}

// SetGroup sets the "group" field.
func (tuo *TaskUpdateOne) SetGroup(s string) *TaskUpdateOne {
	tuo.mutation.SetGroup(s)
	return tuo
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableGroup(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetGroup(*s)
	}
	return tuo
}

// SetCorps sets the "corps" field.
func (tuo *TaskUpdateOne) SetCorps(s string) *TaskUpdateOne {
	tuo.mutation.SetCorps(s)
	return tuo
}

// SetNillableCorps sets the "corps" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCorps(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetCorps(*s)
	}
	return tuo
}

// ClearCorps clears the value of the "corps" field.
func (tuo *TaskUpdateOne) ClearCorps() *TaskUpdateOne {
	tuo.mutation.ClearCorps()
	return tuo
}

// SetFood sets the "food" field.
func (tuo *TaskUpdateOne) SetFood(s string) *TaskUpdateOne {
	tuo.mutation.SetFood(s)
	return tuo
}

// SetNillableFood sets the "food" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableFood(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetFood(*s)
	}
	return tuo
}

// ClearFood clears the value of the "food" field.
func (tuo *TaskUpdateOne) ClearFood() *TaskUpdateOne {
	tuo.mutation.ClearFood()
	return tuo
}

// SetGender sets the "gender" field.
func (tuo *TaskUpdateOne) SetGender(s string) *TaskUpdateOne {
	tuo.mutation.SetGender(s)
	return tuo
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableGender(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetGender(*s)
	}
	return tuo
}

// ClearGender clears the value of the "gender" field.
func (tuo *TaskUpdateOne) ClearGender() *TaskUpdateOne {
	tuo.mutation.ClearGender()
	return tuo
}

// SetGeneration sets the "generation" field.
func (tuo *TaskUpdateOne) SetGeneration(s string) *TaskUpdateOne {
	tuo.mutation.SetGeneration(s)
	return tuo
}

// SetNillableGeneration sets the "generation" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableGeneration(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetGeneration(*s)
	}
	return tuo
}

// ClearGeneration clears the value of the "generation" field.
func (tuo *TaskUpdateOne) ClearGeneration() *TaskUpdateOne {
	tuo.mutation.ClearGeneration()
	return tuo
}

// SetRegion sets the "region" field.
func (tuo *TaskUpdateOne) SetRegion(s string) *TaskUpdateOne {
	tuo.mutation.SetRegion(s)
	return tuo
}

// SetNillableRegion sets the "region" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableRegion(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetRegion(*s)
	}
	return tuo
}

// SetRegisteredAt sets the "registered_at" field.
func (tuo *TaskUpdateOne) SetRegisteredAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetRegisteredAt(t)
	return tuo
}

// SetNillableRegisteredAt sets the "registered_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableRegisteredAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetRegisteredAt(*t)
	}
	return tuo
}

// ClearRegisteredAt clears the value of the "registered_at" field.
func (tuo *TaskUpdateOne) ClearRegisteredAt() *TaskUpdateOne {
	tuo.mutation.ClearRegisteredAt()
	return tuo
}

// SetTaskRecordID sets the "task_record" edge to the TaskRecord entity by ID.
func (tuo *TaskUpdateOne) SetTaskRecordID(id int) *TaskUpdateOne {
	tuo.mutation.SetTaskRecordID(id)
	return tuo
}

// SetNillableTaskRecordID sets the "task_record" edge to the TaskRecord entity by ID if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableTaskRecordID(id *int) *TaskUpdateOne {
	if id != nil {
		tuo = tuo.SetTaskRecordID(*id)
	}
	return tuo
}

// SetTaskRecord sets the "task_record" edge to the TaskRecord entity.
func (tuo *TaskUpdateOne) SetTaskRecord(t *TaskRecord) *TaskUpdateOne {
	return tuo.SetTaskRecordID(t.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// ClearTaskRecord clears the "task_record" edge to the TaskRecord entity.
func (tuo *TaskUpdateOne) ClearTaskRecord() *TaskUpdateOne {
	tuo.mutation.ClearTaskRecord()
	return tuo
}

// Where appends a list predicates to the TaskUpdate builder.
func (tuo *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.SetField(task.FieldType, field.TypeString, value)
	}
	if tuo.mutation.TypeCleared() {
		_spec.ClearField(task.FieldType, field.TypeString)
	}
	if value, ok := tuo.mutation.RowNum(); ok {
		_spec.SetField(task.FieldRowNum, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedRowNum(); ok {
		_spec.AddField(task.FieldRowNum, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(task.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Phone(); ok {
		_spec.SetField(task.FieldPhone, field.TypeString, value)
	}
	if value, ok := tuo.mutation.PayAmount(); ok {
		_spec.SetField(task.FieldPayAmount, field.TypeFloat64, value)
	}
	if value, ok := tuo.mutation.AddedPayAmount(); ok {
		_spec.AddField(task.FieldPayAmount, field.TypeFloat64, value)
	}
	if tuo.mutation.PayAmountCleared() {
		_spec.ClearField(task.FieldPayAmount, field.TypeFloat64)
	}
	if value, ok := tuo.mutation.PaidAt(); ok {
		_spec.SetField(task.FieldPaidAt, field.TypeString, value)
	}
	if tuo.mutation.PaidAtCleared() {
		_spec.ClearField(task.FieldPaidAt, field.TypeString)
	}
	if value, ok := tuo.mutation.Group(); ok {
		_spec.SetField(task.FieldGroup, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Corps(); ok {
		_spec.SetField(task.FieldCorps, field.TypeString, value)
	}
	if tuo.mutation.CorpsCleared() {
		_spec.ClearField(task.FieldCorps, field.TypeString)
	}
	if value, ok := tuo.mutation.Food(); ok {
		_spec.SetField(task.FieldFood, field.TypeString, value)
	}
	if tuo.mutation.FoodCleared() {
		_spec.ClearField(task.FieldFood, field.TypeString)
	}
	if value, ok := tuo.mutation.Gender(); ok {
		_spec.SetField(task.FieldGender, field.TypeString, value)
	}
	if tuo.mutation.GenderCleared() {
		_spec.ClearField(task.FieldGender, field.TypeString)
	}
	if value, ok := tuo.mutation.Generation(); ok {
		_spec.SetField(task.FieldGeneration, field.TypeString, value)
	}
	if tuo.mutation.GenerationCleared() {
		_spec.ClearField(task.FieldGeneration, field.TypeString)
	}
	if value, ok := tuo.mutation.Region(); ok {
		_spec.SetField(task.FieldRegion, field.TypeString, value)
	}
	if value, ok := tuo.mutation.RegisteredAt(); ok {
		_spec.SetField(task.FieldRegisteredAt, field.TypeTime, value)
	}
	if tuo.mutation.RegisteredAtCleared() {
		_spec.ClearField(task.FieldRegisteredAt, field.TypeTime)
	}
	if tuo.mutation.TaskRecordCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   task.TaskRecordTable,
			Columns: []string{task.TaskRecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(taskrecord.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TaskRecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   task.TaskRecordTable,
			Columns: []string{task.TaskRecordColumn},
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
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
