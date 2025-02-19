package interfaces

import (
	"reliquarium/src/domain"
)

type CommandType string

const (
	Pause         CommandType = "pause"
	TakeSnapshots CommandType = "takeSnapshots"
	Resume        CommandType = "resume"
)

type ServiceRegistry interface {
	ConfirmServiceCommand(serviceName domain.ServiceName, command CommandType) bool
	IsCommandConfirmed(command CommandType) bool
	GetUnconfirmedServices(command CommandType) []domain.ServiceName
	ResetServiceStates()
}
