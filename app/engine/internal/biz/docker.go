package biz

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/go-kratos/kratos/v2/log"
)

type Container struct {
	Id              string
	Hostname        string
	Domainname      string
	User            string
	Name            string
	AttachStdin     bool
	AttachStdout    bool
	AttachStderr    bool
	Tty             bool
	OpenStdin       bool
	StdinOnce       bool
	Env             []string
	Cmd             []string
	Image           string
	Labels          map[string]string
	Volumes         map[string]struct{}
	WorkingDir      string
	Entrypoint      string
	NetworkDisabled bool
	MacAddress      string
	ExposedPorts    map[string]struct{}
	StopSignal      string
	StopTimeout     *int
	HostConfig      map[string]interface{}
	NetworkConfig   map[string]struct{}
}

type DockerRepo interface {
	ListContainers(ctx context.Context, options types.ContainerListOptions) ([]*Container, error)
}

type DockerUseCase struct {
	repo DockerRepo
	log  *log.Helper
}

func NewDockerUseCase(repo DockerRepo, logger log.Logger) *DockerUseCase {
	helper := log.NewHelper(log.With(logger, "module", "docker.useCase"))
	return &DockerUseCase{repo: repo, log: helper}
}

func (uc *DockerUseCase) ListContainers(ctx context.Context, options types.ContainerListOptions) ([]*Container, error) {
	return uc.repo.ListContainers(ctx, options)
}
