package data

import (
	"context"
	"galileo/app/engine/internal/biz"
	. "galileo/app/engine/pkg/constant"
	"galileo/ent"
	"galileo/ent/job"
	"galileo/ent/user"
	"galileo/pkg/ctxdata"
	"galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
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

func (r *schedulerRepo) AddPeriodicJob(ctx context.Context, payload *biz.PeriodicJobPayload, expression string) (*biz.Job, error) {
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
	var err error
	/* TODO: 增加逻辑判断-是否已存在未调度且未删除的任务 */
	jobList, err := r.ListJobOnceTask(ctx, payload.Task)
	if err != nil {
		return nil, err
	}
	if len(jobList) >= 0 {
		lo.ForEach(jobList, func(item *biz.Job, _ int) {
			if !item.Active && item.DeletedBy == "" && item.DeletedAt.IsZero() {
				err = errResponse.SetCustomizeErrMsg(errResponse.ReasonParamsError, "task related job is already exists")
			}
		})
		if err != nil {
			return nil, err
		}
	}
	r.log.Debugf("jobList: %v", jobList)
	/* 方法调用顺序：1.创建记录 2.添加如调度器 3.调度成功后的任务信息写入 */
	tx, err := r.data.entDB.Tx(ctx)
	if err != nil {
		return nil, err
	}
	/* 调用方法创建Job记录，写入数据库 */
	ret, err := r.createJob(ctx, &biz.Job{
		TaskID: payload.Task,
		Worker: payload.Worker,
		Config: payload.Config,
	}, tx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	/* typeName、 payload、 delayTime */
	taskInfo, err := r.data.asynqSrv.NewTask(biz.TypeDelayedJob, payload, asynq.MaxRetry(0), asynq.ProcessIn(delay))
	if err != nil {
		return nil, rollback(tx, err)
	}
	/* 更新任务记录的 entry_id、payload、type */
	updateJob, err := tx.Job.UpdateOneID(ret.ID).
		SetEntryID(taskInfo.ID).
		SetPayload(taskInfo.Payload).
		SetType(taskInfo.Type).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	if err := tx.Commit(); err != nil {
		return nil, rollback(tx, err)
	}
	return &biz.Job{
		ID:      updateJob.ID,
		EntryID: updateJob.EntryID,
		Payload: updateJob.Payload,
		Type:    updateJob.Type,
		TaskID:  updateJob.TaskID,
		Worker:  updateJob.Worker,
		Config:  updateJob.Config,
	}, nil
}

func (r *schedulerRepo) createJob(ctx context.Context, job *biz.Job, tx *ent.Tx) (*biz.Job, error) {
	uid := ctxdata.UserIdFromMetaData(ctx)
	ret, err := tx.Job.Create().
		SetCreatedBy(uid).
		SetWorker(job.Worker).
		SetConfig(job.Config).
		SetTaskID(job.TaskID).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Job{
		ID:        ret.ID,
		Payload:   ret.Payload,
		Type:      ret.Type,
		CreatedAt: ret.CreatedAt,
		EntryID:   ret.EntryID,
		Worker:    ret.Worker,
		TaskID:    ret.TaskID,
	}, nil
}

func (r *schedulerRepo) RemovePeriodicJob(ctx context.Context, entryId string) error {
	return r.data.asynqSrv.RemovePeriodicTask(entryId)
}

func (r *schedulerRepo) RemoveDelayedJob(ctx context.Context, qName, jobID string) error {
	/* Redis通过删除对应的键值对实现延时任务的删除 */
	scheduledKey := biz.ScheduledKey(qName)
	jobKey := biz.JobKey(qName, jobID)
	/* Redis事务化操作 */
	txf := func(tx *redis.Tx) error {
		_, err := tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			/* 删除key：asynq:{qName}:scheduled中的指定member */
			cmd := r.data.redisCli.ZRem(ctx, scheduledKey, jobID)
			switch {
			case cmd.Err() != nil:
				return cmd.Err()
			case cmd.Val() <= 0:
				return errResponse.SetCustomizeErrMsg(errResponse.ReasonSystemError, "redis ZRem operation failed")
			}
			/* 删除key：asynq:{qName}:t: */
			if err := r.data.redisCli.Del(ctx, jobKey).Err(); err != nil {
				return err
			}
			return nil
		})
		return err
	}
	/* 乐观锁限制在jobKey没有被修改时进行操作 */
	for i := 0; i < RedisMaxRetries; i++ {
		err := r.data.redisCli.Watch(ctx, txf, jobKey)
		switch err {
		case nil:
			return nil
		case redis.TxFailedErr:
			continue
		}
	}
	return errResponse.SetCustomizeErrMsg(errResponse.ReasonSystemError, "max retries exceeded")
}

func (r *schedulerRepo) syncScheduledJobList(ctx context.Context, qName string) error {
	/* 同步scheduled作业队列 */
	/* 获取type为job:delayed、active为false且未被删除的作业列表 */
	jobList, err := r.data.entDB.Job.Query().
		Where(job.And(
			job.Type(biz.TypeDelayedJob),
			job.Active(false),
			job.DeletedAtIsNil(),
			job.DeletedByIsNil(),
		)).All(ctx)
	if err != nil {
		return err
	}
	r.log.Debugf("jobList: %v", jobList)
	/* 获取Redis中的scheduled中的Member列表 */
	scheduledKey := biz.ScheduledKey(qName)
	ret, err := r.data.redisCli.ZRange(ctx, scheduledKey, 0, -1).Result()
	r.log.Debugf("range list: %v", ret)
	/* 对比两个列表是否存在缺失的作业，重新加入调度列表 */
	return nil
}

func (r *schedulerRepo) ListScheduledJob(ctx context.Context, qName string) ([]*biz.Job, error) {
	/* 获取Redis中的scheduled下存在的作业 */
	key := biz.ScheduledKey(qName)
	/* Zrange */
	res, err := r.data.redisCli.ZRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	r.log.Debugf("redis result: %v", res)
	return nil, nil
}

func (r *schedulerRepo) ListPeriodicJob(ctx context.Context) ([]*biz.Job, error) {
	return nil, nil
}

func (r *schedulerRepo) ListJobOnceTask(ctx context.Context, taskID int64) ([]*biz.Job, error) {
	/* 声明返回的作业列表变量 */
	var jobList []*biz.Job
	/* 获取关联id的作业列表 */
	jobs, err := r.data.entDB.Job.Query().Where(job.TaskID(taskID)).All(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, errResponse.SetErrByReason(errResponse.ReasonRecordNotFound)
	case err != nil:
		return nil, err
	}
	/* 初始化创建者和删除者的用户ID列表 */
	createdByIDs := make([]uint32, len(jobs))
	deletedByIDs := make([]uint32, len(jobs))
	for k, v := range jobs {
		createdByIDs[k] = v.CreatedBy
		deletedByIDs[k] = v.DeletedBy
	}
	/* 查询创建者和删除者的用户名 */
	users, err := r.data.entDB.User.Query().Where(user.Or(user.IDIn(createdByIDs...), user.IDIn(deletedByIDs...))).All(ctx)
	if err != nil {
		return nil, err
	}
	/* 创建用户ID到用户名的映射 */
	userMap := make(map[uint32]string)
	lo.ForEach(users, func(item *ent.User, _ int) {
		userMap[item.ID] = item.Username
	})
	/* 遍历生成返回的作业列表 */
	lo.ForEach(jobs, func(item *ent.Job, _ int) {
		createdBy := userMap[item.CreatedBy]
		deletedBy := userMap[item.DeletedBy]
		jobList = append(jobList, &biz.Job{
			ID:        item.ID,
			Payload:   item.Payload,
			Type:      item.Type,
			CreatedAt: item.CreatedAt,
			CreatedBy: createdBy,
			UpdatedAt: item.UpdatedAt,
			Worker:    item.Worker,
			DeletedAt: item.DeletedAt,
			DeletedBy: deletedBy,
			EntryID:   item.EntryID,
			Config:    item.Config,
			TaskID:    item.TaskID,
			Active:    item.Active,
		})
	})
	return jobList, nil
}
