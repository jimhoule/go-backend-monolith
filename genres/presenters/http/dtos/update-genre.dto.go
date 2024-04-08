package dtos

import (
	"app/translations/application/payloads"
)

type UpdateGenreDto struct {
	UpdateLabelTranslationPayloads []*payloads.UpdateTranslationPayload `json:"updateLabelTranslationPayloads"`
}