package payloads

type CreateTranslationPayload struct {
	LanguageId string `json:"languageId"`
	Text       string `json:"text"`
	Type       string `json:"type"`
}