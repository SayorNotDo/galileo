package service

import (
	"bytes"
	"context"
	v1 "galileo/api/core/v1"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
	"mime/multipart"
	"path"
)

func (c *CoreService) InspectContainer(ctx context.Context, req *v1.InspectContainerRequest) (*v1.ContainerInfo, error) {
	ret, err := c.ec.InspectContainer(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ContainerInfo{
		Id:     ret.Id,
		Name:   ret.Name,
		Image:  ret.Image,
		Labels: ret.Labels,
	}, err
}

/*
接口上传的docker-compose\Dockerfile文件   对象存储
*/

func (c *CoreService) UploadEngineFile(ctx http.Context) error {
	/* 接口参数获取（文件类型） */
	fileName := ctx.Request().FormValue("fileName")
	if fileName == "" {
		return SetCustomizeErrInfoByReason(ReasonFileNameMissing)
	}
	file, fileHeader, _ := ctx.Request().FormFile("file")
	if file == nil {
		return SetCustomizeErrInfoByReason(ReasonFileMissing)
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	if fileHeader.Size > 1024*1024*5 {
		return SetCustomizeErrInfoByReason(ReasonFileOverLimitSize)
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}
	/* 基于文件名进行相应操作
	无需校验文件正确性，运行时决定
	*/
	switch fileName {
	case "Dockerfile", "docker-compose":
		url, err := c.ec.UploadEngineFile(ctx, fileName, path.Ext(fileHeader.Filename), buf.Bytes())
		if err != nil {
			return err
		}
		return ctx.Result(2000, map[string]string{
			"url": url,
		})
	default:
		return SetCustomizeErrInfoByReason(ReasonParamsError)
	}
}
