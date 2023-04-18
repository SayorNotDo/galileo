package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TestCaseSuite holds the schema definition for the TestCaseSuite entity.
type TestCaseSuite struct {
	ent.Schema
}

// Fields of the TestCaseSuite.
func (TestCaseSuite) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").Unique().NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Uint32("created_by").Immutable(),
	}
}

// Edges of the TestCaseSuite.
func (TestCaseSuite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("testcases", TestCase.Type),
	}
}
