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
		field.Uint32("id"),
		field.String("name").Unique().NotEmpty(),
		field.String("identifier").Unique(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Uint32("created_by").Immutable(),
		field.Time("updated_at").UpdateDefault(time.Now),
		field.Uint32("update_by").Optional().Nillable(),
		field.Time("deleted_at").Optional().Nillable(),
		field.Uint32("deleted_by").Optional().Nillable(),
		field.Int8("status").Default(0),
		field.Text("description").Optional().Nillable(),
		field.Text("remark").Optional().Nillable(),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return nil
}
