package models

import "app/translations/domain/models"

type Genre struct {
	Id    string              `json:"id"`
	Label *models.Translation `json:"label,omitempty"`
}