package service

import (
	"bytes"
	"context"
	pb "galileo/api/management/testcase/v1"
	"galileo/app/management/internal/biz"
	"galileo/pkg/ctxdata"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
	"mime/multipart"
	"path"
)

type TestcaseService struct {
	pb.UnimplementedTestcaseServer

	uc     *biz.TestcaseUseCase
	logger *log.Helper
}

func NewTestcaseService(uc *biz.TestcaseUseCase, logger log.Logger) *TestcaseService {
	return &TestcaseService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "management.testcaseService")),
	}
}

func NewTestcase(name string, caseType int8, priority int8, description string, url string) (biz.Testcase, error) {
	if len(name) <= 0 {
		return biz.Testcase{}, SetCustomizeErrMsg(ReasonParamsError, "testcase name must not be empty")
	}
	return biz.Testcase{
		Name:        name,
		Type:        caseType,
		Priority:    priority,
		Description: description,
		Url:         url,
	}, nil
}

func (s *TestcaseService) UploadTestcaseFile(ctx http.Context) (err error) {
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
	log.Debug("--------------------------------\n")
	url, err := s.uc.UploadTestcaseFile(ctx, fileName, path.Ext(fileHeader.Filename), buf.Bytes())
	if err != nil {
		return SetCustomizeErrMsg(ReasonSystemError, err.Error())
	}
	log.Debug("--------------------------------\n")
	return ctx.Result(20000, url)
}

func (s *TestcaseService) CreateTestcase(ctx context.Context, req *pb.CreateTestcaseRequest) (*pb.CreateTestcaseReply, error) {
	newTestcase, err := NewTestcase(req.Name, int8(req.Type), int8(req.Priority), req.Description, req.Url)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonParamsError, err.Error())
	}
	if ret, _ := s.uc.TestcaseByName(ctx, req.Name); ret != nil {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "duplicated testcase name")
	}
	uid := ctx.Value(ctxdata.UserIdKey).(uint32)
	newTestcase.CreatedBy = uid
	res, err := s.uc.CreateTestcase(ctx, &newTestcase)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonUnknownError, err.Error())
	}
	return &pb.CreateTestcaseReply{Id: res.Id, CreatedAt: res.CreatedAt.Unix()}, nil
}

func (s *TestcaseService) UpdateTestcase(ctx context.Context, req *pb.UpdateTestcaseRequest) (*pb.UpdateTestcaseReply, error) {
	updateTestcase, err := NewTestcase(req.Name, int8(req.Type), int8(req.Priority), req.Description, req.Url)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonParamsError, err.Error())
	}
	if ret, _ := s.uc.TestcaseByName(ctx, req.Name); ret != nil {
		return nil, SetCustomizeErrMsg(ReasonParamsError, "duplicated testcase name")
	}
	uid := ctx.Value(ctxdata.UserIdKey).(uint32)
	updateTestcase.Id = req.Id
	updateTestcase.UpdatedBy = uid
	ok, err := s.uc.UpdateTestcase(ctx, &updateTestcase)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonUnknownError, err.Error())
	}
	return &pb.UpdateTestcaseReply{
		Success: ok,
	}, nil
}
func (s *TestcaseService) DeleteTestcase(ctx context.Context, req *pb.DeleteTestcaseRequest) (*pb.DeleteTestcaseReply, error) {
	return &pb.DeleteTestcaseReply{}, nil
}

func (s *TestcaseService) GetTestcaseById(ctx context.Context, req *pb.GetTestcaseRequest) (*pb.GetTestcaseReply, error) {
	queryTestcase, err := s.uc.TestcaseById(ctx, req.Id)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonRecordNotFound, err.Error())
	}
	return &pb.GetTestcaseReply{
		Name:        queryTestcase.Name,
		CreatedBy:   queryTestcase.CreatedBy,
		CreatedAt:   queryTestcase.CreatedAt.Unix(),
		UpdateAt:    queryTestcase.UpdateAt.Unix(),
		UpdateBy:    queryTestcase.UpdatedBy,
		Type:        int32(queryTestcase.Type),
		Priority:    int32(queryTestcase.Priority),
		Status:      int32(queryTestcase.Status),
		Url:         queryTestcase.Url,
		Description: queryTestcase.Description,
	}, nil
}

func (s *TestcaseService) ListTestcase(ctx context.Context, req *pb.ListTestcaseRequest) (*pb.ListTestcaseReply, error) {
	return &pb.ListTestcaseReply{}, nil
}
