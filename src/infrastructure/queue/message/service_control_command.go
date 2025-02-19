package message

import (
	domainIntf "reliquarium/src/domain/interfaces"
)

type ServiceControlCommand struct {
	Command domainIntf.CommandType `json:"command"`
}

func NewServiceControlCommand(Command domainIntf.CommandType) *ServiceControlCommand {
	return &ServiceControlCommand{
		Command: Command,
	}
}
