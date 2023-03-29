package data

import (
	"context"

	"galileo/app/task/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type taskRepo struct {
	data *Data
	log  *log.Helper
}

// NewTaskRepo .
func NewTaskRepo(data *Data, logger log.Logger) biz.TaskRepo {
	return &taskRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *taskRepo) Save(ctx context.Context, g *biz.Task) (*biz.Task, error) {
	return g, nil
}

func (r *taskRepo) Update(ctx context.Context, g *biz.Task) (*biz.Task, error) {
	return g, nil
}

func (r *taskRepo) FindByID(context.Context, int64) (*biz.Task, error) {
	return nil, nil
}

func (r *taskRepo) ListByHello(context.Context, string) ([]*biz.Task, error) {
	return nil, nil
}

func (r *taskRepo) ListAll(context.Context) ([]*biz.Task, error) {
	return nil, nil
}
