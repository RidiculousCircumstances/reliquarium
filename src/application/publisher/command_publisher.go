package publisher

import (
	"context"
	"reliquarium/src/application/publisher/interfaces"
	"reliquarium/src/infrastructure/queue/message"
)

type CommandPublisher struct {
	client interfaces.CommandClient
}

func NewCommandPublisher(client interfaces.CommandClient) *CommandPublisher {
	return &CommandPublisher{client: client}
}

func (s *CommandPublisher) Publish(ctx context.Context, cmd *message.ServiceControlCommand) error {
	return s.client.SendCommand(ctx, cmd)
}
