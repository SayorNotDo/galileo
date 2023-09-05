package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Comment("创建时间"),
		field.Uint32("created_by").Comment("执行人"),
		field.Time("updated_at").Comment("更新时间"),
		field.Time("deleted_at").Comment("删除时间"),
		field.Uint32("deleted_by").Comment("删除操作执行人"),
		field.Int64("task_id").Comment("关联任务ID"),
		field.Int64("worker").Comment("当前的工作节点"),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return nil
}
