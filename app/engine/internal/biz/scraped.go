package biz

import (
	"context"
	"fmt"
	managementV1 "galileo/api/management/v1"
)

type CronJob struct {
	TaskId    int64 `json:"task_id,omitempty"`
	ManageCli managementV1.ManagementClient
	Worker    string `json:"worker,omitempty"`
}

// Run
/* 定时任务调度器执行函数 */
func (c *CronJob) Run() {
	// TODO: 测试任务触发（更新任务状态）
	fmt.Println("下发测试任务: ", c.TaskId)
	_, err := c.ManageCli.UpdateTaskStatus(context.Background(), &managementV1.UpdateTaskStatusRequest{
		Id:     c.TaskId,
		Status: managementV1.TaskStatus_RUNNING,
	})
	if err != nil {
		fmt.Println("Error updating task status: ", err)
	}
}
