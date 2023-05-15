package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewProjectService,
	NewTaskService,
	NewTestcaseService,
	NewApiService,
	NewManagementService,
)
