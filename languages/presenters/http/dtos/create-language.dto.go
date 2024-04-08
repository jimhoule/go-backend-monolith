package dtos

import "app/translations/application/payloads"

type CreateLanguageDto struct {
	Code                      string                               `json:"code"`
	CreateLabelTranslationPayloads []*payloads.CreateTranslationPayload `json:"createLabelTranslationPayloads"`
}