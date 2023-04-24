// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"galileo/app/file/internal/biz"
	"galileo/app/file/internal/conf"
	"galileo/app/file/internal/data"
	"galileo/app/file/internal/server"
	"galileo/app/file/internal/service"
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
	cmdable := data.NewRedis(confData, logger)
	ossClient, err := data.NewOssClient(confData)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup, err := data.NewData(confData, client, logger, cmdable, ossClient)
	if err != nil {
		return nil, nil, err
	}
	fileRepo := data.NewFileRepo(dataData, logger)
	fileUseCase := biz.NewFileUseCase(fileRepo, logger)
	fileService := service.NewFileService(fileUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, fileService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
