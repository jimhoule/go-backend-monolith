package payloads

import "app/translations/application/payloads"

type CreateGenrePayload struct {
	CreateTranslationPayloads []*payloads.CreateTranslationPayload
}