package interfaces

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewRunnerInterface)
