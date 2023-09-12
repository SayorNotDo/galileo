package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Comment("Job ID"),
		field.Time("created_at").Default(time.Now).Immutable().Comment("创建时间"),
		field.Uint32("created_by").Comment("执行人"),
		field.Time("updated_at").UpdateDefault(time.Now).Comment("更新时间"),
		field.Uint32("worker").Comment("当前的工作节点, 存储格式为IP地址"),
		field.Time("deleted_at").Optional().Comment("删除时间"),
		field.Uint32("deleted_by").Optional().Comment("删除操作执行人"),
		field.UUID("uuid", uuid.UUID{}).Comment("Job execute ID"),
		field.String("config").Optional().Comment("配置信息"),
		field.Int64("task_id").Comment("关联的任务ID"),
		field.Bool("active").Comment("Job 执行情况"),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return nil
}
