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
	err := r.data.asynqSrv.NewPeriodicTask(expression, biz.TypePeriodicJob, payload)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *schedulerRepo) AddDefaultJob(ctx context.Context, payload []byte) (*biz.Job, error) {
	/* 调用asynq添加任务到队列中 */
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

func (r *schedulerRepo) RemovePeriodicJob(ctx context.Context, entryId string) error {
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

func (r *schedulerRepo) ListTimingJob(ctx context.Context) ([]*biz.Job, error) {
	return nil, nil
}
