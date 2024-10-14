package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TaskRecord holds the schema definition for the TaskRecord entity.
type TaskRecord struct {
	ent.Schema
}

// Fields of the TaskRecord.
func (TaskRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("state").Default("IDLE"),
		field.Int("retry_count").Default(0),
	}
}

// Edges of the TaskRecord.
func (TaskRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).
			Ref("task_record").
			Unique().
			Required(),
		edge.To("task_logs", TaskLog.Type),
	}
}
