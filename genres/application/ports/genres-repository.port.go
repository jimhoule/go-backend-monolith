package ports

import "app/genres/domain/models"

type GenresRepositoryPort interface {
	FindAll() ([]*models.Genre, error)
	FindById(id string) (*models.Genre, error)
	Create(genre *models.Genre) (*models.Genre, error)
}