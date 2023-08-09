package data

import (
	"context"
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
		SetRemark(p.Description).Save(ctx)
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
		Remark:      res.Remark,
		Description: res.Description,
		Status:      res.Status,
	}, err
}

func (repo *projectRepo) UpdateProject(ctx context.Context, p *biz.Project) error {
	err := repo.data.entDB.Project.UpdateOneID(p.ID).
		SetName(p.Name).
		SetIdentifier(p.Identifier).
		SetDescription(p.Description).
		SetRemark(p.Remark).
		SetUpdatedBy(p.UpdatedBy).
		Exec(ctx)
	switch {
	case ent.IsNotFound(err):
		return errors.NotFound(ReasonRecordNotFound, err.Error())
	case err != nil:
		return err
	}
	return nil
}
