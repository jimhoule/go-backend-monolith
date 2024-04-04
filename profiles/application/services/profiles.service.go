package services

import (
	"app/profiles/application/payloads"
	"app/profiles/application/ports"
	"app/profiles/domain/factories"
	"app/profiles/domain/models"
)

type ProfilesService struct {
	ProfilesFactory *factories.ProfilesFactory
	ProfilesRepository ports.ProfilesRepositoryPort
}

func (ps *ProfilesService) FindAllByAccountId(accountId string) ([]*models.Profile, error) {
	return ps.ProfilesRepository.FindAllByAccountId(accountId)
}

func (ps *ProfilesService) FindById(id string) (*models.Profile, error) {
	return ps.ProfilesRepository.FindById(id)
}

func (ps *ProfilesService) Update(id string, updateProfilePayload *payloads.UpdateProfilePayload) (*models.Profile, error) {
	return ps.ProfilesRepository.Update(id, &models.Profile{ 
		Name: updateProfilePayload.Name,
		LanguageId: updateProfilePayload.LanguageId,
	})
}

func (ps *ProfilesService) Delete(id string) (string, error) {
	return ps.ProfilesRepository.Delete(id)
}

func (ps *ProfilesService) Create(createProfilePayload *payloads.CreateProfilePayload) (*models.Profile, error) {
	profile := ps.ProfilesFactory.Create(
		createProfilePayload.Name,
		createProfilePayload.AccountId,
		createProfilePayload.LanguageId,
	)

	return ps.ProfilesRepository.Create(profile)
}