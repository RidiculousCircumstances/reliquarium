package interfaces

import (
	"context"
	"reliquarium/src/infrastructure/queue/message"
)

type HandleSnapshot interface {
	Process(ctx context.Context, payload *message.TakeSnapshotPayload) error
}

type OrchestrateServices interface {
	Process(ctx context.Context, event *message.CommandConfirmationEvent) error
}
