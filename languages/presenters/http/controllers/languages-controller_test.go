package controllers

import (
	"app/languages/application/payloads"
	languageServices "app/languages/application/services"
	languagesFactories "app/languages/domain/factories"
	"app/languages/domain/models"
	languagesRepositories "app/languages/infrastructures/persistence/fake/repositories"
	"app/languages/presenters/http/dtos"
	"app/router/mock"
	transactionsServices "app/transactions/application/services"
	transactionsRepositories "app/transactions/infrastructures/persistence/fake/repositories"
	translationsServices "app/translations/application/services"
	translationsFactories "app/translations/domain/factories"
	translationsRepositories "app/translations/infrastructures/persistence/fake/repositories"
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
		LanguagesService: &languageServices.LanguagesService{
			LanguagesFactory: &languagesFactories.LanguagesFactory{
				UuidService: uuid.GetService(),
			},
			LanguagesRepository: &languagesRepositories.FakeLanguagesRepository{},
			TranslationsService: &translationsServices.TranslationsService{
				TranslationsFactory: &translationsFactories.TranslationsFactory{},
				TranslationsRepository: &translationsRepositories.FakeTranslationsRepository{},
			},
			TransactionsService: &transactionsServices.TransactionsService{
				TransactionsRepository: &transactionsRepositories.FakeTransactionsRepository{},
			},
		},
	}

	createLanguage := func() (*models.Language, error) {
		return laguagesController.LanguagesService.Create(&payloads.CreateLanguagePayload{
			Code: "Fake code",
		})
	}

	return laguagesController, languagesRepositories.ResetFakeLanguagesRepository, createLanguage
}

func TestCreateLanguageController(t *testing.T) {
	languagesController, reset, _ := getTestContext()
	defer reset()

	// Creates request body
	requestBody, err := json.Marshal(dtos.CreateLanguageDto{
		Code: "Fake code",
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

	if languages[0].Id != newLanguage.Id {
		t.Errorf("Expected Languages id to equal New Language id but got %v", languages[0].Id)
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

	if language.Id != newLanguage.Id {
		t.Errorf("Expected Language id to equal New Language id but got %v", language.Id)
	}
}

func TestUpdateLanguageController(t *testing.T) {
	languagesController, reset, createLanguage := getTestContext()
	defer reset()

	newLanguage, _ := createLanguage()

	// Creates request body
	updatedCode := "Updated fake language code"
	requestBody, err := json.Marshal(dtos.UpdateLanguageDto{
		Code: updatedCode,
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

	if updatedCode != language.Code {
		t.Errorf("Expected updated code to equal Language code but got %s", updatedCode)
	}

	// Validates new language
	if updatedCode != newLanguage.Code {
		t.Errorf("Expected updatedCode code to equal New Language code but got %s", updatedCode)
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