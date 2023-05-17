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
func wireApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, confService *conf.Service, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	client, err := data.NewEntDB(confData)
	if err != nil {
		return nil, nil, err
	}
	cmdable := data.NewRedis(confData, logger)
	discovery := data.NewDiscovery(registry)
	fileClient := data.NewFileServiceClient(confService, discovery)
	dataData, cleanup, err := data.NewData(confData, client, logger, cmdable, fileClient)
	if err != nil {
		return nil, nil, err
	}
	projectRepo := data.NewProjectRepo(dataData, logger)
	projectUseCase := biz.NewProjectUseCase(projectRepo, logger)
	projectService := service.NewProjectService(projectUseCase, logger)
	taskRepo := data.NewTaskRepo(dataData, logger)
	engineClient := data.NewEngineServiceClient(confService, discovery)
	taskUseCase := biz.NewTaskUseCase(taskRepo, engineClient, logger)
	taskService := service.NewTaskService(taskUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, projectService, taskService, logger)
	testcaseRepo := data.NewTestCaseRepo(dataData, logger)
	testcaseUseCase := biz.NewTestcaseUseCase(testcaseRepo, logger)
	testcaseService := service.NewTestcaseService(testcaseUseCase, logger)
	apiRepo := data.NewApiRepo(dataData, logger)
	apiUseCase := biz.NewApiUseCase(apiRepo, logger)
	apiService := service.NewApiService(apiUseCase, logger)
	managementUseCase := biz.NewManagementUseCase(logger)
	managementService := service.NewManagementService(managementUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, auth, projectService, testcaseService, taskService, apiService, managementService, logger)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
