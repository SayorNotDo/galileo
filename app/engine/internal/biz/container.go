package biz

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type Container struct {
	Id              string              `json:"id"`
	Created         string              `json:"created"`
	Hostname        string              `json:"hostname"`
	Domainname      string              `json:"domain_name"`
	User            string              `json:"user"`
	Name            string              `json:"name"`
	AttachStdin     bool                `json:"attach_stdin,omitempty"`
	AttachStdout    bool                `json:"attach_stdout,omitempty"`
	AttachStderr    bool                `json:"attach_stderr,omitempty"`
	Tty             bool                `json:"tty,omitempty"`
	RestartCount    int                 `json:"restart_count,omitempty"`
	OpenStdin       bool                `json:"open_stdin,omitempty"`
	StdinOnce       bool                `json:"stdin_once,omitempty"`
	Env             []string            `json:"env"`
	Cmd             []string            `json:"cmd"`
	Image           string              `json:"image"`
	Labels          map[string]string   `json:"labels"`
	Volumes         map[string]struct{} `json:"volumes"`
	WorkingDir      string              `json:"working_dir"`
	Entrypoint      []string            `json:"entrypoint"`
	NetworkDisabled bool                `json:"network_disabled"`
	MacAddress      string              `json:"mac_address"`
	ExposedPorts    nat.PortSet         `json:"exposed_ports"`
	StopSignal      string              `json:"stop_signal"`
	StopTimeout     *int                `json:"stop_timeout"`
	Driver          string              `json:"driver"`
	Platform        string              `json:"platform"`
	SizeRw          *int64              `json:"size_rw,omitempty"`
	SizeRootFs      *int64              `json:"size_root_fs,omitempty"`
}

func NewContainer(data types.ContainerJSON) *Container {
	return &Container{
		Id:              data.ID,
		Created:         data.Created,
		Hostname:        data.Config.Hostname,
		Domainname:      data.Config.Domainname,
		User:            data.Config.User,
		Name:            data.Name,
		AttachStdin:     data.Config.AttachStdin,
		AttachStdout:    data.Config.AttachStdout,
		AttachStderr:    data.Config.AttachStderr,
		Tty:             data.Config.Tty,
		RestartCount:    data.RestartCount,
		OpenStdin:       data.Config.OpenStdin,
		StdinOnce:       data.Config.StdinOnce,
		Env:             data.Config.Env,
		Cmd:             data.Config.Cmd,
		Image:           data.Config.Image,
		Labels:          data.Config.Labels,
		Volumes:         data.Config.Volumes,
		WorkingDir:      data.Config.WorkingDir,
		Entrypoint:      data.Config.Entrypoint,
		NetworkDisabled: data.Config.NetworkDisabled,
		MacAddress:      data.Config.MacAddress,
		ExposedPorts:    data.Config.ExposedPorts,
		StopSignal:      data.Config.StopSignal,
		StopTimeout:     data.Config.StopTimeout,
		Driver:          data.Driver,
		Platform:        data.Platform,
		SizeRw:          data.SizeRw,
		SizeRootFs:      data.SizeRootFs,
	}
}

type ContainerConfig struct {
	Config        *container.Config
	HostConfig    *container.HostConfig
	NetworkConfig *network.NetworkingConfig
	Platform      *v1.Platform
	containerName string
}
