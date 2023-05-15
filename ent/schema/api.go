package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Api holds the schema definition for the Api entity.
type Api struct {
	ent.Schema
}

func (Api) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "api"},
	}
}

func (Api) Indexes() []ent.Index {
	return []ent.Index{
		// 唯一约束索引
		index.Fields("url", "type").
			Unique(),
	}
}

// Fields of the Api.
func (Api) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").Unique().NotEmpty(),
		field.String("url").Optional(),
		field.Int8("type").Optional(),
		field.Int8("status").Default(0),
		field.String("headers").Optional().Nillable(),
		field.String("body").Optional().Nillable(),
		field.String("query_params").Optional(),
		field.String("response").Optional(),
		field.String("module").Optional(),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Uint32("created_by").Immutable(),
		field.String("include_files").Optional().Nillable(),
		field.Time("update_at").UpdateDefault(time.Now).Optional().Nillable(),
		field.Uint32("update_by"),
		field.Time("deleted_at").Optional().Nillable(),
		field.Uint32("deleted_by").Optional().Nillable(),
	}
}

// Edges of the Api.
func (Api) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("statistics", ApiStatistics.Type).
			Unique(),
	}
}