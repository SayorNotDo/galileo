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
		field.Uint32("id").Comment("主键ID"),
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Comment("通用唯一识别码"),
		field.Time("created_at").
			Default(time.Now).Immutable().Comment("用户创建时间"),
		field.Bool("active").Default(true).Comment("用户的活跃状态：1 活跃，2 非活跃"),
		field.String("username").Unique().Comment("用户名"),
		field.String("chineseName").Optional().Comment("中文名，可选"),
		field.String("password").Comment("用户密码"),
		field.String("phone").Unique().Comment("电话号码"),
		field.String("email").Unique().Comment("电子邮箱"),
		field.Text("avatar").Optional().Comment("头像"),
		field.String("location").Optional().Comment("所处地理位置"),
		field.Time("last_login_time").Optional().Comment("最后登录时间"),
		field.Time("updated_at").Optional().UpdateDefault(time.Now).Comment("用户更新时间"),
		field.Time("deleted_at").Optional().Comment("用户删除时间"),
		field.Uint32("deleted_by").Optional().Comment("删除操作人"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
