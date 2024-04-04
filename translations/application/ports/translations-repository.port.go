package ports

import "app/translations/domain/models"

type TranslationsRepositoryPort interface {
	FindAll() ([]*models.Translation, error)
	FindByCompositeId(entityId string, languageCode string) (*models.Translation, error)
	Create(translation *models.Translation) (*models.Translation, error)
}