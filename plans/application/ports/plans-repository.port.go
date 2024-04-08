package ports

import (
	"app/plans/domain/models"
	"context"
)

type PlansRepositoryPort interface {
	FindAll() ([]*models.Plan, error)
	FindById(id string) (*models.Plan, error)
	Create(ctx context.Context, plan *models.Plan) (*models.Plan, error)
}