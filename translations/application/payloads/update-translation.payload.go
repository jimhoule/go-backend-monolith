package payloads

type UpdateTranslationPayload struct {
	LanguageId string `json:"languageId"`
	Text       string `json:"text"`
}