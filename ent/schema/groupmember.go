package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// GroupMember holds the schema definition for the GroupMember entity.
type GroupMember struct {
	ent.Schema
}

func (GroupMember) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_member"},
	}
}

// Fields of the GroupMember.
func (GroupMember) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("group_id"),
		field.Int64("user_id"),
		field.Uint8("role").Default(0),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Uint32("created_by"),
		field.Time("deleted_at").Optional(),
		field.Uint32("deleted_by").Optional(),
	}
}

func (GroupMember) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("group_id", "user_id").Unique(),
	}
}

// Edges of the GroupMember.
func (GroupMember) Edges() []ent.Edge {
	return nil
}
