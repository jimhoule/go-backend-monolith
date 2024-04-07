package factories

import "app/translations/domain/models"

type TranslationsFactory struct{}

func (tf *TranslationsFactory) Create(entityId string, languageId string, text string) *models.Translation {
	return &models.Translation{
		EntityId: entityId,
		LanguageId: languageId,
		Text: text,
	}
}