package services

import (
	"app/plans/application/payloads"
	"app/plans/application/ports"
	"app/plans/domain/factories"
	"app/plans/domain/models"
	transactionsServices "app/transactions/application/services"
	"app/translations/application/constants"
	translationsServices "app/translations/application/services"
	"context"
)

type PlansService struct {
	PlansFactory *factories.PlansFactory
	PlansRepository ports.PlansRepositoryPort
	TranslationsService *translationsServices.TranslationsService
	TransactionsService *transactionsServices.TransactionsService
}

func (ps *PlansService) FindAll() ([]*models.Plan, error) {
	plans, err := ps.PlansRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, plan := range plans {
		labelTranslations, err := ps.TranslationsService.FindAllByEntityIdAndType(plan.Id, constants.TanslationTypeLabel)
		if err != nil {
			return nil, err
		}

		descriptionTranslations, err := ps.TranslationsService.FindAllByEntityIdAndType(plan.Id, constants.TanslationTypeDescription)
		if err != nil {
			return nil, err
		}

		plan.Labels = labelTranslations
		plan.Descriptions = descriptionTranslations
	}

	return plans, nil
}

func (ps *PlansService) FindById(id string) (*models.Plan, error) {
	plan, err := ps.PlansRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	labelTranslations, err := ps.TranslationsService.FindAllByEntityIdAndType(plan.Id, constants.TanslationTypeLabel)
	if err != nil {
		return nil, err
	}

	descriptionTranslations, err := ps.TranslationsService.FindAllByEntityIdAndType(plan.Id, constants.TanslationTypeDescription)
	if err != nil {
		return nil, err
	}

	plan.Labels = labelTranslations
	plan.Descriptions = descriptionTranslations

	return plan, nil
}

func (ps *PlansService) Create(createPlanPayload *payloads.CreatePlanPayload) (*models.Plan, error) {
	plan, err := ps.TransactionsService.Execute(
		context.Background(),
		func(ctx context.Context) (any, error) {
			plan := ps.PlansFactory.Create(createPlanPayload.Price)

			// Creates plan
			_, err := ps.PlansRepository.Create(ctx, plan)
			if err != nil {
				return nil, err
			}

			// Creates label translations
			labelTranslations, err := ps.TranslationsService.CreateBatch(
				ctx,
				plan.Id,
				constants.TanslationTypeLabel,
				createPlanPayload.CreateLabelTranslationPayloads,
			)
			if err != nil {
				return nil, err
			}

			// Creates description translations
			descriptionTranslations, err := ps.TranslationsService.CreateBatch(
				ctx,
				plan.Id,
				constants.TanslationTypeLabel,
				createPlanPayload.CreateDescriptionTranslationPayloads,
			)
			if err != nil {
				return nil, err
			}

			plan.Labels = labelTranslations
			plan.Descriptions = descriptionTranslations

			return plan, nil
		},
	)
	if err != nil {
		return nil, err
	}

	return plan.(*models.Plan), nil
}