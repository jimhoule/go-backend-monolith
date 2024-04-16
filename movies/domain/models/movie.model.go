package models

import "app/translations/domain/models"

type Movie struct {
	Id           string                `json:"id"`
	Titles       []*models.Translation `json:"titles"`
	Descriptions []*models.Translation `json:"descriptions,omitempty"`
	GenreId      string                `json:"genreId"`
}