package biz

import (
	"context"
	"galileo/app/core/internal/pkg/constant"
	"galileo/app/core/internal/pkg/middleware/auth"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"time"
)

type CoreRepo interface {
	DataReportTrack(ctx context.Context, data []map[string]interface{}, claims *auth.ReportClaims) error
}

type User struct {
	Id          uint32
	Phone       string
	ChineseName string
	Username    string
	Nickname    string
	Gender      string
	Email       string
	Role        int32
	Password    string
	Status      bool
	CreatedAt   time.Time
	DeletedAt   time.Time
	DeletedBy   uint32
}

type CoreUseCase struct {
	repo CoreRepo
	log  *log.Helper
}

func NewCoreUseCase(repo CoreRepo, logger log.Logger) *CoreUseCase {
	helper := log.NewHelper(log.With(logger, "module", "core.useCase"))
	return &CoreUseCase{
		repo: repo,
		log:  helper,
	}
}

func newReportClaim(machine string) *auth.ReportClaims {
	return &auth.ReportClaims{
		Machine: machine,
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),
			ExpiresAt: jwt2.NewNumericDate(time.Now().AddDate(1, 0, 0)),
			IssuedAt:  jwt2.NewNumericDate(time.Now()),
			Issuer:    "Galileo",
		},
	}
}

func (c *CoreUseCase) DataReportTrack(ctx context.Context, data []map[string]interface{}, claims *auth.ReportClaims) error {
	return c.repo.DataReportTrack(ctx, data, claims)
}

func (c *CoreUseCase) ExecuteToken(ctx context.Context, machine string) (string, error) {
	if len(machine) <= 0 {
		return "", SetErrByReason(ReasonParamsError)
	}
	claims := newReportClaim(machine)
	token, err := auth.GenerateExecuteToken(*claims, constant.ReportKey)
	if err != nil {
		return "", SetErrByReason(ReasonSystemError)
	}
	return token, nil
}
