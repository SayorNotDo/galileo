package data

import (
	"context"
	"galileo/ent/task"
	"time"

	"galileo/app/management/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type taskRepo struct {
	data *Data
	log  *log.Helper
}

// NewTaskRepo .
func NewTaskRepo(data *Data, logger log.Logger) biz.TaskRepo {
	return &taskRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "management.dataService.taskRepo")),
	}
}

func (r *taskRepo) UpdateTask(ctx context.Context, task *biz.Task) (bool, error) {
	if _, err := r.data.entDB.Task.UpdateOneID(task.Id).
		SetName(task.Name).
		SetType(task.Type).
		SetRank(task.Rank).
		SetAssignee(task.Assignee).
		SetConfig(task.Config).
		SetStartTime(task.StartTime).
		SetDeadline(task.Deadline).
		SetDescription(task.Description).
		ClearTestcaseSuite().
		AddTestcaseSuiteIDs(task.TestcaseSuites...).
		Save(ctx); err != nil {
		return false, err
	}
	return true, nil
}

func (r *taskRepo) TaskByName(ctx context.Context, name string) (*biz.Task, error) {
	queryTask, err := r.data.entDB.Task.Query().Where(task.Name(name)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Task{
		Id:        queryTask.ID,
		Name:      queryTask.Name,
		Rank:      queryTask.Rank,
		Status:    queryTask.Status,
		Type:      queryTask.Type,
		CreatedAt: queryTask.CreatedAt,
		CreatedBy: queryTask.CreatedBy,
	}, nil
}

func (r *taskRepo) TaskByID(ctx context.Context, id int64) (*biz.Task, error) {
	queryTask, err := r.data.entDB.Task.Query().Where(task.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	testcaseSuites := queryTask.Edges.TestcaseSuite
	suites := make([]int64, 0)
	for _, o := range testcaseSuites {
		suites = append(suites, o.ID)
	}
	return &biz.Task{
		Id:             queryTask.ID,
		Name:           queryTask.Name,
		Rank:           queryTask.Rank,
		Status:         queryTask.Status,
		Type:           queryTask.Type,
		CreatedAt:      queryTask.CreatedAt,
		CreatedBy:      queryTask.CreatedBy,
		TestcaseSuites: suites,
	}, nil
}

func (r *taskRepo) CreateTask(ctx context.Context, task *biz.Task) (*biz.Task, error) {
	createTask, err := r.data.entDB.Task.Create().
		SetName(task.Name).
		SetType(task.Type).
		SetRank(task.Rank).
		SetCreatedBy(task.CreatedBy).
		SetDescription(task.Description).
		AddTestcaseSuiteIDs(task.TestcaseSuites...).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Task{
		Id:        createTask.ID,
		CreatedAt: createTask.CreatedAt,
	}, nil
}

func (r *taskRepo) SoftDeleteTask(ctx context.Context, uid uint32, id int64) (bool, error) {
	err := r.data.entDB.Task.UpdateOneID(id).
		SetDeletedAt(time.Now()).
		SetDeletedBy(uid).
		Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *taskRepo) TaskDetailById(ctx context.Context, id int64) (*biz.Task, error) {
	queryTask, err := r.data.entDB.Task.Query().Where(task.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Task{
		Id:          queryTask.ID,
		Name:        queryTask.Name,
		CreatedAt:   queryTask.CreatedAt,
		CreatedBy:   queryTask.CreatedBy,
		UpdatedAt:   queryTask.UpdatedAt,
		CompletedAt: queryTask.CompletedAt,
		Status:      queryTask.Status,
		Type:        queryTask.Type,
		Rank:        queryTask.Rank,
		Description: queryTask.Description,
		DeletedAt:   queryTask.DeletedAt,
		DeletedBy:   queryTask.DeletedBy,
	}, nil
}

func (r *taskRepo) IsTaskDeleted(ctx context.Context, id int64) (bool, error) {
	queryTask, err := r.data.entDB.Task.Query().Where(task.ID(id)).Only(ctx)
	if err != nil {
		return false, err
	}
	if !queryTask.CompletedAt.IsZero() && queryTask.DeletedBy > 0 {
		return true, nil
	}
	return false, err
}

func (r *taskRepo) ListTask(ctx context.Context, pageNum, pageSize int) ([]*biz.Task, error) {
	return nil, nil
}

func (r *taskRepo) ListAll(context.Context) ([]*biz.Task, error) {
	return nil, nil
}

func (r *taskRepo) CountAllTask(ctx context.Context) (int, error) {
	count, err := r.data.entDB.Task.Query().Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}
