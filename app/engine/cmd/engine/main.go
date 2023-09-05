package main

import (
	"flag"
	"galileo/app/engine/internal/conf"
	"galileo/pkg/utils/bootstrap"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/tx7do/kratos-transport/transport/asynq"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Service = bootstrap.NewServiceInfo(
		"galileo.engine.scheduler",
		"engine.v1",
		"")

	// Flags is the config flag.
	Flags = bootstrap.NewCommandFlags()
)

func init() {
	Flags.Init()

	json.MarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}
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

func newApp(logger log.Logger, gs *grpc.Server, aq *asynq.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			aq,
		),
		kratos.Registrar(rr),
	)
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

	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Service, rc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
