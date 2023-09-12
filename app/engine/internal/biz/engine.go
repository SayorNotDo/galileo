package biz

import (
	"context"
	"time"
)

type SchedulerRepo interface {
	ListTimingJob(ctx context.Context) ([]*Job, error)
	CreateJob(ctx context.Context, job *Job) (*Job, error)
	AddPeriodicJob(ctx context.Context, payload []byte, expression string) (*Job, error)
	AddDefaultJob(ctx context.Context, payload []byte) (*Job, error)
	AddDelayedJob(ctx context.Context, payload []byte, delay time.Duration) (*Job, error)
}

func (sc *SchedulerUseCase) AddPeriodicJob(ctx context.Context, payload []byte, expression string) (*Job, error) {
	return sc.repo.AddPeriodicJob(ctx, payload, expression)
}

func (sc *SchedulerUseCase) AddDefaultJob(ctx context.Context, payload []byte) (*Job, error) {
	return sc.repo.AddDefaultJob(ctx, payload)
}

func (sc *SchedulerUseCase) AddDelayedJob(ctx context.Context, payload []byte, delay time.Duration) (*Job, error) {
	return sc.repo.AddDelayedJob(ctx, payload, delay)
}

func (sc *SchedulerUseCase) CreateJob(ctx context.Context, job *Job) (*Job, error) {
	return sc.repo.CreateJob(ctx, job)
}

func (sc *SchedulerUseCase) TimingTaskJob(ctx context.Context) ([]*Job, error) {
	return sc.repo.ListTimingJob(ctx)
}
