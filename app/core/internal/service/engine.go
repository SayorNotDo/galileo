package service

import (
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/transport/http"
	"mime/multipart"
)

func (c *EngineService) UploadEngineFile(ctx http.Context) error {
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
	//buf := bytes.NewBuffer(nil)
	return nil
}
