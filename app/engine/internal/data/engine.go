package data

import (
	"context"
	taskV1 "galileo/api/management/task/v1"
	"galileo/app/engine/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type engineRepo struct {
	data *Data
	log  *log.Helper
}

func NewEngineRepo(data *Data, logger log.Logger) biz.EngineRepo {
	return &engineRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "engine.engineRepo")),
	}
}

func (r *engineRepo) UpdateTaskStatus(ctx context.Context, status taskV1.TaskStatus) (bool, error) {
	ret, err := r.data.tc.UpdateTaskStatus(ctx, &taskV1.UpdateTaskStatusRequest{Id: 7, Status: status})
	if err != nil {
		return false, err
	}
	return ret.Success, nil
}
