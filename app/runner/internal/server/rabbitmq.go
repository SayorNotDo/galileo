package server

import (
	"context"
	"galileo/app/runner/internal/conf"
	"galileo/app/runner/internal/service"
	"galileo/pkg/transport/server/rabbitmq"
	"github.com/go-kratos/kratos/v2/log"
)

func NewRabbitServer(c *conf.Server, _ log.Logger, runner *service.RunnerService) *rabbitmq.Server {
	ctx := context.Background()

	srv := rabbitmq.NewServer(
		rabbitmq.WithAddress(c.Rabbitmq.Addr),
		rabbitmq.WithCodec("json"),
	)

	registerRabbitmqSubscriber(ctx, srv, runner)

	return srv
}

func registerRabbitmqSubscriber(ctx context.Context, srv *rabbitmq.Server, runner *service.RunnerService) {
	_ = srv.RegisterSubscriber(ctx,
		"logger.runner.ts",
		registerRunnerDataHandler(runner.InsertRunnerData),
		runnerDataCreator)
}
