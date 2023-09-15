package biz

import (
	managementV1 "galileo/api/management/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"time"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewDockerUseCase,
	NewEngineUseCase,
)

type Task struct {
	Id             int64
	Name           string
	Rank           int8
	Type           int8
	Status         managementV1.TaskStatus
	Worker         string
	Config         string
	Frequency      int8
	ScheduleTime   time.Time
	TestcaseSuites []int32
	ExecuteId      int64
}

type SchedulerUseCase struct {
	repo SchedulerRepo
	log  *log.Helper
}

func NewEngineUseCase(repo SchedulerRepo, logger log.Logger) *SchedulerUseCase {
	helper := log.NewHelper(log.With(logger, "module", "engine.useCase"))
	return &SchedulerUseCase{
		repo: repo,
		log:  helper,
	}
}

type DockerUseCase struct {
	repo DockerRepo
	log  *log.Helper
}

func NewDockerUseCase(repo DockerRepo, logger log.Logger) *DockerUseCase {
	helper := log.NewHelper(log.With(logger, "module", "docker.useCase"))
	return &DockerUseCase{
		repo: repo,
		log:  helper,
	}
}
