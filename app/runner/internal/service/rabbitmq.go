package service

import (
	"context"
	"fmt"
	"galileo/app/runner/internal/biz"
	"galileo/pkg/transport/broker"
)

func (rs *RunnerService) InsertRunnerData(ctx context.Context, topic string, headers broker.Headers, msg *biz.Runner) error {
	fmt.Println("InsertData() Topic: ", topic)
	return nil
}
