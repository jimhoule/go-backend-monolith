package ports

import (
	"app/translations/domain/models"
	"context"
)

type TranslationsRepositoryPort interface {
	FindAll() ([]*models.Translation, error)
	FindAllByEntityId(entityId string) ([]*models.Translation, error)
	FindByCompositeId(entityId string, languageId string) (*models.Translation, error)
	UpsertBatch(ctx context.Context, translations []*models.Translation) ([]*models.Translation, error)
	DeleteBatch(ctx context.Context, entityId string) (string, error)
	CreateBatch(ctx context.Context, translations []*models.Translation) ([]*models.Translation, error)
}