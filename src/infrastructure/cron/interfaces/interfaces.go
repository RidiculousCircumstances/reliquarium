package interfaces

import "context"

type CronTask interface {
	Run(ctx context.Context) error
}
