package service

import (
	"context"
	"galileo/app/management/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "galileo/api/management/testcase/v1"
)

type TestcaseService struct {
	pb.UnimplementedTestcaseServer

	uc     *biz.TestCaseUseCase
	logger *log.Helper
}

func NewTestcaseService(logger log.Logger) *TestcaseService {
	return &TestcaseService{
		logger: log.NewHelper(log.With(logger, "module", "management.testcaseService")),
	}
}

func (s *TestcaseService) CreateTestcase(ctx context.Context, req *pb.CreateTestcaseRequest) (*pb.CreateTestcaseReply, error) {
	return &pb.CreateTestcaseReply{}, nil
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
