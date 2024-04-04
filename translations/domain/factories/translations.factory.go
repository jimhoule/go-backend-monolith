package factories

import "app/translations/domain/models"

type TranslationsFactory struct{}

func (tf *TranslationsFactory) Create(entityId string, languageCode string, text string) *models.Translation {
	return &models.Translation{
		EntityId: entityId,
		LanguageCode: languageCode,
		Text: text,
	}
}