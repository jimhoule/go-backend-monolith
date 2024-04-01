package controllers

import (
	accountsService "app/accounts/application/services"
	"app/accounts/domain/factories"
	"app/accounts/persistence/fake/repositories"
	"app/authentication/application/payloads"
	authenticationService "app/authentication/application/services"
	"app/authentication/presenters/http/dtos"
	"app/crypto"
	"app/tokens"
	"app/uuid"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestContext() (*AuthenticationController, func(), func() (*authenticationService.Tokens, error)) {
	authenticationController := &AuthenticationController{
		AuthenticationService: authenticationService.AuthenticationService{
			AccountsService: accountsService.AccountsService{
				AccountsFactory: factories.AccountsFactory{
					UuidService:   uuid.GetService(),
					CryptoService: crypto.GetService(),
				},
				AccountsRepository: &repositories.FakeAccountsRepository{},
			},
			TokensService: tokens.GetService(),
			CryptoService: crypto.GetService(),
		},
	}

	register := func() (*authenticationService.Tokens, error) {
		return authenticationController.AuthenticationService.Register(payloads.RegisterPayload{
			FirstName: "Dummy first name",
			LastName: "Dummy last name",
			Email: "dummy@dummy.com",
			Password: "1234",
			PlanId: "dummyPlanId",
		})
	}

	return authenticationController, repositories.ResetFakeAccountsRepository, register
}

func TestRegisterController(t *testing.T) {
	authenticationController, reset, _ := getTestContext()
	defer reset()

	// Creates request body
	requestBody, err := json.Marshal(dtos.RegisterDto{
		FirstName: "Dummy first name",
		LastName: "Dummy last name",
		Email: "dummy@dummy.com",
		Password: "1234",
		PlanId: "dummyPlanId",
	})
	if err != nil {
		t.Errorf("Expected to create a request body but got %v", err)
		return
	}

	// Creates request
	request, err := http.NewRequest(http.MethodPost, "/authentication/register", bytes.NewReader(requestBody))
	if err != nil {
		t.Errorf("Expected to create a new request but got %v", err)
		return
	}

	// Creates response recorder (which satisfies http.ResponseWriter) to record the response
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(authenticationController.Register)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates the status code
	if responseRecorder.Code != http.StatusCreated {
		t.Errorf("Expected http.StatusCreated but got %d", responseRecorder.Code)
		return
	}
}

func TestLoginController(t *testing.T) {
	authenticationController, reset, register := getTestContext()
	defer reset()

	register()

	// Creates request body
	requestBody, err := json.Marshal(dtos.LoginDto{
		Email: "dummy@dummy.com",
		Password: "1234",
	})
	if err != nil {
		t.Errorf("Expected to create request body but got %v", err)
		return
	}

	// Creates request
	request, err := http.NewRequest(http.MethodPost, "/authentication/login", bytes.NewReader(requestBody))
	if err != nil {
		t.Errorf("Expected to create request but got %v", err)
		return
	}

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(authenticationController.Login)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %v", responseRecorder.Code)
	}
}