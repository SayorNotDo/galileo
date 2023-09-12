package data

import (
	"context"
	managementV1 "galileo/api/management/v1"
	"galileo/app/engine/internal/biz"
	"galileo/pkg/ctxdata"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"github.com/robfig/cron/v3"
	"time"
)

type schedulerRepo struct {
	data *Data
	log  *log.Helper
}

func NewEngineRepo(data *Data, logger log.Logger) biz.SchedulerRepo {
	return &schedulerRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "engine.Repo")),
	}
}

func (r *schedulerRepo) AddPeriodicJob(ctx context.Context, payload []byte, expression string) (*biz.Job, error) {
	err := r.data.asynqSrv.NewPeriodicTask(expression, biz.TypePeriodicJob, payload)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *schedulerRepo) AddDefaultJob(ctx context.Context, payload []byte) (*biz.Job, error) {
	err := r.data.asynqSrv.NewTask(biz.TypeDefaultJob, payload)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *schedulerRepo) AddDelayedJob(ctx context.Context, payload []byte, delay time.Duration) (*biz.Job, error) {
	/* typeName、 payload、 delayTime */
	err := r.data.asynqSrv.NewTask(biz.TypeDelayedJob, payload, asynq.ProcessIn(delay))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *schedulerRepo) TaskByID(ctx context.Context, id int64) (*biz.Task, error) {
	//res, err := r.data.ManageCli.TaskByID(ctx, &managementV1.TaskByIDRequest{Id: id})
	//if err != nil {
	//	return nil, err
	//}
	//return &biz.Task{
	//	Name:           res.Name,
	//	Type:           int8(res.Type),
	//	Rank:           int8(res.Rank),
	//	Status:         res.Status,
	//	Worker:         res.Worker,
	//	Config:         res.Config,
	//	Frequency:      res.Frequency,
	//	ScheduleTime:   res.ScheduleTime.AsTime(),
	//	TestcaseSuites: res.TestcaseSuite,
	//}, nil
	return nil, nil
}

func (r *schedulerRepo) CreateJob(ctx context.Context, task *biz.Job) (*biz.Job, error) {
	uid := ctxdata.UserIdFromMetaData(ctx)
	ret, err := r.data.entDB.Job.Create().
		SetCreatedBy(uid).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Job{
		Id: ret.ID,
	}, nil
}

func (r *schedulerRepo) TimingTaskList(ctx context.Context, status []managementV1.TaskStatus) ([]*biz.Task, error) {
	///* 获取状态为待办的定时/延时类型任务列表 */
	//res, err := r.data.ManageCli.ListTimingTask(ctx, &managementV1.ListTimingTaskRequest{
	//	Status: status,
	//})
	//if err != nil {
	//	log.Error(err)
	//	return nil, err
	//}
	///* 组织结构体返回 */
	//rv := make([]*biz.Task, 0)
	//for _, v := range res.TaskList {
	//	rv = append(rv, &biz.Task{
	//		Id:           v.Id,
	//		Name:         v.Name,
	//		Rank:         int8(v.Rank),
	//		Type:         int8(v.Type),
	//		Status:       v.Status,
	//		Worker:       v.Worker,
	//		Config:       v.Config,
	//		Frequency:    v.Frequency,
	//		ScheduleTime: v.ScheduleTime.AsTime(),
	//	})
	//}
	//return rv, nil
	return nil, nil
}

func (r *schedulerRepo) GetCronJobByTaskId(ctx context.Context, taskId int64) (cron.Entry, error) {
	//entries := r.data.cron.Entries()
	//for _, entry := range entries {
	//	if entry.Job.(*biz.CronJob).TaskId == taskId {
	//		return entry, nil
	//	}
	//}
	//return cron.Entry{}, errResponse.SetCustomizeErrMsg(errResponse.ReasonRecordNotFound, "entry not found")
	return cron.Entry{}, nil
}

func (r *schedulerRepo) RemoveCronJob(ctx context.Context, taskId int64) error {
	//entry, err := r.GetCronJobByTaskId(ctx, taskId)
	//if err != nil {
	//	return err
	//}
	//r.data.cron.Remove(entry.ID)
	return nil
}
