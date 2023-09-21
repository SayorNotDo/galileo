package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// MetaEvent holds the schema definition for the MetaEvent entity.
type MetaEvent struct {
	ent.Schema
}

func (MetaEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "meta_event"},
	}
}

// Fields of the MetaEvent.
func (MetaEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("event_name").Unique().NotEmpty(),
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Comment("通用唯一识别码"),
		field.String("display_name").Unique().NotEmpty(),
		field.String("event_desc").Optional(),
		field.String("remark").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Optional().UpdateDefault(time.Now),
	}
}

// Edges of the MetaEvent.
func (MetaEvent) Edges() []ent.Edge {
	return nil
}
