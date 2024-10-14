package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tracker holds the schema definition for the Tracker entity.
type Tracker struct {
	ent.Schema
}

// Fields of the Tracker.
func (Tracker) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("service").
			Values("MEMBER").  // Enum 필드의 값 정의
			Default("MEMBER"), // 기본값 설정

		field.Time("last_one"),
		field.Time("version"),
	}
}

// Edges of the Tracker.
func (Tracker) Edges() []ent.Edge {
	return nil
}
