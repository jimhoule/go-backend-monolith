package controllers

import (
	"app/languages/application/payloads"
	"app/languages/application/services"
	"app/languages/domain/factories"
	"app/languages/domain/models"
	"app/languages/persistence/fake/repositories"
	"app/languages/presenters/http/dtos"
	"app/router/mock"
	"app/uuid"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestContext() (*LanguagesController, func(), func() (*models.Language, error)) {
	laguagesController := &LanguagesController{
		LanguagesService: &services.LanguagesService{
			LanguagesFactory: &factories.LanguagesFactory{
				UuidService: uuid.GetService(),
			},
			LanguagesRepository: &repositories.FakeLanguagesRepository{},
		},
	}

	createLanguage := func() (*models.Language, error) {
		return laguagesController.LanguagesService.Create(&payloads.CreateLanguagePayload{
			Code: "Fake code",
			Title: "Fake title",
		})
	}

	return laguagesController, repositories.ResetFakeLanguagesRepository, createLanguage
}

func TestCreateLanguageController(t *testing.T) {
	languagesController, reset, _ := getTestContext()
	defer reset()

	// Creates request body
	requestBody, err := json.Marshal(dtos.CreateLanguageDto{
		Code: "Fake code",
		Title: "Fake title",
	})
	if err != nil {
		t.Errorf("Expected to create request body but got %v", err)
		return
	}

	// Creates request
	request, err := http.NewRequest(http.MethodPost, "/languages", bytes.NewReader(requestBody))
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(languagesController.Create)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusCreated {
		t.Errorf("Expected http.StatusCreated but got %d", responseRecorder.Code)
	}
}

func TestFindAllLanguagesController(t *testing.T) {
	languagesController, reset, createLanguage := getTestContext()
	defer reset()

	newLanguage, _ := createLanguage()

	// Creates request
	request, err := http.NewRequest(http.MethodGet, "/languages", nil)
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(languagesController.FindAll)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
		return
	}

	// Validates response body
	var languages []*models.Language
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &languages)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	// NOTE: Dereferencing pointers to compare their values and not their memory addresses
	if *languages[0] != *newLanguage {
		t.Errorf("Expected first element of Languages slice to equal New Language but got %v", *languages[0])
	}
}

func TestFindLanguageByIdController(t *testing.T) {
	languagesController, reset, createLanguage := getTestContext()
	defer reset()

	newLanguage, _ := createLanguage()

	// Creates request
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/languages/%s", newLanguage.Id), nil)
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"id": newLanguage.Id,
	}
	request = mock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(languagesController.FindById)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
		return
	}

	// Validates response body
	var language *models.Language
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &language)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	// NOTE: Dereferencing pointers to compare their values and not their memory addresses
	if *language != *newLanguage {
		t.Errorf("Expected Language to equal New Language but got %v", *language)
	}
}

func TestUpdateLanguageController(t *testing.T) {
	languagesController, reset, createLanguage := getTestContext()
	defer reset()

	newLanguage, _ := createLanguage()

	// Creates request body
	updatedTitle := "Updated fake language title"
	requestBody, err := json.Marshal(dtos.UpdateLanguageDto{
		Code: newLanguage.Code,
		Title: updatedTitle,
	})
	if err != nil {
		t.Errorf("Expected to create request body but got %v", err)
		return
	}

	// Creates request
	request, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("/languages/%s", newLanguage.Id), bytes.NewReader(requestBody))
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"id": newLanguage.Id,
	}
	request = mock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(languagesController.Update)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
	}

	// Validates response body
	var language *models.Language
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &language)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	if language.Title != updatedTitle {
		t.Errorf("Expected Language title to equal updated ttle but got %s", language.Title)
	}

	// Validates new language
	if newLanguage.Title != updatedTitle {
		t.Errorf("Expected New Language title to equal updated title but got %s", newLanguage.Title)
		return
	}
}

func TestDeleteLanguageController(t *testing.T) {
	languagesController, reset, createLanguage := getTestContext()
	defer reset()

	newLanguage, _ := createLanguage()

	// Creates request
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/languages/%s", newLanguage.Id), nil)
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"id": newLanguage.Id,
	}
	request = mock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(languagesController.Delete)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusNoContent {
		t.Errorf("Expected http.StatusNoContent but got %d", responseRecorder.Code)
	}
}