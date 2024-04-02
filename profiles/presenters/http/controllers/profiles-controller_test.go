package controllers

import (
	"app/profiles/application/payloads"
	"app/profiles/application/services"
	"app/profiles/domain/factories"
	"app/profiles/domain/models"
	"app/profiles/persistence/fake/repositories"
	"app/profiles/presenters/http/dtos"
	"app/router/mock"
	"app/uuid"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestContext() (*ProfilesController, func(), func() (*models.Profile, error)) {
	profilesController := &ProfilesController{
		ProfilesService: &services.ProfilesService{
			ProfilesFactory: &factories.ProfilesFactory{
				UuidService: uuid.GetService(),
			},
			ProfilesRepository: &repositories.FakeProfilesRepository{},
		},
	}

	createProfile := func() (*models.Profile, error) {
		return profilesController.ProfilesService.Create(&payloads.CreateProfilePayload{
			Name: "Fake profile name",
			AccountId: "fakeAccoutId",
		})
	}

	return profilesController, repositories.ResetFakeProfilesRepository, createProfile
}

func TestCreateProfileController(t *testing.T) {
	profilesController, reset, _ := getTestContext()
	defer reset()

	// Creates request body
	requestBody, err := json.Marshal(dtos.CreateProfileDto{
		Name: "Fake profile name",
		AccountId: "fakeAccountId",
	})
	if err != nil {
		t.Errorf("Expected to create request body but got %v", err)
		return
	}

	// Creates request
	request, err := http.NewRequest(http.MethodPost, "/profiles", bytes.NewReader(requestBody))
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(profilesController.Create)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusCreated {
		t.Errorf("Expected http.StatusCreated but got %d", responseRecorder.Code)
	}
}

func TestFindAllProfilesByAccountIdController(t *testing.T) {
	profilesController, reset, createProfile := getTestContext()
	defer reset()

	newProfile, _ := createProfile()

	// Creates request
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("profiles/account/%s", newProfile.AccountId), nil)
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"accountId": newProfile.AccountId,
	}
	request = mock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(profilesController.FindAllByAccountId)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
		return
	}

	// Validates response body
	var profiles []*models.Profile
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &profiles)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	// NOTE: Dereferencing pointers to compare their values and not their memory addresses
	if *profiles[0] != *newProfile {
		t.Errorf("Expected first element of Profile slice to equal New Profile but got %v", *profiles[0])
	}
}

func TestFindProfileByIdController(t *testing.T) {
	profilesController, reset, createProfile := getTestContext()
	defer reset()

	newProfile, _ := createProfile()

	// Creates request
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("profiles/%s", newProfile.Id), nil)
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"id": newProfile.Id,
	}
	request = mock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(profilesController.FindById)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
		return
	}

	// Validates response body
	var profile *models.Profile
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &profile)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	// NOTE: Dereferencing pointers to compare their values and not their memory addresses
	if *profile != *newProfile {
		t.Errorf("Expected Profile to equal New Profile but got %v", *profile)
	}
}

func TestUpdateProfileController(t *testing.T) {
	profilesController, reset, createProfile := getTestContext()
	defer reset()

	newProfile, _ := createProfile()

	// Creates request body
	updatedName := "Updated fake profile name"
	requestBody, err := json.Marshal(dtos.UpdateProfileDto{
		Name: updatedName,
	})
	if err != nil {
		t.Errorf("Expected to create request body but got %v", err)
		return
	}

	// Creates request
	request, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("profiles/%s", newProfile.Id), bytes.NewReader(requestBody))
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"id": newProfile.Id,
	}
	request = mock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(profilesController.Update)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
	}

	// Validates response body
	var profile *models.Profile
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &profile)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	if profile.Name != updatedName {
		t.Errorf("Expected Profile name to equal updated name but got %s", profile.Name)
	}

	// Validates new profile
	if newProfile.Name != updatedName {
		t.Errorf("Expected New Profile name to equal updated name but got %s", newProfile.Name)
		return
	}
}

func TestDeleteProfileController(t *testing.T) {
	profilesController, reset, createProfile := getTestContext()
	defer reset()

	newProfile, _ := createProfile()

	// Creates request
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("profiles/%s", newProfile.Id), nil)
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"id": newProfile.Id,
	}
	request = mock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(profilesController.Delete)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusNoContent {
		t.Errorf("Expected http.StatusNoContent but got %d", responseRecorder.Code)
	}
}