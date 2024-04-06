package ports

import (
	"app/translations/domain/models"
	"context"
)

type TranslationsRepositoryPort interface {
	ExecuteTransaction(ctx context.Context, executeQuery func(ctx context.Context) (any, error)) (any, error)
	FindAll() ([]*models.Translation, error)
	FindAllByEntityId(entityId string) ([]*models.Translation, error)
	FindByCompositeId(entityId string, languageCode string) (*models.Translation, error)
	UpdateBatch(ctx context.Context, translations []*models.Translation) ([]*models.Translation, error)
	DeleteBatch(ctx context.Context, entityId string) (string, error)
	CreateBatch(ctx context.Context, translations []*models.Translation) ([]*models.Translation, error)
}