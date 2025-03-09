package consumer

import (
	"context"
	"encoding/json"
	loggerIntf "reliquarium/shared/logger/interfaces"
	consumerIntf "reliquarium/src/application/consumer/interfaces"
	"reliquarium/src/infrastructure/queue/message"
)

type OrchestrateSnapshotMessageHandler struct {
	logger              loggerIntf.Logger
	orchestrateServices consumerIntf.OrchestrateServices
}

func NewOrchestrateSnapshotMessageHandler(
	logger loggerIntf.Logger,
	orchestrateServices consumerIntf.OrchestrateServices,
) *OrchestrateSnapshotMessageHandler {
	return &OrchestrateSnapshotMessageHandler{
		logger:              logger,
		orchestrateServices: orchestrateServices,
	}
}

func (mh *OrchestrateSnapshotMessageHandler) Handle(ctx context.Context, msg []byte) (bool, error) {
	var reliquariumCommand message.CommandConfirmationEvent
	err := json.Unmarshal(msg, &reliquariumCommand)
	if err != nil {
		mh.logger.Error("[OrchestrateSnapshotMessageHandler] failed unmarshalling message")
		return false, err
	}

	// Вызываем use case
	err = mh.orchestrateServices.Process(ctx, &reliquariumCommand)
	if err != nil {
		return false, err
	}

	return true, nil
}
