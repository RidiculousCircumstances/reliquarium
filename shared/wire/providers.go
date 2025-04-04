package wire

import (
	"github.com/RidiculousCircumstances/netherway/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"reliquarium/config/env"
	lifecycleIntf "reliquarium/shared/lifecycle/interfaces"
	"reliquarium/shared/logger"
	loggerIntf "reliquarium/shared/logger/interfaces"
	consumerIntf "reliquarium/src/application/consumer/interfaces"
	appIntf "reliquarium/src/application/interfaces"
	"reliquarium/src/application/publisher"
	"reliquarium/src/application/publisher/interfaces"
	"reliquarium/src/application/service/storage"
	storageIntf "reliquarium/src/application/service/storage/interfaces"
	"reliquarium/src/application/usecase"
	domainIntf "reliquarium/src/domain/interfaces"
	"reliquarium/src/infrastructure/cron"
	"reliquarium/src/infrastructure/queue"
	"reliquarium/src/infrastructure/service"
	infraStorageIntf "reliquarium/src/infrastructure/storage"
)

func ProvideLogger(config *env.Config) (loggerIntf.Logger, error) {
	return logger.NewZapLogger(config)
}

func ProvideLifecycleHooks(
	env *env.Config,
	handlerCfgs []queue.HandlerConfig,
	logger loggerIntf.Logger,
	broker netherway.MessageBroker,
	cronTaskConfigs []cron.CronTaskConfig,
) []lifecycleIntf.Hook {
	return []lifecycleIntf.Hook{
		queue.NewConsumerHook(handlerCfgs, logger, broker),
		cron.NewCronHook(logger, cronTaskConfigs),
	}
}

func ProvideMessageBroker(cfg *env.Config, logger loggerIntf.Logger) netherway.MessageBroker {
	pubFactory := func(conn netherway.Connection) netherway.Publisher {
		return netherway.NewPublisher(conn, logger)
	}

	subFactory := func(conn netherway.Connection) netherway.Subscriber {
		return netherway.NewSubscriber(conn, logger)
	}

	connFactory := netherway.NewConnFactory(netherway.ConnConfig{
		AmqpURI:                  cfg.RabbitMQURL,
		PublisherChannelPoolSize: 10,
		Exchanges: []netherway.Exchange{
			{
				Name: cfg.ReliquariumCommandExchangeName,
				Type: "fanout",
			},
		},
	})
	return netherway.NewMessageBroker(netherway.Config{
		PublisherFactory:  pubFactory,
		SubscriberFactory: subFactory,
		Logger:            logger,
		ConnFactory:       connFactory,
	})
}

func ProvideStorageHelper(cfg *env.Config) *storage.Helper {
	return storage.NewHelper(cfg.SnapshotBucket)
}

func ProvideCreateSnapshot(logger loggerIntf.Logger) consumerIntf.HandleSnapshot {
	return usecase.NewHandleSnapshot(logger)
}

func ProvideOrchestrateServices(
	logger loggerIntf.Logger,
	registry domainIntf.ServiceRegistry,
	publisher appIntf.CommandPublisher,
	handleSnapshot consumerIntf.HandleSnapshot,
) consumerIntf.OrchestrateServices {
	return usecase.NewOrchestrateServices(logger, registry, publisher, handleSnapshot)
}

func ProvideHandlerConfigs(
	cfg *env.Config,
	logger loggerIntf.Logger,
	orchestrateServices consumerIntf.OrchestrateServices,
) []queue.HandlerConfig {
	return queue.BuildHandlerConfigs(cfg, logger, orchestrateServices)
}

func ProvideCronTaskConfigs(
	cfg *env.Config,
	logger loggerIntf.Logger,
	commandPublisher appIntf.CommandPublisher,
) []cron.CronTaskConfig {
	return cron.BuildCronTaskConfigs(cfg, logger, commandPublisher)
}

func ProvideServiceCommandClient(broker netherway.MessageBroker) interfaces.CommandClient {
	return service.NewClient(broker)
}

func ProvideCommandPublisher(client interfaces.CommandClient) appIntf.CommandPublisher {
	return publisher.NewCommandPublisher(client)
}

func ProvideStorageClient(cfg *env.Config) (storageIntf.StorageClient, error) {
	baseMinioClient, err := minio.New(cfg.StorageEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.StorageAccessKeyId, cfg.StorageSecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	minioClient, err := infraStorageIntf.NewMinioClient(baseMinioClient)

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

//func ProvideStorageService(client storageIntf.StorageClient) actionIntf.StorageService {
//	return storage.NewService(client)
//}

func ProvideShutdownRegistry() domainIntf.ServiceRegistry {
	return service.NewRegistry([]domainIntf.ServiceName{
		domainIntf.Eden,
		//domainIntf.TraceFace,
		//domainIntf.StreamMiner,
	})
}
