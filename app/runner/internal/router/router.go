package router

import (
	"galileo/app/runner/internal/service"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/errors"
)

type RunnerHTTPServer interface {
}

func RegisterHTTPServer(runner *service.RunnerService) *gin.Engine {
	router := gin.New()

	loader := router.Group("/v1/api/")
	loader.GET("helloworld/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		if name == "error" {
			kgin.Error(ctx, errors.Unauthorized("UNAUTHORIZED", "no authorization"))
		} else {
			ctx.JSON(200, map[string]string{"welcome": name})
		}
	})
	return router
}
