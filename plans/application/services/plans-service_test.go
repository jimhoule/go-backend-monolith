package services

import (
	"app/plans/application/payloads"
	"app/plans/domain/factories"
	"app/plans/domain/models"
	"app/plans/persistence/fake/repositories"
	"app/uuid"
	"testing"
)

func getTestContext() (*PlansService, func(), func() (*models.Plan, error)) {
	plansService := &PlansService{
		PlansFactory: &factories.PlansFactory{
			UuidService: uuid.GetService(),
		},
		PlansRepository: &repositories.FakePlansRepository{},
	}

	createPlan := func() (*models.Plan, error) {
		return plansService.Create(&payloads.CreatePlanPayload{
			Price: 10.50,
		})
	}

	return plansService, repositories.ResetFakePlansRepository, createPlan
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
	newAccount, _ := createPlan()

	plans, err := plansService.FindAll()
	if err != nil {
		t.Errorf("Expected to find all Plans but got %v", err)
		return
	}

	if len(plans) != 1 {
		t.Errorf("Expected slice of Plans with a length of 1 but got %d", len(plans))
		return
	}

	if plans[0] != newAccount {
		t.Errorf("Expected first Plan of slice to be equal to New Plan but got %v", plans[0])
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

	if newPlan != plan {
		t.Errorf("Expected Plan to be equal to New Plan but got %v", plan)
		return
	}
}