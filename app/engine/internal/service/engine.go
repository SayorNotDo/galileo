package service

import (
	"context"
	v1 "galileo/api/engine/v1"
	taskV1 "galileo/api/management/task/v1"
	"galileo/app/engine/internal/biz"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
)

type EngineService struct {
	v1.UnimplementedEngineServer

	uc  *biz.EngineUseCase
	log *log.Helper
}

func NewEngineService(uc *biz.EngineUseCase, logger log.Logger) *EngineService {
	return &EngineService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "engine.Service")),
	}
}

func (s *EngineService) RunJob(ctx context.Context, req *v1.RunJobRequest) (*v1.RunJobReply, error) {
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
	// TODO: send job command
	// TODO: get callback
	// TODO: update task status
	if ok, err := s.uc.UpdateTaskStatus(ctx, req.TaskId, taskV1.TaskStatus_RUNNING); !ok {
		return nil, SetCustomizeErrMsg(ReasonRecordNotFound, err.Error())
	}
	// TODO: return response
	return &v1.RunJobReply{}, nil
}

func (s *EngineService) BuildContainer(ctx context.Context, req *empty.Empty) (*v1.BuildContainerReply, error) {
	container, err := s.uc.BuildContainer(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.BuildContainerReply{
		Id:    container.Id,
		Name:  container.Name,
		Image: container.Image,
	}, nil
}

func (s *EngineService) TestEngine(ctx context.Context, req *empty.Empty) (*v1.TestEngineReply, error) {
	return &v1.TestEngineReply{Hello: "Galileo"}, nil
}

func (s *EngineService) CronJobScheduler(ctx context.Context) {
	// TODO: Scheduler get all cron jobs
}

func (s *EngineService) AddCronJob(ctx context.Context, req *v1.AddCronJobRequest) (*v1.AddCronJobReply, error) {
	//res, err := s.uc.TaskByID(ctx, req.TaskId)
	//if err != nil {
	//	return nil, err
	//}
	//scheduleTime := res.ScheduleTime
	//fmt.Println(scheduleTime)
	return &v1.AddCronJobReply{}, nil
}
