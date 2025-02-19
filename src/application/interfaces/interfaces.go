package interfaces

import (
	"context"
	"reliquarium/src/infrastructure/queue/message"
)

type CommandPublisher interface {
	Publish(ctx context.Context, profile *message.ServiceControlCommand) error
}
