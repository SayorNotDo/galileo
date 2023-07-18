package biz

import (
	"context"
	engineV1 "galileo/api/engine/v1"
	v1 "galileo/api/management/task/v1"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Task is a Task model.
type Task struct {
	Id              int64
	Name            string
	Rank            int8
	Type            int8
	Status          v1.TaskStatus
	CreatedAt       time.Time
	CreatedBy       uint32
	Assignee        uint32
	Config          string
	UpdatedAt       time.Time
	StatusUpdatedAt time.Time
	CompletedAt     time.Time
	DeletedAt       time.Time
	StartTime       time.Time
	Deadline        time.Time
	DeletedBy       uint32
	Description     string
	TestcaseSuites  []int64
}

// TaskRepo is a Task repo.
type TaskRepo interface {
	ListTask(ctx context.Context, pageNum, pageSize int) ([]*Task, error)
	CreateTask(ctx context.Context, task *Task) (*Task, error)
	TaskByName(ctx context.Context, name string) (*Task, error)
	TaskByID(ctx context.Context, id int64) (*Task, error)
	UpdateTask(ctx context.Context, task *Task) (bool, error)
	UpdateTaskStatus(ctx context.Context, updateTask *Task) (*Task, error)
}

// TaskUseCase is a Task useCase.
type TaskUseCase struct {
	engine engineV1.EngineClient
	repo   TaskRepo
	log    *log.Helper
}

// NewTaskUseCase new a Task useCase.
func NewTaskUseCase(repo TaskRepo, engine engineV1.EngineClient, logger log.Logger) *TaskUseCase {
	return &TaskUseCase{
		engine: engine,
		repo:   repo,
		log:    log.NewHelper(log.With(logger, "module", "management.taskUseCase")),
	}
}

func (uc *TaskUseCase) CreateTask(ctx context.Context, task *Task) (*Task, error) {
	return uc.repo.CreateTask(ctx, task)
}

func (uc *TaskUseCase) TaskByName(ctx context.Context, name string) (*Task, error) {
	return uc.repo.TaskByName(ctx, name)
}

func (uc *TaskUseCase) TaskByID(ctx context.Context, id int64) (*Task, error) {
	return uc.repo.TaskByID(ctx, id)
}

func (uc *TaskUseCase) UpdateTask(ctx context.Context, task *Task) (bool, error) {
	return uc.repo.UpdateTask(ctx, task)
}

func (uc *TaskUseCase) UpdateTaskStatus(ctx context.Context, updateTask *Task) (*Task, error) {
	return uc.repo.UpdateTaskStatus(ctx, updateTask)
}
