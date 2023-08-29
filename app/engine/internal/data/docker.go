package data

import (
	"context"
	"fmt"
	"galileo/app/engine/internal/biz"
	"galileo/pkg/errResponse"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/volume"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"gopkg.in/yaml.v2"
)

type dockerRepo struct {
	data *Data
	log  *log.Helper
}

type ComposeConfig struct {
	Version  string                              `yaml:"version"`
	Services map[string]container.Config         `yaml:"services"`
	Networks map[string]network.NetworkingConfig `yaml:"networks"`
	Volumes  map[string]volume.Volume            `yaml:"volumes"`
}

func NewDockerRepo(data *Data, logger log.Logger) biz.DockerRepo {
	return &dockerRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "docker.Repo")),
	}
}

func (repo *dockerRepo) GetDockerInfo(ctx context.Context) {
	info, err := repo.data.dockerCli.Info(ctx)
	if err != nil {
		return
	}
	repo.log.Info(info)
}

func (repo *dockerRepo) ListContainers(ctx context.Context, options types.ContainerListOptions) ([]*biz.Container, error) {
	_, err := repo.data.dockerCli.ContainerList(ctx, options)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (repo *dockerRepo) CreateContainer(ctx context.Context, c *biz.Container) (id string, warnings []string, err error) {
	/* 新建容器 */
	/* 参数：容器配置、HOST配置、网络配置、平台、容器名 */
	/* 载入 docker-compose.yml 文件 */
	res, err := repo.data.dockerCli.ContainerCreate(ctx,
		&container.Config{},
		&container.HostConfig{},
		&network.NetworkingConfig{},
		&v1.Platform{},
		"containerName")
	if err != nil {
		return
	}
	/* TODO: 容器信息入库管理 */
	/* 获取容器信息 */
	containerInfo, err := repo.InspectContainer(ctx, id)
	if err != nil {
		return "", nil, err
	}
	/* 容器记录入库 */
	/* 数据类型转换 */
	var labels []string
	for k := range containerInfo.Labels {
		labels = append(labels, fmt.Sprintf("%s=%s", k, containerInfo.Labels[k]))
	}
	var volumes []string
	for v := range containerInfo.Volumes {
		volumes = append(volumes, v)
	}
	var exposedPorts []string
	for port := range containerInfo.ExposedPorts {
		exposedPorts = append(exposedPorts, fmt.Sprintf("%s/%s", port.Port(), port.Proto()))
	}
	/* 获取 docker-compose.yml 和 Dockerfile 的 Oss url */
	if _, err := repo.data.entDB.Container.Create().
		SetID(containerInfo.Id).
		SetHostname(containerInfo.Hostname).
		SetDomainname(containerInfo.Domainname).
		SetUser(containerInfo.User).
		SetEnv(containerInfo.Env).
		SetCmd(containerInfo.Cmd).
		SetImage(containerInfo.Image).
		SetLabels(labels).
		SetVolumes(volumes).
		SetWorkingDir(containerInfo.WorkingDir).
		SetEntrypoint(containerInfo.Entrypoint).
		SetMACAddress(containerInfo.MacAddress).
		SetExposePorts(exposedPorts).
		Save(ctx); err != nil {
		return "", nil, err
	}
	return res.ID, res.Warnings, nil
}

func (repo *dockerRepo) InspectContainer(ctx context.Context, id string) (*biz.Container, error) {
	res, err := repo.data.dockerCli.ContainerInspect(ctx, id)
	if err != nil {
		return nil, err
	}
	return biz.NewContainer(res), nil
}

func (repo *dockerRepo) OssGetFile(url string) ([]byte, error) {
	return nil, nil
}

// ParseComposeFile parses a compose file and returns []*biz.ContainerConfig
func (repo *dockerRepo) ParseComposeFile(ctx context.Context, fp []byte) ([]*biz.ContainerConfig, error) {
	/* yaml解析库 gopkg.in/yaml.v2 */
	/* 单个文件存在复数service */
	var composeConfig ComposeConfig
	err := yaml.Unmarshal(fp, &composeConfig)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errResponse.SetCustomizeErrMsg(errResponse.ReasonUnknownError, err.Error())
	}
	fmt.Println("----------------------------------------------------------------here")
	fmt.Printf("version: %v\nservices: %v\nnetworks: %v\nvolumes:%v\n", composeConfig.Version, composeConfig.Services, composeConfig.Networks, composeConfig.Volumes)
	for k, v := range composeConfig.Services {
		fmt.Printf("%v\n", k)
		fmt.Printf("%v\n", v)
	}
	fmt.Println("----------------------------------------------------------------end")
	return nil, nil
}
