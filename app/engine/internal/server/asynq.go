package server

import (
	"galileo/app/engine/internal/biz"
	"galileo/app/engine/internal/conf"
	"galileo/pkg/asynq"
	"github.com/go-kratos/kratos/v2/log"
)

func NewAsynqServer(c *conf.Data, logger log.Logger) *asynq.Server {
	l := log.NewHelper(log.With(logger, "module", "engine.asynq"))
	srv := asynq.NewServer(
		asynq.WithAddress(c.Redis.Addr),
	)
	var err error
	err = asynq.RegisterSubscriber(srv, biz.TypeDelayedJob, biz.HandleDelayedJob)
	err = asynq.RegisterSubscriber(srv, biz.TypeDefaultJob, biz.HandleDefaultJob)
	//err = asynq.RegisterSubscriber(srv, biz.TypePeriodicJob, biz.HandlePeriodicJob)
	if err != nil {
		l.Info("RegisterSubscriber failed: ", err)
		panic(err)
	}
	return srv
}
