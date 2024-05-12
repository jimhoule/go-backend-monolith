package services

import "app/app/application/ports"

type AppService struct {
	AppRepository ports.AppRepositoryPort
}

func (as *AppService) Diagnose() error {
	return as.AppRepository.Diagnose()
}