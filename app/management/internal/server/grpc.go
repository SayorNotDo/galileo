package server

import (
	projectV1 "galileo/api/management/project/v1"
	taskV1 "galileo/api/management/task/v1"
	"galileo/app/management/internal/conf"
	"galileo/app/management/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(tr *conf.Trace, c *conf.Server, project *service.ProjectService, task *service.TaskService, logger log.Logger) *grpc.Server {
	err := initTracer(tr.Endpoint)
	if err != nil {
		panic(err)
	}
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			tracing.Server(),
			metadata.Server(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	projectV1.RegisterProjectServer(srv, project)
	taskV1.RegisterTaskServer(srv, task)
	return srv
}
