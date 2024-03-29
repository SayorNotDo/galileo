package service

import (
	"context"
	"galileo/api/management/v1"
	"galileo/app/management/internal/biz"
	. "galileo/app/management/internal/pkg/constant"
	"galileo/pkg/ctxdata"
	. "galileo/pkg/errResponse"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (s *ManagementService) CreateTask(ctx context.Context, req *v1.CreateTaskRequest) (*v1.CreateTaskReply, error) {
	/* 工厂模式新建 Task 校验传入参数的合法性 */
	createTask, err := biz.NewTask(
		req.Name,
		req.Rank,
		req.Type,
		req.Description,
		req.TestcaseSuiteId,
		req.Assignee,
		req.Deadline.AsTime(),
	)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonParamsError, err.Error())
	}

	/* 获取创建任务的用户ID */
	uid := ctxdata.GetUserId(ctx)
	createTask.CreatedBy = uid

	/* 当任务类型为定时任务或延时任务时，校验设定时间的准确性 */
	switch req.Type {
	case PERIODIC, DELAYED:
		if !req.ScheduleTime.IsValid() || req.ScheduleTime.AsTime().Sub(time.Now()) < 0 {
			return nil, SetCustomizeErrMsg(ReasonParamsError, "illegal schedule time")
		}
		createTask.ScheduleTime = req.ScheduleTime.AsTime()
		createTask.Frequency = int8(req.Frequency)
	}
	s.logger.Debug("createTask: %v", createTask)

	/* 查询数据库中是否存在同名任务（同名且未删除） */
	if queryTask, _ := s.sc.QueryTaskByName(ctx, createTask.Name); queryTask != nil {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "duplicated task name")
	}

	/* 当传入的测试计划ID不为空时，查询测试计划是否存在 */
	if queryPlan, _ := s.uc.GetTestPlan(ctx, req.TestPlanId); queryPlan == nil {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "test plan is not exists")
	}
	createTask.TestPlanId = req.TestPlanId

	/* 创建任务 */
	ret, err := s.sc.CreateTask(ctx, &createTask)
	if err != nil {
		return nil, err
	}
	return &v1.CreateTaskReply{
		Id:        ret.Id,
		Name:      ret.Name,
		Status:    ret.Status,
		CreatedAt: timestamppb.New(ret.CreatedAt),
	}, nil
}

// ExecuteTask 执行任务
func (s *ManagementService) ExecuteTask(ctx context.Context, req *v1.ExecuteTaskRequest) (*empty.Empty, error) {
	err := s.sc.ExecuteTask(ctx, req.TaskID, req.Worker, []byte(req.Config))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *ManagementService) TaskInfo(ctx context.Context, req *v1.TaskInfoRequest) (*v1.Task, error) {
	var reply *biz.TaskInfo
	var err error
	switch ctxdata.MethodFromContext(ctx) {
	case "PUT":
		newTask, err := biz.NewTask(req.Name, req.Rank, req.Type, req.Description, req.TestcaseSuiteId, req.Assignee, req.Deadline.AsTime())
		if err != nil {
			return nil, err
		}
		reply, err = s.sc.UpdateTask(ctx, &newTask)
	case "GET":
		reply, err = s.sc.GetTask(ctx, req.Id)
	}
	if err != nil {
		return nil, err
	}
	return &v1.Task{
		Id:              reply.Id,
		Name:            reply.Name,
		Rank:            int32(reply.Rank),
		Assignee:        reply.Assignee,
		Type:            int32(reply.Type),
		Frequency:       int32(reply.Frequency),
		ScheduleTime:    timestamppb.New(reply.ScheduleTime),
		Description:     reply.Description,
		Status:          reply.Status,
		StartTime:       timestamppb.New(reply.StartTime),
		CompletedAt:     timestamppb.New(reply.CompletedAt),
		UpdatedAt:       timestamppb.New(reply.UpdatedAt),
		UpdatedBy:       reply.UpdatedBy,
		CreatedAt:       timestamppb.New(reply.CreatedAt),
		CreatedBy:       reply.CreatedBy,
		StatusUpdatedAt: timestamppb.New(reply.StatusUpdatedAt),
		Deadline:        timestamppb.New(reply.Deadline),
		Testplan:        reply.Testplan,
	}, nil
}

func (s *ManagementService) DeleteTask(ctx context.Context, req *v1.DeleteTaskRequest) (*v1.DeleteTaskReply, error) {
	return &v1.DeleteTaskReply{}, nil
}

// ListTimingTask 获取数据库中的延时/定时任务·
//func (s *ManagementService) ListTimingTask(ctx context.Context, req *v1.ListTimingTaskRequest) (*v1.ListTimingTaskReply, error) {
//	ret, err := s.sc.ListTimingTask(ctx, req.Status)
//	if err != nil {
//		return nil, err
//	}
//	return &v1.ListTimingTaskReply{
//		TaskList: ret,
//	}, nil
//}

// UpdateTaskStatus 更新数据库中指定任务的状态
// func (s *ManagementService) UpdateTaskStatus(ctx context.Context, req *v1.UpdateTaskStatusRequest) (*v1.UpdateTaskStatusReply, error) {
// 	/* 获取需要更新状态的任务 */
// 	queryTask, err := s.sc.TaskByID(ctx, req.Id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	/* 校验接口请求的状态变更值：不可变更状态为NEW，不可重复请求相同的状态 */
// 	if req.Status == v1.TaskStatus_NEW {
// 		return nil, SetCustomizeErrMsg(ReasonParamsError, "task's status cannot be convert to NEW")
// 	} else if req.Status == queryTask.Status {
// 		return nil, SetCustomizeErrMsg(ReasonParamsError, "task's status already changed")
// 	}

// 	/* 当更改的状态为RUNNING、同时获取的任务状态为NEW时，记录当前时间作为任务开始时间 */
// 	if req.Status == v1.TaskStatus_RUNNING && queryTask.Status == v1.TaskStatus_NEW {
// 		queryTask.StartTime = time.Now()
// 	}
// 	/* 更新状态 */
// 	queryTask.Status = req.Status
// 	ret, err := s.sc.UpdateTaskStatus(ctx, &biz.Task{
// 		Id:        req.Id,
// 		Status:    req.Status,
// 		StartTime: queryTask.StartTime,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	///* 当设置状态为EXCEPTION，且任务属于某一测试计划，需要变更测试计划的状态 */
// 	//if ret.Status == v1.TaskStatus_EXCEPTION && ret.Testplan != "" {
// 	//	/* 基于ID获取对应的测试计划 */
// 	//	queryPlan, err := s.uc.GetTestPlanById(ctx, ret.TestPlanId)
// 	//	if err != nil {
// 	//		return nil, err
// 	//	}
// 	//	queryPlan.Status = v1.Status_BLOCKED
// 	//	if err := s.uc.UpdateTestPlan(ctx, queryPlan); err != nil {
// 	//		return nil, err
// 	//	}
// 	//}
// 	return &v1.UpdateTaskStatusReply{
// 		StatusUpdatedAt: timestamppb.New(ret.StatusUpdatedAt),
// 		Status:          ret.Status,
// 	}, nil
// }

// GetTaskProgress returns the progress of a task
/* TODO: 基于redis实现任务进度的同步管理 */
//func (s *ManagementService) GetTaskProgress(ctx context.Context, req *v1.TaskProgressRequest) (*v1.TaskProgressReply, error) {
//	/* 接口参数获取指定任务信息 */
//	task, err := s.sc.TaskByID(ctx, req.Id)
//	if err != nil {
//		return nil, err
//	}
//	/* 基于任务信息构建key(规则: taskProgress:taskName:ExecuteId)，查询redis数据库中的缓存 */
//	taskKey := NewTaskProgressKey(task.Name, 0)
//	ret, err := s.sc.RedisLRangeTask(ctx, taskKey)
//	if err != nil {
//		return nil, err
//	}
//	/* 通过缓存计算当前测试进度、详情 */
//	fmt.Println(ret)
//	return &v1.TaskProgressReply{}, nil
//}

/* TODO: 重置任务时，清除指定Key的Redis缓存 */
