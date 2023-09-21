package data

import (
	"fmt"
	"strings"
)

/* TODO：DryRun 生成提交查询的SQL */

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

func (b *AnalysisSQLBuilder) WithSELECT(columns ...string) *AnalysisSQLBuilder {
	b.analysisSQL.selectCols = columns
	return b
}

func (b *AnalysisSQLBuilder) WithFROM(tableName string) *AnalysisSQLBuilder {
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

func (b *AnalysisSQLBuilder) Build() string {
	/* 初始化SQL查询语句 */
	sql := "SELECT "
	if len(b.analysisSQL.selectCols) > 0 {
		sql += strings.Join(b.analysisSQL.selectCols, ", ")
	} else {
		sql += "*"
	}
	/* 判断是否存在嵌套的条件 */
	if b.analysisSQL.nestedWhere != nil {
		/* 递归生成嵌套的SQL语句 */
		fmt.Println("nestedWhere is not nil")
	}
	/* 查询表 */
	sql += " FROM " + b.analysisSQL.tableName
	fmt.Println(b.analysisSQL)
	return sql
}
