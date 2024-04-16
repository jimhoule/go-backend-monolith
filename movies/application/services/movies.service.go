package services

import (
	"app/movies/application/payloads"
	"app/movies/application/ports"
	"app/movies/domain/factories"
	"app/movies/domain/models"
	transactionsServices "app/transactions/application/services"
	"app/translations/application/constants"
	translationsServices "app/translations/application/services"
	"context"
)

type MoviesService struct {
	MoviesFactory *factories.MoviesFactory
	MoviesRepository ports.MoviesRepositoryPort
	MoviesStorage ports.MoviesStoragePort
	TranslationsService *translationsServices.TranslationsService
	TransactionsService *transactionsServices.TransactionsService
}

func (ms *MoviesService) FindAll() ([]*models.Movie, error) {
	movies, err := ms.MoviesRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, movie := range movies {
		titleTranslations, err := ms.TranslationsService.FindAllByEntityIdAndType(movie.Id, constants.TanslationTypeTitle)
		if err != nil {
			return nil, err
		}

		descriptionTranslations, err := ms.TranslationsService.FindAllByEntityIdAndType(movie.Id, constants.TanslationTypeDescription)
		if err != nil {
			return nil, err
		}

		movie.Titles = titleTranslations
		movie.Descriptions = descriptionTranslations
	}

	return movies, nil
}

func (ms *MoviesService) FindById(id string) (*models.Movie, error) {
	movie, err := ms.MoviesRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	titleTranslations, err := ms.TranslationsService.FindAllByEntityIdAndType(movie.Id, constants.TanslationTypeTitle)
	if err != nil {
		return nil, err
	}

	descriptionTranslations, err := ms.TranslationsService.FindAllByEntityIdAndType(movie.Id, constants.TanslationTypeDescription)
	if err != nil {
		return nil, err
	}

	movie.Titles = titleTranslations
	movie.Descriptions = descriptionTranslations

	return movie, nil
}

func (ms *MoviesService) Update(id string, updateMoviePayload *payloads.UpdateMoviePayload) (*models.Movie, error) {
	movie, err := ms.TransactionsService.Execute(
		context.Background(),
		func(ctx context.Context) (any, error) {
			// Updates movie
			movie, err := ms.MoviesRepository.Update(ctx, id, &models.Movie{
				GenreId: updateMoviePayload.GenreId,
			})
			if err != nil {
				return nil, err
			}

			//Updates translations
			titleTranslations, err := ms.TranslationsService.UpsertBatch(
				ctx,
				movie.Id,
				constants.TanslationTypeTitle,
				updateMoviePayload.UpdateTitleTranslationPayloads,
			)
			if err != nil {
				return nil, err
			}

			descriptionTranslations, err := ms.TranslationsService.UpsertBatch(
				ctx,
				movie.Id,
				constants.TanslationTypeDescription,
				updateMoviePayload.UpdateTitleTranslationPayloads,
			)
			if err != nil {
				return nil, err
			}

			movie.Titles = titleTranslations
			movie.Descriptions = descriptionTranslations

			return movie, nil
		},
	)
	if err != nil {
		return nil, err
	}

	return movie.(*models.Movie), nil
}

func (ms *MoviesService) Delete(id string) (string, error) {
	_, err := ms.TransactionsService.Execute(
		context.Background(),
		func(ctx context.Context) (any, error) {
			// Deletes movie
			_, err := ms.MoviesRepository.Delete(ctx, id)
			if err != nil {
				return nil, err
			}

			//Deletes translations
			_, err = ms.TranslationsService.DeleteBatch(ctx, id)
			if err != nil {
				return nil, err
			}

			_, err = ms.TranslationsService.DeleteBatch(ctx, id)
			if err != nil {
				return nil, err
			}


			return id, nil
		},
	)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (ms *MoviesService) Create(createMoviePayload *payloads.CreateMoviePayload) (*models.Movie, error) {
	movie, err := ms.TransactionsService.Execute(
		context.Background(),
		func(ctx context.Context) (any, error) {
			movie := ms.MoviesFactory.Create(createMoviePayload.GenreId)

			// Creates movie
			_, err := ms.MoviesRepository.Create(ctx, movie)
			if err != nil {
				return nil, err
			}

			//Updates translations
			titleTranslations, err := ms.TranslationsService.CreateBatch(
				ctx,
				movie.Id,
				constants.TanslationTypeTitle,
				createMoviePayload.CreateTitleTranslationPayloads,
			)
			if err != nil {
				return nil, err
			}

			descriptionTranslations, err := ms.TranslationsService.CreateBatch(
				ctx,
				movie.Id,
				constants.TanslationTypeDescription,
				createMoviePayload.CreateTitleTranslationPayloads,
			)
			if err != nil {
				return nil, err
			}

			movie.Titles = titleTranslations
			movie.Descriptions = descriptionTranslations

			return movie, nil
		},
	)
	if err != nil {
		return nil, err
	}

	return movie.(*models.Movie), nil
}

func (ms *MoviesService) Upload(filePath string, file []byte) (bool, error) {
	return ms.MoviesStorage.Upload(filePath, file)
}

func (ms *MoviesService) Download(filePath string) ([]byte, error) {
	return ms.MoviesStorage.Download(filePath)
}