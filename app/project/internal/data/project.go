package data

import (
	"galileo/app/project/internal/biz"
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
		log:  log.NewHelper(logger),
	}
}
