package biz

import (
	"context"
	"fmt"
	managementV1 "galileo/api/management/v1"
	"galileo/pkg/errResponse"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
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
		Worker: c.Worker,
	})
	if err != nil {
		fmt.Println("Error updating task status: ", err)
	}
}

// JobCommand send a job command to the executor through http request
func JobCommand(schema, worker string, taskId int64) error {
	// 创建HTTP客户端
	client := &http.Client{}

	// TODO: 创建发送至Worker的HTTP请求
	/* 验证参数worker是否为合法 */
	params := strings.Split(worker, ":")
	ip := params[0]
	port := params[1]
	if net.ParseIP(ip) == nil {
		return errResponse.SetCustomizeErrMsg(errResponse.ReasonParamsError, "Invalid IP address")
	}
	if _, err := net.LookupPort("tcp", port); err != nil {
		return errResponse.SetCustomizeErrMsg(errResponse.ReasonParamsError, "Invalid Port")
	}
	/* 拼接请求地址: schema://ip:port/job */
	url := fmt.Sprintf("%s%s/job", schema, worker)

	/* 构建请求体 taskId, token */
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

	// 获取响应内容
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
