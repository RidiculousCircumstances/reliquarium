package consumer

import (
	"context"
	"encoding/json"
	loggerIntf "reliquarium/shared/logger/interfaces"
	"reliquarium/src/application/consumer/interfaces"
	"reliquarium/src/infrastructure/queue/message"
)

type SnapshotTriggerMessageHandler struct {
	logger         loggerIntf.Logger
	createSnapshot interfaces.CreateSnapshot
}

func NewSnapshotTriggerMessageHandler(logger loggerIntf.Logger, createSnapshot interfaces.CreateSnapshot) *SnapshotTriggerMessageHandler {
	return &SnapshotTriggerMessageHandler{
		logger:         logger,
		createSnapshot: createSnapshot,
	}
}

func (mh *SnapshotTriggerMessageHandler) Handle(ctx context.Context, msg []byte) (bool, error) {
	var reliquariumCommand message.PauseConfirmationEvent
	err := json.Unmarshal(msg, &reliquariumCommand)
	if err != nil {
		mh.logger.Error("[SnapshotTriggerMessageHandler] failed unmarshalling message")
		return false, err
	}

	return true, nil
}
