package biz

import (
	"context"
	v1 "galileo/api/engine/v1"
	"github.com/docker/docker/api/types"
)

type DockerRepo interface {
	ListContainers(ctx context.Context, options types.ContainerListOptions) ([]*Container, error)
	InspectContainer(ctx context.Context, id string) (*Container, error)
	CreateContainer(ctx context.Context, container *Container) (id string, warnings []string, err error)
	ParseComposeFile(ctx context.Context, fp []byte) ([]*ContainerConfig, error)
	GetDockerInfo(ctx context.Context)
}

func (dc *DockerUseCase) ListContainers(ctx context.Context, options types.ContainerListOptions) ([]*Container, error) {
	return dc.repo.ListContainers(ctx, options)
}

func (dc *DockerUseCase) ParseComposeFile(ctx context.Context, fp []byte) ([]*ContainerConfig, error) {
	return dc.repo.ParseComposeFile(ctx, fp)
}

func (dc *DockerUseCase) CreateContainer(ctx context.Context, container *Container) (id string, warnings []string, err error) {
	return dc.repo.CreateContainer(ctx, container)

}

func (dc *DockerUseCase) InspectContainer(ctx context.Context, id string) (*Container, error) {
	return dc.repo.InspectContainer(ctx, id)
}

func ContainerMessageBody(c *Container) *v1.Container {
	var portsMap []*v1.Container_ExposedPorts
	for port, _ := range c.ExposedPorts {
		containerPort := &v1.Container_ExposedPorts{
			Port: port.Port(),
			Type: port.Proto(),
		}
		portsMap = append(portsMap, containerPort)
	}
	return &v1.Container{
		Id:           c.Id,
		Name:         c.Name,
		Created:      c.Created,
		Hostname:     c.Hostname,
		Domainname:   c.Domainname,
		Image:        c.Image,
		Cmd:          c.Cmd,
		ExposedPorts: portsMap,
	}
}

func (dc *DockerUseCase) ParseDockerfile(ctx context.Context, fp string) (map[string]interface{}, error) {
	return nil, nil
}
