package payloads

import "app/translations/application/payloads"

type UpdateGenrePayload struct {
	UpdateLabelTranslationPayloads []*payloads.UpdateTranslationPayload
}