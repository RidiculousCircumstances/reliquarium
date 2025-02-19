package message

import (
	"reliquarium/src/domain"
	domainIntf "reliquarium/src/domain/interfaces"
)

type CommandConfirmationEvent struct {
	ServiceName      domain.ServiceName     `json:"service_name"`
	Command          domainIntf.CommandType `json:"command"`
	TakeSnapshotData *TakeSnapshotPayload   `json:"takeSnapshotData,omitempty"`
}
type TakeSnapshotPayload struct {
	SnapshotStorageKey string `json:"snapshotStorageKey,omitempty"`
}

func NewCommandConfirmationEvent(serviceName domain.ServiceName, command domainIntf.CommandType, snapshotStorageKey string) *CommandConfirmationEvent {
	event := &CommandConfirmationEvent{
		ServiceName: serviceName,
		Command:     command,
	}

	if command == domainIntf.TakeSnapshots {
		event.TakeSnapshotData = &TakeSnapshotPayload{SnapshotStorageKey: snapshotStorageKey}
	}

	return event
}
