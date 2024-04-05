package services

import (
	"app/translations/application/payloads"
	"app/translations/application/ports"
	"app/translations/domain/factories"
	"app/translations/domain/models"
	"context"
)

type TranslationsService struct {
	TranslationsFactory *factories.TranslationsFactory
	TranslationsRepository ports.TranslationsRepositoryPort
}

func (ts *TranslationsService) ExecuteTransaction(
	ctx context.Context,
	executeQuery func(ctx context.Context) (any, error),
) (any, error) {
	return ts.TranslationsRepository.ExecuteTransaction(ctx, executeQuery)
}

func (ts *TranslationsService) FindAll() ([]*models.Translation, error) {
	return ts.TranslationsRepository.FindAll()
}

func (ts *TranslationsService) FindAllByEntityId(entityId string) ([]*models.Translation, error) {
	return ts.TranslationsRepository.FindAllByEntityId(entityId)
}

func (ts *TranslationsService) FindByCompositeId(entityId string, languageCode string) (*models.Translation, error) {
	return ts.TranslationsRepository.FindByCompositeId(entityId, languageCode)
}

func (ts *TranslationsService) Create(ctx context.Context, createTranslationPayloads []*payloads.CreateTranslationPayload) ([]*models.Translation, error) {
	translations := []*models.Translation{}
	for _, createTranslationPayload := range createTranslationPayloads {
		translation := ts.TranslationsFactory.Create(
			createTranslationPayload.EntityId,
			createTranslationPayload.LanguageCode,
			createTranslationPayload.Text,
		)

		translations = append(translations, translation)
	}

	return ts.TranslationsRepository.Create(ctx, translations)
}