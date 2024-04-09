package services

import (
	"app/genres/application/payloads"
	"app/genres/domain/factories"
	"app/genres/domain/models"
	genresRepositories "app/genres/persistence/fake/repositories"
	transactionsServices "app/transactions/application/services"
	transactionsRepositories "app/transactions/persistence/fake/repositories"
	translationsServices "app/translations/application/services"
	translationsFactories "app/translations/domain/factories"
	translationsRepositories "app/translations/persistence/fake/repositories"
	"app/uuid"
	"testing"
)

func getTestContext() (*GenresService, func(), func() (*models.Genre, error)) {
	genresService := &GenresService{
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
	}

	createGenre := func() (*models.Genre, error) {
		return genresService.Create(&payloads.CreateGenrePayload{})
	}

	return genresService, genresRepositories.ResetFakeGenresRepository, createGenre
}

func TestCreateGenreService(t *testing.T) {
	_, reset, createGenre := getTestContext()
	defer reset()

	_, err := createGenre()
	if err != nil {
		t.Errorf("Expected to create Genre but got %v", err)
	}
}

func TestFindAllGenresService(t *testing.T) {
	genresService, reset, createGenre := getTestContext()
	defer reset()

	newGenre, _ := createGenre()

	genres, err := genresService.FindAll()
	if err != nil {
		t.Errorf("Expected slice of Genres but got %v", err)
		return
	}

	if genres[0].Id != newGenre.Id {
		t.Errorf("Expected Genre id to be equal to New Genre id but got %v", genres[0].Id)
	}
}

func TestFindGenreByIdService(t *testing.T) {
	genresService, reset, createGenre := getTestContext()
	defer reset()

	newGenre, _ := createGenre()

	genre, err := genresService.FindById(newGenre.Id)
	if err != nil {
		t.Errorf("Expected Genre but got %v", err)
		return
	}

	if genre.Id != newGenre.Id {
		t.Errorf("Expected Genre id to equal New Genre id but got %v", genre.Id)
	}
}

func TestUpdateGenreService(t *testing.T) {
	genresService, reset, createGenre := getTestContext()
	defer reset()

	newGenre, _ := createGenre()

	genre, err := genresService.Update(newGenre.Id, &payloads.UpdateGenrePayload{})
	if err != nil {
		t.Errorf("Expected Genre but got %v", err)
		return
	}

	if genre.Id != newGenre.Id {
		t.Errorf("Expected Genre id to equal New Genre id code but got %s", genre.Id)
		return
	}
}

func TestDeleteGenreService(t *testing.T) {
	genresService, reset, createGenre := getTestContext()
	defer reset()

	newGenre, _ := createGenre()

	genreId, err := genresService.Delete(newGenre.Id)
	if err != nil {
		t.Errorf("Expected Genre id but got %v", err)
		return
	}

	if genreId != newGenre.Id {
		t.Errorf("Expected Genre id to equal New Genre id but got %s", genreId)
	}
}