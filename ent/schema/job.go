package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

func (Job) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "job"},
	}
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Comment("Job ID"),
		field.Time("created_at").Default(time.Now).Immutable().Comment("创建时间"),
		field.Uint32("created_by").Immutable().Comment("执行人"),
		field.Bytes("payload").Optional().Comment("Payload"),
		field.String("type").Optional().Comment("任务类型"),
		field.Time("updated_at").Optional().UpdateDefault(time.Now).Comment("更新时间"),
		field.Uint32("worker").Comment("当前的工作节点, 存储格式为IP地址"),
		field.Time("deleted_at").Optional().Comment("删除时间"),
		field.Uint32("deleted_by").Optional().Comment("删除操作执行人"),
		field.String("entry_id").Optional().Comment("Entry ID"),
		field.Bytes("config").Optional().Comment("配置信息"),
		field.Int64("task_id").Comment("关联的任务ID"),
		field.Bool("active").Default(false).Comment("Job 执行情况"),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return nil
}
