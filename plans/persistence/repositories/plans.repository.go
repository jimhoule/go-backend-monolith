package repositories

import "app/plans/domain/models"

type PlansRepository interface {
	FindAll() ([]*models.Plan, error)
	FindById(id string) (*models.Plan, error)
	Create(plan *models.Plan) (*models.Plan, error)
}