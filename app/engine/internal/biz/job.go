package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"time"
)

const (
	TypeDefaultJob  = "job:default"
	TypeDelayedJob  = "job:delayed"
	TypePeriodicJob = "job:periodic"
)

type DelayedJobPayload struct {
	Task      int64         `json:"task"`
	Worker    int64         `json:"worker"`
	DelayTime time.Duration `json:"delay_time"`
}

type DefaultJobPayload struct {
	Task   int64 `json:"task"`
	Worker int64 `json:"worker"`
}

type PeriodicJobPayload struct {
	Task     int64  `json:"task"`
	Worker   int64  `json:"worker"`
	Schedule string `json:"schedule"`
}

func NewPeriodicJobPayload(task int64, worker int64, schedule string) ([]byte, error) {
	payload, err := json.Marshal(PeriodicJobPayload{
		Task:   task,
		Worker: worker,
	})
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func NewDefaultJobPayload(task int64, worker int64) ([]byte, error) {
	payload, err := json.Marshal(DefaultJobPayload{
		Task:   task,
		Worker: worker,
	})
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func NewDelayedJobPayload(task int64, worker int64, delayTime time.Duration) ([]byte, error) {
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
