package payloads

import "app/translations/application/payloads"

type UpdateLanguagePayload struct {
	Code                      string
	UpdateLabelTranslationPayloads []*payloads.UpdateTranslationPayload
}