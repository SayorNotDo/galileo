package biz

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/go-kratos/kratos/v2/log"
)

type Container struct {
	Id              string              `json:"id"`
	Hostname        string              `json:"hostname"`
	Domainname      string              `json:"domain_name"`
	User            string              `json:"user"`
	Name            string              `json:"name"`
	AttachStdin     bool                `json:"attach_stdin,omitempty"`
	AttachStdout    bool                `json:"attach_stdout,omitempty"`
	AttachStderr    bool                `json:"attach_stderr,omitempty"`
	Tty             bool                `json:"tty,omitempty"`
	OpenStdin       bool                `json:"open_stdin,omitempty"`
	StdinOnce       bool                `json:"stdin_once,omitempty"`
	Env             []string            `json:"env"`
	Cmd             []string            `json:"cmd"`
	Image           string              `json:"image"`
	Labels          map[string]string   `json:"labels"`
	Volumes         map[string]struct{} `json:"volumes"`
	WorkingDir      string              `json:"working_dir"`
	Entrypoint      string              `json:"entrypoint"`
	NetworkDisabled bool                `json:"network_disabled"`
	MacAddress      string              `json:"mac_address"`
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
