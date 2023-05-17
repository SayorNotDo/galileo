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
		field.String("url").NotEmpty(),
		field.Int8("type").NonNegative(),
		field.Int8("status").Default(0),
		field.String("headers").Optional(),
		field.String("body").Optional(),
		field.String("label").Optional(),
		field.String("query_params").Optional(),
		field.String("response").Optional(),
		field.String("module").Optional(),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Uint32("created_by").Immutable(),
		field.String("include_files").Optional(),
		field.Time("update_at").UpdateDefault(time.Now).Optional(),
		field.Uint32("update_by").Optional(),
		field.Time("deleted_at").Optional(),
		field.Uint32("deleted_by").Optional(),
	}
}

// Edges of the Api.
func (Api) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("statistics", ApiStatistics.Type).
			Unique(),
	}
}
