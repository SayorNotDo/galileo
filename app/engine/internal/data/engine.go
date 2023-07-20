package data

import (
	"context"
	"fmt"
	taskV1 "galileo/api/management/task/v1"
	"galileo/app/engine/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
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

func (r *engineRepo) AddCronJob(ctx context.Context) {
	if _, err := r.data.cron.AddJob("@every 10s", &biz.CronJob{TaskId: 17}); err != nil {
		return
	}
	r.data.cron.Start()
}

func (r *engineRepo) TimingTaskList(ctx context.Context) ([]*biz.Task, error) {
	res, err := r.data.taskCli.ListTimingTask(ctx, &empty.Empty{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	rv := make([]*biz.Task, 0)
	for _, v := range res.TaskList {
		rv = append(rv, &biz.Task{
			Id:           v.Id,
			Name:         v.Name,
			Rank:         int8(v.Rank),
			Type:         int8(v.Type),
			Status:       v.Status,
			Worker:       v.Worker,
			Config:       v.Config,
			Frequency:    v.Frequency,
			ScheduleTime: v.ScheduleTime.AsTime(),
		})
	}
	return rv, nil
}

func (r *engineRepo) GetCronJobList(ctx context.Context) []*biz.CronJob {
	entries := r.data.cron.Entries()
	rv := make([]*biz.CronJob, 0)
	r.log.Debug("GetCronEntries================================>>")
	for _, entry := range entries {
		if cronJob, ok := entry.Job.(*biz.CronJob); ok {
			fmt.Printf("Job related task id: %v\n", cronJob.TaskId)
			rv = append(rv, &biz.CronJob{
				TaskId: cronJob.TaskId,
			})
		}
	}
	return rv
}
