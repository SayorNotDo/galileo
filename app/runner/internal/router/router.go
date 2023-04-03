package router

import (
	"galileo/app/runner/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"time"
)

type RunnerHTTPServer interface {
}

func RegisterHTTPServer(runner *service.RunnerService) *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Upgrade", "Origin", "Connection", "Accept-Encoding", "Accept-Language", "Host", "x-requested-with", "CurrentTeamID"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(kgin.Middlewares(recovery.Recovery()))
	router.GET("/helloworld/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		if name == "error" {
			kgin.Error(ctx, errors.Unauthorized("UNAUTHORIZED", "no authorization"))
		} else {
			ctx.JSON(200, map[string]string{"welcome": name})
		}
	})
	return router
}
