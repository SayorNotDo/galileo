package server

import (
	"context"
	"encoding/json"
	"fmt"
	"galileo/app/runner/internal/biz"
	"galileo/pkg/transport/broker"
	"github.com/go-kratos/kratos/v2/log"
)

func runnerCreator() broker.Any     { return nil }
func runnerDataCreator() broker.Any { return &biz.Runner{} }

type runnerHandler func(_ context.Context, topic string, headers broker.Headers, msg *biz.Runner) error

type runnerDataHandler func(_ context.Context, topic string, headers broker.Headers, msg *biz.Runner) error

func RegisterRunnerDataJsonHandler(fnc runnerDataHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		switch t := event.Message().Body.(type) {
		case *biz.Runner:
			if err := fnc(ctx, event.Topic(), event.Message().Headers, t); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		return nil
	}
}

func registerRunnerDataHandler(fnc runnerDataHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		var r biz.Runner
		switch t := event.Message().Body.(type) {
		case *biz.Runner:
		case []byte:
			if err := json.Unmarshal(t, &r); err != nil {
				log.Error("json Unmarshal failed: ", err.Error())
			}
		case string:
			if err := json.Unmarshal([]byte(t), &r); err != nil {
				log.Error("json Unmarshal failed: ", err.Error())
			}
		default:
			log.Error("unknown type Unmarshal failed:", t)
			return fmt.Errorf("unsupported type: %T", t)
		}
		if err := fnc(ctx, event.Topic(), event.Message().Headers, &r); err != nil {
			return err
		}
		return nil
	}
}

func registerRunnerHandler(fnc runnerHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		return nil
	}
}
