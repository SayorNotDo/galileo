package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Task is a Task model.
type Task struct {
	Id            int64
	Name          string
	Rank          int8
	Type          int8
	Status        int8
	CreatedAt     time.Time
	CreatedBy     uint32
	CompleteAt    time.Time
	UpdateAt      time.Time
	DeletedAt     time.Time
	DeletedBy     uint32
	Description   string
	TestcaseSuite []int
}

// TaskRepo is a Task repo.
type TaskRepo interface {
	ListTask(ctx context.Context, pageNum, pageSize int) ([]*Task, error)
	CreateTask(ctx context.Context, task *Task) (*Task, error)
}

// TaskUseCase is a Task useCase.
type TaskUseCase struct {
	repo TaskRepo
	log  *log.Helper
}

// NewTaskUseCase new a Task useCase.
func NewTaskUseCase(repo TaskRepo, logger log.Logger) *TaskUseCase {
	return &TaskUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "management.taskUseCase")),
	}
}

func (uc *TaskUseCase) CreateTask(ctx context.Context, task *Task) (*Task, error) {
	return uc.repo.CreateTask(ctx, task)
}
