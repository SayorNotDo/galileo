package service

import (
	"context"

	pb "galileo/api/engine/v1"
)

type EngineService struct {
	pb.UnimplementedEngineServer
}

func NewEngineService() *EngineService {
	return &EngineService{}
}

func (s *EngineService) CreateEngine(ctx context.Context, req *pb.CreateEngineRequest) (*pb.CreateEngineReply, error) {
	return &pb.CreateEngineReply{}, nil
}
func (s *EngineService) UpdateEngine(ctx context.Context, req *pb.UpdateEngineRequest) (*pb.UpdateEngineReply, error) {
	return &pb.UpdateEngineReply{}, nil
}
func (s *EngineService) DeleteEngine(ctx context.Context, req *pb.DeleteEngineRequest) (*pb.DeleteEngineReply, error) {
	return &pb.DeleteEngineReply{}, nil
}
func (s *EngineService) GetEngine(ctx context.Context, req *pb.GetEngineRequest) (*pb.GetEngineReply, error) {
	return &pb.GetEngineReply{}, nil
}
func (s *EngineService) ListEngine(ctx context.Context, req *pb.ListEngineRequest) (*pb.ListEngineReply, error) {
	return &pb.ListEngineReply{}, nil
}
