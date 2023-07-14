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

func NewTask(name string, rank, taskType int32, description string, testcaseSuites []int64) (biz.Task, error) {
	if len(name) <= 0 {
		return biz.Task{}, errors.New("task name must not be empty")
	}
	if rank < 0 || taskType < 0 {
		return biz.Task{}, errors.New("invalid parameter")
	}
	return biz.Task{
		Name:           name,
		Rank:           int8(rank),
		Type:           int8(taskType),
		Description:    description,
		TestcaseSuites: testcaseSuites,
	}, nil
}

func (s *TaskService) CreateTask(ctx context.Context, req *v1.CreateTaskRequest) (*v1.CreateTaskReply, error) {
	createTask, err := NewTask(req.Name, req.Rank, req.Type, req.Description, req.TestcaseSuiteId)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonParamsError, err.Error())
	}
	if queryTask, _ := s.uc.TaskByName(ctx, createTask.Name); queryTask != nil {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "duplicated task name")
	}
	uid := ctx.Value(ctxdata.UserIdKey)
	createTask.CreatedBy = uid.(uint32)
	ret, err := s.uc.CreateTask(ctx, &createTask)
	if err != nil {
		return nil, err
	}
	return &v1.CreateTaskReply{
		Id:        ret.Id,
		Status:    v1.TaskStatus(ret.Status),
		CreatedAt: timestamppb.New(ret.CreatedAt),
	}, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, req *v1.UpdateTaskRequest) (*empty.Empty, error) {
	ret, err := s.uc.TaskByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	updateTask, err := NewTask(req.Name, req.Rank, req.Type, req.Description, req.TestcaseSuiteId)
	if err != nil {
		return nil, err
	}
	updateTask.Id = ret.Id
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
	}, nil
}
func (s *TaskService) ListTask(ctx context.Context, req *v1.ListTaskRequest) (*v1.ListTaskReply, error) {
	return &v1.ListTaskReply{}, nil
}
