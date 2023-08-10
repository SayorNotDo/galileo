package biz

import (
	"context"
	v1 "galileo/api/management/v1"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ManagementRepo interface {
	CreateTestPlan(context.Context, *TestPlan) (*TestPlan, error)
	UpdateTestPlan(context.Context, *TestPlan) error
	GetTestPlanById(ctx context.Context, id int64) (*TestPlan, error)
}

type BaseInfo struct {
	ApiCount       int64
	ApiCaseCount   int64
	CronJobCount   int64
	SceneCaseCount int64
}

type TestPlan struct {
	Id              int64
	Name            string
	CreatedAt       time.Time
	CreatedBy       uint32
	UpdatedAt       time.Time
	UpdatedBy       uint32
	Description     string
	StartTime       time.Time
	Deadline        time.Time
	StatusUpdatedAt time.Time
	Status          v1.PlanStatus
	Tasks           []int64
	ProjectId       int64
}

type ManagementUseCase struct {
	repo   ManagementRepo
	logger *log.Helper
}

func NewManagementUseCase(logger log.Logger) *ManagementUseCase {
	return &ManagementUseCase{
		logger: log.NewHelper(log.With(logger, "module", "management.UseCase")),
	}
}

func (uc *ManagementUseCase) BaseInformation(ctx context.Context) (*BaseInfo, error) {
	return &BaseInfo{}, nil
}
func (uc *ManagementUseCase) CreateTestPlan(ctx context.Context, testPlan *TestPlan) (*TestPlan, error) {
	return uc.repo.CreateTestPlan(ctx, testPlan)
}

func (uc *ManagementUseCase) GetTestPlanById(ctx context.Context, id int64) (*TestPlan, error) {
	return uc.repo.GetTestPlanById(ctx, id)
}

func (uc *ManagementUseCase) UpdateTestPlan(ctx context.Context, testPlan *TestPlan) error {
	return uc.repo.UpdateTestPlan(ctx, testPlan)
}

func NewTestPlan(userId uint32, name, description string, startTime, deadline *timestamppb.Timestamp) (TestPlan, error) {
	if len(name) <= 0 {
		return TestPlan{}, SetCustomizeErrMsg(ReasonParamsError, "test plan name cannot be empty")
	} else if len(description) <= 0 {
	} else if userId <= 0 {
		return TestPlan{}, SetErrByReason(ReasonUserUnauthorized)
	}
	return TestPlan{
		Name:        name,
		CreatedBy:   userId,
		Description: description,
		StartTime:   startTime.AsTime(),
		Deadline:    deadline.AsTime(),
	}, nil
}
