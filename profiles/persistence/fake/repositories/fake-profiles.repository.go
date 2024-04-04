package repositories

import (
	"app/profiles/domain/models"
	"fmt"
)

var profiles []*models.Profile = []*models.Profile{}

func ResetFakeProfilesRepository() {
	profiles = []*models.Profile{}
}

type FakeProfilesRepository struct {}

func (fpr *FakeProfilesRepository) FindAllByAccountId(accountId string) ([]*models.Profile, error) {
	accountProfiles := []*models.Profile{}
	for _, profile := range profiles {
		if profile.AccountId == accountId {
			accountProfiles = append(accountProfiles, profile)
		}
	}

	return accountProfiles, nil
}

func (fpr *FakeProfilesRepository) FindById(id string) (*models.Profile, error) {
	for _, profile := range profiles {
		if profile.Id == id {
			return profile, nil
		}
	}

	return nil, fmt.Errorf("the profile with id %s does not exist", id)
}

func (fpr *FakeProfilesRepository) Update(id string, profile *models.Profile) (*models.Profile, error) {
	for _, profileToUpdate := range profiles {
		if profileToUpdate.Id == id {
			profileToUpdate.Name = profile.Name
			profileToUpdate.LanguageId = profile.LanguageId
			return profileToUpdate, nil
		}
	}

	return nil, fmt.Errorf("the profile with id %s does not exist", id)
}

func (fpr *FakeProfilesRepository) Delete(id string) (string, error) {

	for index, profile := range profiles {
		if profile.Id == id {
			/*
			 * NOTES:
			 *	- profiles[:index] yields the slice elements before
			 *	- profiles[index+1:]... yields the slice elements after
			 *	- The 2 slices are then merged together by append()
			 */
			profiles = append(profiles[:index], profiles[index + 1:]...)
		}
	}

	return id, nil
}

func (fpr *FakeProfilesRepository) Create(profile *models.Profile) (*models.Profile, error) {
	profiles = append(profiles, profile)

	return profile, nil
}