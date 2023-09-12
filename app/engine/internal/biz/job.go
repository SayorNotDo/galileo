package biz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"galileo/pkg/errResponse"
	"github.com/avast/retry-go"
	"github.com/hibiken/asynq"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	TypeDefaultJob  = "job:default"
	TypeDelayedJob  = "job:delayed"
	TypePeriodicJob = "job:periodic"
)

type DelayedJobPayload struct {
	Task      int64         `json:"task"`
	Worker    uint32        `json:"worker"`
	DelayTime time.Duration `json:"delay_time"`
}

type DefaultJobPayload struct {
	Task   int64  `json:"task"`
	Worker uint32 `json:"worker"`
	Config []byte `json:"config"`
}

type PeriodicJobPayload struct {
	Task     int64  `json:"task"`
	Worker   uint32 `json:"worker"`
	Schedule string `json:"schedule"`
}

type Config struct {
	Port     int    `json:"port"`
	Protocol string `json:"schema"`
	Params   []byte `json:"params"`
}

func NewPeriodicJobPayload(task int64, worker uint32, schedule string) ([]byte, error) {
	payload, err := json.Marshal(PeriodicJobPayload{
		Task:   task,
		Worker: worker,
	})
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func NewDefaultJobPayload(task int64, worker uint32, config []byte) ([]byte, error) {
	payload, err := json.Marshal(DefaultJobPayload{
		Task:   task,
		Worker: worker,
		Config: config,
	})
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func NewDelayedJobPayload(task int64, worker uint32, delayTime time.Duration) ([]byte, error) {
	payload, err := json.Marshal(DelayedJobPayload{
		Task:      task,
		Worker:    worker,
		DelayTime: delayTime,
	})
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func HandlePeriodicJob(ctx context.Context, task *asynq.Task) error {
	var p PeriodicJobPayload
	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("DelayedJob: task_id=%d, worker_id=%d", p.Task, p.Worker)
	return nil
}

func HandleDefaultJob(ctx context.Context, task *asynq.Task) error {
	var p DefaultJobPayload
	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("DelayedJob: task_id=%d, worker_id=%d", p.Task, p.Worker)
	/* Default Task */
	/* 定义重试策略 */
	retryStrategy := []retry.Option{
		retry.Delay(100 * time.Millisecond),
		retry.Attempts(5),
		retry.LastErrorOnly(true),
	}
	err := retry.Do(func() error {
		if err := JobRunCommand(p.Task, p.Worker, p.Config); err != nil {
			return err
		}
		return nil
	}, retryStrategy...)
	if err != nil {
		return err
	}
	/* 数据库写入Job记录 */
	return nil
}

func HandleDelayedJob(ctx context.Context, task *asynq.Task) error {
	var p DelayedJobPayload
	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("DelayedJob: task_id=%d, worker_id=%d", p.Task, p.Worker)
	/* Delay Task */
	return nil
}

// JobRunCommand send a job command to worker through http request
func JobRunCommand(taskId int64, worker uint32, config []byte) error {
	/* 创建HTTP客户端 */
	c := &http.Client{}

	/* 转化uint32的ip为net.IP类型，同时校验合法性 */
	ip := intToIP(worker)
	if ip == nil {
		return errors.New("invalid ip address")
	}

	/* 解析生成Config结构体 */
	var conf *Config
	if err := json.Unmarshal(config, &conf); err != nil {
		return err
	}
	protocol, port := conf.Protocol, conf.Port

	/* 验证端口合法性 */
	if _, err := net.LookupPort("tcp", strconv.Itoa(port)); err != nil {
		return errResponse.SetCustomizeErrMsg(errResponse.ReasonParamsError, "Invalid Port")
	}

	/* 生成发送http请求的地址 */
	var url string
	switch protocol {
	case "http":
		url = fmt.Sprintf("%s://%d:%d/job", protocol, worker, port)
	}

	/* 构建请求体内容 */
	payload := strings.NewReader("taskId=%s" + strconv.FormatInt(taskId, 10))
	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return err
	}
	/* 发起HTTP请求 */
	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	/* 获取响应内容 */
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	/* 基于响应内容判断调用是否成功 */
	var result struct {
		Code    int
		Message string
		Data    []byte
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}
	if result.Code != 200 {
		return errResponse.SetCustomizeErrMsg(errResponse.ReasonSystemError, result.Message)
	}
	return nil
}

func intToIP(ipInt uint32) net.IP {
	ip := make(net.IP, 4)
	ip[0] = byte(ipInt & 0xFF)
	ip[1] = byte((ipInt >> 8) & 0xFF)
	ip[2] = byte((ipInt >> 16) & 0xFF)
	ip[3] = byte((ipInt >> 24) & 0xFF)
	if ip.To4() == nil {
		return nil
	}
	return ip
}
