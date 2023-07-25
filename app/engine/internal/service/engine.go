package service

import (
	"context"
	"fmt"
	v1 "galileo/api/engine/v1"
	taskV1 "galileo/api/management/task/v1"
	"galileo/app/engine/internal/biz"
	. "galileo/pkg/errResponse"
	"galileo/pkg/utils/snowflake"
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

func (s *EngineService) TestEngine(ctx context.Context, req *empty.Empty) (*v1.TestEngineReply, error) {
	return &v1.TestEngineReply{Hello: "Galileo"}, nil
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

func (s *EngineService) CronJobScheduler(ctx context.Context) {
	// 设置循环频率为 5 min
	interval := 5 * time.Minute
	// 循环逻辑
	for {
		select {
		case <-ctx.Done():
			s.log.Info("*--------------------------------*CronJobScheduler Stopped*--------------------------------*")
			return
		default:
			s.log.Info("*--------------------------------*Running loop routine*--------------------------------*")
			taskList, err := s.uc.TimingTaskList(ctx)
			if err != nil {
				s.log.Error("Error getting timing task list")
			}
			s.log.Debug("Getting Timing TaskList: ", taskList)
			// TODO: 检查当前定时任务的调度列表
			cronJobList := s.uc.GetCronJobList(ctx)
			for _, v := range cronJobList {
				s.log.Debug("Getting Cron Job: ", v.TaskId)
			}
			// TODO: 遍历定时任务、延时任务的列表记录，若记录不在定时任务中，则重新调度，若调度列表存在过期任务，则移除
			timingTaskList, err := s.uc.TimingTaskList(ctx)
			if err != nil {
				s.log.Error("Getting Timing Task List Error")
			}
			for _, task := range timingTaskList {
				fmt.Println("Getting Timing Task: ", task)
			}

		}
		time.Sleep(interval)
	}
}

func (s *EngineService) AddCronJob(ctx context.Context, req *v1.AddCronJobRequest) (*v1.AddCronJobReply, error) {
	/* 构建Task对象 */
	task := &biz.Task{
		Id:           req.TaskId,
		Type:         int8(req.Type),
		ScheduleTime: req.ScheduleTime.AsTime(),
		Frequency:    taskV1.Frequency(req.Frequency),
	}
	/* 增加定时任务到调度列表 */
	_, err := s.uc.AddCronJob(ctx, task)
	if err != nil {
		return nil, err
	}
	/* 雪花算法生成执行ID */
	sn, err := snowflake.NewSnowflake(int64(0), int64(0))
	if err != nil {
		return nil, err
	}
	executeId := sn.NextVal()
	return &v1.AddCronJobReply{
		ExecuteId: executeId,
	}, nil
}

func (s *EngineService) UpdateCronJob(ctx context.Context, req *v1.UpdateCronJobRequest) (*empty.Empty, error) {
	/* 构建更新 Task */
	task := &biz.Task{
		Id:           req.TaskId,
		Type:         int8(req.Type),
		ScheduleTime: req.ScheduleTime.AsTime(),
		Frequency:    taskV1.Frequency(req.Frequency),
	}
	/* 查询定时任务调度列表 */
	cronJobList := s.uc.GetCronJobList(ctx)
	for _, v := range cronJobList {
		if v.TaskId == task.Id {
			/* 更新定时任务逻辑：旧任务移除 */
			err := s.uc.RemoveCronJob(ctx, task.Id)
			if err != nil {
				return nil, err
			}
			/* 添加新的定时任务 */
			if _, err := s.uc.AddCronJob(ctx, task); err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}
func (s *EngineService) RunJob(ctx context.Context, req *v1.RunJobRequest) (*v1.RunJobReply, error) {
	s.log.Info("*--------------------------------*Run Job*--------------------------------*")
	res, err := s.uc.TaskByID(ctx, req.TaskId)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonRecordNotFound, err.Error())
	}
	//if err := biz.JobCommand(req.Schema, req.Worker, req.TaskId); err != nil {
	//	return nil, err
	//}
	/* 雪花算法生成执行ID */
	if req.Type == 0 {
		sn, err := snowflake.NewSnowflake(int64(0), int64(0))
		if err != nil {
			return nil, err
		}
		res.ExecuteId = sn.NextVal()
	}
	return &v1.RunJobReply{
		ExecuteId: res.ExecuteId,
	}, nil
}
