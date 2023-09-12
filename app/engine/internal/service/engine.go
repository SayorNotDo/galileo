package service

import (
	"context"
	v1 "galileo/api/engine/v1"
	"galileo/app/engine/internal/biz"
	"time"
)

// AddPeriodicJob
/* 添加任务到调度列表 */
func (s *EngineService) AddPeriodicJob(ctx context.Context, req *v1.AddPeriodicJobRequest) (*v1.AddPeriodicJobReply, error) {
	/* 生成定时规则 */
	expression := biz.NewScheduleExpression(req.ScheduleTime.AsTime(), req.Frequency.String())
	/* 生成定时任务payload */
	payload, err := biz.NewPeriodicJobPayload(req.TaskId, req.Worker, expression)
	if err != nil {
		return nil, err
	}
	/* 增加定时任务到调度列表 */
	_, err = s.sc.AddPeriodicJob(ctx, payload, expression)
	if err != nil {
		return nil, err
	}
	return &v1.AddPeriodicJobReply{}, nil
}

// AddDefaultJob
/* 添加默认任务到default队列 */
func (s *EngineService) AddDefaultJob(ctx context.Context, req *v1.AddDefaultJobRequest) (*v1.AddDefaultJobReply, error) {
	/* 构建默认任务的Payload */
	payload, err := biz.NewDefaultJobPayload(req.TaskId, req.Worker, req.Config)
	if err != nil {
		return nil, err
	}
	/* 调用添加默认任务至队列中 */
	_, err = s.sc.AddDefaultJob(ctx, payload)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// AddDelayedJob
/* 添加延时任务到delayed队列 */
func (s *EngineService) AddDelayedJob(ctx context.Context, req *v1.AddDelayedJobRequest) (*v1.AddDelayedJobReply, error) {
	/* 构建延迟任务的Payload */
	delay := req.DelayedTime.AsTime().Sub(time.Now())
	payload, err := biz.NewDelayedJobPayload(req.TaskId, req.Worker, delay)
	if err != nil {
		return nil, err
	}
	/* 调用添加延迟任务到队列中 */
	_, err = s.sc.AddDelayedJob(ctx, payload, delay)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
