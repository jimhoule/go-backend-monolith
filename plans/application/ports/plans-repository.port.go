package ports

import "app/plans/domain/models"

type PlansRepositoryPort interface {
	FindAll() ([]*models.Plan, error)
	FindById(id string) (*models.Plan, error)
	Create(plan *models.Plan) (*models.Plan, error)
}