package models

import "app/translations/domain/models"

type Genre struct {
	Id    string                 `json:"id"`
	Labels []*models.Translation `json:"labels,omitempty"`
}