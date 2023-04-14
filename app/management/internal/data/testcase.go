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

func NewTestCaseRepo(data *Data, logger log.Logger) biz.TestCaseRepo {
	return &testcaseRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "management.TestCaseRepo")),
	}
}

func (r *testcaseRepo) CreateTestCase(ctx context.Context, testCase *biz.TestCase) (*biz.TestCase, error) {
	return nil, nil
}
