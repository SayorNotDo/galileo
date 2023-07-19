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

func (r *engineRepo) UpdateTaskStatus(ctx context.Context, id int64, status taskV1.TaskStatus) (bool, error) {
	//ret, err := r.data.tc.UpdateTask(ctx, &taskV1.UpdateTaskRequest{Id: id, Status: status})
	//if err != nil {
	//	return false, err
	//}
	//return ret.Success, nil
	return false, nil
}

func (r *engineRepo) TaskByID(ctx context.Context, id int64) (*biz.Task, error) {
	res, err := r.data.taskCli.TaskByID(ctx, &taskV1.TaskByIDRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &biz.Task{
		Name:           res.Name,
		Type:           int8(res.Type),
		Rank:           int8(res.Rank),
		Status:         res.Status,
		Worker:         res.Worker,
		Config:         res.Config,
		Frequency:      res.Frequency,
		ScheduleTime:   res.ScheduleTime.AsTime(),
		TestcaseSuites: res.TestcaseSuiteId,
	}, nil
}
