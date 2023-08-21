package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ApiTag holds the schema definition for the ApiTag entity.
type ApiTag struct {
	ent.Schema
}

func (ApiTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "api_tag"},
	}
}

// Fields of the ApiTag.
func (ApiTag) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("name"),
		field.Int32("sort"),
		field.Time("created_at"),
		field.Uint32("created_by"),
		field.Time("update_at"),
		field.Uint32("update_by"),
		field.String("description"),
	}
}

// Edges of the ApiTag.
func (ApiTag) Edges() []ent.Edge {
	return nil
}
