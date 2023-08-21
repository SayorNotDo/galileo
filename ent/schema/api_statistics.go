package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ApiStatistics holds the schema definition for the ApiStatistics entity.
type ApiStatistics struct {
	ent.Schema
}

func (ApiStatistics) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "api_statistics"},
	}
}

// Fields of the ApiStatistics.
func (ApiStatistics) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.Int32("call_count"),
		field.Int32("success_count"),
		field.Int32("failure_count"),
		field.Float("avg_response_time"),
		field.Float("max_response_time"),
		field.Float("min_response_time"),
		field.Float("avg_traffic"),
		field.Float("max_traffic"),
		field.Float("min_traffic"),
		field.String("description"),
		field.Time("created_at"),
		field.Time("update_at"),
		field.Int32("api_id").Positive(),
	}
}

// Edges of the ApiStatistics.
func (ApiStatistics) Edges() []ent.Edge {
	return []ent.Edge{}
}
