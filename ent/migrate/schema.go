// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "created_at", Type: field.TypeTime},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// ProjectColumns holds the columns for the "project" table.
	ProjectColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "identifier", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "update_by", Type: field.TypeUint32, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUint32, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 0},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "remark", Type: field.TypeString, Nullable: true, Size: 2147483647},
	}
	// ProjectTable holds the schema information for the "project" table.
	ProjectTable = &schema.Table{
		Name:       "project",
		Columns:    ProjectColumns,
		PrimaryKey: []*schema.Column{ProjectColumns[0]},
	}
	// TaskColumns holds the columns for the "task" table.
	TaskColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "rank", Type: field.TypeInt8, Default: 0},
		{Name: "type", Type: field.TypeInt8},
		{Name: "status", Type: field.TypeInt8, Default: 0},
		{Name: "complete_at", Type: field.TypeTime, Nullable: true},
		{Name: "update_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUint32, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
	}
	// TaskTable holds the schema information for the "task" table.
	TaskTable = &schema.Table{
		Name:       "task",
		Columns:    TaskColumns,
		PrimaryKey: []*schema.Column{TaskColumns[0]},
	}
	// TestCasesColumns holds the columns for the "test_cases" table.
	TestCasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "update_by", Type: field.TypeUint32, Nullable: true},
		{Name: "update_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 0},
		{Name: "type", Type: field.TypeInt8, Default: 0},
		{Name: "priority", Type: field.TypeInt8, Default: 0},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUint32, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "url", Type: field.TypeString, Nullable: true},
	}
	// TestCasesTable holds the schema information for the "test_cases" table.
	TestCasesTable = &schema.Table{
		Name:       "test_cases",
		Columns:    TestCasesColumns,
		PrimaryKey: []*schema.Column{TestCasesColumns[0]},
	}
	// TestCaseSuitesColumns holds the columns for the "test_case_suites" table.
	TestCaseSuitesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "task_testcase_suite", Type: field.TypeInt64, Nullable: true},
	}
	// TestCaseSuitesTable holds the schema information for the "test_case_suites" table.
	TestCaseSuitesTable = &schema.Table{
		Name:       "test_case_suites",
		Columns:    TestCaseSuitesColumns,
		PrimaryKey: []*schema.Column{TestCaseSuitesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "test_case_suites_task_testcaseSuite",
				Columns:    []*schema.Column{TestCaseSuitesColumns[4]},
				RefColumns: []*schema.Column{TaskColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "chinese_name", Type: field.TypeString, Nullable: true},
		{Name: "nickname", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "role", Type: field.TypeUint8, Default: 3},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUint32, Nullable: true},
		{Name: "is_deleted", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "group_user", Type: field.TypeInt, Nullable: true},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_groups_user",
				Columns:    []*schema.Column{UserColumns[16]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TestCaseSuiteTestcasesColumns holds the columns for the "test_case_suite_testcases" table.
	TestCaseSuiteTestcasesColumns = []*schema.Column{
		{Name: "test_case_suite_id", Type: field.TypeInt64},
		{Name: "test_case_id", Type: field.TypeInt64},
	}
	// TestCaseSuiteTestcasesTable holds the schema information for the "test_case_suite_testcases" table.
	TestCaseSuiteTestcasesTable = &schema.Table{
		Name:       "test_case_suite_testcases",
		Columns:    TestCaseSuiteTestcasesColumns,
		PrimaryKey: []*schema.Column{TestCaseSuiteTestcasesColumns[0], TestCaseSuiteTestcasesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "test_case_suite_testcases_test_case_suite_id",
				Columns:    []*schema.Column{TestCaseSuiteTestcasesColumns[0]},
				RefColumns: []*schema.Column{TestCaseSuitesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "test_case_suite_testcases_test_case_id",
				Columns:    []*schema.Column{TestCaseSuiteTestcasesColumns[1]},
				RefColumns: []*schema.Column{TestCasesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroupsTable,
		ProjectTable,
		TaskTable,
		TestCasesTable,
		TestCaseSuitesTable,
		UserTable,
		TestCaseSuiteTestcasesTable,
	}
)

func init() {
	ProjectTable.Annotation = &entsql.Annotation{
		Table: "project",
	}
	TaskTable.Annotation = &entsql.Annotation{
		Table: "task",
	}
	TestCaseSuitesTable.ForeignKeys[0].RefTable = TaskTable
	UserTable.ForeignKeys[0].RefTable = GroupsTable
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
	TestCaseSuiteTestcasesTable.ForeignKeys[0].RefTable = TestCaseSuitesTable
	TestCaseSuiteTestcasesTable.ForeignKeys[1].RefTable = TestCasesTable
}
