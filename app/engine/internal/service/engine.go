package service

import (
	"context"
	"fmt"
	v1 "galileo/api/engine/v1"
	taskV1 "galileo/api/management/task/v1"
	"galileo/app/engine/internal/biz"
	. "galileo/pkg/errResponse"
	"galileo/pkg/utils/snowflake"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/samber/lo"
	"time"
)

func (s *EngineService) Cron(ctx context.Context) {
	/* 设置循环频率为 5 min */
	interval := 5 * time.Minute
	/* 循环逻辑 */
	for {
		select {
		/* 上下文返回截止信号时退出循环 */
		case <-ctx.Done():
			return
		default:
			s.Scheduler(ctx)
		}
		time.Sleep(interval)
	}
}

func (s *EngineService) Scheduler(ctx context.Context) {
	s.log.Info("*** Running Scheduler ***")
	/* 获取处于新建状态 (NEW) 的定时任务列表 */
	timingTaskList, err := s.uc.TimingTaskList(ctx, []taskV1.TaskStatus{taskV1.TaskStatus_NEW})
	if err != nil {
		/* TODO：增加重试逻辑 */
		s.log.Error("Error getting timing task list")
	}
	s.log.Info("Getting Timing TaskList: ", timingTaskList)
	cronJobList := s.uc.GetCronJobList(ctx)
	/* 检查新建的任务是否处于调度列表，如果不在则放入调度列表 */
	for _, task := range timingTaskList {
		fmt.Println("Getting Timing Task: ", task)
		_, exist := lo.Find(cronJobList, func(item *biz.CronJob) bool {
			return item.TaskId == task.Id
		})
		if !exist {
			/* 调用函数将任务重新加入调度列表 */
			if _, err := s.uc.AddCronJob(ctx, &biz.Task{
				Id:           task.Id,
				Type:         task.Type,
				ScheduleTime: task.ScheduleTime,
				Frequency:    task.Frequency,
			}); err != nil {
				s.log.Info("Error adding Timing Task: ", err)
			}
		} else {
			/* 调整cronJobList调度列表 */
			cronJobList = lo.DropWhile(cronJobList, func(item *biz.CronJob) bool {
				return item.TaskId == task.Id
			})
		}
	}
	/* 从调度列表中删除无效的任务 */
	for _, item := range cronJobList {
		if err := s.uc.RemoveCronJob(ctx, item.TaskId); err != nil {
			s.log.Error("Error delete invalid timing task list")
		}
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
