package service

import (
	"context"
	v1 "galileo/api/engine/v1"
	managementV1 "galileo/api/management/v1"
	"galileo/app/engine/internal/biz"
	. "galileo/pkg/errResponse"
	"galileo/pkg/utils/snowflake"
	"github.com/golang/protobuf/ptypes/empty"
)

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
/* 添加延时任务到delayed队列 */
func (s *EngineService) AddDelayedJob(ctx context.Context, req *v1.AddDelayedJobRequest) (*v1.AddDelayedJobReply, error) {
	/* 构建Task对象 */
	//task := &biz.Task{
	//	Id: req.TaskId,
	//}
	return nil, nil
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
