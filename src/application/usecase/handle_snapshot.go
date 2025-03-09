package usecase

import (
	"context"
	loggerIntf "reliquarium/shared/logger/interfaces"
	"reliquarium/src/infrastructure/queue/message"
)

type HandleSnapshot struct {
	logger loggerIntf.Logger
}

func NewHandleSnapshot(logger loggerIntf.Logger) *HandleSnapshot {
	return &HandleSnapshot{logger: logger}
}

func (u *HandleSnapshot) Process(ctx context.Context, payload *message.TakeSnapshotPayload) error {
	u.logger.Info("Processing snapshot results")
	return nil
}
