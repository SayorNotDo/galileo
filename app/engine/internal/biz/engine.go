package biz

import (
	"context"
	managementV1 "galileo/api/management/v1"
	"time"
)

type SchedulerRepo interface {
	TaskByID(ctx context.Context, id int64) (*Task, error)
	TimingTaskList(ctx context.Context, status []managementV1.TaskStatus) ([]*Task, error)
	CreateJob(ctx context.Context, job *Job) (*Job, error)
	RemoveCronJob(ctx context.Context, taskId int64) error
	AddPeriodicJob(ctx context.Context, payload []byte, expression string) (*Job, error)
	AddDefaultJob(ctx context.Context, payload []byte) (*Job, error)
	AddDelayedJob(ctx context.Context, payload []byte, delay time.Duration) (*Job, error)
}

func (sc *SchedulerUseCase) AddPeriodicJob(ctx context.Context, payload []byte, expression string) (*Job, error) {
	return sc.repo.AddPeriodicJob(ctx, payload, expression)
}

func (sc *SchedulerUseCase) AddDefaultJob(ctx context.Context, payload []byte) (*Job, error) {
	return sc.repo.AddDefaultJob(ctx, payload)
}

func (sc *SchedulerUseCase) AddDelayedJob(ctx context.Context, payload []byte, delay time.Duration) (*Job, error) {
	return sc.repo.AddDelayedJob(ctx, payload, delay)
}

func (sc *SchedulerUseCase) TaskByID(ctx context.Context, id int64) (*Task, error) {
	return sc.repo.TaskByID(ctx, id)
}

func (sc *SchedulerUseCase) CreateJob(ctx context.Context, job *Job) (*Job, error) {
	return sc.repo.CreateJob(ctx, job)
}

func (sc *SchedulerUseCase) TimingTaskList(ctx context.Context, status []managementV1.TaskStatus) ([]*Task, error) {
	return sc.repo.TimingTaskList(ctx, status)
}

func (sc *SchedulerUseCase) RemoveCronJob(ctx context.Context, taskId int64) error {
	return sc.repo.RemoveCronJob(ctx, taskId)
}
