package services

import (
	genrePayloads "app/genres/application/payloads"
	"app/genres/application/ports"
	"app/genres/domain/factories"
	"app/genres/domain/models"
	"app/translations/application/services"
	"context"
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

// NOTE: No need to use a transaction here because Genres table only contains the Id so the only things we need to update in this case are the translations
func (gs *GenresService) Update(id string, updateGenrePayload *genrePayloads.UpdateGenrePayload) (*models.Genre, error) {
	genre, err := gs.FindById(id)
	if err != nil {
		return nil, err
	}

	translations, err := gs.TranslationsService.UpdateBatch(
		context.Background(),
		id,
		updateGenrePayload.UpdateTranslationPayloads,
	)
	if err != nil {
		return nil, err
	}

	genre.Labels = translations

	return genre, nil
}

func (gs *GenresService) Delete(id string) (string, error) {
	deletedGenreId, err := gs.TranslationsService.ExecuteTransaction(
		context.Background(),
		func(ctx context.Context) (any, error) {
			// Deletes genre
			_, err := gs.GenresRepository.Delete(ctx, id)
			if err != nil {
				return "", err
			}

			// Deletes all translations
			_, err = gs.TranslationsService.DeleteBatch(ctx, id)
			if err != nil {
				return "", err
			}

			return id, nil
		},
	)

	return deletedGenreId.(string), err
}

func (gs *GenresService) Create(createGenrePayload *genrePayloads.CreateGenrePayload) (*models.Genre, error) {
	genre, err := gs.TranslationsService.ExecuteTransaction(
		context.Background(),
		func(ctx context.Context) (any, error) {
			// Creates genre
			genre := gs.GenresFactory.Create()
			_, err := gs.GenresRepository.Create(ctx, genre)
			if err != nil {
				return nil, err
			}

			// Creates all translations
			translations, err := gs.TranslationsService.CreateBatch(ctx, genre.Id, createGenrePayload.CreateTranslationPayloads)
			if err != nil {
				return nil, err
			}

			genre.Labels = translations

			return genre, nil
		},
	)

	return genre.(*models.Genre), err
}