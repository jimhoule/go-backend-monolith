package services

import (
	genrePayloads "app/genres/application/payloads"
	"app/genres/application/ports"
	"app/genres/domain/factories"
	"app/genres/domain/models"
	translationPayloads "app/translations/application/payloads"
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
		translation, err := gs.TranslationsService.FindByCompositeId(genre.Id, "en")
		if err != nil {
			return nil, err
		}

		genre.Label = translation
	}

	return genres, nil
}

func (gs *GenresService) FindById(id string) (*models.Genre, error) {
	genre, err := gs.GenresRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	translation, err := gs.TranslationsService.FindByCompositeId(genre.Id, "en")
	if err != nil {
		return nil, err
	}

	genre.Label = translation

	return genre, nil
}

func (gs *GenresService) Create(createGenrePayload *genrePayloads.CreateGenrePayload) (*models.Genre, error) {
	// Creates genre
	genre := gs.GenresFactory.Create()
	_, err := gs.GenresRepository.Create(genre)
	if err != nil {
		return nil, err
	}

	// Creates translation of genre label
	translation, err := gs.TranslationsService.Create(&translationPayloads.CreateTranslationPayload{
		EntityId: genre.Id,
		LanguageCode: "en",
		Text: "english text",
	})
	if err != nil {
		return nil, err
	}

	genre.Label = translation

	return genre, nil
}