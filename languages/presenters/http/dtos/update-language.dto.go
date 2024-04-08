package dtos

import "app/translations/application/payloads"

type UpdateLanguageDto struct {
	Code                      string                               `json:"code"`
	UpdateLabelTranslationPayloads []*payloads.UpdateTranslationPayload `json:"updateLabelTranslationPayloads"`
}