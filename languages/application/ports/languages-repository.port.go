package ports

import (
	"app/languages/domain/models"
	"context"
)

type LanguagesRepositoryPort interface {
	FindAll() ([]*models.Language, error)
	FindById(id string) (*models.Language, error)
	Update(ctx context.Context, id string, language *models.Language) (*models.Language, error)
	Delete(id string) (string, error)
	Create(ctx context.Context, language *models.Language) (*models.Language, error)
}