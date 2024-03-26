package services

import (
	"app/plans/application/payloads"
	"app/plans/application/ports"
	"app/plans/domain/factories"
	"app/plans/domain/models"
)

type PlansService struct {
	PlansFactory factories.PlansFactory
	PlansRepository ports.PlansRepository
}

func (ps *PlansService) FindAll() ([]*models.Plan, error) {
	return ps.PlansRepository.FindAll()
}

func (ps *PlansService) FindById(id string) (*models.Plan, error) {
	return ps.PlansRepository.FindById(id)
}

func (ps *PlansService) Create(createPlanPayload payloads.CreatePlanPayload) (*models.Plan, error) {
	plan := ps.PlansFactory.Create(
		createPlanPayload.Name,
		createPlanPayload.Description,
		createPlanPayload.Price,
	)

	return ps.PlansRepository.Create(plan)
}