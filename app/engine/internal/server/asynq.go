package server

import (
	"galileo/app/engine/internal/biz"
	"galileo/app/engine/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/asynq"
)

func NewAsynqServer(c *conf.Data, logger log.Logger) *asynq.Server {
	srv := asynq.NewServer(
		asynq.WithAddress(c.Redis.Addr),
	)
	var err error
	err = srv.HandleFunc(biz.TypeDelayedJob, biz.HandleDelayedJob)
	err = srv.HandleFunc(biz.TypeDefaultJob, biz.HandleDefaultJob)
	err = srv.HandleFunc(biz.TypePeriodicJob, biz.HandlePeriodicJob)
	if err != nil {
		panic(err)
	}
	return srv
}
