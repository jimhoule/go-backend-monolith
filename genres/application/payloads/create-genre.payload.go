package payloads

import "app/translations/application/payloads"

type CreateGenrePayload struct {
	CreateLabelTranslationPayloads []*payloads.CreateTranslationPayload
}