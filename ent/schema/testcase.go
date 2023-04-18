package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TestCase holds the schema definition for the TestCase entity.
type TestCase struct {
	ent.Schema
}

// Fields of the TestCase.
func (TestCase) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").Unique().NotEmpty(),
		field.Uint32("created_by").Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Uint32("update_by").Optional().Nillable(),
		field.Time("update_at").UpdateDefault(time.Now).Optional().Nillable(),
		field.Int8("status").Default(0),
		field.Int8("type").Default(0),
		field.Int8("priority").Default(0),
		field.Time("deleted_at").Optional(),
		field.Uint32("deleted_by").Optional(),
		field.String("description").Optional(),
		field.String("url").Optional(),
	}
}

// Edges of the TestCase.
func (TestCase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("testcaseSuites", TestCaseSuite.Type).Ref("testcases"),
	}
}
