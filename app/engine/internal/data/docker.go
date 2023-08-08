package data

import (
	"context"
	"fmt"
	"galileo/app/engine/internal/biz"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type dockerRepo struct {
	data *Data
	log  *log.Helper
}

func NewDockerRepo(data *Data, logger log.Logger) biz.DockerRepo {
	return &dockerRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "docker.Repo")),
	}
}

func (r *dockerRepo) ListContainers(ctx context.Context, options types.ContainerListOptions) ([]*biz.Container, error) {
	_, err := r.data.dockerCli.ContainerList(ctx, options)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *dockerRepo) CreateContainers(ctx context.Context) {

	_, err := r.data.dockerCli.ContainerCreate(ctx,
		&container.Config{},
		&container.HostConfig{},
		&network.NetworkingConfig{},
		&v1.Platform{},
		"containerName")
	if err != nil {
		return
	}
}

func (r *dockerRepo) InspectContainer(ctx context.Context, id string) (err error) {
	res, err := r.data.dockerCli.ContainerInspect(ctx, id)
	if err != nil {
		return err
	}
	containerJsonBase := res.ContainerJSONBase
	mount := res.Mounts
	config := res.Config
	networkSettings := res.NetworkSettings
	fmt.Printf("containerJsonbase: %v\n mount: %v\n config: %v\n networkSetting: %v\n", containerJsonBase, mount, config, networkSettings)
	return
}
