package models

import "app/translations/domain/models"

type Language struct {
	Id     string                `json:"id"`
	Code   string                `json:"code"`
	Labels []*models.Translation `json:"labels"`
}