package interfaces

import (
	"galileo/app/runner/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type RunnerUseCase struct {
	runnerService *service.RunnerService
	log           *log.Helper
}

func NewRunnerUseCase(us *service.RunnerService, logger log.Logger) *RunnerUseCase {
	return &RunnerUseCase{runnerService: us, log: log.NewHelper(log.With(logger, "logger"))}
}

func (c *RunnerUseCase) SayHi(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "hello world",
	})
}
