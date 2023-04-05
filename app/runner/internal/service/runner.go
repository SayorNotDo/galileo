package interfaces

import (
	"galileo/app/runner/internal/biz"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	stdHttp "net/http"
)

type RunnerService struct {
	runnerUseCase *biz.RunnerUseCase
	log           *log.Helper
}

func NewRunnerInterface(uc *biz.RunnerUseCase, logger log.Logger) *RunnerService {
	return &RunnerService{runnerUseCase: uc, log: log.NewHelper(log.With(logger, "logger", "runnerService"))}
}

func (uc *RunnerService) SayHi(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code":    stdHttp.StatusOK,
		"reason":  "success",
		"message": "hello world",
		"data":    map[string]interface{}{},
	})
}
