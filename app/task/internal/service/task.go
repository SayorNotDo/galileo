package service

import (
	"context"
	"galileo/app/task/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "galileo/api/task/v1"
)

type TaskService struct {
	pb.UnimplementedTaskServer

	uc     *biz.TaskUseCase
	logger *log.Helper
}

func NewTaskService(uc *biz.TaskUseCase, logger log.Logger) *TaskService {
	return &TaskService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "task")),
	}
}

func (s *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskReply, error) {
	return &pb.CreateTaskReply{}, nil
}
func (s *TaskService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskReply, error) {
	return &pb.UpdateTaskReply{}, nil
}
func (s *TaskService) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskReply, error) {
	return &pb.DeleteTaskReply{}, nil
}
func (s *TaskService) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskReply, error) {
	return &pb.GetTaskReply{}, nil
}
func (s *TaskService) ListTask(ctx context.Context, req *pb.ListTaskRequest) (*pb.ListTaskReply, error) {
	return &pb.ListTaskReply{}, nil
}
