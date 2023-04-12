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

	runnerDataUseCase *biz.RunnerDataUseCase
}

func NewRunnerService(runner *biz.RunnerUseCase, runnerData *biz.RunnerDataUseCase, logger log.Logger) *RunnerService {
	return &RunnerService{
		runnerUseCase:     runner,
		runnerDataUseCase: runnerData,
		log:               log.NewHelper(log.With(logger, "module", "service.runnerService")),
	}
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
	err := rs.runnerDataUseCase.InsertRunnerData(ctx, &biz.Runner{})
	if err != nil {
		return
	}
	ctx.JSON(200, gin.H{
		"code":    stdHttp.StatusOK,
		"reason":  "success",
		"message": "hello world",
		"data":    map[string]interface{}{},
	})
}
