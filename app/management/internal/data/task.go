package data

import (
	"context"
	engineV1 "galileo/api/engine/v1"
	taskV1 "galileo/api/management/task/v1"
	"galileo/app/management/internal/biz"
	. "galileo/app/management/internal/pkg/constant"
	"galileo/ent"
	"galileo/ent/task"
	. "galileo/pkg/errResponse"
	. "galileo/pkg/factory"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/redis/go-redis/v9"
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

func (r *taskRepo) UpdateTask(ctx context.Context, task *biz.Task) (bool, error) {
	ret, err := r.data.entDB.Task.UpdateOneID(task.Id).
		SetName(task.Name).
		SetRank(task.Rank).
		SetType(task.Type).
		SetDescription(task.Description).
		SetAssignee(task.Assignee).
		ClearTestcaseSuite().
		AddTestcaseSuiteIDs(task.TestcaseSuites...).
		SetDeadline(task.Deadline).
		SetConfig(task.Config).
		Save(ctx)
	if err != nil {
		return false, err
	}
	if ret.Type == 1 || ret.Type == 2 {
		_, err := r.data.engineCli.UpdateCronJob(ctx, &engineV1.UpdateCronJobRequest{
			TaskId:       ret.ID,
			Type:         int32(ret.Type),
			ScheduleTime: timestamppb.New(ret.ScheduleTime),
			Frequency:    engineV1.Frequency(engineV1.Frequency_value[ret.Frequency]),
		})
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (r *taskRepo) TaskByName(ctx context.Context, name string) (*biz.Task, error) {
	queryTask, err := r.data.entDB.Task.Query().Where(task.Name(name)).Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, errors.NotFound(ReasonRecordNotFound, err.Error())
	case err != nil:
		return nil, err
	}
	testcaseSuites := queryTask.Edges.TestcaseSuite
	suites := make([]int64, 0)
	for _, o := range testcaseSuites {
		suites = append(suites, o.ID)
	}
	return &biz.Task{
		Id:             queryTask.ID,
		Name:           queryTask.Name,
		Rank:           queryTask.Rank,
		Status:         taskV1.TaskStatus(queryTask.Status),
		Type:           queryTask.Type,
		CreatedAt:      queryTask.CreatedAt,
		CreatedBy:      queryTask.CreatedBy,
		StartTime:      queryTask.StartTime,
		TestcaseSuites: suites,
	}, nil
}

func (r *taskRepo) TaskByID(ctx context.Context, id int64) (*biz.Task, error) {
	queryTask, err := r.data.entDB.Task.Query().Where(task.ID(id)).Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, errors.NotFound(ReasonRecordNotFound, err.Error())
	case err != nil:
		return nil, err
	}
	testcaseSuites := queryTask.Edges.TestcaseSuite
	suites := make([]int64, 0)
	for _, o := range testcaseSuites {
		suites = append(suites, o.ID)
	}
	return &biz.Task{
		Id:             queryTask.ID,
		Name:           queryTask.Name,
		Rank:           queryTask.Rank,
		Status:         taskV1.TaskStatus(queryTask.Status),
		Type:           queryTask.Type,
		CreatedAt:      queryTask.CreatedAt,
		CreatedBy:      queryTask.CreatedBy,
		StartTime:      queryTask.StartTime,
		TestcaseSuites: suites,
	}, nil
}

func (r *taskRepo) CreateTask(ctx context.Context, task *biz.Task) (*biz.Task, error) {
	/* 创建事务 */
	tx, err := r.data.entDB.Tx(ctx)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonSystemError, err.Error())
	}
	/* 写入任务 */
	createTask, err := tx.Task.Create().
		SetName(task.Name).
		SetType(task.Type).
		SetRank(task.Rank).
		SetAssignee(task.Assignee).
		SetConfig(task.Config).
		SetWorker(task.Worker).
		SetScheduleTime(task.ScheduleTime).
		SetFrequency(task.Frequency).
		SetStartTime(task.StartTime).
		SetDeadline(task.Deadline).
		SetCreatedBy(task.CreatedBy).
		SetTestplanID(task.TestPlanId).
		SetDescription(task.Description).
		AddTestcaseSuiteIDs(task.TestcaseSuites...).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	/* 当创建的任务为延时任务、定时任务时，调用Engine服务添加到定时任务列表 */
	if createTask.Type == 1 || createTask.Type == 2 {
		/* 添加定时任务 */
		res, err := r.data.engineCli.AddCronJob(ctx, &engineV1.AddCronJobRequest{
			TaskId:       createTask.ID,
			Type:         int32(createTask.Type),
			ScheduleTime: timestamppb.New(createTask.ScheduleTime),
			Frequency:    engineV1.Frequency(taskV1.Frequency_value[createTask.Frequency]),
		})
		if err != nil {
			return nil, rollback(tx, err)
		}
		/* 添加成功后基于返回更新对应任务的ExecuteId */
		if _, err := tx.Task.UpdateOneID(createTask.ID).
			SetExecuteID(res.ExecuteId).
			SetUpdatedBy(createTask.CreatedBy).
			Save(ctx); err != nil {
			return nil, rollback(tx, err)
		}
	}
	/* 提交事务，失败则回滚 */
	if err := tx.Commit(); err != nil {
		return nil, rollback(tx, err)
	}
	return &biz.Task{
		Id:        createTask.ID,
		CreatedAt: createTask.CreatedAt,
		Status:    taskV1.TaskStatus(createTask.Status),
	}, nil
}

func (r *taskRepo) SoftDeleteTask(ctx context.Context, uid uint32, id int64) (bool, error) {
	err := r.data.entDB.Task.UpdateOneID(id).
		SetDeletedAt(time.Now()).
		SetDeletedBy(uid).
		Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *taskRepo) TaskDetailById(ctx context.Context, id int64) (*biz.Task, error) {
	queryTask, err := r.data.entDB.Task.Query().Where(task.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Task{
		Id:              queryTask.ID,
		Name:            queryTask.Name,
		Type:            queryTask.Type,
		Rank:            queryTask.Rank,
		Status:          taskV1.TaskStatus(queryTask.Status),
		CreatedAt:       queryTask.CreatedAt,
		CreatedBy:       queryTask.CreatedBy,
		Assignee:        queryTask.Assignee,
		Worker:          queryTask.Worker,
		Config:          queryTask.Config,
		Frequency:       queryTask.Frequency,
		ScheduleTime:    queryTask.ScheduleTime,
		UpdatedAt:       queryTask.UpdatedAt,
		UpdatedBy:       queryTask.UpdatedBy,
		StatusUpdatedAt: queryTask.StatusUpdatedAt,
		CompletedAt:     queryTask.CompletedAt,
		StartTime:       queryTask.StartTime,
		Deadline:        queryTask.Deadline,
		DeletedAt:       queryTask.DeletedAt,
		DeletedBy:       queryTask.DeletedBy,
		Description:     queryTask.Description,
		ExecuteId:       queryTask.ExecuteID,
	}, nil
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

func (r *taskRepo) ListTimingTask(ctx context.Context, status []taskV1.TaskStatus) ([]*taskV1.TaskInfo, error) {
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
	rv := make([]*taskV1.TaskInfo, 0)
	for _, v := range ret {
		rv = append(rv, &taskV1.TaskInfo{
			Id:           v.ID,
			Name:         v.Name,
			Rank:         int32(v.Rank),
			Type:         int32(v.Type),
			Assignee:     v.Assignee,
			Status:       taskV1.TaskStatus(v.Status),
			Worker:       v.Worker,
			ScheduleTime: timestamppb.New(v.ScheduleTime),
			Frequency:    taskV1.Frequency(taskV1.Frequency_value[v.Frequency]),
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

func (r *taskRepo) UpdateTaskStatus(ctx context.Context, updateTask *biz.Task) (*biz.Task, error) {
	/* 创建事务 */
	tx, err := r.data.entDB.Tx(ctx)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonSystemError, err.Error())
	}
	/* 更新任务的状态、开始时间、状态更新时间 */
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
	// 当设置状态为RUNNING时，调用Engine服务进行测试任务的下发，参数：taskId，Worker
	switch updateTask.Status {
	case taskV1.TaskStatus_RUNNING:
		reply, err := r.data.engineCli.RunJob(ctx,
			&engineV1.RunJobRequest{
				TaskId: updateTask.Id,
				Worker: updateTask.Worker,
				Schema: "http://",
				Type:   int32(updateTask.Type),
			})
		if err != nil {
			return nil, rollback(tx, err)
		}
		/* 下发测试任务成功后更新execute_id、worker */
		_, err = tx.Task.UpdateOneID(updateTask.Id).
			SetExecuteID(reply.ExecuteId).
			SetWorker(updateTask.Worker).
			Save(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}
	/* 当设置状态为FINISHED时，设置Redis缓存过期时间 */
	case taskV1.TaskStatus_FINISHED:
		key := NewTaskProgressKey(updateTask.Name, updateTask.ExecuteId)
		err := r.SetTaskInfoExpiration(ctx, key, int64(TaskExpiration))
		if err != nil {
			return nil, rollback(tx, err)
		}
	}
	/* 提交事务 */
	if err := tx.Commit(); err != nil {
		return nil, rollback(tx, err)
	}
	return &biz.Task{
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
