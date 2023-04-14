// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"galileo/app/management/internal/biz"
	"galileo/app/management/internal/conf"
	"galileo/app/management/internal/data"
	"galileo/app/management/internal/server"
	"galileo/app/management/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	client, err := data.NewEntDB(confData)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup, err := data.NewData(confData, client, logger)
	if err != nil {
		return nil, nil, err
	}
	projectRepo := data.NewProjectRepo(dataData, logger)
	projectUseCase := biz.NewProjectUseCase(projectRepo, logger)
	projectService := service.NewProjectService(projectUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, projectService, logger)
	testcaseService := service.NewTestcaseService(logger)
	httpServer := server.NewHTTPServer(confServer, projectService, testcaseService, logger)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}