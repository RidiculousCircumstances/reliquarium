package cron

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"go.uber.org/zap"
	loggerIntf "reliquarium/shared/logger/interfaces"
)

// CronHook - Хук для управления задачами по расписанию
type CronHook struct {
	scheduler *gocron.Scheduler
	tasks     []CronTaskConfig
	logger    loggerIntf.Logger
}

// NewCronHook - Конструктор для создания CronHook
func NewCronHook(logger loggerIntf.Logger, taskConfigs []CronTaskConfig) *CronHook {
	return &CronHook{
		scheduler: gocron.NewScheduler(time.UTC),
		tasks:     taskConfigs,
		logger:    logger,
	}
}

func (c *CronHook) Setup(ctx context.Context) error {
	return nil
}

// Start - Запускает Cron-задачи с индивидуальными расписаниями
func (c *CronHook) Start(ctx context.Context) error {
	for _, taskConfig := range c.tasks {
		schedule := taskConfig.Schedule
		task := taskConfig.Task

		_, err := c.scheduler.Every(1).Day().At(schedule).Do(func() {
			c.logger.Info("Executing cron task", zap.String("schedule", schedule))

			// Запускаем задачу в отдельной горутине
			go func() {
				if err := task.Run(ctx); err != nil {
					c.logger.Error("Failed to execute cron task", zap.String("schedule", schedule), zap.Error(err))
				}
			}()
		})
		if err != nil {
			return err
		}
	}

	c.scheduler.StartAsync()
	c.logger.Info("Cron scheduler started")
	return nil
}

// Shutdown - Завершает работу Cron-задач
func (c *CronHook) Shutdown(ctx context.Context) error {
	c.scheduler.Stop()
	c.logger.Info("Cron scheduler stopped")
	return nil
}
