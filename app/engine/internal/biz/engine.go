package biz

import (
	"context"
	managementV1 "galileo/api/management/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
	"time"
)

type Task struct {
	Id             int64
	Name           string
	Rank           int8
	Type           int8
	Status         managementV1.TaskStatus
	Worker         string
	Config         string
	Frequency      int8
	ScheduleTime   time.Time
	TestcaseSuites []int32
	ExecuteId      int64
}

type EngineRepo interface {
	TaskByID(ctx context.Context, id int64) (*Task, error)
	AddCronJob(ctx context.Context, task *Task) (cron.EntryID, error)
	TimingTaskList(ctx context.Context, status []managementV1.TaskStatus) ([]*Task, error)
	GetCronJobList(ctx context.Context) []*CronJob
	RemoveCronJob(ctx context.Context, taskId int64) error
	AddPeriodicJob(ctx context.Context, payload []byte, expression string) (*Job, error)
	AddDefaultJob(ctx context.Context, payload []byte) (*Job, error)
	AddDelayedJob(ctx context.Context, payload []byte, delay time.Duration) (*Job, error)
}

type EngineUseCase struct {
	repo EngineRepo
	log  *log.Helper
}

func NewEngineUseCase(repo EngineRepo, logger log.Logger) *EngineUseCase {
	helper := log.NewHelper(log.With(logger, "module", "engine.useCase"))
	return &EngineUseCase{repo: repo, log: helper}
}

func (c *EngineUseCase) AddPeriodicJob(ctx context.Context, payload []byte, expression string) (*Job, error) {
	return c.repo.AddPeriodicJob(ctx, payload, expression)
}

func (c *EngineUseCase) AddDefaultJob(ctx context.Context, payload []byte) (*Job, error) {
	return c.repo.AddDefaultJob(ctx, payload)
}

func (c *EngineUseCase) AddDelayedJob(ctx context.Context, payload []byte, delay time.Duration) (*Job, error) {
	return c.repo.AddDelayedJob(ctx, payload, delay)
}

func (c *EngineUseCase) TaskByID(ctx context.Context, id int64) (*Task, error) {
	return c.repo.TaskByID(ctx, id)
}

func (c *EngineUseCase) AddCronJob(ctx context.Context, task *Task) (cron.EntryID, error) {
	return c.repo.AddCronJob(ctx, task)
}

func (c *EngineUseCase) TimingTaskList(ctx context.Context, status []managementV1.TaskStatus) ([]*Task, error) {
	return c.repo.TimingTaskList(ctx, status)
}

func (c *EngineUseCase) GetCronJobList(ctx context.Context) []*CronJob {
	return c.repo.GetCronJobList(ctx)
}

func (c *EngineUseCase) RemoveCronJob(ctx context.Context, taskId int64) error {
	return c.repo.RemoveCronJob(ctx, taskId)
}
