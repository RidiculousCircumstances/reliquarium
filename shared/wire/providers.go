package wire

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"reliquarium/config/env"
	brokerLib "reliquarium/shared/broker"
	"reliquarium/shared/broker/amqp"
	brokerIntf "reliquarium/shared/broker/interfaces"
	lifecycleIntf "reliquarium/shared/lifecycle/interfaces"
	"reliquarium/shared/logger"
	loggerIntf "reliquarium/shared/logger/interfaces"
	"reliquarium/src/application/consumer"
	"reliquarium/src/application/service/storage"
	storageIntf "reliquarium/src/application/service/storage/interfaces"
	"reliquarium/src/infrastructure/queue"
	infraStorageIntf "reliquarium/src/infrastructure/storage"
	"time"
)

func ProvideLogger(config *env.Config) (loggerIntf.Logger, error) {
	return logger.NewZapLogger(config)
}

func ProvideLifecycleHooks(
	env *env.Config,
	handlerCfgs []queue.HandlerConfig,
	logger loggerIntf.Logger,
	broker brokerIntf.MessageBroker,
) []lifecycleIntf.Hook {
	return []lifecycleIntf.Hook{
		queue.NewConsumerHook(handlerCfgs, logger, broker),
	}
}

func ProvideMessageBroker(cfg *env.Config, logger loggerIntf.Logger) brokerIntf.MessageBroker {
	pubFactory := func(conn brokerIntf.Connection) brokerIntf.Publisher {
		return brokerLib.NewPublisher(conn, logger)
	}

	subFactory := func(conn brokerIntf.Connection) brokerIntf.Subscriber {
		return brokerLib.NewSubscriber(conn, logger)
	}

	connFactory := amqp.NewConnFactory(amqp.ConnConfig{
		AmqpURI:                   cfg.RabbitMQURL,
		PublisherChannelPoolSize:  10,
		SubscriberChannelPoolSize: 10,
		Exchanges: []amqp.Exchange{
			{
				Name: cfg.ReliquariumCommandExchangeName,
				Type: "fanout",
			},
		},
	})

	return brokerLib.NewMessageBroker(brokerLib.Config{
		PublisherFactory:  pubFactory,
		SubscriberFactory: subFactory,
		Logger:            logger,
		ConnFactory:       connFactory,
	})
}

func NewRedisClient(cfg *env.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:               cfg.RedisAddr,     // Хост и порт Redis, например, "localhost:6379"
		Password:           cfg.RedisPassword, // Пароль для подключения (если есть)
		DB:                 cfg.RedisDB,       // Номер базы данных (по умолчанию 0)
		DialTimeout:        5 * time.Second,   // Таймаут для подключения
		ReadTimeout:        3 * time.Second,   // Таймаут чтения
		WriteTimeout:       3 * time.Second,   // Таймаут записи
		PoolSize:           10,                // Размер пула соединений
		MinIdleConns:       2,                 // Минимальное количество свободных соединений
		IdleTimeout:        5 * time.Minute,   // Таймаут для свободных соединений
		IdleCheckFrequency: 1 * time.Minute,   // Частота проверки свободных соединений
		TLSConfig:          nil,               // Настройки TLS (если требуется защищённое соединение)
	})
}

func ProvideStorageHelper(cfg *env.Config) *storage.Helper {
	return storage.NewHelper(cfg.SnapshotBucket)
}

func ProvideHandlerConfigs(
	cfg *env.Config,
	logger loggerIntf.Logger,
	snapshotTriggerHandler consumer.SnapshotTriggerMessageHandler,
) []queue.HandlerConfig {
	return queue.BuildHandlerConfigs(cfg, logger, snapshotTriggerHandler)
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

func ProvideStorageService(client storageIntf.StorageClient) actionIntf.StorageService {
	return storage.NewService(client)
}
