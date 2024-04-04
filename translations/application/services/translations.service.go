package services

import (
	"app/translations/application/payloads"
	"app/translations/application/ports"
	"app/translations/domain/factories"
	"app/translations/domain/models"
)

type TranslationsService struct {
	TranslationsFactory *factories.TranslationsFactory
	TranslationsRepository ports.TranslationsRepositoryPort
}

func (ts *TranslationsService) FindAll() ([]*models.Translation, error) {
	return ts.TranslationsRepository.FindAll()
}

func (ts *TranslationsService) FindByCompositeId(entityId string, languageCode string) (*models.Translation, error) {
	return ts.TranslationsRepository.FindByCompositeId(entityId, languageCode)
}

func (ts *TranslationsService) Create(createTranslationPayload *payloads.CreateTranslationPayload) (*models.Translation, error) {
	translation := ts.TranslationsFactory.Create(
		createTranslationPayload.EntityId,
		createTranslationPayload.LanguageCode,
		createTranslationPayload.Text,
	)

	return ts.TranslationsRepository.Create(translation)
}