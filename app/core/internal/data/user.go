package data

import (
	"context"
	v1 "galileo/api/core/v1"
	"galileo/app/core/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "user.Repo")),
	}
}

func (repo userRepo) GetUserProjectList(ctx context.Context) ([]*v1.ProjectInfo, error) {
	/* 初始化返回值 */
	var projectList []*v1.ProjectInfo
	repo.log.Info("--------------------------------userRepo.GetUserProjectList")
	res, err := repo.data.projectCli.GetUserProjectList(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}
	for _, v := range res.ProjectList {
		projectList = append(projectList, &v1.ProjectInfo{
			Id:          v.Id,
			Name:        v.Name,
			Identifier:  v.Identifier,
			CreatedAt:   v.CreatedAt,
			CreatedBy:   v.CreatedBy,
			UpdatedBy:   v.UpdatedBy,
			UpdatedAt:   v.UpdatedAt,
			DeletedAt:   v.DeletedAt,
			DeletedBy:   v.DeletedBy,
			StartTime:   v.StartTime,
			Deadline:    v.Deadline,
			Description: v.Description,
			Remark:      v.Remark,
		})
	}
	return nil, nil
}
