package service

import (
	"github.com/RidiculousCircumstances/netherway/v2"
	"golang.org/x/net/context"
	"reliquarium/src/infrastructure/queue/message"
)

type Client struct {
	broker   netherway.MessageBroker
	exchange string
}

func NewClient(broker netherway.MessageBroker) *Client {
	return &Client{
		broker:   broker,
		exchange: "reliquarium_command_exchange",
	}
}

func (c *Client) SendCommand(ctx context.Context, cmd *message.ServiceControlCommand) error {
	return c.broker.Publish(ctx, c.exchange, "", cmd)
}
