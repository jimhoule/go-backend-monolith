package services

import (
	"app/plans/application/payloads"
	"app/plans/domain/factories"
	"app/plans/domain/models"
	plansRepositories "app/plans/infrastructures/persistence/fake/repositories"
	transactionsServices "app/transactions/application/services"
	transactionsRepositories "app/transactions/infrastructures/persistence/fake/repositories"
	translationsServices "app/translations/application/services"
	translationsFactories "app/translations/domain/factories"
	translationsRepositories "app/translations/infrastructures/persistence/fake/repositories"
	"app/uuid"
	"testing"
)

func getTestContext() (*PlansService, func(), func() (*models.Plan, error)) {
	plansService := &PlansService{
		PlansFactory: &factories.PlansFactory{
			UuidService: uuid.GetService(),
		},
		PlansRepository: &plansRepositories.FakePlansRepository{},
		TranslationsService: &translationsServices.TranslationsService{
			TranslationsFactory: &translationsFactories.TranslationsFactory{},
			TranslationsRepository: &translationsRepositories.FakeTranslationsRepository{},
		},
		TransactionsService: &transactionsServices.TransactionsService{
			TransactionsRepository: &transactionsRepositories.FakeTransactionsRepository{},
		},
	}

	createPlan := func() (*models.Plan, error) {
		return plansService.Create(&payloads.CreatePlanPayload{
			Price: 10.50,
		})
	}

	return plansService, plansRepositories.ResetFakePlansRepository, createPlan
}

func TestCreatePlanService(t *testing.T) {
	_, reset, createPlan := getTestContext()
	defer reset()

	_, err := createPlan()
	if err != nil {
		t.Errorf("Expected to create a Plan but got %v", err)
		return
	}
}

func TestFindAllPlansService(t *testing.T) {
	plansService, reset, createPlan := getTestContext()
	defer reset()
	newPlan, _ := createPlan()

	plans, err := plansService.FindAll()
	if err != nil {
		t.Errorf("Expected to find all Plans but got %v", err)
		return
	}

	if len(plans) != 1 {
		t.Errorf("Expected slice of Plans with a length of 1 but got %d", len(plans))
		return
	}

	if plans[0].Id != newPlan.Id {
		t.Errorf("Expected Plan id to be equal to New Plan id but got %v", plans[0].Id)
		return
	}
}

func TestFindPlanByIdService(t *testing.T) {
	plansService, reset, createPlan := getTestContext()
	defer reset()
	newPlan, _ := createPlan()

	plan, err := plansService.FindById(newPlan.Id)
	if err != nil {
		t.Errorf("Expected to find a Plan by id but got %v", err)
		return
	}

	if plan.Id != newPlan.Id {
		t.Errorf("Expected Plan id to be equal to New Plan id but got %v", plan.Id)
		return
	}
}