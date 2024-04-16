package dtos

import "app/translations/application/payloads"

type CreateMovieDto struct {
	GenreId                              string                               `json:"genreId"`
	CreateTitleTranslationPayloads       []*payloads.CreateTranslationPayload `json:"createTitleTranslationPayloads"`
	CreateDescriptionTranslationPayloads []*payloads.CreateTranslationPayload `json:"createDescriptionTranslationPayloads"`
}