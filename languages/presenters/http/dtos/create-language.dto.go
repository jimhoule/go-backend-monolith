package dtos

import "app/translations/application/payloads"

type CreateLanguageDto struct {
	Code                      string                               `json:"code"`
	CreateTranslationPayloads []*payloads.CreateTranslationPayload `json:"createTranslationPayloads"`
}