package service

import (
	"context"
	"fmt"
	pb "galileo/api/management/testcase/v1"
	"galileo/app/management/internal/biz"
	"galileo/pkg/ctxdata"
	"galileo/pkg/errResponse"
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

func (s *TestcaseService) CreateTestcase(ctx context.Context, req *pb.CreateTestcaseRequest) (*pb.CreateTestcaseReply, error) {
	uid := ctx.Value(ctxdata.UserIdKey).(uint32)
	newTestcase := &biz.Testcase{
		Name:        req.Name,
		CreatedBy:   uid,
		CaseType:    int16(req.Type),
		Priority:    int8(req.Priority),
		Description: req.Description,
		Url:         req.Url,
	}
	res, err := s.uc.CreateTestcase(ctx, newTestcase)
	if err != nil {
		println("--------------------------------")
		fmt.Printf("%v", errResponse.SetCustomizeErrMsg("ReasonParamsError", err.Error()))
		println("--------------------------------")
		return nil, errResponse.SetCustomizeErrMsg("PARAMS_ERROR", err.Error())
	}
	return &pb.CreateTestcaseReply{Id: res.Id, CreatedAt: res.CreatedAt.Unix()}, nil
}
func (s *TestcaseService) UpdateTestcase(ctx context.Context, req *pb.UpdateTestcaseRequest) (*pb.UpdateTestcaseReply, error) {

	return &pb.UpdateTestcaseReply{}, nil
}
func (s *TestcaseService) DeleteTestcase(ctx context.Context, req *pb.DeleteTestcaseRequest) (*pb.DeleteTestcaseReply, error) {
	return &pb.DeleteTestcaseReply{}, nil
}
func (s *TestcaseService) GetTestcase(ctx context.Context, req *pb.GetTestcaseRequest) (*pb.GetTestcaseReply, error) {
	return &pb.GetTestcaseReply{}, nil
}
func (s *TestcaseService) ListTestcase(ctx context.Context, req *pb.ListTestcaseRequest) (*pb.ListTestcaseReply, error) {
	return &pb.ListTestcaseReply{}, nil
}
