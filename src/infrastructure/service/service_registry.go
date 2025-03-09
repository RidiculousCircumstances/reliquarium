package service

import (
	"reliquarium/src/domain"
	domainIntf "reliquarium/src/domain/interfaces"
	"sync"
)

type Registry struct {
	services map[domainIntf.ServiceName]*domain.Service
	mu       sync.RWMutex
}

func NewRegistry(serviceNames []domainIntf.ServiceName) *Registry {
	registry := &Registry{
		services: make(map[domainIntf.ServiceName]*domain.Service),
	}
	for _, name := range serviceNames {
		registry.services[name] = domain.NewService(name)
	}
	return registry
}

// Подтверждение выполнения команды сервисом
func (r *Registry) ConfirmServiceCommand(serviceName domainIntf.ServiceName, command domainIntf.CommandType) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if service, exists := r.services[serviceName]; exists {
		service.ConfirmCommand(command)
		return true
	}
	return false
}

// Проверяет, подтвердили ли все сервисы выполнение конкретной команды (pause, takeSnapshots, resume)
func (r *Registry) IsCommandConfirmed(command domainIntf.CommandType) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, service := range r.services {
		if !service.IsCommandConfirmed(command) {
			return false
		}
	}
	return true
}

// Возвращает список сервисов, которые не подтвердили выполнение команды
func (r *Registry) GetUnconfirmedServices(command domainIntf.CommandType) []domainIntf.ServiceName {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var unconfirmed []domainIntf.ServiceName
	for name, service := range r.services {
		if !service.IsCommandConfirmed(command) {
			unconfirmed = append(unconfirmed, name)
		}
	}
	return unconfirmed
}

// Сбрасывает состояние всех сервисов перед новым процессом
func (r *Registry) ResetServiceStates() {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, service := range r.services {
		service.ResetState()
	}
}
