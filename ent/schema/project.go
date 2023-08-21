package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

func (Project) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "project"},
	}
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("name").Unique().NotEmpty(),
		field.String("identifier").Unique().NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Uint32("created_by").Immutable(),
		field.Time("updated_at").Optional().UpdateDefault(time.Now),
		field.Uint32("updated_by").Optional(),
		field.Time("deleted_at").Optional(),
		field.Uint32("deleted_by").Optional(),
		field.Int8("status").Default(0),
		field.Time("start_time").Optional(),
		field.Time("deadline").Optional(),
		field.Text("description").Optional(),
		field.Text("remark").Optional(),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{}
}
