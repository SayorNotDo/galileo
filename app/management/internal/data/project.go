package data

import (
	"context"
	"galileo/app/management/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Project struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"varchar(255"`
	Identifier string    `json:"identifier" gorm:"varchar(255)"`
	Status     uint32    `json:"status" gorm:"default:1"`
	CreatedBy  uint32    `json:"created_by"`
	CreateAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdateBy   uint32    `json:"update_by"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeleteBy   uint32    `json:"delete_by"`
	DeleteAt   time.Time `json:"delete_at" gorm:"default:null"`
	Remark     string    `json:"remark"`
}

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
