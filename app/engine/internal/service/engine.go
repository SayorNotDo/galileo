package service

import (
	"context"
	pb "galileo/api/engine/v1"
	taskV1 "galileo/api/management/task/v1"
	"galileo/app/engine/internal/biz"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
)

type EngineService struct {
	pb.UnimplementedEngineServer

	uc  *biz.EngineUseCase
	log *log.Helper
}

func NewEngineService(uc *biz.EngineUseCase, logger log.Logger) *EngineService {
	return &EngineService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "engine.Service")),
	}
}

func (s *EngineService) RunJob(ctx context.Context, req *pb.RunJobRequest) (*pb.RunJobReply, error) {
	// TODO: query task from db
	task, err := s.uc.TaskByID(ctx, req.TaskId)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonRecordNotFound, err.Error())
	}
	// TODO: validate job configuration
	testcaseSuites := task.TestcaseSuites
	println("testcaseSuites: ", testcaseSuites)
	// TODO: Run testcase
	// TODO: validate worker environment
	// TODO: send job command through rabbitmq
	// TODO: get callback from rabbitmq
	// TODO: update task status
	if ok, err := s.uc.UpdateTaskStatus(ctx, req.TaskId, taskV1.TaskStatus_RUNNING); !ok {
		return nil, SetCustomizeErrMsg(ReasonRecordNotFound, err.Error())
	}
	// TODO: return response
	return &pb.RunJobReply{
		Success: true,
	}, nil
}

func (s *EngineService) BuildContainer(ctx context.Context, req *empty.Empty) (*pb.BuildContainerReply, error) {
	container, err := s.uc.BuildContainer(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.BuildContainerReply{
		Id:    container.Id,
		Name:  container.Name,
		Image: container.Image,
	}, nil
}

func (s *EngineService) TestEngine(ctx context.Context, req *empty.Empty) (*pb.TestEngineReply, error) {
	return &pb.TestEngineReply{Hello: "Galileo"}, nil
}
