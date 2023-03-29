package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"net/url"
	"time"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "task"},
	}
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Uint32("created_by"),
		field.Int8("rank").Default(1),
		field.Int16("type"),
		field.Int16("status"),
		field.Time("complete_at").Optional().Nillable(),
		field.Time("update_at").UpdateDefault(time.Now),
		field.Uint32("update_by"),
		field.Bool("is_deleted").Optional().Nillable(),
		field.Time("deleted_at").Optional().Nillable(),
		field.Uint32("deleted_by").Optional().Nillable(),
		field.Text("description"),
		field.JSON("url", &url.URL{}),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}
