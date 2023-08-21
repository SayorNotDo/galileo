package service

import (
	"bytes"
	"context"
	v1 "galileo/api/management/v1"
	"galileo/app/management/internal/biz"
	"galileo/pkg/ctxdata"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"mime/multipart"
	"path"
)

func NewTestcase(name string, caseType int8, priority int8, description string, url string) (biz.Testcase, error) {
	if len(name) <= 0 {
		return biz.Testcase{}, SetCustomizeErrMsg(ReasonParamsError, "testcase name cannot be empty")
	} else if caseType < 0 {
		return biz.Testcase{}, SetCustomizeErrMsg(ReasonParamsError, "illegal testcase type")
	} else if priority < 0 {
		return biz.Testcase{}, SetCustomizeErrMsg(ReasonParamsError, "illegal testcase priority")
	}
	return biz.Testcase{
		Name:        name,
		Type:        caseType,
		Priority:    priority,
		Description: description,
		Url:         url,
	}, nil
}

// UploadTestcaseFile Upload Testcase file get url
func (s *ManagementService) UploadTestcaseFile(ctx http.Context) (err error) {
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
	if err := s.tc.TestcaseValidator(path.Ext(fileHeader.Filename), buf.String()); err != nil {
		return err
	}
	url, err := s.tc.UploadTestcaseFile(ctx, fileName, path.Ext(fileHeader.Filename), buf.Bytes())
	if err != nil {
		return SetCustomizeErrMsg(ReasonSystemError, err.Error())
	}
	return ctx.Result(20000, map[string]string{
		"url": url,
	})
}

// CreateTestcase creates a new Testcase, returns id, create time
func (s *ManagementService) CreateTestcase(ctx context.Context, req *v1.CreateTestcaseRequest) (*v1.CreateTestcaseReply, error) {
	newTestcase, err := NewTestcase(req.Name, int8(req.Type), int8(req.Priority), req.Description, req.Url)
	if err != nil {
		return nil, err
	}
	if ret, _ := s.tc.TestcaseByName(ctx, req.Name); ret != nil {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "duplicated testcase name")
	}
	uid := ctx.Value(ctxdata.UserIdKey).(uint32)
	newTestcase.CreatedBy = uid
	res, err := s.tc.CreateTestcase(ctx, &newTestcase)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonUnknownError, err.Error())
	}
	return &v1.CreateTestcaseReply{Id: res.Id, CreatedAt: timestamppb.New(res.CreatedAt)}, nil
}

func (s *ManagementService) UpdateTestcase(ctx context.Context, req *v1.UpdateTestcaseRequest) (*v1.UpdateTestcaseReply, error) {
	if _, err := s.tc.TestcaseById(ctx, req.Id); err != nil {
		return nil, SetCustomizeErrMsg(ReasonRecordNotFound, err.Error())
	}
	updateTestcase, err := NewTestcase(req.Name, int8(req.Type), int8(req.Priority), req.Description, req.Url)
	if err != nil {
		return nil, err
	}
	if ret, _ := s.tc.TestcaseByName(ctx, req.Name); ret != nil {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "duplicated testcase name")
	}
	uid := ctx.Value(ctxdata.UserIdKey).(uint32)
	updateTestcase.Id = req.Id
	updateTestcase.UpdatedBy = uid
	ok, err := s.tc.UpdateTestcase(ctx, &updateTestcase)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonUnknownError, err.Error())
	}
	return &v1.UpdateTestcaseReply{
		Success: ok,
	}, nil
}

func (s *ManagementService) DeleteTestcase(ctx context.Context, req *v1.DeleteTestcaseRequest) (*v1.DeleteTestcaseReply, error) {
	return &v1.DeleteTestcaseReply{}, nil
}

func (s *ManagementService) GetTestcaseById(ctx context.Context, req *v1.GetTestcaseRequest) (*v1.GetTestcaseReply, error) {
	queryTestcase, err := s.tc.TestcaseById(ctx, req.Id)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonRecordNotFound, err.Error())
	}
	return &v1.GetTestcaseReply{
		TestcaseInfo: &v1.TestcaseInfo{
			Name:        queryTestcase.Name,
			CreatedBy:   queryTestcase.CreatedBy,
			CreatedAt:   timestamppb.New(queryTestcase.CreatedAt),
			UpdatedAt:   timestamppb.New(queryTestcase.UpdatedAt),
			UpdatedBy:   queryTestcase.UpdatedBy,
			Type:        int32(queryTestcase.Type),
			Priority:    int32(queryTestcase.Priority),
			Status:      int32(queryTestcase.Status),
			Url:         queryTestcase.Url,
			Description: queryTestcase.Description,
		},
	}, nil
}

func (s *ManagementService) ListTestcase(ctx context.Context, req *v1.ListTestcaseRequest) (*v1.ListTestcaseReply, error) {
	return &v1.ListTestcaseReply{}, nil
}

func (s *ManagementService) LoadFramework(ctx context.Context, req *v1.LoadFrameworkRequest) (*v1.LoadFrameworkReply, error) {
	fpath, lang, config := req.Path, req.Lang, req.Config
	log.Debugf("LoadFrameWork request path: %s, lang: %d, config: %s", fpath, lang, string(config))
	// run docker return container-id
	// initialize container based on request param: config & language
	return &v1.LoadFrameworkReply{Success: true, Worker: "docker-container-id"}, nil
}

func (s *ManagementService) CreateTestcaseSuite(ctx context.Context, req *v1.CreateTestcaseSuiteRequest) (*v1.CreateTestcaseSuiteReply, error) {
	if len(req.Name) <= 0 {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "suite name cannot be empty")
	}
	res, err := s.tc.CreateTestcaseSuite(ctx, req.Name, req.TestcaseList)
	if err != nil {
		return nil, err
	}
	return &v1.CreateTestcaseSuiteReply{
		Id:        res.Id,
		CreatedAt: timestamppb.New(res.CreatedAt),
	}, nil
}

func (s *ManagementService) GetTestcaseSuiteById(ctx context.Context, req *v1.GetTestcaseSuiteRequest) (*v1.GetTestcaseSuiteReply, error) {
	ret, err := s.tc.GetTestcaseSuiteById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	rv := make([]*v1.TestcaseInfo, 0)
	lo.ForEach(ret.TestcaseList, func(item *biz.Testcase, _ int) {
		rv = append(rv, &v1.TestcaseInfo{
			Id:          item.Id,
			Name:        item.Name,
			CreatedAt:   timestamppb.New(item.CreatedAt),
			CreatedBy:   item.CreatedBy,
			UpdatedAt:   timestamppb.New(ret.UpdatedAt),
			UpdatedBy:   item.UpdatedBy,
			Type:        int32(item.Type),
			Priority:    int32(item.Priority),
			Status:      int32(item.Status),
			Url:         item.Url,
			Description: item.Description,
		})
	})
	return &v1.GetTestcaseSuiteReply{
		Id:           ret.Id,
		Name:         ret.Name,
		CreatedAt:    timestamppb.New(ret.CreatedAt),
		CreatedBy:    ret.CreatedBy,
		UpdatedAt:    timestamppb.New(ret.UpdatedAt),
		UpdatedBy:    ret.UpdatedBy,
		TestcaseInfo: rv,
	}, nil
}

func (s *ManagementService) DebugTestcase(ctx context.Context, req *v1.DebugTestcaseRequest) (*v1.DebugTestcaseReply, error) {
	/* 1.调用TestcaseValidator */
	if err := s.tc.TestcaseValidator(".json", req.Content); err != nil {
		return nil, err
	}
	/* 2.MOCK环境运行测试用例 */
	err, ret := s.tc.DebugTestcase(ctx, req.Content)
	if err != nil {
		return nil, err
	}
	/* 3.返回调试结果 */
	return &v1.DebugTestcaseReply{
		Result: ret,
	}, nil
}
