package dtos

import "app/translations/application/payloads"

type UpdateMovieDto struct {
	GenreId                              string                               `json:"genreId"`
	UpdateTitleTranslationPayloads       []*payloads.UpdateTranslationPayload `json:"updateTitleTranslationPayloads"`
	UpdateDescriptionTranslationPayloads []*payloads.UpdateTranslationPayload `json:"updateDescriptionTranslationPayloads"`
}