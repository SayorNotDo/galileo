package service

import (
	"context"
	v1 "galileo/api/engine/v1"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/samber/lo"
)

func (s *EngineService) ListContainers(ctx context.Context, req *v1.ListContainerRequest) (*v1.ListContainersReply, error) {
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
	_, err := s.dc.ListContainers(ctx, types.ContainerListOptions{
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

func (s *EngineService) CreateContainer(ctx context.Context) {}
