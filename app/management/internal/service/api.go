package service

import (
	"bytes"
	"context"
	v1 "galileo/api/management/v1"
	"galileo/app/management/internal/biz"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
	"mime/multipart"
	"path"
)

func NewApi(req *v1.ApiInfo) (biz.Api, error) {
	if len(req.Name) <= 0 {
		return biz.Api{}, SetCustomizeErrMsg(ReasonParamsError, "name is required")
	}
	if len(req.Url) <= 0 {
		return biz.Api{}, SetCustomizeErrMsg(ReasonParamsError, "url is required")
	}
	if req.Type <= 0 {
		return biz.Api{}, SetCustomizeErrMsg(ReasonParamsError, "type is required")
	}
	return biz.Api{
		Name:        req.Name,
		Url:         req.Url,
		Module:      req.Module,
		Type:        int8(req.Type.Number()),
		Body:        req.Body,
		Label:       req.Label,
		QueryParams: req.QueryParams,
		Response:    req.Response,
		Description: req.Description,
	}, nil
}

func (s *ManagementService) CreateApi(ctx context.Context, req *v1.CreateApiRequest) (*v1.CreateApiReply, error) {
	newApi, err := NewApi(req.ApiInfo)
	if err != nil {
		return nil, err
	}
	if ret, _ := s.ac.ApiByName(ctx, newApi.Name); ret != nil {
		return nil, SetCustomizeErrInfoByReason(ReasonRecordExists)
	}
	// check if record already exists
	if ok, err := s.ac.IsApiDuplicated(ctx, newApi.Type, newApi.Url); err != nil {
		return nil, SetCustomizeErrInfo(err)
	} else if !ok {
		return nil, SetCustomizeErrInfoByReason(ReasonRecordExists)
	}
	userId := ctx.Value("x-md-global-userId").(uint32)
	newApi.CreatedBy = userId
	ret, err := s.ac.CreateApi(ctx, &newApi)
	if err != nil {
		return nil, err
	}
	return &v1.CreateApiReply{
		Id:        ret.ID,
		CreatedAt: ret.CreatedAt.Unix(),
	}, nil
}

func (s *ManagementService) UpdateApi(ctx context.Context, req *v1.UpdateApiRequest) (*v1.UpdateApiReply, error) {
	_, err := s.ac.ApiById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	updateApi, err := NewApi(req.ApiInfo)
	if err != nil {
		return nil, err
	}
	if _, err := s.ac.UpdateApi(ctx, &updateApi); err != nil {
		return nil, err
	}
	return &v1.UpdateApiReply{}, nil
}

// UploadApiFile Upload Api File get url
func (s *ManagementService) UploadApiFile(ctx http.Context) (err error) {
	fileName := ctx.Request().FormValue("fileName")
	file, fileHeader, _ := ctx.Request().FormFile("file")
	if fileName == "" {
		return SetCustomizeErrInfoByReason(ReasonFileMissing)
	}
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
	url, err := s.ac.UploadInterfaceFile(ctx, fileName, path.Ext(fileHeader.Filename), buf.Bytes())
	if err != nil {
		return SetCustomizeErrMsg(ReasonSystemError, err.Error())
	}
	return ctx.Result(20000, url)
}
