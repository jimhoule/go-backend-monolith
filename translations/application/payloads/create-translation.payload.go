package payloads

type CreateTranslationPayload struct {
	EntityId     string
	LanguageCode string
	Text         string
}