package data

import (
	"context"
	"galileo/app/management/internal/biz"
	"galileo/ent"
	"galileo/ent/project"
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
		log:  log.NewHelper(log.With(logger, "module", "management.projectRepo")),
	}
}

func (repo *projectRepo) CreateProject(ctx context.Context, p *biz.Project) (*biz.Project, error) {
	res, err := repo.data.entDB.Project.Create().
		SetName(p.Name).
		SetCreatedBy(p.CreatedBy).
		SetIdentifier(p.Identifier).
		SetDescription(p.Description).
		SetRemark(p.Description).Save(ctx)
	if err != nil {
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
	if err != nil {
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

func (repo *projectRepo) UpdateProject(ctx context.Context, p *biz.Project) (bool, error) {
	err := repo.data.entDB.Project.UpdateOneID(p.ID).
		SetName(p.Name).
		SetIdentifier(p.Identifier).
		SetDescription(p.Description).
		SetRemark(p.Remark).
		SetUpdatedBy(p.UpdatedBy).
		Exec(ctx)
	switch {
	case ent.IsNotFound(err):
		return false, errors.NotFound("Project Not Found", err.Error())
	case err != nil:
		return false, err
	}
	return true, nil
}
