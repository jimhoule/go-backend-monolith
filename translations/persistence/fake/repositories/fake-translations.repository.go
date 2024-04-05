package repositories

import (
	"app/translations/domain/models"
	"context"
	"fmt"
)

var translations []*models.Translation = []*models.Translation{}

func ResetFakeTranslationRepository() {
	translations = []*models.Translation{}
}

type FakeTranslationsRepository struct {}

func (ftr *FakeTranslationsRepository) ExecuteTransaction(
	ctx context.Context,
	executeQuery func(ctx context.Context) (any, error),
) (any, error) {
	return executeQuery(ctx)
}

func (ftr *FakeTranslationsRepository) FindAll() ([]*models.Translation, error) {
	return translations, nil
}

func (ftr *FakeTranslationsRepository) FindAllByEntityId(entityId string) ([]*models.Translation, error) {
	entityTranslations := []*models.Translation{}
	for _, translation := range translations {
		if translation.EntityId == entityId {
			entityTranslations = append(entityTranslations, translation)
		}
	}

	return entityTranslations, nil
}

func (ftr *FakeTranslationsRepository) FindByCompositeId(entityId string, languageCode string) (*models.Translation, error) {
	for _, translation := range translations {
		if translation.EntityId == entityId && translation.LanguageCode == languageCode {
			return translation, nil
		}
	}

	return nil, fmt.Errorf("the translation with composite id (%s, %s) does not exist", entityId, languageCode)
}

func (ftr *FakeTranslationsRepository) Create(ctx context.Context, newTranslations []*models.Translation) ([]*models.Translation, error) {
	translations = append(translations, newTranslations...)

	return translations, nil
}