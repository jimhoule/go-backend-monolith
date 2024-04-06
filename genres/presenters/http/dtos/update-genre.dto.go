package dtos

import (
	"app/translations/application/payloads"
)

type UpdateGenreDto struct {
	UpdateTranslationPayloads []*payloads.UpdateTranslationPayload `json:"updateTranslationPayloads"`
}