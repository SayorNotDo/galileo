package biz

import (
	"context"
	"time"
)

type SchedulerRepo interface {
	ListTimingJob(ctx context.Context) ([]*Job, error)
	CreateJob(ctx context.Context, job *Job) (*Job, error)
	AddPeriodicJob(ctx context.Context, payload *PeriodicJobPayload, expression string) (*Job, error)
	AddDefaultJob(ctx context.Context, payload *DefaultJobPayload) (*Job, error)
	AddDelayedJob(ctx context.Context, payload *DelayedJobPayload, delay time.Duration) (*Job, error)
	RemoveDelayedJob(ctx context.Context, qName, jobID string) error
}

func (sc *SchedulerUseCase) AddPeriodicJob(ctx context.Context, payload *PeriodicJobPayload, expression string) (*Job, error) {
	return sc.repo.AddPeriodicJob(ctx, payload, expression)
}

func (sc *SchedulerUseCase) AddDefaultJob(ctx context.Context, payload *DefaultJobPayload) (*Job, error) {
	return sc.repo.AddDefaultJob(ctx, payload)
}

func (sc *SchedulerUseCase) AddDelayedJob(ctx context.Context, payload *DelayedJobPayload, delay time.Duration) (*Job, error) {
	return sc.repo.AddDelayedJob(ctx, payload, delay)
}

func (sc *SchedulerUseCase) RemoveDelayedJob(ctx context.Context, qName, jobID string) error {
	return sc.repo.RemoveDelayedJob(ctx, qName, jobID)
}

func (sc *SchedulerUseCase) CreateJob(ctx context.Context, job *Job) (*Job, error) {
	return sc.repo.CreateJob(ctx, job)
}

func (sc *SchedulerUseCase) TimingTaskJob(ctx context.Context) ([]*Job, error) {
	return sc.repo.ListTimingJob(ctx)
}
