package job

import (
	"context"
	"reliquarium/src/application/interfaces"
	domainIntf "reliquarium/src/domain/interfaces"
	"reliquarium/src/infrastructure/queue/message"
)

type SnapshotTriggerJob struct {
	shutdownTrigger interfaces.CommandPublisher
}

func NewSnapshotTriggerJob(shutdownTrigger interfaces.CommandPublisher) *SnapshotTriggerJob {
	return &SnapshotTriggerJob{shutdownTrigger: shutdownTrigger}
}

func (j *SnapshotTriggerJob) Run(ctx context.Context) error {
	return j.shutdownTrigger.Publish(ctx, message.NewServiceControlCommand(domainIntf.Pause))
}
