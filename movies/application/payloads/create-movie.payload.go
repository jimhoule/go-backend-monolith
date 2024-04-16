package payloads

import "app/translations/application/payloads"

type CreateMoviePayload struct {
	GenreId                              string
	CreateTitleTranslationPayloads       []*payloads.CreateTranslationPayload
	CreateDescriptionTranslationPayloads []*payloads.CreateTranslationPayload
}