package service

import (
	"context"
	v1 "galileo/api/engine/v1"
	taskV1 "galileo/api/management/task/v1"
	"galileo/app/engine/internal/biz"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	"time"
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
	// TODO: 获取任务配置等信息
	task, err := s.uc.TaskByID(ctx, req.TaskId)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonRecordNotFound, err.Error())
	}
	// TODO: 验证配置文件
	testcaseSuites := task.TestcaseSuites
	println("testcaseSuites: ", testcaseSuites)
	// TODO: 运行用例
	// TODO: 验证用例环境
	// TODO: 下发执行命令
	// TODO: 回调
	if ok, err := s.uc.UpdateTaskStatus(ctx, req.TaskId, taskV1.TaskStatus_RUNNING); !ok {
		return nil, SetCustomizeErrMsg(ReasonRecordNotFound, err.Error())
	}
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
	// 设置循环频率为 5 min
	interval := 30 * time.Second
	// 循环逻辑
	for {
		select {
		case <-ctx.Done():
			s.log.Info("*--------------------------------*CronJobScheduler Stopped*--------------------------------*")
			return
		default:
			s.log.Info("*--------------------------------*Running loop routine*--------------------------------*")
			// TODO: 获取数据库中类型为定时任务、延时任务的列表记录
			taskList, err := s.uc.TimingTaskList(ctx)
			if err != nil {
				s.log.Error("Error getting timing task list")
			}
			s.log.Debug("Getting Timing TaskList: ", taskList)
			// TODO: 检查当前定时任务的调度列表
			s.uc.GetCronEntries(ctx)
			// TODO: 遍历定时任务、延时任务的列表记录，若记录不在定时任务中，则重新调度，若调度列表存在过期任务，则移除
		}
		time.Sleep(interval)
	}
}

func (s *EngineService) AddCronJob(ctx context.Context, req *v1.AddCronJobRequest) (*v1.AddCronJobReply, error) {
	//res, err := s.uc.TaskByID(ctx, req.TaskId)
	//if err != nil {
	//	return nil, err
	//}
	//scheduleTime := res.ScheduleTime
	//fmt.Println(scheduleTime)
	s.uc.AddCronJob(ctx)
	return &v1.AddCronJobReply{}, nil
}
