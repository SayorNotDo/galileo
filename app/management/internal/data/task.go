package data

import (
	"context"
	engineV1 "galileo/api/engine/v1"
	taskV1 "galileo/api/management/v1"
	"galileo/app/management/internal/biz"
	. "galileo/app/management/internal/pkg/constant"
	"galileo/ent"
	"galileo/ent/task"
	"galileo/ent/testplan"
	"galileo/ent/user"
	"galileo/pkg/ctxdata"
	. "galileo/pkg/errResponse"
	. "galileo/pkg/factory"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

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
		log:  log.NewHelper(log.With(logger, "module", "management.dataService.taskRepo")),
	}
}

func (r *taskRepo) UpdateTask(ctx context.Context, task *biz.Task) (*biz.TaskInfo, error) {
	/* 创建事务 */
	queryTask, err := r.GetTask(ctx, task.Id)
	if err != nil {
		return nil, err
	}
	if queryTask.Type != task.Type {
		/* TODO：基于原任务类型与修改后的任务类型，更新调度器中的任务记录 */
		/* 移除调度器中未执行的任务 */
		/* 基于修改后的任务类型处理逻辑 */
	}
	_, err = r.data.entDB.Task.UpdateOneID(task.Id).
		SetName(task.Name).
		SetRank(task.Rank).
		SetType(task.Type).
		SetDescription(task.Description).
		SetAssignee(task.Assignee).
		SetTestcaseSuite(task.TestcaseSuite).
		SetDeadline(task.Deadline).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *taskRepo) QueryTaskByName(ctx context.Context, name string) (*biz.TaskInfo, error) {
	/* 通过Name查询DeletedBy、DeletedAt为空的记录 */
	queryTask, err := r.data.entDB.Task.Query().
		Where(
			task.Name(name),
			task.DeletedByIsNil(),
			task.DeletedAtIsNil(),
		).
		Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, errors.NotFound(ReasonRecordNotFound, err.Error())
	case err != nil:
		return nil, err
	}
	return &biz.TaskInfo{
		Id:            queryTask.ID,
		Name:          queryTask.Name,
		Rank:          queryTask.Rank,
		Status:        taskV1.TaskStatus(queryTask.Status),
		Type:          queryTask.Type,
		StartTime:     queryTask.StartTime,
		TestcaseSuite: queryTask.TestcaseSuite,
	}, nil
}

func (r *taskRepo) GetTask(ctx context.Context, id int64) (*biz.TaskInfo, error) {
	var queryTask *biz.TaskInfo
	ret, err := r.data.entDB.Task.Query().Where(task.ID(id)).Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, errors.NotFound(ReasonRecordNotFound, err.Error())
	case err != nil:
		return nil, err
	}
	queryTask = &biz.TaskInfo{
		Id:              ret.ID,
		Name:            ret.Name,
		Rank:            ret.Rank,
		Type:            ret.Type,
		Frequency:       ret.Frequency,
		ScheduleTime:    ret.ScheduleTime,
		Status:          taskV1.TaskStatus(ret.Status),
		CreatedAt:       ret.CreatedAt,
		UpdatedAt:       ret.UpdatedAt,
		StatusUpdatedAt: ret.StatusUpdatedAt,
		StartTime:       ret.StartTime,
		Deadline:        ret.Deadline,
		Description:     ret.Description,
		TestcaseSuite:   ret.TestcaseSuite,
	}
	/* 获取任务关联的测试计划相关信息 */
	planName, err := r.data.entDB.TestPlan.Query().Where(testplan.ID(ret.TestplanID)).Only(ctx)
	switch {
	case ent.IsNotFound(err):
	case err != nil:
		return nil, err
	default:
		queryTask.Testplan = planName.Name
	}
	/* TODO: 获取测试用例集合的具体信息 */
	users, err := r.data.entDB.User.Query().Where(user.IDIn(ret.CreatedBy, ret.Assignee, ret.UpdatedBy)).All(ctx)
	if err != nil {
		return nil, err
	}
	lo.ForEach(users, func(item *ent.User, _ int) {
		switch item.ID {
		case ret.CreatedBy:
			queryTask.CreatedBy = item.Username
		case ret.Assignee:
			queryTask.Assignee = item.Username
		case ret.UpdatedBy:
			queryTask.UpdatedBy = item.Username
		}
	})
	return queryTask, nil
}

func (r *taskRepo) CreateTask(ctx context.Context, task *biz.Task) (*biz.TaskInfo, error) {
	/* 创建SQL事务 */
	tx, err := r.data.entDB.Tx(ctx)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonSystemError, err.Error())
	}

	/* 创建任务记录 */
	createTask, err := tx.Task.Create().
		SetName(task.Name).
		SetType(task.Type).
		SetRank(task.Rank).
		SetAssignee(task.Assignee).
		SetScheduleTime(task.ScheduleTime).
		SetFrequency(task.Frequency).
		SetStartTime(task.StartTime).
		SetDeadline(task.Deadline).
		SetCreatedBy(task.CreatedBy).
		SetTestplanID(task.TestPlanId).
		SetDescription(task.Description).
		SetTestcaseSuite(task.TestcaseSuite).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	/* 提交事务，失败则回滚 */
	if err := tx.Commit(); err != nil {
		return nil, rollback(tx, err)
	}
	return &biz.TaskInfo{
		Id:        createTask.ID,
		Name:      createTask.Name,
		Status:    taskV1.TaskStatus(createTask.Status),
		CreatedAt: createTask.CreatedAt,
	}, nil
}

func (r *taskRepo) ExecuteTask(ctx context.Context, taskId int64, worker uint32, config []byte) error {
	/* 获取待执行任务的信息 */
	r.log.Debugf("------------->>>>> get user id: %v", ctxdata.GetUserId(ctx))
	r.log.Debugf("ExecuteTask taskId: %d, worker: %d, config: %v", taskId, worker, config)
	executeTask, err := r.GetTask(ctx, taskId)
	if err != nil {
		return err
	}

	switch executeTask.Type {
	case DEFAULT:
		_, err := r.data.engineCli.AddDefaultJob(ctx, &engineV1.AddDefaultJobRequest{
			TaskId: taskId,
			Worker: worker,
			Config: config,
		})
		if err != nil {
			return err
		}
	/* 当创建的任务为延时任务、定时任务时，调用Engine服务添加到定时任务列表 */
	// TODO: 重构延时任务、定时任务添加逻辑
	/* 基于任务类型调用Engine的Job服务添加至不同队列 */
	case DELAYED:
		/* 添加延时任务 */
		if _, err := r.data.engineCli.AddDelayedJob(ctx, &engineV1.AddDelayedJobRequest{
			TaskID:      taskId,
			Worker:      worker,
			Config:      config,
			DelayedTime: timestamppb.New(time.Now().Add(time.Second * 60)),
		}); err != nil {
			return err
		}
	case PERIODIC:
		/* 添加定时任务 */
		if _, err := r.data.engineCli.AddPeriodicJob(ctx, &engineV1.AddPeriodicJobRequest{}); err != nil {
			return err
		}
	}
	return nil
}
func (r *taskRepo) SoftDeleteTask(ctx context.Context, uid uint32, taskId int64) (bool, error) {
	err := r.data.entDB.Task.UpdateOneID(taskId).
		SetDeletedAt(time.Now()).
		SetDeletedBy(uid).
		Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *taskRepo) IsTaskDeleted(ctx context.Context, id int64) (bool, error) {
	queryTask, err := r.data.entDB.Task.Query().Where(task.ID(id)).Only(ctx)
	if err != nil {
		return false, err
	}
	if !queryTask.CompletedAt.IsZero() && queryTask.DeletedBy > 0 {
		return true, nil
	}
	return false, err
}

func (r *taskRepo) ListTimingTask(ctx context.Context, status []taskV1.TaskStatus) ([]*taskV1.Task, error) {
	var numList []int8
	/* 枚举状态值转化为Int8类型 */
	for _, t := range status {
		numList = append(numList, int8(t.Number()))
	}
	/* 查询 */
	ret, err := r.data.entDB.Task.Query().
		Where(task.TypeIn(1, 2)).
		Where(task.StatusIn(numList...)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*taskV1.Task, 0)
	for _, v := range ret {
		rv = append(rv, &taskV1.Task{
			Id:           v.ID,
			Name:         v.Name,
			Rank:         int32(v.Rank),
			Type:         int32(v.Type),
			Status:       taskV1.TaskStatus(v.Status),
			ScheduleTime: timestamppb.New(v.ScheduleTime),
			Deadline:     timestamppb.New(v.Deadline),
			StartTime:    timestamppb.New(v.StartTime),
		})
	}
	return rv, nil
}

func (r *taskRepo) ListAll(context.Context) ([]*biz.Task, error) {
	return nil, nil
}

func (r *taskRepo) CountAllTask(ctx context.Context) (int, error) {
	count, err := r.data.entDB.Task.Query().Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *taskRepo) RedisLRangeTask(ctx context.Context, key string) ([]string, error) {
	valList, err := r.data.redisCli.LRange(ctx, key, 0, -1).Result()
	switch {
	case err == redis.Nil:
		return nil, SetCustomizeErrMsg(ReasonRecordNotFound, "key not found")
	case err != nil:
		return nil, err
	}
	/* 返回值：进度（已执行次数 / 需执行次数），详情 */
	return valList, nil
}

func (r *taskRepo) UpdateTaskStatus(ctx context.Context, updateTask *biz.Task) (*biz.TaskInfo, error) {
	/*
		任务的状态：
		1.NEW
		2.PENDING
		3.RUNNING
		4.PAUSED
		5.CLOSED
		6.EXCEPTION
		7.CANCELLED
		8.FINISHED
	*/
	/* 创建事务 */
	tx, err := r.data.entDB.Tx(ctx)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonSystemError, err.Error())
	}
	/* 修改任务的状态、开始时间、状态更新时间 */
	ret, err := tx.Task.UpdateOneID(updateTask.Id).
		SetStatus(int8(updateTask.Status.Number())).
		SetStartTime(updateTask.StartTime).
		SetStatusUpdatedAt(time.Now()).
		Save(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, rollback(tx, errors.NotFound(ReasonRecordNotFound, err.Error()))
	case err != nil:
		return nil, rollback(tx, err)
	}
	/* 当设置状态为RUNNING时，调用Engine服务进行测试任务的下发，参数：taskId，Worker */
	switch updateTask.Status {
	case taskV1.TaskStatus_RUNNING:
	/* 当设置状态为EXCEPTION时 */
	/* 当设置状态为FINISHED时，设置Redis缓存过期时间 */
	case taskV1.TaskStatus_FINISHED:
		key := NewTaskProgressKey(updateTask.Name, 0)
		err := r.SetTaskInfoExpiration(ctx, key, int64(TaskExpiration))
		if err != nil {
			return nil, rollback(tx, err)
		}
	}
	/* 提交事务 */
	if err := tx.Commit(); err != nil {
		return nil, rollback(tx, err)
	}
	return &biz.TaskInfo{
		Status:          taskV1.TaskStatus(ret.Status),
		StatusUpdatedAt: ret.StatusUpdatedAt,
	}, nil

}

// TODO: 服务端主动获取执行机状态(WebSocket，KIM：心跳保持)
/* 方案二具体步骤：
1. 服务端向执行机发送WS协议升级信息
2. 客户端发起WS连接
3. 服务端获取执行机当前处理信息
*/

func (r *taskRepo) SetTaskInfoExpiration(ctx context.Context, key string, expiration int64) error {
	/* 事务函数：乐观锁实现任务数据缓存过期时间设置 */
	txf := func(tx *redis.Tx) error {
		cmd := r.data.redisCli.Get(ctx, key)
		if cmd.Err() != nil {
			return cmd.Err()
		}
		val := cmd.Val()

		_, err := tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			expire := pipe.TTL(ctx, key).Val()
			if expire == -1 {
				pipe.SetXX(ctx, key, val, time.Duration(expiration))
				return nil
			}
			pipe.SetXX(ctx, key, val, expire)
			return nil
		})
		return err
	}
	for i := 0; i < RedisMaxRetries; i++ {
		err := r.data.redisCli.Watch(ctx, txf, key)
		switch err {
		case nil:
			return nil
		case redis.TxFailedErr:
			continue
		}
	}
	return SetCustomizeErrMsg(ReasonSystemError, "max retries exceeded")
}
