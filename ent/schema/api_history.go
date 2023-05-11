package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ApiHistory holds the schema definition for the ApiHistory entity.
type ApiHistory struct {
	ent.Schema
}

func (ApiHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "api_history"},
	}
}

// Fields of the ApiHistory.
func (ApiHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("version"),
		field.String("query_params"),
		field.Time("created_at"),
		field.Uint32("created_by"),
		field.String("description"),
	}
}

// Edges of the ApiHistory.
func (ApiHistory) Edges() []ent.Edge {
	return nil
}
