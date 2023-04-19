package service

import (
	"context"
	"errors"
	"galileo/api/management/task/v1"
	"galileo/app/management/internal/biz"
	"galileo/pkg/ctxdata"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
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
		Status:    int32(ret.Status),
		CreatedAt: ret.CreatedAt.Unix(),
	}, nil
}

func (s *TaskService) UpdateTaskStatus(ctx context.Context, req *v1.UpdateTaskStatusRequest) (*v1.UpdateTaskStatusReply, error) {
	println("----------------------------------------------------------------")
	return &v1.UpdateTaskStatusReply{}, nil
}
func (s *TaskService) DeleteTask(ctx context.Context, req *v1.DeleteTaskRequest) (*v1.DeleteTaskReply, error) {
	return &v1.DeleteTaskReply{}, nil
}
func (s *TaskService) GetTask(ctx context.Context, req *v1.GetTaskRequest) (*v1.GetTaskReply, error) {
	return &v1.GetTaskReply{}, nil
}
func (s *TaskService) ListTask(ctx context.Context, req *v1.ListTaskRequest) (*v1.ListTaskReply, error) {
	return &v1.ListTaskReply{}, nil
}

func NewTask(name string, rank, status int32, description string, testcaseSuites []int64) (biz.Task, error) {
	if len(name) <= 0 {
		return biz.Task{}, errors.New("testcase name must not be empty")
	}
	if rank < 0 || status < 0 {
		return biz.Task{}, errors.New("invalid parameter")
	}
	return biz.Task{
		Name:           name,
		Rank:           int8(rank),
		Status:         int8(status),
		Description:    description,
		TestcaseSuites: testcaseSuites,
	}, nil
}
