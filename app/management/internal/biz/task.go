package biz

import (
	"context"
	"errors"
	engineV1 "galileo/api/engine/v1"
	v1 "galileo/api/management/v1"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// Task is a Task model.
type Task struct {
	Id            int64
	Name          string
	TestPlanId    int64
	Rank          int8
	Type          int8
	Status        v1.TaskStatus
	CreatedBy     uint32
	Assignee      uint32
	Frequency     int8
	ScheduleTime  time.Time
	UpdatedBy     uint32
	CompletedAt   time.Time
	DeletedAt     time.Time
	StartTime     time.Time
	Deadline      time.Time
	DeletedBy     uint32
	Description   string
	TestcaseSuite []int64
}

type TaskInfo struct {
	Id              int64
	Name            string
	Rank            int8
	Type            int8
	Frequency       int8
	ScheduleTime    time.Time
	Status          v1.TaskStatus
	CreatedAt       time.Time
	CreatedBy       string
	Assignee        string
	UpdatedAt       time.Time
	UpdatedBy       string
	StatusUpdatedAt time.Time
	StartTime       time.Time
	CompletedAt     time.Time
	Deadline        time.Time
	Description     string
	Testplan        string
	TestcaseSuite   []int64
}

// TaskUseCase is a Task use case.
type TaskUseCase struct {
	engine engineV1.EngineClient
	repo   TaskRepo
	log    *log.Helper
}

// TaskRepo is a Task repo.
type TaskRepo interface {
	ListTimingTask(ctx context.Context, status []v1.TaskStatus) ([]*v1.Task, error)
	CreateTask(ctx context.Context, task *Task) (*TaskInfo, error)
	ExecuteTask(ctx context.Context, taskId int64, worker uint32, config []byte) error
	TaskByName(ctx context.Context, name string) (*TaskInfo, error)
	GetTask(ctx context.Context, id int64) (*TaskInfo, error)
	UpdateTask(ctx context.Context, task *Task) (*TaskInfo, error)
	UpdateTaskStatus(ctx context.Context, updateTask *Task) (*TaskInfo, error)
	RedisLRangeTask(ctx context.Context, key string) ([]string, error)
	SetTaskInfoExpiration(ctx context.Context, key string, expiration int64) error
}

// NewTaskUseCase new a Task useCase.
func NewTaskUseCase(repo TaskRepo, engine engineV1.EngineClient, logger log.Logger) *TaskUseCase {
	return &TaskUseCase{
		engine: engine,
		repo:   repo,
		log:    log.NewHelper(log.With(logger, "module", "management.taskUseCase")),
	}
}

func NewTask(name string, rank, taskType int32, description string, testcaseSuite []int64, assignee uint32, deadline time.Time) (Task, error) {
	if len(name) <= 0 {
		return Task{}, errors.New("task name must not be empty")
	} else if rank < 0 || taskType < 0 {
		return Task{}, errors.New("invalid parameter")
	}
	return Task{
		Name:          name,
		Rank:          int8(rank),
		Type:          int8(taskType),
		Description:   description,
		TestcaseSuite: testcaseSuite,
		Assignee:      assignee,
		Deadline:      deadline,
	}, nil
}

func (uc *TaskUseCase) CreateTask(ctx context.Context, task *Task) (*TaskInfo, error) {
	return uc.repo.CreateTask(ctx, task)
}

func (uc *TaskUseCase) ExecuteTask(ctx context.Context, taskId int64, worker uint32, config []byte) error {
	return uc.repo.ExecuteTask(ctx, taskId, worker, config)
}

func (uc *TaskUseCase) TaskByName(ctx context.Context, name string) (*TaskInfo, error) {
	return uc.repo.TaskByName(ctx, name)
}

func (uc *TaskUseCase) GetTask(ctx context.Context, id int64) (*TaskInfo, error) {
	return uc.repo.GetTask(ctx, id)
}

func (uc *TaskUseCase) UpdateTask(ctx context.Context, task *Task) (*TaskInfo, error) {
	return uc.repo.UpdateTask(ctx, task)
}

func (uc *TaskUseCase) UpdateTaskStatus(ctx context.Context, updateTask *Task) (*TaskInfo, error) {
	return uc.repo.UpdateTaskStatus(ctx, updateTask)
}

func (uc *TaskUseCase) ListTimingTask(ctx context.Context, status []v1.TaskStatus) ([]*v1.Task, error) {
	return uc.repo.ListTimingTask(ctx, status)
}

func (uc *TaskUseCase) RedisLRangeTask(ctx context.Context, key string) ([]string, error) {
	return uc.repo.RedisLRangeTask(ctx, key)
}

func (uc *TaskUseCase) SetTaskInfoExpiration(ctx context.Context, key string, expiration int64) error {
	return uc.repo.SetTaskInfoExpiration(ctx, key, expiration)
}
