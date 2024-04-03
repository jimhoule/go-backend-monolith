package factories

import (
	"app/languages/domain/models"
	"app/uuid/services"
)

type LanguagesFactory struct {
	UuidService services.UuidService
}

func (lf *LanguagesFactory) Create(code string, title string) *models.Language {
	return &models.Language{
		Id: lf.UuidService.Generate(),
		Code: code,
		Title: title,
	}
}