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

type AnalysisSQL struct {
	Sql string
}

type AnalysisSQLBuilder struct {
	analysisSQL AnalysisSQL
}

type AnalysisSqlBuilderOption func(*AnalysisSQLBuilder)

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

func (b *AnalysisSQLBuilder) Build() AnalysisSQL {
	return b.analysisSQL
}
