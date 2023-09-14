package data

import (
	"context"
	"galileo/app/engine/internal/biz"
	"galileo/pkg/ctxdata"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
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
	entryId, err := r.data.asynqSrv.NewPeriodicTask(expression, biz.TypePeriodicJob, payload)
	if err != nil {
		return nil, err
	}
	r.log.Info("addPeriodicJob EntryId: %v", entryId)
	return nil, nil
}

func (r *schedulerRepo) AddDefaultJob(ctx context.Context, payload *biz.DefaultJobPayload) (*biz.Job, error) {
	/* 调用asynq添加任务到队列中 */
	taskInfo, err := r.data.asynqSrv.NewTask(biz.TypeDefaultJob, payload, asynq.MaxRetry(0))
	r.log.Debugf("addDefaultJob: %v", taskInfo)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *schedulerRepo) AddDelayedJob(ctx context.Context, payload *biz.DelayedJobPayload, delay time.Duration) (*biz.Job, error) {
	/* typeName、 payload、 delayTime */
	taskInfo, err := r.data.asynqSrv.NewTask(biz.TypeDelayedJob, payload, asynq.MaxRetry(0), asynq.ProcessIn(delay))
	r.log.Debugf("addDefaultJob: %v", taskInfo)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *schedulerRepo) RemovePeriodicJob(ctx context.Context, entryId string) error {
	return r.data.asynqSrv.RemovePeriodicTask(entryId)
}

func (r *schedulerRepo) RemoveDelayedJob(ctx context.Context) error {
	return nil
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

func (r *schedulerRepo) GetQueue(ctx context.Context) []string {
	return nil
}

func (r *schedulerRepo) ListTimingJob(ctx context.Context) ([]*biz.Job, error) {
	return nil, nil
}
