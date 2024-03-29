// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// APIColumns holds the columns for the "api" table.
	APIColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "url", Type: field.TypeString},
		{Name: "type", Type: field.TypeInt8},
		{Name: "status", Type: field.TypeInt8, Default: 0},
		{Name: "headers", Type: field.TypeString, Nullable: true},
		{Name: "body", Type: field.TypeString, Nullable: true},
		{Name: "label", Type: field.TypeString, Nullable: true},
		{Name: "query_params", Type: field.TypeString, Nullable: true},
		{Name: "response", Type: field.TypeString, Nullable: true},
		{Name: "module", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "include_files", Type: field.TypeString, Nullable: true},
		{Name: "update_at", Type: field.TypeTime, Nullable: true},
		{Name: "update_by", Type: field.TypeUint32, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUint32, Nullable: true},
	}
	// APITable holds the schema information for the "api" table.
	APITable = &schema.Table{
		Name:       "api",
		Columns:    APIColumns,
		PrimaryKey: []*schema.Column{APIColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "api_url_type",
				Unique:  true,
				Columns: []*schema.Column{APIColumns[2], APIColumns[3]},
			},
		},
	}
	// APICategoryColumns holds the columns for the "api_category" table.
	APICategoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// APICategoryTable holds the schema information for the "api_category" table.
	APICategoryTable = &schema.Table{
		Name:       "api_category",
		Columns:    APICategoryColumns,
		PrimaryKey: []*schema.Column{APICategoryColumns[0]},
	}
	// APIHistoryColumns holds the columns for the "api_history" table.
	APIHistoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "version", Type: field.TypeInt32},
		{Name: "query_params", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "description", Type: field.TypeString},
	}
	// APIHistoryTable holds the schema information for the "api_history" table.
	APIHistoryTable = &schema.Table{
		Name:       "api_history",
		Columns:    APIHistoryColumns,
		PrimaryKey: []*schema.Column{APIHistoryColumns[0]},
	}
	// APIStatisticsColumns holds the columns for the "api_statistics" table.
	APIStatisticsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "call_count", Type: field.TypeInt32},
		{Name: "success_count", Type: field.TypeInt32},
		{Name: "failure_count", Type: field.TypeInt32},
		{Name: "avg_response_time", Type: field.TypeFloat64},
		{Name: "max_response_time", Type: field.TypeFloat64},
		{Name: "min_response_time", Type: field.TypeFloat64},
		{Name: "avg_traffic", Type: field.TypeFloat64},
		{Name: "max_traffic", Type: field.TypeFloat64},
		{Name: "min_traffic", Type: field.TypeFloat64},
		{Name: "description", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "api_id", Type: field.TypeInt32},
	}
	// APIStatisticsTable holds the schema information for the "api_statistics" table.
	APIStatisticsTable = &schema.Table{
		Name:       "api_statistics",
		Columns:    APIStatisticsColumns,
		PrimaryKey: []*schema.Column{APIStatisticsColumns[0]},
	}
	// APITagColumns holds the columns for the "api_tag" table.
	APITagColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "sort", Type: field.TypeInt32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "update_by", Type: field.TypeUint32},
		{Name: "description", Type: field.TypeString},
	}
	// APITagTable holds the schema information for the "api_tag" table.
	APITagTable = &schema.Table{
		Name:       "api_tag",
		Columns:    APITagColumns,
		PrimaryKey: []*schema.Column{APITagColumns[0]},
	}
	// ContainerColumns holds the columns for the "container" table.
	ContainerColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "hostname", Type: field.TypeString},
		{Name: "domainname", Type: field.TypeString},
		{Name: "user", Type: field.TypeString, Nullable: true},
		{Name: "env", Type: field.TypeJSON, Nullable: true},
		{Name: "cmd", Type: field.TypeJSON, Nullable: true},
		{Name: "image", Type: field.TypeString},
		{Name: "labels", Type: field.TypeJSON, Nullable: true},
		{Name: "volumes", Type: field.TypeJSON, Nullable: true},
		{Name: "working_dir", Type: field.TypeString, Nullable: true},
		{Name: "entrypoint", Type: field.TypeJSON, Nullable: true},
		{Name: "mac_address", Type: field.TypeString},
		{Name: "expose_ports", Type: field.TypeJSON, Nullable: true},
		{Name: "compose_file_url", Type: field.TypeString, Nullable: true},
		{Name: "dockerfile_url", Type: field.TypeString, Nullable: true},
	}
	// ContainerTable holds the schema information for the "container" table.
	ContainerTable = &schema.Table{
		Name:       "container",
		Columns:    ContainerColumns,
		PrimaryKey: []*schema.Column{ContainerColumns[0]},
	}
	// GroupColumns holds the columns for the "group" table.
	GroupColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_by", Type: field.TypeUint32, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUint32, Nullable: true},
		{Name: "headcount", Type: field.TypeInt32},
	}
	// GroupTable holds the schema information for the "group" table.
	GroupTable = &schema.Table{
		Name:       "group",
		Columns:    GroupColumns,
		PrimaryKey: []*schema.Column{GroupColumns[0]},
	}
	// GroupMemberColumns holds the columns for the "group_member" table.
	GroupMemberColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "group_id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeUint32},
		{Name: "role", Type: field.TypeUint8, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUint32, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUint32, Nullable: true},
	}
	// GroupMemberTable holds the schema information for the "group_member" table.
	GroupMemberTable = &schema.Table{
		Name:       "group_member",
		Columns:    GroupMemberColumns,
		PrimaryKey: []*schema.Column{GroupMemberColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "groupmember_group_id_user_id",
				Unique:  true,
				Columns: []*schema.Column{GroupMemberColumns[1], GroupMemberColumns[2]},
			},
		},
	}
	// JobColumns holds the columns for the "job" table.
	JobColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "payload", Type: field.TypeBytes, Nullable: true},
		{Name: "type", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "worker", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUint32, Nullable: true},
		{Name: "entry_id", Type: field.TypeString, Nullable: true},
		{Name: "config", Type: field.TypeBytes, Nullable: true},
		{Name: "task_id", Type: field.TypeInt64},
		{Name: "active", Type: field.TypeBool, Default: false},
	}
	// JobTable holds the schema information for the "job" table.
	JobTable = &schema.Table{
		Name:       "job",
		Columns:    JobColumns,
		PrimaryKey: []*schema.Column{JobColumns[0]},
	}
	// MetaEventColumns holds the columns for the "meta_event" table.
	MetaEventColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "event_name", Type: field.TypeString, Unique: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "display_name", Type: field.TypeString, Unique: true},
		{Name: "event_desc", Type: field.TypeString, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// MetaEventTable holds the schema information for the "meta_event" table.
	MetaEventTable = &schema.Table{
		Name:       "meta_event",
		Columns:    MetaEventColumns,
		PrimaryKey: []*schema.Column{MetaEventColumns[0]},
	}
	// ProjectColumns holds the columns for the "project" table.
	ProjectColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "identifier", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_by", Type: field.TypeUint32, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUint32, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 0},
		{Name: "start_time", Type: field.TypeTime, Nullable: true},
		{Name: "deadline", Type: field.TypeTime, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "remark", Type: field.TypeString, Nullable: true, Size: 2147483647},
	}
	// ProjectTable holds the schema information for the "project" table.
	ProjectTable = &schema.Table{
		Name:       "project",
		Columns:    ProjectColumns,
		PrimaryKey: []*schema.Column{ProjectColumns[0]},
	}
	// ProjectMemberColumns holds the columns for the "project_member" table.
	ProjectMemberColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "project_id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeUint32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUint32, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 0},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeUint8, Default: 0},
	}
	// ProjectMemberTable holds the schema information for the "project_member" table.
	ProjectMemberTable = &schema.Table{
		Name:       "project_member",
		Columns:    ProjectMemberColumns,
		PrimaryKey: []*schema.Column{ProjectMemberColumns[0]},
	}
	// TaskColumns holds the columns for the "task" table.
	TaskColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "assignee", Type: field.TypeUint32, Nullable: true},
		{Name: "type", Type: field.TypeInt8, Default: 0},
		{Name: "frequency", Type: field.TypeInt8, Nullable: true},
		{Name: "schedule_time", Type: field.TypeTime, Nullable: true},
		{Name: "rank", Type: field.TypeInt8, Default: 0},
		{Name: "status", Type: field.TypeInt8, Default: 0},
		{Name: "start_time", Type: field.TypeTime, Nullable: true},
		{Name: "completed_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_by", Type: field.TypeUint32, Nullable: true},
		{Name: "status_updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deadline", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUint32, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "testplan_id", Type: field.TypeInt64, Nullable: true},
		{Name: "testcase_suite", Type: field.TypeJSON, Nullable: true},
	}
	// TaskTable holds the schema information for the "task" table.
	TaskTable = &schema.Table{
		Name:       "task",
		Columns:    TaskColumns,
		PrimaryKey: []*schema.Column{TaskColumns[0]},
	}
	// TestPlanColumns holds the columns for the "test_plan" table.
	TestPlanColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_by", Type: field.TypeUint32, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "start_time", Type: field.TypeTime, Nullable: true},
		{Name: "deadline", Type: field.TypeTime, Nullable: true},
		{Name: "status_updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 0},
		{Name: "tasks", Type: field.TypeJSON, Nullable: true},
		{Name: "project_id", Type: field.TypeInt32},
	}
	// TestPlanTable holds the schema information for the "test_plan" table.
	TestPlanTable = &schema.Table{
		Name:       "test_plan",
		Columns:    TestPlanColumns,
		PrimaryKey: []*schema.Column{TestPlanColumns[0]},
	}
	// TestcaseColumns holds the columns for the "testcase" table.
	TestcaseColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeUint32, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 0},
		{Name: "type", Type: field.TypeInt8, Default: 0},
		{Name: "priority", Type: field.TypeInt8, Default: 0},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUint32, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "label", Type: field.TypeString, Nullable: true},
		{Name: "url", Type: field.TypeString, Nullable: true},
	}
	// TestcaseTable holds the schema information for the "testcase" table.
	TestcaseTable = &schema.Table{
		Name:       "testcase",
		Columns:    TestcaseColumns,
		PrimaryKey: []*schema.Column{TestcaseColumns[0]},
	}
	// TestcaseSuiteColumns holds the columns for the "testcase_suite" table.
	TestcaseSuiteColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_by", Type: field.TypeUint32, Nullable: true},
		{Name: "testcases", Type: field.TypeJSON, Nullable: true},
	}
	// TestcaseSuiteTable holds the schema information for the "testcase_suite" table.
	TestcaseSuiteTable = &schema.Table{
		Name:       "testcase_suite",
		Columns:    TestcaseSuiteColumns,
		PrimaryKey: []*schema.Column{TestcaseSuiteColumns[0]},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "chinese_name", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "location", Type: field.TypeString, Nullable: true},
		{Name: "last_login_time", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUint32, Nullable: true},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		APITable,
		APICategoryTable,
		APIHistoryTable,
		APIStatisticsTable,
		APITagTable,
		ContainerTable,
		GroupTable,
		GroupMemberTable,
		JobTable,
		MetaEventTable,
		ProjectTable,
		ProjectMemberTable,
		TaskTable,
		TestPlanTable,
		TestcaseTable,
		TestcaseSuiteTable,
		UserTable,
	}
)

func init() {
	APITable.Annotation = &entsql.Annotation{
		Table: "api",
	}
	APICategoryTable.Annotation = &entsql.Annotation{
		Table: "api_category",
	}
	APIHistoryTable.Annotation = &entsql.Annotation{
		Table: "api_history",
	}
	APIStatisticsTable.Annotation = &entsql.Annotation{
		Table: "api_statistics",
	}
	APITagTable.Annotation = &entsql.Annotation{
		Table: "api_tag",
	}
	ContainerTable.Annotation = &entsql.Annotation{
		Table: "container",
	}
	GroupTable.Annotation = &entsql.Annotation{
		Table: "group",
	}
	GroupMemberTable.Annotation = &entsql.Annotation{
		Table: "group_member",
	}
	JobTable.Annotation = &entsql.Annotation{
		Table: "job",
	}
	MetaEventTable.Annotation = &entsql.Annotation{
		Table: "meta_event",
	}
	ProjectTable.Annotation = &entsql.Annotation{
		Table: "project",
	}
	ProjectMemberTable.Annotation = &entsql.Annotation{
		Table: "project_member",
	}
	TaskTable.Annotation = &entsql.Annotation{
		Table: "task",
	}
	TestPlanTable.Annotation = &entsql.Annotation{
		Table: "test_plan",
	}
	TestcaseTable.Annotation = &entsql.Annotation{
		Table: "testcase",
	}
	TestcaseSuiteTable.Annotation = &entsql.Annotation{
		Table: "testcase_suite",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
}
