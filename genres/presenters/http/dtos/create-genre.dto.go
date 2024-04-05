package dtos

import (
	"app/translations/application/payloads"
)

type CreateGenreDto struct {
	CreateTranslationPayloads []*payloads.CreateTranslationPayload `json:"createTranslationPayloads"`
}