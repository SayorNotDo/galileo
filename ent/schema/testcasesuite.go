package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
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
		field.Int32("id"),
		field.String("name").Unique().NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Uint32("created_by").Immutable(),
		field.Time("updated_at").UpdateDefault(time.Now).Optional(),
		field.Uint32("updated_by").Optional(),
		field.JSON("testcases", []int32{}).Optional(),
	}
}

// Edges of the TestCaseSuite.
func (TestcaseSuite) Edges() []ent.Edge {
	return []ent.Edge{}
}
