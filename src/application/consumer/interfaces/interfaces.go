package interfaces

import "context"

type CreateSnapshot interface {
	Process(ctx context.Context) error
}
