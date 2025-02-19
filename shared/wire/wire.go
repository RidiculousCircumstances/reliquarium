//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"reliquarium/config/env"
	"reliquarium/shared/lifecycle"
)

var LifecycleSet = wire.NewSet(
	lifecycle.NewAppLifecycle,
	ProvideLifecycleHooks,
)

var InfraSet = wire.NewSet(
	ProvideMessageBroker,
	ProvideStorageHelper,
	ProvideStorageClient,
	ProvideShutdownRegistry,
	ProvideCronTaskConfigs,
	ProvideServiceCommandClient,
)

var ApplicationSet = wire.NewSet(
	env.LoadConfig,
	ProvideLogger,
	ProvideApplication,
	ProvideHandlerConfigs,
	ProvideCreateSnapshot,
	ProvideCommandPublisher,
	ProvideOrchestrateServices,
)

func InitializeApp() (Application, error) {
	wire.Build(
		ApplicationSet,
		InfraSet,
		LifecycleSet,
	)

	return Application{}, nil
}
