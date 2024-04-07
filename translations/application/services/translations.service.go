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

func (ts *TranslationsService) FindAll() ([]*models.Translation, error) {
	return ts.TranslationsRepository.FindAll()
}

func (ts *TranslationsService) FindAllByEntityId(entityId string) ([]*models.Translation, error) {
	return ts.TranslationsRepository.FindAllByEntityId(entityId)
}

func (ts *TranslationsService) FindByCompositeId(entityId string, languageId string) (*models.Translation, error) {
	return ts.TranslationsRepository.FindByCompositeId(entityId, languageId)
}

func (ts *TranslationsService) UpsertBatch(
	ctx context.Context,
	entityId string,
	updateTranslationPayloads []*payloads.UpdateTranslationPayload,
) ([]*models.Translation, error) {
	translations := []*models.Translation{}
	for _, updateTranslationPayload := range updateTranslationPayloads {
		translation := &models.Translation{
			EntityId: entityId,
			LanguageId: updateTranslationPayload.LanguageId,
			Text: updateTranslationPayload.Text,
		}

		translations = append(translations, translation)
	}

	return ts.TranslationsRepository.UpsertBatch(ctx, translations)
}

func (ts *TranslationsService) DeleteBatch(ctx context.Context, entityId string) (string, error) {
	return ts.TranslationsRepository.DeleteBatch(ctx, entityId)
}

func (ts *TranslationsService) CreateBatch(
	ctx context.Context,
	entityId string,
	createTranslationPayloads []*payloads.CreateTranslationPayload,
) ([]*models.Translation, error) {
	translations := []*models.Translation{}
	for _, createTranslationPayload := range createTranslationPayloads {
		translation := ts.TranslationsFactory.Create(
			entityId,
			createTranslationPayload.LanguageId,
			createTranslationPayload.Text,
		)

		translations = append(translations, translation)
	}

	return ts.TranslationsRepository.CreateBatch(ctx, translations)
}