package ports

import "app/translations/domain/models"

type TranslationsRepositoryPort interface {
	FindAll() ([]*models.Translation, error)
	FindAllByEntityId(entityId string) ([]*models.Translation, error)
	FindByCompositeId(entityId string, languageCode string) (*models.Translation, error)
	Create(translations []*models.Translation) ([]*models.Translation, error)
}