package biz

import (
	"context"
	taskV1 "galileo/api/management/task/v1"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/go-kratos/kratos/v2/log"
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
	UpdateTaskStatus(ctx context.Context, id int64, status taskV1.TaskStatus) (bool, error)
	TaskByID(ctx context.Context, id int64) (*Task, error)
	AddCronJob(ctx context.Context)
}

type EngineUseCase struct {
	Repo EngineRepo
	log  *log.Helper
}

func NewEngineUseCase(repo EngineRepo, logger log.Logger) *EngineUseCase {
	helper := log.NewHelper(log.With(logger, "module", "engine.useCase"))
	return &EngineUseCase{Repo: repo, log: helper}
}

func (c *EngineUseCase) UpdateTaskStatus(ctx context.Context, id int64, status taskV1.TaskStatus) (bool, error) {
	return c.Repo.UpdateTaskStatus(ctx, id, status)
}

func (c *EngineUseCase) TaskByID(ctx context.Context, id int64) (*Task, error) {
	return c.Repo.TaskByID(ctx, id)
}

func (c *EngineUseCase) AddCronJob(ctx context.Context) {
	c.Repo.AddCronJob(ctx)
}

func (c *EngineUseCase) BuildContainer(context.Context) (*Container, error) {
	// TODO: container building process
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
	io.Copy(os.Stdout, reader)
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
