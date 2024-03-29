// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"galileo/app/analytics/internal/biz"
	"galileo/app/analytics/internal/conf"
	"galileo/app/analytics/internal/data"
	"galileo/app/analytics/internal/server"
	"galileo/app/analytics/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	dataData, err := data.NewData(db, confData, logger)
	if err != nil {
		return nil, nil, err
	}
	analyticsRepo := data.NewAnalyticsRepo(dataData, logger)
	analyticsUseCase := biz.NewAnalyticsUseCase(analyticsRepo, logger)
	analyticsService := service.NewAnalyticsService(analyticsUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, analyticsService, logger)
	app := newApp(logger, httpServer)
	return app, func() {
	}, nil
}
