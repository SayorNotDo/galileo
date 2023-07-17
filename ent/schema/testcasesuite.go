package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TestcaseSuite holds the schema definition for the TestcaseSuite entity.
type TestcaseSuite struct {
	ent.Schema
}

func (TestcaseSuite) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "testcase_suite"},
	}
}

// Fields of the TestCaseSuite.
func (TestcaseSuite) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").Unique().NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Uint32("created_by").Immutable(),
	}
}

// Edges of the TestCaseSuite.
func (TestcaseSuite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("testcase", Testcase.Type),
	}
}
