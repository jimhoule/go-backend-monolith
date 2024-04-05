package payloads

type CreateTranslationPayload struct {
	EntityId     string `json:"entityId"`
	LanguageCode string `json:"languageCode"`
	Text         string `json:"text"`
}