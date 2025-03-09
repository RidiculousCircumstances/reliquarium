package cron

import (
	"reliquarium/config/env"
	"reliquarium/shared/logger/interfaces"
	jobIntf "reliquarium/src/application/interfaces"
	"reliquarium/src/application/job"
	cronIntf "reliquarium/src/infrastructure/cron/interfaces"
)

// CronTaskConfig — структура для хранения параметров задачи Cron
type CronTaskConfig struct {
	Schedule string            // Расписание выполнения (например, "03:00")
	Task     cronIntf.CronTask // Задача, реализующая интерфейс CronTask
}

// BuildCronTaskConfigs — создаёт список задач Cron с их конфигурацией
func BuildCronTaskConfigs(
	env *env.Config,
	logger interfaces.Logger,
	publisher jobIntf.CommandPublisher,
) []CronTaskConfig {
	return []CronTaskConfig{
		{
			Schedule: "0 3 * * *",
			Task:     job.NewSnapshotTriggerJob(publisher),
		},
	}
}
