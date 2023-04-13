package service

import (
	"context"
	"galileo/app/runner/internal/biz"
	"galileo/pkg/transport/broker"
	"github.com/go-kratos/kratos/v2/log"
)

func (rs *RunnerService) InsertRunnerData(ctx context.Context, topic string, headers broker.Headers, msg *biz.Runner) error {
	log.Infof("Topic %s, Headers: %+v, Payload: %+v\n", topic, headers, msg)

	if err := rs.runnerDataUseCase.InsertRunnerData(context.Background(), msg); err != nil {
		log.Debug("InsertRunnerData Insert ", err.Error())
		return err
	}

	return nil
}
