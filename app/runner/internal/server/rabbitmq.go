package server

import (
	"context"
	"galileo/app/runner/internal/conf"
	"galileo/app/runner/internal/service"
	"galileo/pkg/transport/broker"
	rabbitmqBroker "galileo/pkg/transport/broker/rabbitmq"
	"galileo/pkg/transport/server/rabbitmq"
	"github.com/go-kratos/kratos/v2/log"
)

const (
	exchange   = "runner_exchange"
	routingKey = "logger.runner.ts"
	queueName  = "runner_queue"
)

func NewRabbitServer(c *conf.Server, _ log.Logger, runner *service.RunnerService) *rabbitmq.Server {
	ctx := context.Background()

	srv := rabbitmq.NewServer(
		rabbitmq.WithAddress(c.Rabbitmq.Addr),
		rabbitmq.WithCodec("json"),
		rabbitmq.WithExchange(exchange, true),
	)

	registerRabbitmqSubscriber(ctx, srv, runner)

	return srv
}

func registerRabbitmqSubscriber(ctx context.Context, srv *rabbitmq.Server, runner *service.RunnerService) {
	_ = srv.RegisterSubscriber(ctx,
		routingKey,
		registerRunnerDataHandler(runner.InsertRunnerData),
		runnerDataCreator,
		broker.WithQueueName(queueName),
		rabbitmqBroker.WithDurableQueue())
	if err := srv.Start(ctx); err != nil {
		panic(err)
	}
}
