package domain

import (
	"reliquarium/src/domain/interfaces"
	"sync"
)

type Service struct {
	name              interfaces.ServiceName
	confirmedCommands map[interfaces.CommandType]bool
	mu                sync.RWMutex
}

func NewService(name interfaces.ServiceName) *Service {
	return &Service{
		name:              name,
		confirmedCommands: make(map[interfaces.CommandType]bool),
	}
}

// Подтверждает выполнение команды
func (s *Service) ConfirmCommand(command interfaces.CommandType) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.confirmedCommands[command] = true
}

// Проверяет, была ли команда подтверждена
func (s *Service) IsCommandConfirmed(command interfaces.CommandType) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.confirmedCommands[command]
}

// Сбрасывает состояние сервисов перед новым циклом оркестрации
func (s *Service) ResetState() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.confirmedCommands = make(map[interfaces.CommandType]bool)
}
