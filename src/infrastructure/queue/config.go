package queue

import (
	"reliquarium/config/env"
	"reliquarium/shared/broker/interfaces"
	loggerIntf "reliquarium/shared/logger/interfaces"
	"reliquarium/src/application/consumer"
	consumerIntf "reliquarium/src/application/consumer/interfaces"
)

type HandlerConfig struct {
	QueueName    string
	ExchangeName string
	Handler      interfaces.MessageHandler
}

func BuildHandlerConfigs(
	env *env.Config,
	logger loggerIntf.Logger,
	createSnapshot consumerIntf.CreateSnapshot,
) []HandlerConfig {
	return []HandlerConfig{
		{
			QueueName:    env.ReliquariumConfirmationQueueName,
			ExchangeName: env.ReliquariumConfirmationExchangeName,
			Handler:      consumer.NewSnapshotTriggerMessageHandler(logger, createSnapshot),
		},
	}
}
