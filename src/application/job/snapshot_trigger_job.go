package job

import (
	"context"
	"reliquarium/src/application/interfaces"
	domainIntf "reliquarium/src/domain/interfaces"
	"reliquarium/src/infrastructure/queue/message"
)

type SnapshotTriggerJob struct {
	pauseTrigger interfaces.CommandPublisher
}

func NewSnapshotTriggerJob(shutdownTrigger interfaces.CommandPublisher) *SnapshotTriggerJob {
	return &SnapshotTriggerJob{pauseTrigger: shutdownTrigger}
}

func (j *SnapshotTriggerJob) Run(ctx context.Context) error {
	return j.pauseTrigger.Publish(ctx, message.NewServiceControlCommand(domainIntf.Pause))
}
