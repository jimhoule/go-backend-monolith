package models

import "app/translations/domain/models"

type Plan struct {
	Id           string                `json:"id"`
	Price        float32               `json:"price"`
	Labels       []*models.Translation `json:"labels"`
	Descriptions []*models.Translation `json:"descriptions"`
}