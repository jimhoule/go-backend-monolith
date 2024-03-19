package repositories

import "app/plans/domain/models"

type PlansRepository interface {
	FindAll() ([]*models.Plan, error)
	FindById(id string) (*models.Plan, error)
	Create(planModel *models.Plan) (*models.Plan, error)
}