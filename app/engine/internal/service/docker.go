package service

import (
	"context"
	v1 "galileo/api/engine/v1"
	"galileo/app/engine/internal/biz"
	. "galileo/pkg/errResponse"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
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

func (s *EngineService) InspectContainer(ctx context.Context, req *v1.InspectContainerRequest) (*v1.InspectContainerReply, error) {
	ret, err := s.dc.InspectContainer(ctx, req.Id)
	switch {
	case client.IsErrNotFound(err):
		return nil, SetCustomizeErrMsg(ReasonRecordNotFound, err.Error())
	case err != nil:
		return nil, err
	}
	container := biz.ContainerMessageBody(ret)
	return &v1.InspectContainerReply{
		Container: container,
	}, nil
}

/* TODO: 容器化测试
1. 创建容器化的测试环境
	两次种创建方式:
		a. 基于docker-compose.yml Dockerfile 创建
		b. 简单创建
2. File Sharing 导入测试用例
3. 容器自动化测试执行
4. 容器监控、日志收集
5. 测试结果输出
*/

/* 解析docker-compose.yml
1. 获取上传的docker-compose.yml 文件（直接上传）
2. 使用yaml解析库解析docker-compose.yml文件到Config对象
*/

/* 解析Dockerfile
1. 获取上传的Dockerfile
2. yaml解析库解析
3. 输出解析后的指令
*/

func (s *EngineService) CreateContainer(ctx context.Context, req *v1.CreateContainerRequest) (*v1.CreateContainerReply, error) {
	container, err := s.uc.CreateContainer(ctx, &biz.Container{})
	if err != nil {
		return nil, err
	}
	return &v1.CreateContainerReply{
		Id:      container.Id,
		Warning: "",
	}, nil
}
