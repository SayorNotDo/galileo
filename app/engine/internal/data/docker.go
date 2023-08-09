package data

import (
	"context"
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

func (r *dockerRepo) CreateContainer(ctx context.Context, c *biz.Container) (id string, warnings []string, err error) {
	/* 新建容器 */
	/* 参数：容器配置、HOST配置、网络配置、平台、容器名 */
	/* 载入 docker-compose.yml 文件 */
	res, err := r.data.dockerCli.ContainerCreate(ctx,
		&container.Config{},
		&container.HostConfig{},
		&network.NetworkingConfig{},
		&v1.Platform{},
		"containerName")
	if err != nil {
		return
	}
	return res.ID, res.Warnings, nil
}

func (r *dockerRepo) InspectContainer(ctx context.Context, id string) (*biz.Container, error) {
	res, err := r.data.dockerCli.ContainerInspect(ctx, id)
	if err != nil {
		return nil, err
	}
	return ReContainer(res), nil
}

func ReContainer(data types.ContainerJSON) *biz.Container {
	return &biz.Container{
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

func (r *dockerRepo) ParseComposeFile(ctx context.Context, fp string) (map[string]container.Config, error) {
	/* 读取docker-compose文件 */
	//file, err := os.ReadFile(fp)
	//if err != nil {
	//	return nil, err
	//}
	//defer func(file *os.File) {
	//	err := file.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(file)
	//var config container.Config
	///* 使用yaml解析库解析docker-compose.yml文件到Config对象 */
	//decoder := yaml.NewDecoder(file)
	//if err := decoder.Decode(&config); err != nil {
	//	return nil, err
	//}

	return nil, nil
}
