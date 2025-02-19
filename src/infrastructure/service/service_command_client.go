package service

import (
	"golang.org/x/net/context"
	brokerIntf "reliquarium/shared/broker/interfaces"
	"reliquarium/src/infrastructure/queue/message"
)

type Client struct {
	broker   brokerIntf.MessageBroker
	exchange string
}

func NewClient(broker brokerIntf.MessageBroker) *Client {
	return &Client{
		broker:   broker,
		exchange: "reliquarium_command_exchange",
	}
}

func (c *Client) SendCommand(ctx context.Context, cmd *message.ServiceControlCommand) error {
	return c.broker.Publish(ctx, c.exchange, "", cmd)
}
