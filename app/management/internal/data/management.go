package data

import (
	"context"
	managementV1 "galileo/api/management/v1"
	"galileo/app/management/internal/biz"
	"galileo/ent"
	. "galileo/pkg/errResponse"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

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

func (r *ManagementRepo) CreateTestPlan(ctx context.Context, plan *biz.TestPlan) (*biz.TestPlan, error) {
	tx, err := r.data.entDB.Tx(ctx)
	if err != nil {
		return nil, SetCustomizeErrMsg(ReasonSystemError, err.Error())
	}
	/* 创建测试计划 */
	ret, err := tx.TestPlan.Create().
		SetName(plan.Name).
		SetProjectID(plan.ProjectId).
		SetCreatedBy(plan.CreatedBy).
		SetDescription(plan.Description).
		SetStartTime(plan.StartTime).
		SetDeadline(plan.Deadline).
		SetTasks(plan.Tasks).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	return &biz.TestPlan{
		Id: ret.ID,
	}, nil
}

func (r *ManagementRepo) UpdateTestPlan(ctx context.Context, plan *biz.TestPlan) error {
	tx, err := r.data.entDB.Tx(ctx)
	if err != nil {
		return SetCustomizeErrMsg(ReasonSystemError, err.Error())
	}
	/* 更新测试计划 */
	ret, err := tx.TestPlan.UpdateOneID(plan.Id).
		SetName(plan.Name).
		SetUpdatedBy(plan.UpdatedBy).
		SetDescription(plan.Description).
		Save(ctx)
	switch {
	case ent.IsNotFound(err):
		return rollback(tx, SetErrByReason(ReasonRecordNotFound))
	case err != nil:
		return rollback(tx, SetCustomizeErrMsg(ReasonUnknownError, err.Error()))
	}
	if ret.Status != int8(plan.Status.Number()) {
		ret, err = tx.TestPlan.UpdateOneID(plan.Id).
			SetStatus(int8(plan.Status.Number())).
			SetStatusUpdatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
		/* 基于计划更新的状态执行对应逻辑 */
		switch plan.Status {
		case managementV1.PlanStatus_INPROGRESS:
		case managementV1.PlanStatus_BLOCKED:
		case managementV1.PlanStatus_DELAY:
		case managementV1.PlanStatus_TERMINATED:
		case managementV1.PlanStatus_COMPLETED:
		}
	}

	/* 提交事务 */
	if err := tx.Commit(); err != nil {
		return rollback(tx, err)
	}
	return nil
}
