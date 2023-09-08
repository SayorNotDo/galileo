package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
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
		field.Int64("id").Unique().Comment("任务ID"),
		field.String("name").NotEmpty().Comment("任务名称"),
		field.Time("created_at").Default(time.Now).Immutable().Comment("任务创建时间"),
		field.Uint32("created_by").Immutable().Comment("任务创建人"),
		field.Uint32("assignee").Optional().Comment("任务经办人"),
		field.Int8("type").Default(0).Comment("任务类型：实时任务 0、延时任务 1、定时任务 2、响应式任务 3"),
		field.Int8("frequency").Optional().Comment("定时任务的频率，仅用于定时任务"),
		field.Time("schedule_time").Optional().Comment("预期调度时间: 实时任务为空，延时任务取年月日时分秒，定时任务取 时分秒+频率"),
		field.Int8("rank").Default(0).Comment("任务优先级"),
		field.Int8("status").Default(0).Comment("任务状态"),
		field.Time("start_time").Optional().Comment("任务开始时间"),
		field.Time("completed_at").Optional().Comment("任务完成时间"),
		field.Time("updated_at").Optional().UpdateDefault(time.Now).Comment("数据库记录更新时间"),
		field.Uint32("updated_by").Optional().Comment("任务更新人"),
		field.Time("status_updated_at").Optional().Comment("任务状态更新时间"),
		field.Time("deadline").Optional().Comment("任务截止日期"),
		field.Time("deleted_at").Optional().Comment("任务删除时间"),
		field.Uint32("deleted_by").Optional().Comment("任务删除人"),
		field.Text("description").Optional().Comment("任务描述"),
		field.Int64("testplan_id").Optional().Comment("任务所属的测试计划"),
		field.JSON("testcase_suite", []int64{}).Optional().Comment("测试用例集合"),
	}
}

func (Task) Indexes() []ent.Index {
	return nil
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}
