package service

import "github.com/google/wire"

// ProviderSet is scheduler providers.
var ProviderSet = wire.NewSet(
	NewProjectService,
	NewTaskService,
	NewTestcaseService,
	NewApiService,
	NewManagementService,
)
