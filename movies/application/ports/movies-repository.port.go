package ports

import (
	"app/movies/domain/models"
	"context"
)

type MoviesRepositoryPort interface {
	FindAll() ([]*models.Movie, error)
	FindById(id string) (*models.Movie, error)
	Update(ctx context.Context, id string, movie *models.Movie) (*models.Movie, error)
	Delete(ctx context.Context, id string) (string, error)
	Create(ctx context.Context, movie *models.Movie) (*models.Movie, error)
}