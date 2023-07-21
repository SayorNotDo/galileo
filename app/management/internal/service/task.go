package service

import (
	"context"
	"errors"
	"galileo/api/management/task/v1"
	"galileo/app/management/internal/biz"
	"galileo/pkg/ctxdata"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type TaskService struct {
	v1.UnimplementedTaskServer

	uc     *biz.TaskUseCase
	logger *log.Helper
}

func NewTaskService(uc *biz.TaskUseCase, logger log.Logger) *TaskService {
	return &TaskService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "management.taskService")),
	}
}

func NewTask(name string, rank, taskType int32, description string, testcaseSuites []int64, scheduleTime time.Time, worker string, assignee uint32, frequency v1.Frequency, deadline time.Time, config string) (biz.Task, error) {
	if len(name) <= 0 {
		return biz.Task{}, errors.New("task name must not be empty")
	} else if rank < 0 || taskType < 0 {
		return biz.Task{}, errors.New("invalid parameter")
	} else if taskType == 1 || taskType == 2 {
		if scheduleTime.Unix() == 0 {
			return biz.Task{}, errors.New("schedule time can not be empty")
		}
	}
	return biz.Task{
		Name:           name,
		Rank:           int8(rank),
		Type:           int8(taskType),
		Description:    description,
		TestcaseSuites: testcaseSuites,
		Worker:         worker,
		Assignee:       assignee,
		ScheduleTime:   scheduleTime,
		Frequency:      frequency.String(),
		Deadline:       deadline,
		Config:         config,
	}, nil
}

func (s *TaskService) CreateTask(ctx context.Context, req *v1.CreateTaskRequest) (*v1.CreateTaskReply, error) {
	/* 工厂模式新建 Task 校验传入参数的合法性 */
	createTask, err := NewTask(
		req.Name,
		req.Rank,
		req.Type,
		req.Description,
		req.TestcaseSuiteId,
		req.ScheduleTime.AsTime(),
		req.Worker,
		req.Assignee,
		req.Frequency,
		req.Deadline.AsTime(),
		req.Config,
	)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonParamsError, err.Error())
	}
	/* 查询数据库中是否存在同名的 Task */
	if queryTask, _ := s.uc.TaskByName(ctx, createTask.Name); queryTask != nil {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "duplicated task name")
	}
	/* 获取创建任务的用户ID */
	uid := ctx.Value(ctxdata.UserIdKey)
	createTask.CreatedBy = uid.(uint32)
	/* 创建任务 */
	ret, err := s.uc.CreateTask(ctx, &createTask)
	if err != nil {
		return nil, err
	}
	return &v1.CreateTaskReply{
		Id:        ret.Id,
		Status:    ret.Status,
		CreatedAt: timestamppb.New(ret.CreatedAt),
	}, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, req *v1.UpdateTaskRequest) (*empty.Empty, error) {
	/* 查询 Task */
	ret, err := s.uc.TaskByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	/* 基于更新内容新建 Task */
	updateTask, err := NewTask(
		req.Name,
		req.Rank,
		req.Type,
		req.Description,
		req.TestcaseSuiteId,
		req.ScheduleTime.AsTime(),
		req.Worker,
		req.Assignee,
		req.Frequency,
		req.Deadline.AsTime(),
		req.Config,
	)
	if err != nil {
		return nil, err
	}
	/* 赋值 Task 更新的ID */
	updateTask.Id = ret.Id
	/* 更新 */
	if _, err := s.uc.UpdateTask(ctx, &updateTask); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, req *v1.DeleteTaskRequest) (*v1.DeleteTaskReply, error) {
	return &v1.DeleteTaskReply{}, nil
}

func (s *TaskService) TaskByID(ctx context.Context, req *v1.TaskByIDRequest) (*v1.GetTaskReply, error) {
	task, err := s.uc.TaskByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetTaskReply{
		Name:            task.Name,
		Type:            int32(task.Type),
		Rank:            int32(task.Rank),
		Description:     task.Description,
		TestcaseSuiteId: task.TestcaseSuites,
		Status:          task.Status,
		Config:          task.Config,
		Worker:          task.Worker,
		ScheduleTime:    timestamppb.New(task.ScheduleTime),
		Deadline:        timestamppb.New(task.Deadline),
		StartTime:       timestamppb.New(task.StartTime),
		Assignee:        task.Assignee,
	}, nil
}

// ListTimingTask 获取数据库中所有的延时/定时任务
func (s *TaskService) ListTimingTask(ctx context.Context, req *empty.Empty) (*v1.ListTimingTaskReply, error) {
	ret, err := s.uc.ListTimingTask(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ListTimingTaskReply{
		TaskList: ret,
	}, nil
}

func (s *TaskService) UpdateTaskStatus(ctx context.Context, req *v1.UpdateTaskStatusRequest) (*v1.UpdateTaskStatusReply, error) {
	/* 获取需要更新状态的测试任务 */
	queryTask, err := s.uc.TaskByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	/* 校验接口请求的状态变更值：不可变更状态为NEW，不可重复请求相同的状态 */
	if req.Status == v1.TaskStatus_NEW {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "task's status cannot be convert to NEW")
	} else if req.Status == queryTask.Status {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "task's status already changed")
	}
	/* 当获取的任务状态为NEW时，记录当前时间作为任务开始时间 */
	if queryTask.Status == v1.TaskStatus_NEW {
		queryTask.StartTime = time.Now()
	}
	if req.Status == v1.TaskStatus_RUNNING {
		queryTask.Worker = req.Worker
	}
	/* 更新状态 */
	queryTask.Status = req.Status
	ret, err := s.uc.UpdateTaskStatus(ctx, queryTask)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateTaskStatusReply{
		StatusUpdatedAt: timestamppb.New(ret.StatusUpdatedAt),
		Status:          ret.Status,
	}, nil
}
