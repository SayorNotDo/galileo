// Code generated by ent, DO NOT EDIT.

package ent

import (
	"galileo/ent/api"
	"galileo/ent/apistatistics"
	"galileo/ent/container"
	"galileo/ent/group"
	"galileo/ent/groupmember"
	"galileo/ent/job"
	"galileo/ent/project"
	"galileo/ent/projectmember"
	"galileo/ent/schema"
	"galileo/ent/task"
	"galileo/ent/testcase"
	"galileo/ent/testcasesuite"
	"galileo/ent/testplan"
	"galileo/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	apiFields := schema.Api{}.Fields()
	_ = apiFields
	// apiDescName is the schema descriptor for name field.
	apiDescName := apiFields[1].Descriptor()
	// api.NameValidator is a validator for the "name" field. It is called by the builders before save.
	api.NameValidator = apiDescName.Validators[0].(func(string) error)
	// apiDescURL is the schema descriptor for url field.
	apiDescURL := apiFields[2].Descriptor()
	// api.URLValidator is a validator for the "url" field. It is called by the builders before save.
	api.URLValidator = apiDescURL.Validators[0].(func(string) error)
	// apiDescType is the schema descriptor for type field.
	apiDescType := apiFields[3].Descriptor()
	// api.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	api.TypeValidator = apiDescType.Validators[0].(func(int8) error)
	// apiDescStatus is the schema descriptor for status field.
	apiDescStatus := apiFields[4].Descriptor()
	// api.DefaultStatus holds the default value on creation for the status field.
	api.DefaultStatus = apiDescStatus.Default.(int8)
	// apiDescCreatedAt is the schema descriptor for created_at field.
	apiDescCreatedAt := apiFields[12].Descriptor()
	// api.DefaultCreatedAt holds the default value on creation for the created_at field.
	api.DefaultCreatedAt = apiDescCreatedAt.Default.(func() time.Time)
	// apiDescUpdateAt is the schema descriptor for update_at field.
	apiDescUpdateAt := apiFields[15].Descriptor()
	// api.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	api.UpdateDefaultUpdateAt = apiDescUpdateAt.UpdateDefault.(func() time.Time)
	apistatisticsFields := schema.ApiStatistics{}.Fields()
	_ = apistatisticsFields
	// apistatisticsDescAPIID is the schema descriptor for api_id field.
	apistatisticsDescAPIID := apistatisticsFields[13].Descriptor()
	// apistatistics.APIIDValidator is a validator for the "api_id" field. It is called by the builders before save.
	apistatistics.APIIDValidator = apistatisticsDescAPIID.Validators[0].(func(int32) error)
	containerFields := schema.Container{}.Fields()
	_ = containerFields
	// containerDescHostname is the schema descriptor for hostname field.
	containerDescHostname := containerFields[1].Descriptor()
	// container.HostnameValidator is a validator for the "hostname" field. It is called by the builders before save.
	container.HostnameValidator = containerDescHostname.Validators[0].(func(string) error)
	// containerDescDomainname is the schema descriptor for domainname field.
	containerDescDomainname := containerFields[2].Descriptor()
	// container.DomainnameValidator is a validator for the "domainname" field. It is called by the builders before save.
	container.DomainnameValidator = containerDescDomainname.Validators[0].(func(string) error)
	// containerDescImage is the schema descriptor for image field.
	containerDescImage := containerFields[6].Descriptor()
	// container.ImageValidator is a validator for the "image" field. It is called by the builders before save.
	container.ImageValidator = containerDescImage.Validators[0].(func(string) error)
	// containerDescMACAddress is the schema descriptor for mac_address field.
	containerDescMACAddress := containerFields[11].Descriptor()
	// container.MACAddressValidator is a validator for the "mac_address" field. It is called by the builders before save.
	container.MACAddressValidator = containerDescMACAddress.Validators[0].(func(string) error)
	// containerDescID is the schema descriptor for id field.
	containerDescID := containerFields[0].Descriptor()
	// container.IDValidator is a validator for the "id" field. It is called by the builders before save.
	container.IDValidator = containerDescID.Validators[0].(func(string) error)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[1].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	// groupDescCreatedAt is the schema descriptor for created_at field.
	groupDescCreatedAt := groupFields[5].Descriptor()
	// group.DefaultCreatedAt holds the default value on creation for the created_at field.
	group.DefaultCreatedAt = groupDescCreatedAt.Default.(func() time.Time)
	// groupDescUpdatedAt is the schema descriptor for updated_at field.
	groupDescUpdatedAt := groupFields[6].Descriptor()
	// group.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	group.UpdateDefaultUpdatedAt = groupDescUpdatedAt.UpdateDefault.(func() time.Time)
	groupmemberFields := schema.GroupMember{}.Fields()
	_ = groupmemberFields
	// groupmemberDescRole is the schema descriptor for role field.
	groupmemberDescRole := groupmemberFields[3].Descriptor()
	// groupmember.DefaultRole holds the default value on creation for the role field.
	groupmember.DefaultRole = groupmemberDescRole.Default.(uint8)
	// groupmemberDescCreatedAt is the schema descriptor for created_at field.
	groupmemberDescCreatedAt := groupmemberFields[4].Descriptor()
	// groupmember.DefaultCreatedAt holds the default value on creation for the created_at field.
	groupmember.DefaultCreatedAt = groupmemberDescCreatedAt.Default.(func() time.Time)
	jobFields := schema.Job{}.Fields()
	_ = jobFields
	// jobDescCreatedAt is the schema descriptor for created_at field.
	jobDescCreatedAt := jobFields[1].Descriptor()
	// job.DefaultCreatedAt holds the default value on creation for the created_at field.
	job.DefaultCreatedAt = jobDescCreatedAt.Default.(func() time.Time)
	// jobDescUpdatedAt is the schema descriptor for updated_at field.
	jobDescUpdatedAt := jobFields[3].Descriptor()
	// job.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	job.UpdateDefaultUpdatedAt = jobDescUpdatedAt.UpdateDefault.(func() time.Time)
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescName is the schema descriptor for name field.
	projectDescName := projectFields[1].Descriptor()
	// project.NameValidator is a validator for the "name" field. It is called by the builders before save.
	project.NameValidator = projectDescName.Validators[0].(func(string) error)
	// projectDescIdentifier is the schema descriptor for identifier field.
	projectDescIdentifier := projectFields[2].Descriptor()
	// project.IdentifierValidator is a validator for the "identifier" field. It is called by the builders before save.
	project.IdentifierValidator = projectDescIdentifier.Validators[0].(func(string) error)
	// projectDescCreatedAt is the schema descriptor for created_at field.
	projectDescCreatedAt := projectFields[3].Descriptor()
	// project.DefaultCreatedAt holds the default value on creation for the created_at field.
	project.DefaultCreatedAt = projectDescCreatedAt.Default.(func() time.Time)
	// projectDescUpdatedAt is the schema descriptor for updated_at field.
	projectDescUpdatedAt := projectFields[5].Descriptor()
	// project.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	project.UpdateDefaultUpdatedAt = projectDescUpdatedAt.UpdateDefault.(func() time.Time)
	// projectDescStatus is the schema descriptor for status field.
	projectDescStatus := projectFields[9].Descriptor()
	// project.DefaultStatus holds the default value on creation for the status field.
	project.DefaultStatus = projectDescStatus.Default.(int8)
	projectmemberFields := schema.ProjectMember{}.Fields()
	_ = projectmemberFields
	// projectmemberDescCreatedAt is the schema descriptor for created_at field.
	projectmemberDescCreatedAt := projectmemberFields[2].Descriptor()
	// projectmember.DefaultCreatedAt holds the default value on creation for the created_at field.
	projectmember.DefaultCreatedAt = projectmemberDescCreatedAt.Default.(func() time.Time)
	// projectmemberDescStatus is the schema descriptor for status field.
	projectmemberDescStatus := projectmemberFields[6].Descriptor()
	// projectmember.DefaultStatus holds the default value on creation for the status field.
	projectmember.DefaultStatus = projectmemberDescStatus.Default.(int8)
	// projectmemberDescRole is the schema descriptor for role field.
	projectmemberDescRole := projectmemberFields[9].Descriptor()
	// projectmember.DefaultRole holds the default value on creation for the role field.
	projectmember.DefaultRole = projectmemberDescRole.Default.(uint8)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescName is the schema descriptor for name field.
	taskDescName := taskFields[1].Descriptor()
	// task.NameValidator is a validator for the "name" field. It is called by the builders before save.
	task.NameValidator = taskDescName.Validators[0].(func(string) error)
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskFields[2].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescType is the schema descriptor for type field.
	taskDescType := taskFields[5].Descriptor()
	// task.DefaultType holds the default value on creation for the type field.
	task.DefaultType = taskDescType.Default.(int8)
	// taskDescRank is the schema descriptor for rank field.
	taskDescRank := taskFields[8].Descriptor()
	// task.DefaultRank holds the default value on creation for the rank field.
	task.DefaultRank = taskDescRank.Default.(int8)
	// taskDescStatus is the schema descriptor for status field.
	taskDescStatus := taskFields[9].Descriptor()
	// task.DefaultStatus holds the default value on creation for the status field.
	task.DefaultStatus = taskDescStatus.Default.(int8)
	// taskDescUpdatedAt is the schema descriptor for updated_at field.
	taskDescUpdatedAt := taskFields[12].Descriptor()
	// task.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	task.UpdateDefaultUpdatedAt = taskDescUpdatedAt.UpdateDefault.(func() time.Time)
	testplanFields := schema.TestPlan{}.Fields()
	_ = testplanFields
	// testplanDescName is the schema descriptor for name field.
	testplanDescName := testplanFields[1].Descriptor()
	// testplan.NameValidator is a validator for the "name" field. It is called by the builders before save.
	testplan.NameValidator = testplanDescName.Validators[0].(func(string) error)
	// testplanDescCreatedAt is the schema descriptor for created_at field.
	testplanDescCreatedAt := testplanFields[2].Descriptor()
	// testplan.DefaultCreatedAt holds the default value on creation for the created_at field.
	testplan.DefaultCreatedAt = testplanDescCreatedAt.Default.(func() time.Time)
	// testplanDescUpdatedAt is the schema descriptor for updated_at field.
	testplanDescUpdatedAt := testplanFields[4].Descriptor()
	// testplan.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	testplan.UpdateDefaultUpdatedAt = testplanDescUpdatedAt.UpdateDefault.(func() time.Time)
	// testplanDescStatus is the schema descriptor for status field.
	testplanDescStatus := testplanFields[10].Descriptor()
	// testplan.DefaultStatus holds the default value on creation for the status field.
	testplan.DefaultStatus = testplanDescStatus.Default.(int8)
	// testplanDescProjectID is the schema descriptor for project_id field.
	testplanDescProjectID := testplanFields[12].Descriptor()
	// testplan.ProjectIDValidator is a validator for the "project_id" field. It is called by the builders before save.
	testplan.ProjectIDValidator = testplanDescProjectID.Validators[0].(func(int32) error)
	testcaseFields := schema.Testcase{}.Fields()
	_ = testcaseFields
	// testcaseDescName is the schema descriptor for name field.
	testcaseDescName := testcaseFields[1].Descriptor()
	// testcase.NameValidator is a validator for the "name" field. It is called by the builders before save.
	testcase.NameValidator = testcaseDescName.Validators[0].(func(string) error)
	// testcaseDescCreatedAt is the schema descriptor for created_at field.
	testcaseDescCreatedAt := testcaseFields[3].Descriptor()
	// testcase.DefaultCreatedAt holds the default value on creation for the created_at field.
	testcase.DefaultCreatedAt = testcaseDescCreatedAt.Default.(func() time.Time)
	// testcaseDescUpdatedAt is the schema descriptor for updated_at field.
	testcaseDescUpdatedAt := testcaseFields[5].Descriptor()
	// testcase.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	testcase.UpdateDefaultUpdatedAt = testcaseDescUpdatedAt.UpdateDefault.(func() time.Time)
	// testcaseDescStatus is the schema descriptor for status field.
	testcaseDescStatus := testcaseFields[6].Descriptor()
	// testcase.DefaultStatus holds the default value on creation for the status field.
	testcase.DefaultStatus = testcaseDescStatus.Default.(int8)
	// testcaseDescType is the schema descriptor for type field.
	testcaseDescType := testcaseFields[7].Descriptor()
	// testcase.DefaultType holds the default value on creation for the type field.
	testcase.DefaultType = testcaseDescType.Default.(int8)
	// testcaseDescPriority is the schema descriptor for priority field.
	testcaseDescPriority := testcaseFields[8].Descriptor()
	// testcase.DefaultPriority holds the default value on creation for the priority field.
	testcase.DefaultPriority = testcaseDescPriority.Default.(int8)
	testcasesuiteFields := schema.TestcaseSuite{}.Fields()
	_ = testcasesuiteFields
	// testcasesuiteDescName is the schema descriptor for name field.
	testcasesuiteDescName := testcasesuiteFields[1].Descriptor()
	// testcasesuite.NameValidator is a validator for the "name" field. It is called by the builders before save.
	testcasesuite.NameValidator = testcasesuiteDescName.Validators[0].(func(string) error)
	// testcasesuiteDescCreatedAt is the schema descriptor for created_at field.
	testcasesuiteDescCreatedAt := testcasesuiteFields[2].Descriptor()
	// testcasesuite.DefaultCreatedAt holds the default value on creation for the created_at field.
	testcasesuite.DefaultCreatedAt = testcasesuiteDescCreatedAt.Default.(func() time.Time)
	// testcasesuiteDescUpdatedAt is the schema descriptor for updated_at field.
	testcasesuiteDescUpdatedAt := testcasesuiteFields[4].Descriptor()
	// testcasesuite.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	testcasesuite.UpdateDefaultUpdatedAt = testcasesuiteDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUUID is the schema descriptor for uuid field.
	userDescUUID := userFields[1].Descriptor()
	// user.DefaultUUID holds the default value on creation for the uuid field.
	user.DefaultUUID = userDescUUID.Default.(func() uuid.UUID)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[2].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescActive is the schema descriptor for active field.
	userDescActive := userFields[3].Descriptor()
	// user.DefaultActive holds the default value on creation for the active field.
	user.DefaultActive = userDescActive.Default.(bool)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[12].Descriptor()
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
