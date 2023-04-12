package main

import (
	"flag"
	"galileo/app/runner/internal/conf"
	"galileo/pkg/transport/server/rabbitmq"
	"galileo/pkg/utils/bootstrap"
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Service = bootstrap.NewServiceInfo(
		"kratos.runner.api",
		"0.0.0",
		"",
	)

	// Flags is the config flag.
	Flags = bootstrap.NewCommandFlags()
)

func init() {
	Flags.Init()
}

func newApp(logger log.Logger, hs *http.Server, rs *rabbitmq.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			rs,
		),
		kratos.Registrar(rr),
	)
}

func loadConfig() (*conf.Bootstrap, *conf.Registry) {
	c := bootstrap.NewConfigProvider(Flags.ConfigType, Flags.ConfigHost, Flags.Conf, Service.Name)

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}

	return &bc, &rc
}

func main() {
	flag.Parse()

	logger := bootstrap.NewLoggerProvider(&Service)

	bc, rc := loadConfig()
	if bc == nil || rc == nil {
		panic("load config failed")
	}

	err := bootstrap.NewTracerProvider(bc.Trace.Endpoint, Flags.Env, &Service)
	if err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Auth, bc.Data, logger, rc)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
