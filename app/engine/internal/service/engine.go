package service

import (
	"context"
	v1 "galileo/api/engine/v1"
)

// AddPeriodicJob
/* 添加任务到调度列表 */
func (s *EngineService) AddPeriodicJob(ctx context.Context, req *v1.AddPeriodicJobRequest) (*v1.AddPeriodicJobReply, error) {
	/* 构建Task对象 */
	//task := &biz.Task{
	//	Id:           req.TaskId,
	//	Type:         int8(req.Type),
	//	ScheduleTime: req.ScheduleTime.AsTime(),
	//	Frequency:    int8(req.Frequency),
	//}
	/* 增加定时任务到调度列表 */
	//_, err := s.ec.AddCronJob(ctx, task)
	//if err != nil {
	//	return nil, err
	//}
	return &v1.AddPeriodicJobReply{}, nil
}

// AddDefaultJob
/* 添加默认任务到default队列 */
func (s *EngineService) AddDefaultJob(ctx context.Context, req *v1.AddDefaultJobRequest) (*v1.AddDefaultJobReply, error) {
	/* 构建默认任务的Payload */
	//payload, err := biz.NewDefaultJobPayload(req.TaskId, req.Worker, req.Config)
	//if err != nil {
	//	return nil, err
	//}
	///* 调用添加默认任务至队列中 */
	return nil, nil
}

// AddDelayedJob
/* 添加延时任务到delayed队列 */
func (s *EngineService) AddDelayedJob(ctx context.Context, req *v1.AddDelayedJobRequest) (*v1.AddDelayedJobReply, error) {
	/* 构建延迟任务的Payload */
	//delay := req.DelayedTime.AsTime().Sub(time.Now())
	//payload, err := biz.NewDelayedJobPayload(req.TaskId, req.Worker, delay)
	//if err != nil {
	//	return nil, err
	//}
	/* 调用添加延迟任务到队列中 */
	//_, err = s.ec.AddDelayedJob(ctx, payload, delay)
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}

//func (s *EngineService) UpdateCronJob(ctx context.Context, req *v1.UpdateCronJobRequest) (*empty.Empty, error) {
//	/* 构建更新 Task */
//	task := &biz.Task{
//		Id:           req.TaskId,
//		Type:         int8(req.Type),
//		ScheduleTime: req.ScheduleTime.AsTime(),
//		Frequency:    int8(req.Frequency),
//	}
//	/* 更新定时任务逻辑：旧任务移除 */
//	err := s.ec.RemoveCronJob(ctx, task.Id)
//	if err != nil {
//		return nil, err
//	}
//	/* 添加新的定时任务 */
//	if _, err := s.ec.AddCronJob(ctx, task); err != nil {
//		return nil, err
//	}
//	return nil, nil
//}
