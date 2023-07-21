package data

import (
	"context"
	"fmt"
	taskV1 "galileo/api/management/task/v1"
	"galileo/app/engine/internal/biz"
	"galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/robfig/cron/v3"
	"time"
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

func (r *engineRepo) AddCronJob(ctx context.Context, task *biz.Task) (cron.EntryID, error) {
	t := task.ScheduleTime
	if delay := t.Sub(time.Now()); delay < 0 {
		return 0, errResponse.SetCustomizeErrMsg(errResponse.ReasonParamsError, "Invalid schedule time")
	}
	cronExpression := fmt.Sprintf("%d %d %d %d %d *", t.Second(), t.Minute(), t.Hour(), t.Day(), t.Month())
	entryId, err := r.data.cron.AddJob(cronExpression, &biz.CronJob{
		TaskId:  task.Id,
		TaskCli: r.data.taskCli,
		Worker:  task.Worker,
	})
	if err != nil {
		return 0, err
	}
	return entryId, nil
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

func (r *engineRepo) RemoveCronJob(ctx context.Context, taskId int64) error {
	entries := r.data.cron.Entries()
	for _, entry := range entries {
		if entry.Job.(*biz.CronJob).TaskId == taskId {
			r.data.cron.Remove(entry.ID)
			return nil
		}
	}
	return errResponse.SetCustomizeErrMsg(errResponse.ReasonRecordNotFound, "entry not found")
}
