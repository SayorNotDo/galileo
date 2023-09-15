package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"regexp"
	"time"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group"},
	}
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("name").Match(regexp.MustCompile("[a-zA-Z_]+$")),
		field.Text("avatar").Optional(),
		field.Text("description").Optional(),
		field.Uint32("created_by"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").Optional().UpdateDefault(time.Now),
		field.Uint32("updated_by").Optional(),
		field.Time("deleted_at").Optional(),
		field.Uint32("deleted_by").Optional(),
		field.Int32("headcount"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{}
}
