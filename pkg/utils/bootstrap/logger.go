package bootstrap

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"os"
)

func NewLoggerProvider(serviceInfo *ServiceInfo) log.Logger {
	l := log.NewStdLogger(os.Stdout)
	return log.With(
		l,
		"scheduler.id", serviceInfo.Id,
		"scheduler.name", serviceInfo.Name,
		"scheduler.version", serviceInfo.Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}
