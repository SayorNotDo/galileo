package biz

import "fmt"

type CronJob struct {
	TaskId int64 `json:"task_id,omitempty"`
}

func (c *CronJob) Run() {
	fmt.Println("执行自定义任务: ", c.TaskId)
}
