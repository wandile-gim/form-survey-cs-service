package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("type").Optional().Comment("업무 타입"),
		field.Int("row_num").Unique().Comment("행 번호"),
		field.String("name").Default("unknown"),
		field.String("phone").Comment("전화번호"),
		field.String("group").Comment("소속"),
		field.String("corps").Optional(),
		field.String("food").Optional(),
		field.String("gender").Optional(),
		field.String("generation").Optional(),
		field.String("region").Comment("지역"),
		field.Time("registered_at").Optional(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("task_record", TaskRecord.Type).
			Unique(),
	}
}
