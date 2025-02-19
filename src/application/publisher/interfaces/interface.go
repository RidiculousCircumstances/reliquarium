package interfaces

import (
	"context"
	"reliquarium/src/infrastructure/queue/message"
)

type CommandClient interface {
	SendCommand(ctx context.Context, cmd *message.ServiceControlCommand) error
}
