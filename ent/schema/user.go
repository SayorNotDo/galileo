package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").Unique().Immutable(),
		field.Time("created_at").
			Default(time.Now).Immutable(),
		field.Bool("active").Default(true),
		field.String("username").Unique(),
		field.String("nickname").Unique(),
		field.String("password"),
		field.String("email").Unique(),
		field.Text("avatar"),
		field.Uint8("role"),
		field.Time("updated_at").UpdateDefault(time.Now),
		field.Time("deleted_at").Default(nil),
		field.Uint32("deleted_by"),
		field.Time("last_login_at"),
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
