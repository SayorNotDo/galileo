package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type TestCase struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	CreatedBy uint32    `json:"created_by,omitempty"`
	UpdateAt  time.Time `json:"update_at,omitempty"`
	UpdatedBy uint32    `json:"updated_by,omitempty"`
	Status    int8      `json:"status,omitempty"`
	CaseType  int16     `json:"case_type,omitempty"`
	Priority  int8      `json:"priority,omitempty"`
	DeletedBy uint32    `json:"deleted_by,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Url       string    `json:"url,omitempty"`
}

type TestCaseRepo interface {
	CreateTestCase(ctx context.Context, testCase *TestCase) (*TestCase, error)
}

type TestCaseUseCase struct {
	repo TestCaseRepo
	log  *log.Helper
}

func NewTestCaseUseCase(repo TestCaseRepo, logger log.Logger) *TestCaseUseCase {
	return &TestCaseUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "management.testcaseUseCase")),
	}
}

func (uc *TestCaseUseCase) CreateTestCase(ctx context.Context, testCase *TestCase) (*TestCase, error) {
	return uc.repo.CreateTestCase(ctx, testCase)
}
