package payloads

import "app/translations/application/payloads"

type CreateLanguagePayload struct {
	Code                      string
	CreateLabelTranslationPayloads []*payloads.CreateTranslationPayload
}