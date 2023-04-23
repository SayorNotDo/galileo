package biz

import (
	"context"
	taskV1 "galileo/api/management/task/v1"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Task struct {
	Id             int64
	Name           string
	Rank           int8
	Type           int8
	Status         int8
	CreatedAt      time.Time
	CreatedBy      uint32
	UpdateAt       time.Time
	CompleteAt     time.Time
	DeletedAt      time.Time
	DeletedBy      uint32
	Description    string
	TestcaseSuites []int64
}

type EngineRepo interface {
	UpdateTaskStatus(ctx context.Context, id int64, status taskV1.TaskStatus) (bool, error)
	TaskByID(ctx context.Context, id int64) (*Task, error)
}

type EngineUseCase struct {
	Repo EngineRepo
	log  *log.Helper
}

func NewEngineUseCase(repo EngineRepo, logger log.Logger) *EngineUseCase {
	helper := log.NewHelper(log.With(logger, "module", "engine.useCase"))
	return &EngineUseCase{Repo: repo, log: helper}
}

func (c *EngineUseCase) UpdateTaskStatus(ctx context.Context, id int64, status taskV1.TaskStatus) (bool, error) {
	return c.Repo.UpdateTaskStatus(ctx, id, status)
}

func (c *EngineUseCase) TaskByID(ctx context.Context, id int64) (*Task, error) {
	return c.Repo.TaskByID(ctx, id)
}
