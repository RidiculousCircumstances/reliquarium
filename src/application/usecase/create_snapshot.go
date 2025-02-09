package usecase

import "context"

type CreateSnapshot struct{}

func (u *CreateSnapshot) Process(ctx context.Context) error {
	return nil
}
