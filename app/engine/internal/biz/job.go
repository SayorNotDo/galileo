package biz

import (
	"context"
	"fmt"
	managementV1 "galileo/api/management/v1"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type CronJob struct {
	TaskId    int32 `json:"task_id,omitempty"`
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
		Worker: c.Worker,
	})
	if err != nil {
		fmt.Println("Error updating task status: ", err)
	}
}

func JobCommand(schema, worker string, taskId int32) error {
	// 创建HTTP客户端
	client := &http.Client{}

	// TODO: 创建发送至Worker的HTTP请求
	url := fmt.Sprintf("%s%s/job", schema, worker)
	/* body: taskId, token, */
	payload := strings.NewReader("taskId=%s" + strconv.FormatInt(int64(taskId), 10))
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println("Error Creating Request: ", err)
		return err
	}

	// TODO: 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Sending Request: ", err)
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error Closing Body: ", err)
		}
	}(resp.Body)

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Reading Response: ", err)
		return err
	}

	// 打印响应内容
	fmt.Println(string(body))
	// TODO: 根据响应的内容返回是否调用成功
	/* 响应格式
	{"code": 200, "message":"", "data": {}}
	*/
	return nil
}
