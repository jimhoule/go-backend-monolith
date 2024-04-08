package dtos

import (
	"app/translations/application/payloads"
)

type CreateGenreDto struct {
	CreateLabelTranslationPayloads []*payloads.CreateTranslationPayload `json:"createLabelTranslationPayloads"`
}