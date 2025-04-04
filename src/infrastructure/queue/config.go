package queue

import (
	"github.com/RidiculousCircumstances/netherway/v2"
	"reliquarium/config/env"
	loggerIntf "reliquarium/shared/logger/interfaces"
	"reliquarium/src/application/consumer"
	consumerIntf "reliquarium/src/application/consumer/interfaces"
)

type HandlerConfig struct {
	QueueName    string
	ExchangeName string
	Handler      netherway.MessageHandler
}

func BuildHandlerConfigs(
	env *env.Config,
	logger loggerIntf.Logger,
	orchestrateServices consumerIntf.OrchestrateServices,
) []HandlerConfig {
	return []HandlerConfig{
		{
			QueueName:    env.ReliquariumConfirmationQueueName,
			ExchangeName: env.ReliquariumConfirmationExchangeName,
			Handler:      consumer.NewOrchestrateSnapshotMessageHandler(logger, orchestrateServices),
		},
	}
}
