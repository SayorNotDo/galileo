package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// ProjectMember holds the schema definition for the ProjectMember entity.
type ProjectMember struct {
	ent.Schema
}

func (ProjectMember) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "project_member",
		},
	}
}

// Fields of the ProjectMember.
func (ProjectMember) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("project_id"),
		field.Uint32("user_id"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Uint32("created_by").Immutable(),
		field.Time("deleted_at").Optional(),
		field.Uint32("deleted_by").Optional(),
		field.Int8("status").Default(0),
		field.String("description").Optional(),
		field.String("remark").Optional(),
		field.Uint8("role").Default(0),
	}
}

// Edges of the ProjectMember.
func (ProjectMember) Edges() []ent.Edge {
	return nil
}
