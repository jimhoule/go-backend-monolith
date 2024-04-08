package dtos

import "app/translations/application/payloads"

type CreatePlanDto struct {
	Price                                float32                              `jons:"price"`
	CreateLabelTranslationPayloads       []*payloads.CreateTranslationPayload `json:"createLabelTranslationPayloads"`
	CreateDescriptionTranslationPayloads []*payloads.CreateTranslationPayload `json:"createDescriptionTranslationPayloads"`
}