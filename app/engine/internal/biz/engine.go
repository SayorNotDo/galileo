package biz

import (
	"context"
	taskV1 "galileo/api/management/task/v1"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
	"io"
	"os"
	"time"
)

type Task struct {
	Id             int64
	Name           string
	Rank           int8
	Type           int8
	Status         taskV1.TaskStatus
	Worker         string
	Config         string
	Frequency      taskV1.Frequency
	ScheduleTime   time.Time
	TestcaseSuites []int64
	ExecuteId      int64
}

type Container struct {
	Id    string
	Name  string
	Image string
}

type EngineRepo interface {
	TaskByID(ctx context.Context, id int64) (*Task, error)
	AddCronJob(ctx context.Context, task *Task) (cron.EntryID, error)
	TimingTaskList(ctx context.Context) ([]*Task, error)
	GetCronJobList(ctx context.Context) []*CronJob
	RemoveCronJob(ctx context.Context, taskId int64) error
}

type EngineUseCase struct {
	Repo EngineRepo
	log  *log.Helper
}

func NewEngineUseCase(repo EngineRepo, logger log.Logger) *EngineUseCase {
	helper := log.NewHelper(log.With(logger, "module", "engine.useCase"))
	return &EngineUseCase{Repo: repo, log: helper}
}

func (c *EngineUseCase) TaskByID(ctx context.Context, id int64) (*Task, error) {
	return c.Repo.TaskByID(ctx, id)
}

func (c *EngineUseCase) AddCronJob(ctx context.Context, task *Task) (cron.EntryID, error) {
	return c.Repo.AddCronJob(ctx, task)
}

func (c *EngineUseCase) BuildContainer(context.Context) (*Container, error) {
	// TODO: 容器构建过程
	ctx := context.Background()
	log.Debug("Building container------------------------")
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		return nil, err
	}
	_, _ = io.Copy(os.Stdout, reader)
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
	}, nil, nil, nil, "")
	if err != nil {
		return nil, err
	}
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return nil, err
	}
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return nil, err
	}
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	return &Container{Id: resp.ID}, nil
}

func (c *EngineUseCase) TimingTaskList(ctx context.Context) ([]*Task, error) {
	return c.Repo.TimingTaskList(ctx)
}

func (c *EngineUseCase) GetCronJobList(ctx context.Context) []*CronJob {
	return c.Repo.GetCronJobList(ctx)
}

func (c *EngineUseCase) RemoveCronJob(ctx context.Context, taskId int64) error {
	return c.Repo.RemoveCronJob(ctx, taskId)
}
