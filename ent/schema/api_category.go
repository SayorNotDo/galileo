package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// ApiCategory holds the schema definition for the ApiCategory entity.
type ApiCategory struct {
	ent.Schema
}

func (ApiCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "api_category"},
	}
}

// Fields of the ApiCategory.
func (ApiCategory) Fields() []ent.Field {
	return nil
}

// Edges of the ApiCategory.
func (ApiCategory) Edges() []ent.Edge {
	return nil
}
