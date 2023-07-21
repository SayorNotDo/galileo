package biz

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type CronJob struct {
	TaskId int64  `json:"task_id,omitempty"`
	Worker string `json:"worker,omitempty"`
	Schema string `json:"schema,omitempty"`
}

func (c *CronJob) Run() {
	// TODO: 测试任务触发（下发命令）
	fmt.Println("下发测试任务: ", c.TaskId)
	err := JobCommand(c.Schema, c.Worker, c.TaskId)
	if err != nil {
		fmt.Println("Send Job Command Error: ", err)
		return
	}
}

func JobCommand(schema, worker string, taskId int64) error {
	// 创建HTTP客户端
	client := &http.Client{}

	// TODO: 创建发送至Worker的HTTP请求
	url := fmt.Sprintf("%s%s/job", schema, worker)
	payload := strings.NewReader("taskId=%s" + strconv.FormatInt(taskId, 10))
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
