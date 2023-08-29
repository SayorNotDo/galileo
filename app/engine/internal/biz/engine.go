package biz

import (
	"context"
	managementV1 "galileo/api/management/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
	"time"
)

type Task struct {
	Id             int32
	Name           string
	Rank           int8
	Type           int8
	Status         managementV1.TaskStatus
	Worker         string
	Config         string
	Frequency      managementV1.Frequency
	ScheduleTime   time.Time
	TestcaseSuites []int32
	ExecuteId      int64
}

type EngineRepo interface {
	TaskByID(ctx context.Context, id int32) (*Task, error)
	AddCronJob(ctx context.Context, task *Task) (cron.EntryID, error)
	TimingTaskList(ctx context.Context, status []managementV1.TaskStatus) ([]*Task, error)
	GetCronJobList(ctx context.Context) []*CronJob
	RemoveCronJob(ctx context.Context, taskId int32) error
}

type EngineUseCase struct {
	repo EngineRepo
	log  *log.Helper
}

func NewEngineUseCase(repo EngineRepo, logger log.Logger) *EngineUseCase {
	helper := log.NewHelper(log.With(logger, "module", "engine.useCase"))
	return &EngineUseCase{repo: repo, log: helper}
}

func (c *EngineUseCase) TaskByID(ctx context.Context, id int32) (*Task, error) {
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

func (c *EngineUseCase) RemoveCronJob(ctx context.Context, taskId int32) error {
	return c.repo.RemoveCronJob(ctx, taskId)
}

func (c *EngineUseCase) ParseDockerfile(ctx context.Context, fp string) (map[string]interface{}, error) {
	return nil, nil
}
