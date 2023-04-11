package service

import (
	"galileo/app/runner/internal/biz"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	stdHttp "net/http"
)

type RunnerService struct {
	runnerUseCase *biz.RunnerUseCase
	log           *log.Helper

	runnerData *biz.RunnerUseCase
}

func NewRunnerService(uc *biz.RunnerUseCase, logger log.Logger) *RunnerService {
	return &RunnerService{
		runnerUseCase: uc,
		log:           log.NewHelper(log.With(logger, "logger", "runnerService"))}
}

func (rs *RunnerService) SayHi(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code":    stdHttp.StatusOK,
		"reason":  "success",
		"message": "hello world",
		"data":    map[string]interface{}{},
	})
}

func (rs *RunnerService) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code":    stdHttp.StatusOK,
		"reason":  "success",
		"message": "hello world",
		"data":    map[string]interface{}{},
	})
}
