// Code generated by ent, DO NOT EDIT.

package testcasesuite

import (
	"time"
)

const (
	// Label holds the string label denoting the testcasesuite type in the database.
	Label = "testcase_suite"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// EdgeTestcase holds the string denoting the testcase edge name in mutations.
	EdgeTestcase = "testcase"
	// Table holds the table name of the testcasesuite in the database.
	Table = "testcase_suite"
	// TestcaseTable is the table that holds the testcase relation/edge. The primary key declared below.
	TestcaseTable = "testcase_suite_testcase"
	// TestcaseInverseTable is the table name for the Testcase entity.
	// It exists in this package in order to avoid circular dependency with the "testcase" package.
	TestcaseInverseTable = "testcase"
)

// Columns holds all SQL columns for testcasesuite fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
	FieldCreatedBy,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "testcase_suite"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"task_testcase_suite",
}

var (
	// TestcasePrimaryKey and TestcaseColumn2 are the table columns denoting the
	// primary key for the testcase relation (M2M).
	TestcasePrimaryKey = []string{"testcase_suite_id", "testcase_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
