package data

import (
	"context"
	"galileo/app/engine/internal/biz"
	"github.com/docker/docker/api/types"
	"github.com/go-kratos/kratos/v2/log"
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
