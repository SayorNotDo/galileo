package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Task is a Task model.
type Task struct {
	Hello string
}

// TaskRepo is a Task repo.
type TaskRepo interface {
	Save(context.Context, *Task) (*Task, error)
	Update(context.Context, *Task) (*Task, error)
	FindByID(context.Context, int64) (*Task, error)
	ListByHello(context.Context, string) ([]*Task, error)
	ListAll(context.Context) ([]*Task, error)
}

// TaskUseCase is a Task useCase.
type TaskUseCase struct {
	repo TaskRepo
	log  *log.Helper
}

// NewTaskUseCase new a Task useCase.
func NewTaskUseCase(repo TaskRepo, logger log.Logger) *TaskUseCase {
	return &TaskUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *TaskUseCase) CreateGreeter(ctx context.Context, g *Task) (*Task, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
