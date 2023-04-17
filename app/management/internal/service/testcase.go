package service

import (
	"context"
	"errors"
	pb "galileo/api/management/testcase/v1"
	"galileo/app/management/internal/biz"
	"galileo/pkg/ctxdata"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
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

func NewTestcase(name string, caseType int16, priority int8, description string, url string) (biz.Testcase, error) {
	if len(name) <= 0 {
		return biz.Testcase{}, errors.New("testcase name must not be empty")
	}
	return biz.Testcase{
		Name:        name,
		CaseType:    caseType,
		Priority:    priority,
		Description: description,
		Url:         url,
	}, nil
}

func (s *TestcaseService) CreateTestcase(ctx context.Context, req *pb.CreateTestcaseRequest) (*pb.CreateTestcaseReply, error) {
	newTestcase, err := NewTestcase(req.Name, int16(req.Type), int8(req.Priority), req.Description, req.Url)
	if err != nil {
		return nil, SetCustomizeErrMsg("PARAMS_ERROR", err.Error())
	}
	if ret, _ := s.uc.TestcaseByName(ctx, req.Name); ret != nil {
		return nil, SetCustomizeErrMsg("PARAMS_ERROR", "duplicated testcase name")
	}
	uid := ctx.Value(ctxdata.UserIdKey).(uint32)
	newTestcase.CreatedBy = uid
	res, err := s.uc.CreateTestcase(ctx, &newTestcase)
	if err != nil {
		return nil, SetCustomizeErrMsg("UNKNOWN_ERROR", err.Error())
	}
	return &pb.CreateTestcaseReply{Id: res.Id, CreatedAt: res.CreatedAt.Unix()}, nil
}

func (s *TestcaseService) UpdateTestcase(ctx context.Context, req *pb.UpdateTestcaseRequest) (*pb.UpdateTestcaseReply, error) {
	updateTestcase, err := NewTestcase(req.Name, int16(req.Type), int8(req.Priority), req.Description, req.Url)
	if err != nil {
		return nil, SetCustomizeErrMsg("PARAMS_ERROR", err.Error())
	}
	if ret, _ := s.uc.TestcaseByName(ctx, req.Name); ret != nil {
		return nil, SetCustomizeErrMsg("PARAMS_ERROR", "duplicated testcase name")
	}
	uid := ctx.Value(ctxdata.UserIdKey).(uint32)
	updateTestcase.Id = req.Id
	updateTestcase.UpdatedBy = uid
	ok, err := s.uc.UpdateTestcase(ctx, &updateTestcase)
	if err != nil {
		return nil, SetCustomizeErrMsg("UNKNOWN_ERROR", err.Error())
	}
	return &pb.UpdateTestcaseReply{
		Success: ok,
	}, nil
}
func (s *TestcaseService) DeleteTestcase(ctx context.Context, req *pb.DeleteTestcaseRequest) (*pb.DeleteTestcaseReply, error) {
	return &pb.DeleteTestcaseReply{}, nil
}
func (s *TestcaseService) GetTestcaseById(ctx context.Context, req *pb.GetTestcaseRequest) (*pb.GetTestcaseReply, error) {

	return &pb.GetTestcaseReply{}, nil
}

func (s *TestcaseService) ListTestcase(ctx context.Context, req *pb.ListTestcaseRequest) (*pb.ListTestcaseReply, error) {
	return &pb.ListTestcaseReply{}, nil
}
