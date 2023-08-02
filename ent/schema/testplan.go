package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// TestPlan holds the schema definition for the TestPlan entity.
type TestPlan struct {
	ent.Schema
}

func (TestPlan) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "test_plan"},
	}
}

// Fields of the TestPlan.
func (TestPlan) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").Unique().NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Uint32("created_by").Immutable(),
		field.Time("updated_at").Optional().UpdateDefault(time.Now),
		field.Uint32("updated_by").Optional(),
		field.Text("description").Optional(),
	}
}

// Edges of the TestPlan.
func (TestPlan) Edges() []ent.Edge {
	return nil
}
