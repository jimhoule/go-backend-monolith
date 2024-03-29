package controllers

import (
	"app/plans/application/payloads"
	"app/plans/application/services"
	"app/plans/domain/factories"
	"app/plans/domain/models"
	"app/plans/persistence/fake/repositories"
	"app/plans/presenters/http/dtos"
	"app/router/mock"
	"app/uuid"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestContext() (*PlansController, func(), func() (*models.Plan, error)) {
	plansController := &PlansController{
		PlansService: services.PlansService{
			PlansFactory: factories.PlansFactory{
				UuidService: uuid.GetService(),
			},
			PlansRepository: &repositories.FakePlansRepository{},
		},
	}

	createPlan := func() (*models.Plan, error) {
		return plansController.PlansService.Create(payloads.CreatePlanPayload{
			Name: "Dummy Plan name",
			Description: "Dummy Plan description",
			Price: 10.50,
		})
	}

	return plansController, repositories.ResetFakePlansRepository, createPlan
}

func TestCreatePlanController(t *testing.T) {
	plansController, reset, _ := getTestContext()
	defer reset()

	// Creates request body
	requestBody, err := json.Marshal(dtos.CreatePlanDto{
		Name: "Dummy Plan name",
		Description: "Dummy Plan description",
		Price: 10.50,
	})
	if err != nil {
		t.Errorf("Expected to create a request body but got %v", err)
		return
	}

	// Creates request
	request, err := http.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(requestBody))
	if err != nil {
		t.Errorf("Expected to create a new request but got %v", err)
		return
	}

	// Creates repsonse recorder (which satisfies http.ResponseWriter) to record the response
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(plansController.Create)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates the status code
	if responseRecorder.Code != http.StatusCreated {
		t.Errorf("Expected http.StatusCreated but got %d", responseRecorder.Code)
		return
	}
}

func TestFindAllPlansController(t *testing.T) {
	plansController, reset, createPlan := getTestContext()
	defer reset()

	newPlan, _ := createPlan()

	// Creates request
	request, err := http.NewRequest(http.MethodGet, "/accounts", nil)
	if err != nil {
		t.Errorf("Expected to create a new request but got %v", err)
		return
	}

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(plansController.FindAll)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %v", responseRecorder.Code)
		return
	}

	// Validates response body
	var plans []*models.Plan
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &plans)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	// NOTE: Dereferences pointers to compares the values and not the memory addresses (memory addresses are different but values are the same)
	if *plans[0] != *newPlan {
		t.Errorf("Expected first element of Plans slice to equal New Plan but got %v", *plans[0])
		return
	}
}

func TestFindPlanByIdController(t *testing.T) {
	plansController, reset, createPlan := getTestContext()
	defer reset()

	newPlan, _ := createPlan()

	// Creates request
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/accounts/%s", newPlan.Id), nil)
	if err != nil {
		t.Errorf("Expected to create a new request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"id": newPlan.Id,
	}
	request = mock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(plansController.FindById)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
		return
	}

	// Validates response body
	var plan *models.Plan
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &plan)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	// NOTE: Dereferencing pointers to compare their values and not their memory addresses
	if *plan != *newPlan {
		t.Errorf("Expected Plan to equal New Plan but got %v", *plan)
		return
	}
}