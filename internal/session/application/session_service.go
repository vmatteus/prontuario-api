package application

import (
	interfaces "prontuario/internal/session/application/interfaces"
	"prontuario/internal/session/domain"
	"prontuario/internal/session/infrastructure"
)

type sessionService struct {
	driver domain.SessionDriverInterface
}

func NewSessionService() interfaces.SessionServiceInterface {
	return &sessionService{driver: infrastructure.GetSessionDriver()}
}

func (service *sessionService) Set(key string, value interface{}) error {
	return nil
}

func (service *sessionService) Get(key string) (interface{}, error) {
	return nil, nil
}

func (service *sessionService) Delete(key string) error {
	return nil
}
