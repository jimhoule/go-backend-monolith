package payloads

type UpdateTranslationPayload struct {
	LanguageCode string `json:"languageCode"`
	Text         string `json:"text"`
}