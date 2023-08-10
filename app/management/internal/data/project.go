package data

import (
	"context"
	v1 "galileo/api/management/project/v1"
	"galileo/app/management/internal/biz"
	"galileo/ent"
	"galileo/ent/project"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type projectRepo struct {
	data *Data
	log  *log.Helper
}

// NewProjectRepo .
func NewProjectRepo(data *Data, logger log.Logger) biz.ProjectRepo {
	return &projectRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "project.Repo")),
	}
}

func (repo *projectRepo) CreateProject(ctx context.Context, p *biz.Project) (*biz.Project, error) {
	res, err := repo.data.entDB.Project.Create().
		SetName(p.Name).
		SetCreatedBy(p.CreatedBy).
		SetIdentifier(p.Identifier).
		SetDescription(p.Description).
		SetRemark(p.Description).
		SetStartTime(p.StartTime).
		SetDeadline(p.Deadline).
		Save(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, errors.NotFound(ReasonRecordNotFound, err.Error())
	case err != nil:
		return nil, err
	}
	return &biz.Project{
		ID: res.ID,
	}, nil
}

func (repo *projectRepo) GetProjectById(ctx context.Context, id int64) (*biz.Project, error) {
	res, err := repo.data.entDB.Project.Query().
		Where(project.ID(id)).
		Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, errors.NotFound(ReasonRecordNotFound, err.Error())
	case err != nil:
		return nil, err
	}
	return &biz.Project{
		ID:          res.ID,
		Identifier:  res.Identifier,
		Name:        res.Name,
		CreatedAt:   res.CreatedAt,
		CreatedBy:   res.CreatedBy,
		UpdatedAt:   res.UpdatedAt,
		UpdatedBy:   res.UpdatedBy,
		StartTime:   res.StartTime,
		Deadline:    res.Deadline,
		Remark:      res.Remark,
		Description: res.Description,
		Status:      v1.ProjectStatus(res.Status),
	}, err
}

func (repo *projectRepo) UpdateProject(ctx context.Context, p *biz.Project) error {
	tx, err := repo.data.entDB.Tx(ctx)
	if err != nil {
		return SetCustomizeErrMsg(ReasonSystemError, err.Error())
	}
	/* 更新项目信息 */
	err = tx.Project.UpdateOneID(p.ID).
		SetName(p.Name).
		SetIdentifier(p.Identifier).
		SetDescription(p.Description).
		SetRemark(p.Remark).
		SetStartTime(p.StartTime).
		SetDeadline(p.Deadline).
		SetUpdatedBy(p.UpdatedBy).
		SetStatus(int8(p.Status.Number())).
		Exec(ctx)
	switch {
	case ent.IsNotFound(err):
		return rollback(tx, errors.NotFound(ReasonRecordNotFound, err.Error()))
	case err != nil:
		return rollback(tx, err)
	}
	/* 基于项目状态的走不同逻辑流 */
	switch p.Status {
	case v1.ProjectStatus_INPROGRESS:
	case v1.ProjectStatus_DELAY:
	case v1.ProjectStatus_TERMINATED:
	case v1.ProjectStatus_COMPLETED:
	case v1.ProjectStatus_BLOCKED:
		/* 项目阻塞时通知相关人员 */
	}
	return nil
}
