package service

import (
	"context"
	"galileo/api/management/task/v1"
	"galileo/app/management/internal/biz"
	"galileo/pkg/ctxdata"
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
	uid := ctx.Value(ctxdata.UserIdKey)
	createTask := &biz.Task{
		CreatedBy: uid.(uint32),
		Name:      req.Name,
		Rank:      int8(req.Rank),
		Type:      int8(req.Type),
	}
	println(createTask)
	return &v1.CreateTaskReply{}, nil
}
func (s *TaskService) UpdateTask(ctx context.Context, req *v1.UpdateTaskRequest) (*v1.UpdateTaskReply, error) {
	return &v1.UpdateTaskReply{}, nil
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
