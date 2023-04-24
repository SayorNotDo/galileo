package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Testcase struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	CreatedBy   uint32    `json:"created_by,omitempty"`
	UpdateAt    time.Time `json:"update_at,omitempty"`
	UpdatedBy   uint32    `json:"updated_by,omitempty"`
	Status      int8      `json:"status,omitempty"`
	Type        int8      `json:"type,omitempty"`
	Priority    int8      `json:"priority,omitempty"`
	DeletedBy   uint32    `json:"deleted_by,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	Description string    `json:"description,omitempty"`
	Url         string    `json:"url,omitempty"`
}

type TestcaseRepo interface {
	CreateTestcase(ctx context.Context, testcase *Testcase) (*Testcase, error)
	UpdateTestcase(ctx context.Context, testcase *Testcase) (bool, error)
	TestcaseById(ctx context.Context, id int64) (*Testcase, error)
	TestcaseByName(ctx context.Context, name string) (*Testcase, error)
	UploadTestcaseFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error)
}

type TestcaseUseCase struct {
	repo TestcaseRepo
	log  *log.Helper
}

func NewTestcaseUseCase(repo TestcaseRepo, logger log.Logger) *TestcaseUseCase {
	return &TestcaseUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "management.testcaseUseCase")),
	}
}

func (uc *TestcaseUseCase) CreateTestcase(ctx context.Context, testcase *Testcase) (*Testcase, error) {
	return uc.repo.CreateTestcase(ctx, testcase)
}

func (uc *TestcaseUseCase) UpdateTestcase(ctx context.Context, testcase *Testcase) (bool, error) {
	return uc.repo.UpdateTestcase(ctx, testcase)
}

func (uc *TestcaseUseCase) TestcaseById(ctx context.Context, id int64) (*Testcase, error) {
	return uc.repo.TestcaseById(ctx, id)
}

func (uc *TestcaseUseCase) TestcaseByName(ctx context.Context, name string) (*Testcase, error) {
	return uc.repo.TestcaseByName(ctx, name)
}

func (uc *TestcaseUseCase) UploadTestcaseFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error) {
	return uc.repo.UploadTestcaseFile(ctx, fileName, fileType, content)
}
