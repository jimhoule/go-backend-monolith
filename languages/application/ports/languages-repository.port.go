package ports

import "app/languages/domain/models"

type LanguagesRepositoryPort interface {
	FindAll() ([]*models.Language, error)
	FindById(id string) (*models.Language, error)
	Update(id string, language *models.Language) (*models.Language, error)
	Delete(id string) (string, error)
	Create(language *models.Language) (*models.Language, error)
}