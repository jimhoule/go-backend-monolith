package dtos

type CreateTranslationDto struct {
	LanguageCode string `json:"languageCode"`
	Text         string `json:"text"`
}