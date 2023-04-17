package data

import (
	"context"
	"galileo/app/management/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type testcaseRepo struct {
	data *Data
	log  *log.Helper
}

func NewTestCaseRepo(data *Data, logger log.Logger) biz.TestcaseRepo {
	return &testcaseRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "management.TestCaseRepo")),
	}
}

func (r *testcaseRepo) CreateTestcase(ctx context.Context, testCase *biz.Testcase) (*biz.Testcase, error) {
	createTestcase, err := r.data.entDB.TestCase.Create().
		SetName(testCase.Name).
		SetCreatedBy(testCase.CreatedBy).
		SetPriority(testCase.Priority).
		SetType(testCase.CaseType).
		SetURL(testCase.Url).
		SetDescription(testCase.Description).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Testcase{
		Id:        createTestcase.ID,
		CreatedAt: createTestcase.CreatedAt,
	}, nil
}

func (r *testcaseRepo) UpdateTestcase(ctx context.Context, testcase *biz.Testcase) (bool, error) {
	err := r.data.entDB.TestCase.UpdateOneID(testcase.Id).
		SetName(testcase.Name).
		SetPriority(testcase.Priority).
		SetDescription(testcase.Description).
		SetType(testcase.CaseType).
		SetURL(testcase.Url).
		Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *testcaseRepo) TestcaseById(ctx context.Context, id int64) (*biz.Testcase, error) {
	return nil, nil
}
