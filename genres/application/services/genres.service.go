package services

import (
	genrePayloads "app/genres/application/payloads"
	"app/genres/application/ports"
	"app/genres/domain/factories"
	"app/genres/domain/models"
	"app/translations/application/services"
)

type GenresService struct {
	GenresFactory *factories.GenresFactory
	GenresRepository ports.GenresRepositoryPort
	TranslationsService *services.TranslationsService
}

func (gs *GenresService) FindAll() ([]*models.Genre, error) {
	genres, err := gs.GenresRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, genre := range genres {
		translations, err := gs.TranslationsService.FindAllByEntityId(genre.Id)
		if err != nil {
			return nil, err
		}

		genre.Labels = translations
	}

	return genres, nil
}

func (gs *GenresService) FindById(id string) (*models.Genre, error) {
	genre, err := gs.GenresRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	translations, err := gs.TranslationsService.FindAllByEntityId(genre.Id)
	if err != nil {
		return nil, err
	}

	genre.Labels = translations

	return genre, nil
}

func (gs *GenresService) Create(createGenrePayload *genrePayloads.CreateGenrePayload) (*models.Genre, error) {
	// Creates genre
	genre := gs.GenresFactory.Create()
	_, err := gs.GenresRepository.Create(genre)
	if err != nil {
		return nil, err
	}

	// Adds Translations entity id
	for _, createTranslationPayload := range createGenrePayload.CreateTranslationPayloads {
		createTranslationPayload.EntityId = genre.Id
	}

	// Creates translations of genre labels
	translations, err := gs.TranslationsService.Create(createGenrePayload.CreateTranslationPayloads)
	if err != nil {
		return nil, err
	}

	genre.Labels = translations

	return genre, nil
}