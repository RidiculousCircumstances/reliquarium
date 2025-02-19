package consumer

import (
	"context"
	"encoding/json"
	loggerIntf "reliquarium/shared/logger/interfaces"
	consumerIntf "reliquarium/src/application/consumer/interfaces"
	"reliquarium/src/infrastructure/queue/message"
)

type SnapshotTriggerMessageHandler struct {
	logger              loggerIntf.Logger
	orchestrateServices consumerIntf.OrchestrateServices
}

func NewSnapshotTriggerMessageHandler(
	logger loggerIntf.Logger,
	orchestrateServices consumerIntf.OrchestrateServices,
) *SnapshotTriggerMessageHandler {
	return &SnapshotTriggerMessageHandler{
		logger:              logger,
		orchestrateServices: orchestrateServices,
	}
}

func (mh *SnapshotTriggerMessageHandler) Handle(ctx context.Context, msg []byte) (bool, error) {
	var reliquariumCommand message.CommandConfirmationEvent
	err := json.Unmarshal(msg, &reliquariumCommand)
	if err != nil {
		mh.logger.Error("[SnapshotTriggerMessageHandler] failed unmarshalling message")
		return false, err
	}

	// Вызываем use case
	err = mh.orchestrateServices.Process(ctx, &reliquariumCommand)
	if err != nil {
		return false, err
	}

	return true, nil
}
