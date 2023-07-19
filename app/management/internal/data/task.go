package data

import (
	"context"
	engineV1 "galileo/api/engine/v1"
	taskV1 "galileo/api/management/task/v1"
	"galileo/app/management/internal/biz"
	"galileo/ent"
	"galileo/ent/task"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/errors"
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
		SetType(task.Type).
		SetRank(task.Rank).
		SetAssignee(task.Assignee).
		SetConfig(task.Config).
		SetDeadline(task.Deadline).
		SetDescription(task.Description).
		ClearTestcaseSuite().
		AddTestcaseSuiteIDs(task.TestcaseSuites...).
		Save(ctx)
	if err != nil {
		return false, err
	}
	if ret.Type == 1 || ret.Type == 2 {
		// TODO: 更新定时任务调度列表
	}
	return true, nil
}

func (r *taskRepo) TaskByName(ctx context.Context, name string) (*biz.Task, error) {
	queryTask, err := r.data.entDB.Task.Query().Where(task.Name(name)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Task{
		Id:        queryTask.ID,
		Name:      queryTask.Name,
		Rank:      queryTask.Rank,
		Status:    taskV1.TaskStatus(queryTask.Status),
		Type:      queryTask.Type,
		CreatedAt: queryTask.CreatedAt,
		CreatedBy: queryTask.CreatedBy,
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
		TestcaseSuites: suites,
	}, nil
}

func (r *taskRepo) CreateTask(ctx context.Context, task *biz.Task) (*biz.Task, error) {
	tx, err := r.data.entDB.Tx(ctx)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonSystemError, err.Error())
	}
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
		SetDescription(task.Description).
		AddTestcaseSuiteIDs(task.TestcaseSuites...).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	// 当创建的任务为延时任务、定时任务时，调用Engine服务添加到定时任务列表
	if createTask.Type == 1 || createTask.Type == 2 {
		res, err := r.data.engineCli.AddCronJob(ctx, &engineV1.AddCronJobRequest{
			TaskId:       createTask.ID,
			Type:         int32(createTask.Type),
			ScheduleTime: timestamppb.New(createTask.ScheduleTime),
			Frequency:    engineV1.Frequency(taskV1.Frequency_value[createTask.Frequency]),
		})
		if err != nil {
			return nil, rollback(tx, err)
		}
		if _, err := tx.Task.UpdateOneID(createTask.ID).
			SetExecuteID(res.ExecuteId).
			SetUpdatedBy(createTask.CreatedBy).
			Save(ctx); err != nil {
			return nil, rollback(tx, err)
		}
	}
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

func (r *taskRepo) ListTask(ctx context.Context, pageNum, pageSize int) ([]*biz.Task, error) {
	return nil, nil
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

func (r *taskRepo) UpdateTaskStatus(ctx context.Context, updateTask *biz.Task) (*biz.Task, error) {
	tx, err := r.data.entDB.Tx(ctx)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonSystemError, err.Error())
	}
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
	if updateTask.Status == taskV1.TaskStatus_RUNNING {
		// 当设置状态为RUNNING时，调用Engine服务进行测试任务的下发，参数：taskId，Config，Worker
		reply, err := r.data.engineCli.RunJob(ctx,
			&engineV1.RunJobRequest{
				TaskId: updateTask.Id,
				Conf:   updateTask.Config,
				Worker: updateTask.Worker,
			})
		if err != nil {
			return nil, rollback(tx, err)
		}
		_, err = tx.Task.UpdateOneID(updateTask.Id).
			SetExecuteID(reply.ExecuteId).
			SetWorker(updateTask.Worker).
			Save(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, rollback(tx, err)
	}
	return &biz.Task{
		Status:          taskV1.TaskStatus(ret.Status),
		StatusUpdatedAt: ret.StatusUpdatedAt,
	}, nil

}
