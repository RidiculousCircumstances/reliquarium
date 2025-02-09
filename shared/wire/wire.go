//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"reliquarium/config/env"
	"reliquarium/shared/lifecycle"
	"reliquarium/src/infrastructure/metrics"
)

var LifecycleSet = wire.NewSet(
	lifecycle.NewAppLifecycle,
	ProvideLifecycleHooks,
)

var InfraSet = wire.NewSet(
	ProvideMessageBroker,
	NewRedisClient,
	ProvideSearchAggregateRepo,
	ProvideEdenClient,
	ProvideTraceFaceClient,
	ProvideStorageHelper,
	ProvideStorageClient,
	ProvideRoutes,
	ProvideTgRoutes,
	ProvideTgBotApi,
	ProvideTelegramClient,
	metrics.NewPrometheusHook,
	ProvideMetricsConfig,
	ProvideMetricRegistry,
	ProvideMetricsProvider,
)

var ApplicationSet = wire.NewSet(
	env.LoadConfig,
	ProvideLogger,
	ProvideApplication,
	ProvideHandlerConfigs,
	ProvideAggregateService,
	ProvideSearchProfilesPublisher,
	ProvideSearchProfilesByPhotos,
	ProvideSearchFacesPublisher,
	ProvideSearchPhotosByFaces,
	ProvideSearchResultMapper,
	ProvideGetSearchResult,
	ProvideAggregateResult,
	ProvideStorageService,
	ProvideUploadFileAction,
	ProvideInitSearchAction,
	ProvideGetSearchResultAction,
	ProvideUploadFile,
	ProvideSearchAction,
)

func InitializeApp() (Application, error) {
	wire.Build(
		ApplicationSet,
		InfraSet,
		LifecycleSet,
	)

	return Application{}, nil
}
