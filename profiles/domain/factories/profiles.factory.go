package factories

import (
	"app/profiles/domain/models"
	"app/uuid/services"
)

type ProfilesFactory struct {
	UuidService services.UuidService
}

func (pf *ProfilesFactory) Create(name string, accountId string, languageId string) *models.Profile {
	return &models.Profile{
		Id: pf.UuidService.Generate(),
		Name: name,
		AccountId: accountId,
		LanguageId: languageId,
	}
}