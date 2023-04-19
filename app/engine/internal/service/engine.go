package service

import (
	"context"
	"errors"
	"fmt"
	pb "galileo/api/engine/v1"
	taskV1 "galileo/api/management/task/v1"
	"galileo/app/engine/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type EngineService struct {
	pb.UnimplementedEngineServer

	uc  *biz.EngineUseCase
	log *log.Helper
}

func (s *EngineService) RunJob(ctx context.Context, req *pb.RunJobRequest) (*pb.RunJobReply, error) {
	// TODO: query task from db
	// TODO: validate job configuration
	// TODO: validate worker environment
	// TODO: send job command through rabbitmq
	// TODO: get callback from rabbitmq
	// TODO: update task status
	if ok, err := s.uc.Repo.UpdateTaskStatus(ctx, taskV1.TaskStatus_RUNNING); !ok {
		fmt.Printf("%v", err)
		return nil, errors.New("task status update failed")
	}
	// TODO: return response
	return &pb.RunJobReply{
		Success: true,
	}, nil
}
