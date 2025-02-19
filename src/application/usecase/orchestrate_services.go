package usecase

import (
	"context"
	"errors"
	"go.uber.org/zap"
	loggerIntf "reliquarium/shared/logger/interfaces"
	consumerIntf "reliquarium/src/application/consumer/interfaces"
	appIntf "reliquarium/src/application/interfaces"
	domainIntf "reliquarium/src/domain/interfaces"
	"reliquarium/src/infrastructure/queue/message"
)

var (
	ErrUnknownService = errors.New("unknown service received")
)

type OrchestrateServices struct {
	logger         loggerIntf.Logger
	registry       domainIntf.ServiceRegistry
	publisher      appIntf.CommandPublisher
	handleSnapshot consumerIntf.HandleSnapshot
}

func NewOrchestrateServices(
	logger loggerIntf.Logger,
	registry domainIntf.ServiceRegistry,
	publisher appIntf.CommandPublisher,
	handleSnapshot consumerIntf.HandleSnapshot,
) *OrchestrateServices {
	return &OrchestrateServices{
		logger:         logger,
		registry:       registry,
		publisher:      publisher,
		handleSnapshot: handleSnapshot,
	}
}

// Основная логика обработки команд
func (uc *OrchestrateServices) Process(ctx context.Context, msg *message.CommandConfirmationEvent) error {
	// Подтверждаем команду
	success := uc.registry.ConfirmServiceCommand(msg.ServiceName, msg.Command)
	if !success {
		uc.logger.Error("[OrchestrateServices] Failed to confirm service",
			zap.String("service", string(msg.ServiceName)))
		return ErrUnknownService
	}

	// Ждем подтверждения от всех сервисов
	if !uc.registry.IsCommandConfirmed(msg.Command) {
		return nil
	}

	// Оркестрация действий после подтверждения команды
	switch msg.Command {
	case domainIntf.Pause:
		uc.logger.Info("All services paused, sending TakeSnapshots command")
		return uc.publisher.Publish(ctx, message.NewServiceControlCommand(domainIntf.TakeSnapshots))

	case domainIntf.TakeSnapshots:
		uc.logger.Info("All snapshots taken, resuming services")
		err := uc.publisher.Publish(ctx, message.NewServiceControlCommand(domainIntf.Resume))
		if err != nil {
			return err
		}
		// Обрабатываем сохраненные снапшоты
		return uc.handleSnapshot.Process(ctx, msg.TakeSnapshotData)
	}

	return nil
}
