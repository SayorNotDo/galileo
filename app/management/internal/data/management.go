package data

import "github.com/go-kratos/kratos/v2/log"

type ManagementRepo struct {
	data *Data
	log  *log.Helper
}

func NewManagementRepo(data *Data, logger log.Logger) *ManagementRepo {
	return &ManagementRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "management.Repo")),
	}
}
