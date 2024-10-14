package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TaskLog holds the schema definition for the TaskLog entity.
type TaskLog struct {
	ent.Schema
}

// Fields of the TaskLog.
func (TaskLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("message").Default("").Comment("메시지"),
	}
}

// Edges of the TaskLog.
func (TaskLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task_records", TaskRecord.Type).Ref("task_logs"), // many to many
	}
}
