package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		field.String("name").Unique().NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Uint32("created_by").Immutable(),
		field.Uint32("assignee").Optional(),                         // 任务经办人
		field.Int8("type").Default(0),                               // 任务类型：非定时/定时任务
		field.String("config").Optional(),                           // 任务配置
		field.Int8("rank").Default(0),                               // 任务优先级
		field.Int8("status").Default(0),                             // 任务状态
		field.Time("start_time").Optional(),                         // 任务开始时间
		field.Time("completed_at").Optional(),                       // 任务完成时间
		field.Time("updated_at").Optional().UpdateDefault(time.Now), // 数据库记录更新时间
		field.Time("status_updated_at").Optional(),                  // 任务状态更新时间
		field.Time("deadline").Optional(),                           // 任务截止日期
		field.Time("deleted_at").Optional(),                         // 任务删除时间
		field.Uint32("deleted_by").Optional(),                       // 任务删除人
		field.Text("description").Optional(),                        // 任务描述
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("testcaseSuite", TestCaseSuite.Type),
	}
}
