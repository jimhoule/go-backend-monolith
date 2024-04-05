package dtos

import "app/translations/presenters/http/dtos"

type CreateGenreDto struct {
	CreateTranslationDtos []*dtos.CreateTranslationDto `json:"createTranslationDtos"`
}