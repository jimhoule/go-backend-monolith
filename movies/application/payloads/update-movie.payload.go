package payloads

import "app/translations/application/payloads"

type UpdateMoviePayload struct {
	GenreId                              string
	UpdateTitleTranslationPayloads       []*payloads.UpdateTranslationPayload
	UpdateDescriptionTranslationPayloads []*payloads.UpdateTranslationPayload
}