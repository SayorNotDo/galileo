// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"galileo/app/runner/internal/biz"
	"galileo/app/runner/internal/conf"
	"galileo/app/runner/internal/data"
	"galileo/app/runner/internal/server"
	"galileo/app/runner/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, auth *conf.Auth, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	runnerRepo := data.NewRunnerRepo(dataData, logger)
	runnerUseCase := biz.NewRunnerUserCase(runnerRepo, logger)
	runnerService := service.NewRunnerService(runnerUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, auth, runnerService, logger)
	app := newApp(logger, httpServer)
	return app, func() {
		cleanup()
	}, nil
}