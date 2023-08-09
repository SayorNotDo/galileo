package biz

import "github.com/docker/go-connections/nat"

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
