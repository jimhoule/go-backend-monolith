package services

import (
	"app/plans/domain/factories"
	"app/plans/domain/models"
	"app/plans/persistence/repositories"
)

type PlansService struct {
	PlansFactory factories.PlansFactory
	PlansRepository repositories.PlansRepository
}

func (ps *PlansService) FindAll() ([]*models.Plan, error) {
	return ps.PlansRepository.FindAll()
}

func (ps *PlansService) FindById(id string) (*models.Plan, error) {
	return ps.PlansRepository.FindById(id)
}

func (ps *PlansService) Create(name string, description string, price float32) (*models.Plan, error) {
	plan := ps.PlansFactory.Create(name, description, price)

	return ps.PlansRepository.Create(plan)
}