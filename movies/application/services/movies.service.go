package services

import (
	"app/movies/application/payloads"
	"app/movies/application/ports"
	"app/movies/domain/factories"
	"app/movies/domain/models"
	transactionsServices "app/transactions/application/services"
	transcoderServices "app/transcoder/services"
	"app/translations/application/constants"
	translationsServices "app/translations/application/services"
	"context"
	"fmt"
	"os"
)

type MoviesService struct {
	MoviesFactory *factories.MoviesFactory
	MoviesRepository ports.MoviesRepositoryPort
	MoviesStorage ports.MoviesStoragePort
	TranscoderService transcoderServices.TranscoderService
	TransactionsService *transactionsServices.TransactionsService
	TranslationsService *translationsServices.TranslationsService
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

func (ms *MoviesService) Upload(uploadMoviePayload *payloads.UploadMoviePayload) (bool, error) {
	// Uploads file to storage
	_, err := ms.MoviesStorage.UploadLarge(
		uploadMoviePayload.FilePath,
		uploadMoviePayload.File,
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (ms *MoviesService) TranscodeDashAndUploadVideo(transcodeDashAndUploadMoviePayload *payloads.TranscodeDashAndUploadVideoPayload) (bool, error) {
	// Creates temp dir to store uploaded video
	tempDirPath, err := os.MkdirTemp("", transcodeDashAndUploadMoviePayload.FileName)
	if err != nil {
		return false, err
	}
	defer os.RemoveAll(tempDirPath)

	// Writes video into temp dir
	err = os.WriteFile(
		fmt.Sprintf(
			"%s/%s.%s",
			tempDirPath,
			transcodeDashAndUploadMoviePayload.FileName,
			transcodeDashAndUploadMoviePayload.FileExtension,
		),
		transcodeDashAndUploadMoviePayload.File,
		0777,
	)
	if err != nil {
		return false, err
	}

	// Creates DASH files in temp dir based on video in same temp dir
	err = ms.TranscoderService.TranscodeDash(
		tempDirPath,
		transcodeDashAndUploadMoviePayload.FileName,
		transcodeDashAndUploadMoviePayload.FileExtension,
		transcodeDashAndUploadMoviePayload.OnTranscodingProgressSent,
	)
	if err != nil {
		return false, err
	}

	// Emits start of upload
	transcodeDashAndUploadMoviePayload.OnUploadStarted()

	// Gets temp dir content
	tempDirEntries, err := os.ReadDir(tempDirPath)
	if err != nil {
		return false, err
	}


	for _, tempDirEntry := range tempDirEntries {
		tempDirEntryName := tempDirEntry.Name()

		// Reads file from temp dir
		tempFile, err := os.ReadFile(fmt.Sprintf("%s/%s", tempDirPath, tempDirEntryName))
		if err != nil {
			return false, err
		}

		// Gets storage upload file path
		uploadFilePath := fmt.Sprintf("%s/dash/%s", transcodeDashAndUploadMoviePayload.FileName, tempDirEntryName)
		// NOTE: Sets file path to root folder if temp dir entry is original video
		if tempDirEntryName == fmt.Sprintf("%s.%s", transcodeDashAndUploadMoviePayload.FileName, transcodeDashAndUploadMoviePayload.FileExtension) {
			uploadFilePath = fmt.Sprintf("%s/%s", transcodeDashAndUploadMoviePayload.FileName, tempDirEntryName)
		}

		// Uploads file to storage
		_, err = ms.MoviesStorage.UploadLarge(
			uploadFilePath,
			tempFile,
		)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}