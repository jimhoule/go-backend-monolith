package payloads

import "app/translations/application/payloads"

type CreatePlanPayload struct {
	Price                                float32
	CreateLabelTranslationPayloads       []*payloads.CreateTranslationPayload
	CreateDescriptionTranslationPayloads []*payloads.CreateTranslationPayload
}