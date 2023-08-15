package data

import (
	"galileo/app/analytics/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type analyticsRepo struct {
	data *Data
	log  *log.Helper
}

func NewAnalyticsRepo(data *Data, logger log.Logger) biz.AnalyticsRepo {
	return &analyticsRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "analytics.Repo")),
	}
}

/* TODO：DryRun 生成提交查询的SQL到大数据处 */

type AnalysisSQL struct {
	tableName   string
	selectCols  []string
	joins       []string
	conditions  []string
	groupBy     string
	orderBy     string
	limit       int
	offset      int
	nestedWhere []*AnalysisSQL /* 存储嵌套的条件表达式 */
}

type AnalysisSQLBuilder struct {
	analysisSQL AnalysisSQL
}

type AnalysisSqlBuilderWithOption func(*AnalysisSQLBuilder)

func WithSELECT() AnalysisSqlBuilderWithOption {
	return func(builder *AnalysisSQLBuilder) {
		builder.analysisSQL.selectCols = []string{}
	}
}

func WithFROM() AnalysisSqlBuilderWithOption {
	return func(builder *AnalysisSQLBuilder) {}
}

func Offset() AnalysisSqlBuilderWithOption {
	return func(builder *AnalysisSQLBuilder) {}

}

func NewAnalysisSQLBuilder(options ...AnalysisSqlBuilderWithOption) *AnalysisSQLBuilder {
	builder := &AnalysisSQLBuilder{}
	for _, option := range options {
		option(builder)
	}
	return builder
}

func NewAnalysisSqlBuilder() *AnalysisSQLBuilder {
	return &AnalysisSQLBuilder{}
}

func (b *AnalysisSQLBuilder) WithSELECT() *AnalysisSQLBuilder {
	return b
}

func (b *AnalysisSQLBuilder) WithFROM() *AnalysisSQLBuilder {
	return b
}

func (b *AnalysisSQLBuilder) WithWHERE() *AnalysisSQLBuilder {
	return b
}

func (b *AnalysisSQLBuilder) WithGROUPBY() *AnalysisSQLBuilder {
	return b
}

func (b *AnalysisSQLBuilder) WithORDERBY() *AnalysisSQLBuilder {
	return b
}

func (b *AnalysisSQLBuilder) WithLIMIT() *AnalysisSQLBuilder {
	return b
}

func (b *AnalysisSQLBuilder) WithOFFSET() *AnalysisSQLBuilder {
	return b
}

func (b *AnalysisSQLBuilder) Build() AnalysisSQL {
	return b.analysisSQL
}
