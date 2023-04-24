package data

import (
	"context"
	fileService "galileo/api/file/v1"
	"galileo/app/management/internal/biz"
	"galileo/ent/testcase"
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
		SetType(testCase.Type).
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
		SetUpdateBy(testcase.UpdatedBy).
		SetType(testcase.Type).
		SetURL(testcase.Url).
		Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *testcaseRepo) TestcaseById(ctx context.Context, id int64) (*biz.Testcase, error) {
	queryTestcase, err := r.data.entDB.TestCase.Query().Where(testcase.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Testcase{
		Name:      queryTestcase.Name,
		CreatedAt: queryTestcase.CreatedAt,
		CreatedBy: queryTestcase.CreatedBy,
		UpdatedBy: *queryTestcase.UpdateBy,
		UpdateAt:  *queryTestcase.UpdateAt,
		Status:    queryTestcase.Status,
		Type:      queryTestcase.Type,
		Priority:  queryTestcase.Priority,
		Url:       queryTestcase.URL,
	}, nil
}

func (r *testcaseRepo) TestcaseByName(ctx context.Context, name string) (*biz.Testcase, error) {
	queryTestcase, err := r.data.entDB.TestCase.Query().Where(testcase.Name(name)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Testcase{
		Id:        queryTestcase.ID,
		Name:      queryTestcase.Name,
		CreatedAt: queryTestcase.CreatedAt,
		CreatedBy: queryTestcase.CreatedBy,
		UpdatedBy: *queryTestcase.UpdateBy,
		UpdateAt:  *queryTestcase.UpdateAt,
		Status:    queryTestcase.Status,
		Type:      queryTestcase.Type,
		Priority:  queryTestcase.Priority,
		Url:       queryTestcase.URL,
	}, nil
}

func (r *testcaseRepo) UploadTestcaseFile(ctx context.Context, fileName, fileType string, content []byte) (string, error) {
	res, err := r.data.fileCli.UploadFile(ctx, &fileService.UploadFileRequest{
		Name:    fileName,
		Type:    fileType,
		Content: content})
	if err != nil {
		return "", err
	}
	return res.Url, nil
}
