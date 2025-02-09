package storage

import (
	"reliquarium/src/application/service/storage/interfaces"
)

type Service struct {
	client interfaces.StorageClient
}

func NewService(client interfaces.StorageClient) *Service {
	return &Service{client: client}
}
