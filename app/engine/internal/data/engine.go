package data

import (
	"context"
	managementV1 "galileo/api/management/v1"
	"galileo/app/engine/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
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
		log:  log.NewHelper(log.With(logger, "module", "engine.Repo")),
	}
}

func (r *engineRepo) AddPeriodicJob(ctx context.Context, payload []byte, expression string) (*biz.Job, error) {
	err := r.data.asynqSrv.NewPeriodicTask(expression, biz.TypePeriodicJob, payload)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *engineRepo) AddDefaultJob(ctx context.Context, payload []byte) (*biz.Job, error) {
	err := r.data.asynqSrv.NewTask(biz.TypeDefaultJob, payload)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *engineRepo) AddDelayedJob(ctx context.Context, payload []byte, delay time.Duration) (*biz.Job, error) {
	/* typeName、 payload、 delayTime */
	err := r.data.asynqSrv.NewTask(biz.TypeDelayedJob, payload, asynq.ProcessIn(delay))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *engineRepo) TaskByID(ctx context.Context, id int64) (*biz.Task, error) {
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

func (r *engineRepo) AddCronJob(ctx context.Context, task *biz.Task) (cron.EntryID, error) {
	//t := task.ScheduleTime
	///* 判断调度时间是否有效 */
	//if delay := t.Sub(time.Now()); delay < 0 {
	//	return 0, errResponse.SetCustomizeErrMsg(errResponse.ReasonParamsError, "Invalid schedule time")
	//}
	///* 检查是否已存在定时任务，避免出现重复调度 */
	//if _, err := r.GetCronJobByTaskId(ctx, task.Id); err != nil {
	//	return 0, err
	//}
	///* 构建定时任务调度时间表达式 */
	//cronExpression := fmt.Sprintf("%d %d %d %d %d *", t.Second(), t.Minute(), t.Hour(), t.Day(), t.Month())
	//entryId, err := r.data.cron.AddJob(cronExpression, &biz.CronJob{
	//	TaskId:    task.Id,
	//	ManageCli: r.data.ManageCli,
	//	Worker:    task.Worker,
	//})
	//if err != nil {
	//	return 0, err
	//}
	//return entryId, nil
	return 0, nil
}

func (r *engineRepo) TimingTaskList(ctx context.Context, status []managementV1.TaskStatus) ([]*biz.Task, error) {
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

func (r *engineRepo) GetCronJobList(ctx context.Context) []*biz.CronJob {
	//entries := r.data.cron.Entries()
	//rv := make([]*biz.CronJob, 0)
	//for _, entry := range entries {
	//	if cronJob, ok := entry.Job.(*biz.CronJob); ok {
	//		fmt.Printf("Job related task id: %v\n", cronJob.TaskId)
	//		rv = append(rv, &biz.CronJob{
	//			TaskId: cronJob.TaskId,
	//		})
	//	}
	//}
	//return rv
	return nil
}

func (r *engineRepo) GetCronJobByTaskId(ctx context.Context, taskId int64) (cron.Entry, error) {
	//entries := r.data.cron.Entries()
	//for _, entry := range entries {
	//	if entry.Job.(*biz.CronJob).TaskId == taskId {
	//		return entry, nil
	//	}
	//}
	//return cron.Entry{}, errResponse.SetCustomizeErrMsg(errResponse.ReasonRecordNotFound, "entry not found")
	return cron.Entry{}, nil
}

func (r *engineRepo) RemoveCronJob(ctx context.Context, taskId int64) error {
	//entry, err := r.GetCronJobByTaskId(ctx, taskId)
	//if err != nil {
	//	return err
	//}
	//r.data.cron.Remove(entry.ID)
	return nil
}
