package service

import (
	"context"
	v1 "galileo/api/engine/v1"
	"galileo/app/engine/internal/biz"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
)

type DockerService struct {
	v1.UnimplementedDockerServer

	uc  *biz.DockerUseCase
	log *log.Helper
}

func NewDockerService(uc *biz.DockerUseCase, logger log.Logger) *DockerService {
	return &DockerService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "docker.Service")),
	}
}

func (s *DockerService) ListContainers(ctx context.Context, req *v1.ListContainerRequest) (*v1.ListContainersReply, error) {
	/* queryParams 中的filters 类型为 map[string][]string */
	/* 构建types.ContainerListOptions */
	var filterArgs []filters.KeyValuePair
	for k, stringList := range req.Filters {
		lo.ForEach(stringList.Values, func(v string, _ int) {
			filterArgs = append(filterArgs, filters.KeyValuePair{
				Key:   k,
				Value: v,
			})
		})
	}
	args := filters.NewArgs(filterArgs...)
	_, err := s.uc.ListContainers(ctx, types.ContainerListOptions{
		All:     req.All,
		Size:    req.Size,
		Limit:   int(req.Limit),
		Filters: args,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *DockerService) CreateContainer(ctx context.Context) {}
