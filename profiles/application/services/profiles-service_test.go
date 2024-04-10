package services

import (
	"app/profiles/application/payloads"
	"app/profiles/domain/factories"
	"app/profiles/domain/models"
	"app/profiles/infrastructures/persistence/fake/repositories"
	"app/uuid"
	"testing"
)

func getTestContext() (*ProfilesService, func(), func() (*models.Profile, error)) {
	profilesService := &ProfilesService{
		ProfilesFactory: &factories.ProfilesFactory{
			UuidService: uuid.GetService(),
		},
		ProfilesRepository: &repositories.FakeProfilesRepository{},
	}

	createProfile := func() (*models.Profile, error) {
		return profilesService.Create(&payloads.CreateProfilePayload{
			Name: "Fake profile name",
			AccountId: "fakeAccoutId",
			LanguageId: "fakeLanguageId",
		})
	}

	return profilesService, repositories.ResetFakeProfilesRepository, createProfile
}

func TestCreateProfileService(t *testing.T) {
	_, reset, createProfile := getTestContext()
	defer reset()

	_, err := createProfile()
	if err != nil {
		t.Errorf("Expected to create Profile but got %v", err)
	}
}

func TestFindAllProfilesByAccountIdService(t *testing.T) {
	profilesService, reset, createProfile := getTestContext()
	defer reset()

	newProfile, _ := createProfile()

	profiles, err := profilesService.FindAllByAccountId(newProfile.AccountId)
	if err != nil {
		t.Errorf("Expected slice of Profiles but got %v", err)
		return
	}

	if profiles[0] != newProfile {
		t.Errorf("Expected first element of Profiles slice to be equal to New Profile but got %v", profiles[0])
	}
}

func TestFindProfileByIdService(t *testing.T) {
	profilesService, reset, createProfile := getTestContext()
	defer reset()

	newProfile, _ := createProfile()

	profile, err := profilesService.FindById(newProfile.Id)
	if err != nil {
		t.Errorf("Expected Profile but got %v", err)
		return
	}

	if profile != newProfile {
		t.Errorf("Expected Profile to equal New Profile but got %v", profile)
	}
}

func TestUpdateProfileService(t *testing.T) {
	profilesService, reset, createProfile := getTestContext()
	defer reset()

	newProfile, _ := createProfile()

	updatedName := "Updated fake profile name"
	profile, err := profilesService.Update(newProfile.Id, &payloads.UpdateProfilePayload{ 
		Name: updatedName,
		LanguageId: newProfile.LanguageId,
	})
	if err != nil {
		t.Errorf("Expected Profile but got %v", err)
		return
	}

	if newProfile.Name != updatedName {
		t.Errorf("Expected New Profile name to equal updated name but got %s", newProfile.Name)
		return
	}

	if profile.Name != updatedName {
		t.Errorf("Expected Profile name to equal updated name but got %s", profile.Name)
	}
}

func TestDeleteProfileService(t *testing.T) {
	profilesService, reset, createProfile := getTestContext()
	defer reset()

	newProfile, _ := createProfile()

	profileId, err := profilesService.Delete(newProfile.Id)
	if err != nil {
		t.Errorf("Expected Profile id but got %v", err)
		return
	}

	if newProfile.Id != profileId {
		t.Errorf("Expected New Profile id to equal Profile id but got %s", newProfile.Id)
	}
}