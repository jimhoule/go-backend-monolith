package services

import (
	"app/languages/application/payloads"
	"app/languages/application/ports"
	"app/languages/domain/factories"
	"app/languages/domain/models"
	transactionsServices "app/transactions/application/services"
	"app/translations/application/constants"
	translationsServices "app/translations/application/services"
	"context"
)

type LanguagesService struct {
	LanguagesFactory *factories.LanguagesFactory
	LanguagesRepository ports.LanguagesRepositoryPort
	TranslationsService *translationsServices.TranslationsService
	TransactionsService *transactionsServices.TransactionsService
}

func (ls *LanguagesService) FindAll() ([]*models.Language, error) {
	languages, err := ls.LanguagesRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, language := range languages {
		labelTranslations, err := ls.TranslationsService.FindAllByEntityIdAndType(language.Id, constants.TanslationTypeLabel)
		if err != nil {
			return nil, err
		}

		language.Labels = labelTranslations
	}

	return languages, nil
}

func (ls *LanguagesService) FindById(id string) (*models.Language, error) {
	language, err := ls.LanguagesRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	labelTranslations, err := ls.TranslationsService.FindAllByEntityIdAndType(language.Id, constants.TanslationTypeLabel)
	if err != nil {
		return nil, err
	}

	language.Labels = labelTranslations

	return language, nil
}

func (ls *LanguagesService) Update(id string, updateLanguagePayload *payloads.UpdateLanguagePayload) (*models.Language, error) {
	language, err := ls.TransactionsService.Execute(
		context.Background(),
		func(ctx context.Context) (any, error) {
			// Updates language
			language, err := ls.LanguagesRepository.Update(ctx, id, &models.Language{
				Code: updateLanguagePayload.Code,
			})
			if err != nil {
				return nil, err
			}

			// Upserts translations
			labelTranslations, err := ls.TranslationsService.UpsertBatch(
				ctx,
				language.Id,
				constants.TanslationTypeLabel,
				updateLanguagePayload.UpdateLabelTranslationPayloads,
			)
			if err != nil {
				return nil, err
			}

			language.Labels = labelTranslations

			return language, nil
		},
	)
	if err != nil {
		return nil, err
	}

	return language.(*models.Language), nil
}

func (ls *LanguagesService) Delete(id string) (string, error) {
	/*
	 * NOTES:
	 *	- Deletes language
	 *	- There is no need to delete all translation by language id manually as there is
	 *    an ON DELETE CASCADE constraint on the languageId foreign key in the translations table
	 */
	_, err := ls.LanguagesRepository.Delete(id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (ls *LanguagesService) Create(createLanguagePayload *payloads.CreateLanguagePayload) (*models.Language, error) {
	language, err := ls.TransactionsService.Execute(
		context.Background(),
		func(ctx context.Context) (any, error) {
			language := ls.LanguagesFactory.Create(createLanguagePayload.Code)

			// Creates language
			_, err := ls.LanguagesRepository.Create(ctx, language)
			if err != nil {
				return nil, err
			}


			// Creates translations
			labelTranslations, err := ls.TranslationsService.CreateBatch(
				ctx,
				language.Id,
				constants.TanslationTypeLabel,
				createLanguagePayload.CreateLabelTranslationPayloads,
			)
			if err != nil {
				return nil, err
			}

			language.Labels = labelTranslations

			return language, nil
		},
	)
	if err != nil {
		return nil, err
	}

	return language.(*models.Language), nil
}