package repositories

import (
	"app/translations/domain/models"
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

func (ftr *FakeTranslationsRepository) FindByEntityId(entityId string) ([]*models.Translation, error) {
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

func (ftr *FakeTranslationsRepository) Create(translation *models.Translation) (*models.Translation, error) {
	translations = append(translations, translation)

	return translation, nil
}