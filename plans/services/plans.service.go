package services

import (
	"app/plans/domain/factories"
	"app/plans/domain/models"
	"app/plans/dtos"
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

func (ps *PlansService) Create(createPlanDto dtos.CreatePlanDto) (*models.Plan, error) {
	plan := ps.PlansFactory.Create(
		createPlanDto.Name,
		createPlanDto.Description,
		createPlanDto.Price,
	)

	return ps.PlansRepository.Create(plan)
}