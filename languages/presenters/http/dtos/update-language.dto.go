package dtos

import "app/translations/application/payloads"

type UpdateLanguageDto struct {
	Code                      string                               `json:"code"`
	UpdateTranslationPayloads []*payloads.UpdateTranslationPayload `json:"updateTranslationPayloads"`
}