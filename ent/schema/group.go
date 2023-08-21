package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"regexp"
	"time"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("name").Match(regexp.MustCompile("[a-zA-Z_]+$")),
		field.Uint32("created_by"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{}
}
