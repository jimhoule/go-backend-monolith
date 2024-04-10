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

func (ftr *FakeTranslationsRepository) FindAll() ([]*models.Translation, error) {
	return translations, nil
}

func (ftr *FakeTranslationsRepository) FindAllByEntityIdAndType(entityId string, translationType string) ([]*models.Translation, error) {
	entityTranslations := []*models.Translation{}
	for _, translation := range translations {
		if translation.EntityId == entityId && translation.Type == translationType {
			entityTranslations = append(entityTranslations, translation)
		}
	}

	return entityTranslations, nil
}

func (ftr *FakeTranslationsRepository) FindByCompositeId(entityId string, languageId string, translationType string) (*models.Translation, error) {
	for _, translation := range translations {
		if translation.EntityId == entityId && translation.LanguageId == languageId && translation.Type == translationType {
			return translation, nil
		}
	}

	return nil, fmt.Errorf("the translation with composite id (%s, %s) does not exist", entityId, languageId)
}

func (ftr *FakeTranslationsRepository) UpsertBatch(ctx context.Context, updatedTranslations []*models.Translation) ([]*models.Translation, error) {
	for _, updatedTranslation := range updatedTranslations {
		translation, err := ftr.FindByCompositeId(updatedTranslation.EntityId, updatedTranslation.LanguageId, updatedTranslation.Type)
		// Creates translation if not found
		if err != nil {
			translations = append(translations, updatedTranslation)
			continue
		}

		// Updates translation if found
		translation.Text = updatedTranslation.Text
		translation.Type = updatedTranslation.Type
	}

	return updatedTranslations, nil
}

func (ftr *FakeTranslationsRepository) DeleteBatch(ctx context.Context, entityId string) (string, error) {
	for index, translation := range translations {
		if translation.EntityId == entityId {
			translations = append(translations[:index], translations[index + 1:]...)
		}
	}

	return entityId, nil
}

func (ftr *FakeTranslationsRepository) CreateBatch(ctx context.Context, newTranslations []*models.Translation) ([]*models.Translation, error) {
	translations = append(translations, newTranslations...)

	return translations, nil
}