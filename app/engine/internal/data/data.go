package data

import (
	"context"
	taskV1 "galileo/api/management/task/v1"
	"galileo/app/engine/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	grpcx "google.golang.org/grpc"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEngineRepo, NewTaskServiceClient)

// Data .
type Data struct {
	log     *log.Helper
	taskCli taskV1.TaskClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "engine.DataService"))
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return &Data{log: l}, cleanup, nil
}

func NewTaskServiceClient(sr *conf.Service, rr registry.Discovery) taskV1.TaskClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Task.Endpoint),
		grpc.WithDiscovery(rr),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
			metadata.Client(),
		),
		grpc.WithTimeout(2*time.Second),
		grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)
	if err != nil {
		panic(err)
	}
	c := taskV1.NewTaskClient(conn)
	return c
}
