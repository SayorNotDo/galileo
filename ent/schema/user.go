package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id"),
		field.Time("created_at").
			Default(time.Now).Immutable(),
		field.Bool("active").Default(true),
		field.String("username").Unique(),
		field.String("chineseName").Optional(),
		field.String("nickname").Unique(),
		field.String("password"),
		field.String("phone").Unique(),
		field.String("email").Unique(),
		field.Text("avatar").Optional(),
		field.Uint8("role").Default(3),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
		field.Uint32("deleted_by").Optional().Nillable(),
		field.Bool("is_deleted").Optional().Nillable().Default(false),
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
