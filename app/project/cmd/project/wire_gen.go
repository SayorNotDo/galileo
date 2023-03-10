// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"galileo/app/project/internal/biz"
	"galileo/app/project/internal/conf"
	"galileo/app/project/internal/data"
	"galileo/app/project/internal/server"
	"galileo/app/project/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	client := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db, client)
	if err != nil {
		return nil, nil, err
	}
	projectRepo := data.NewProjectRepo(dataData, logger)
	projectUseCase := biz.NewProjectUseCase(projectRepo, logger)
	projectService := service.NewProjectService(projectUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, projectService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
