// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"galileo/app/core/internal/biz"
	"galileo/app/core/internal/conf"
	"galileo/app/core/internal/data"
	"galileo/app/core/internal/server"
	"galileo/app/core/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, confService *conf.Service, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	userClient := data.NewUserServiceClient(auth, confService, discovery)
	cmdable := data.NewRedis(confData, logger)
	syncProducer := data.NewKafkaProducer(confData, logger)
	dataData, err := data.NewData(confData, userClient, logger, cmdable, syncProducer)
	if err != nil {
		return nil, nil, err
	}
	coreRepo := data.NewCoreRepo(dataData, logger)
	coreUseCase := biz.NewCoreUseCase(coreRepo, logger, auth)
	coreService := service.NewCoreService(coreUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, auth, coreService, logger)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, httpServer, registrar)
	return app, func() {
	}, nil
}
