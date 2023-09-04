package service

import (
	"context"
	"fmt"
	v1 "galileo/api/engine/v1"
	managementV1 "galileo/api/management/v1"
	"galileo/app/engine/internal/biz"
	. "galileo/pkg/errResponse"
	"galileo/pkg/utils/snowflake"
	"github.com/avast/retry-go"
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
			/* 调度器执行 */
			s.Scheduler(ctx)
		}
		time.Sleep(interval)
	}
}

// Scheduler
/* 定时任务调度器的主逻辑函数 */
func (s *EngineService) Scheduler(ctx context.Context) {
	s.log.Info("*** Scheduler ***")
	/* 定义重试策略 */
	retryStrategy := []retry.Option{
		retry.Delay(100 * time.Millisecond),
		retry.Attempts(5),
		retry.LastErrorOnly(true),
	}
	/* 获取处于新建状态 (NEW) 的定时任务列表 */
	var timingTaskList []*biz.Task
	var err error
	err = retry.Do(func() error {
		timingTaskList, err = s.uc.TimingTaskList(ctx, []managementV1.TaskStatus{managementV1.TaskStatus_NEW})
		if err != nil {
			return err
		}
		return nil
	}, retryStrategy...)
	if err != nil {
		s.log.Error("Error getting timing task list after 5 retries")
	}
	s.log.Info("Getting Timing TaskList: ", timingTaskList)
	/* 获取调度器的任务列表 */
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
		s.log.Infof("Removed Timing Task: %v : %v", item.TaskId, item.Worker)
	}

}

// AddCronJob
/* 添加任务到调度列表 */
func (s *EngineService) AddCronJob(ctx context.Context, req *v1.AddCronJobRequest) (*v1.AddCronJobReply, error) {
	/* 构建Task对象 */
	task := &biz.Task{
		Id:           req.TaskId,
		Type:         int8(req.Type),
		ScheduleTime: req.ScheduleTime.AsTime(),
		Frequency:    managementV1.Frequency(req.Frequency),
	}
	/* 增加定时任务到调度列表 */
	_, err := s.uc.AddCronJob(ctx, task)
	if err != nil {
		return nil, err
	}
	/* 雪花算法生成executeId */
	sn, err := snowflake.NewSnowflake(int64(0), int64(0))
	if err != nil {
		return nil, err
	}
	executeId := sn.NextVal()
	return &v1.AddCronJobReply{
		ExecuteId: executeId,
	}, nil
}

// AddDelayedJob
/* 添加延时任务到delayed对立 */
func (s *EngineService) AddDelayedJob(ctx context.Context, req *v1.AddDelayedJobRequest) {
	/* 构建Task对象 */
	//task := &biz.Task{
	//	Id: req.TaskId,
	//}
}

func (s *EngineService) UpdateCronJob(ctx context.Context, req *v1.UpdateCronJobRequest) (*empty.Empty, error) {
	/* 构建更新 Task */
	task := &biz.Task{
		Id:           req.TaskId,
		Type:         int8(req.Type),
		ScheduleTime: req.ScheduleTime.AsTime(),
		Frequency:    managementV1.Frequency(req.Frequency),
	}
	/* 更新定时任务逻辑：旧任务移除 */
	err := s.uc.RemoveCronJob(ctx, task.Id)
	if err != nil {
		return nil, err
	}
	/* 添加新的定时任务 */
	if _, err := s.uc.AddCronJob(ctx, task); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *EngineService) RunJob(ctx context.Context, req *v1.RunJobRequest) (*v1.RunJobReply, error) {
	res, err := s.uc.TaskByID(ctx, req.TaskId)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonRecordNotFound, err.Error())
	}
	if err := biz.JobCommand(req.Schema, req.Worker, req.TaskId); err != nil {
		return nil, err
	}
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
