package service

import (
	"context"
	pb "galileo/api/engine/v1"
	"galileo/app/engine/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type EngineService struct {
	pb.UnimplementedEngineServer

	uc  *biz.EngineUseCase
	log *log.Helper
}

func (s *EngineService) RunJob(ctx context.Context, req *pb.RunJobRequest) (*pb.RunJobReply, error) {
	// TODO: query task from db
	// TODO: validate job configuration
	// TODO: validate worker environment
	// TODO: send job command through rabbitmq
	// TODO: get callback from rabbitmq
	// TODO: return response
	return nil, nil
}
