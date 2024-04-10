package controllers

import (
	"app/genres/application/payloads"
	genreServices "app/genres/application/services"
	"app/genres/domain/factories"
	"app/genres/domain/models"
	genresRepositories "app/genres/infrastructures/persistence/fake/repositories"
	"app/genres/presenters/http/dtos"
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

func getTestContext() (*GenresController, func(), func() (*models.Genre, error)) {
	genresController := &GenresController{
		GenresService: &genreServices.GenresService{
			GenresFactory: &factories.GenresFactory{
				UuidService: uuid.GetService(),
			},
			GenresRepository: &genresRepositories.FakeGenresRepository{},
			TranslationsService: &translationsServices.TranslationsService{
				TranslationsFactory: &translationsFactories.TranslationsFactory{},
				TranslationsRepository: &translationsRepositories.FakeTranslationsRepository{},
			},
			TransactionsService: &transactionsServices.TransactionsService{
				TransactionsRepository: &transactionsRepositories.FakeTransactionsRepository{},
			},
		},
	}

	createGenre := func() (*models.Genre, error) {
		return genresController.GenresService.Create(&payloads.CreateGenrePayload{})
	}

	return genresController, genresRepositories.ResetFakeGenresRepository, createGenre
}

func TestCreateGenreController(t *testing.T) {
	genresController, reset, _ := getTestContext()
	defer reset()

	// Creates request body
	requestBody, err := json.Marshal(dtos.CreateGenreDto{})
	if err != nil {
		t.Errorf("Expected to create request body but got %v", err)
		return
	}

	// Creates request
	request, err := http.NewRequest(http.MethodPost, "/genres", bytes.NewReader(requestBody))
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(genresController.Create)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusCreated {
		t.Errorf("Expected http.StatusCreated but got %d", responseRecorder.Code)
	}
}

func TestFindAllGenresController(t *testing.T) {
	genresController, reset, createGenre := getTestContext()
	defer reset()

	newGenre, _ := createGenre()

	// Creates request
	request, err := http.NewRequest(http.MethodGet, "/genres", nil)
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(genresController.FindAll)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
		return
	}

	// Validates response body
	var genres []*models.Genre
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &genres)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	if genres[0].Id != newGenre.Id {
		t.Errorf("Expected Genres id to equal New Genre id but got %v", genres[0].Id)
	}
}

func TestFindGenreByIdController(t *testing.T) {
	genresController, reset, createGenre := getTestContext()
	defer reset()

	newGenre, _ := createGenre()

	// Creates request
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/genres/%s", newGenre.Id), nil)
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"id": newGenre.Id,
	}
	request = mock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(genresController.FindById)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
		return
	}

	// Validates response body
	var genre *models.Genre
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &genre)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	if genre.Id != newGenre.Id {
		t.Errorf("Expected Genre id to equal New Genre id but got %v", genre.Id)
	}
}

func TestUpdateGenreController(t *testing.T) {
	genresController, reset, createGenre := getTestContext()
	defer reset()

	newGenre, _ := createGenre()

	// Creates request body
	requestBody, err := json.Marshal(dtos.CreateGenreDto{})
	if err != nil {
		t.Errorf("Expected to create request body but got %v", err)
		return
	}

	// Creates request
	request, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("/genres/%s", newGenre.Id), bytes.NewReader(requestBody))
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"id": newGenre.Id,
	}
	request = mock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(genresController.Update)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
	}

	// Validates response body
	var genre *models.Genre
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &genre)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	if genre.Id != newGenre.Id {
		t.Errorf("Expected Genre id to equal New Genre id but got %s", genre.Id)
	}
}

func TestDeleteGenreController(t *testing.T) {
	genresController, reset, createGenre := getTestContext()
	defer reset()

	newGenre, _ := createGenre()

	// Creates request
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/genres/%s", newGenre.Id), nil)
	if err != nil {
		t.Errorf("Expected to create a request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"id": newGenre.Id,
	}
	request = mock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(genresController.Delete)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusNoContent {
		t.Errorf("Expected http.StatusNoContent but got %d", responseRecorder.Code)
	}
}