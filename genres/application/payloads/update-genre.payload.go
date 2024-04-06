package payloads

import "app/translations/application/payloads"

type UpdateGenrePayload struct {
	UpdateTranslationPayloads []*payloads.UpdateTranslationPayload
}