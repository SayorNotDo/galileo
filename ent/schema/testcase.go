package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Testcase holds the schema definition for the Testcase entity.
type Testcase struct {
	ent.Schema
}

func (Testcase) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "testcase"},
	}
}

// Fields of the TestCase.
func (Testcase) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),                                           // 用例ID
		field.String("name").Unique().NotEmpty(),                    // 用例名称
		field.Uint32("created_by").Immutable(),                      // 用例创建人
		field.Time("created_at").Default(time.Now).Immutable(),      // 用例创建时间
		field.Uint32("updated_by").Optional(),                       // 用例更新人
		field.Time("updated_at").Optional().UpdateDefault(time.Now), // 用例更新时间
		field.Int8("status").Default(0),                             // 用例状态：0 新建，1 进行中，2 完成，3 废弃
		field.Int8("type").Default(0),                               // 用例类型
		field.Int8("priority").Default(0),                           // 用例优先级
		field.Time("deleted_at").Optional(),                         // 用例删除时间
		field.Uint32("deleted_by").Optional(),                       // 用例删除的人
		field.Text("description").Optional(),                        // 描述
		field.String("label").Optional(),                            // 用例标签
		field.String("url").Optional(),                              // 用例文件地址
	}
}

// Edges of the TestCase.
func (Testcase) Edges() []ent.Edge {
	return []ent.Edge{}
}
