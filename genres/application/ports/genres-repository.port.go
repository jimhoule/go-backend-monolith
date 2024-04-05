package ports

import (
	"app/genres/domain/models"
	"context"
)

type GenresRepositoryPort interface {
	FindAll() ([]*models.Genre, error)
	FindById(id string) (*models.Genre, error)
	Delete(ctx context.Context, id string) (string, error)
	Create(ctx context.Context, genre *models.Genre) (*models.Genre, error)
}