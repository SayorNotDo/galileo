package server

import (
	"galileo/app/engine/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/asynq"
)

const (
	DefaultJob  = "job:default"
	DelayedJob  = "job:delayed"
	PeriodicJob = "job:periodic"
)

func NewAsynqServer(c *conf.Data, logger log.Logger) *asynq.Server {
	srv := asynq.NewServer(
		asynq.WithAddress(c.Redis.Addr),
	)

	return srv
}
