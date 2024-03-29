package controllers

import (
	accountPayload "app/accounts/application/payloads"
	accountsService "app/accounts/application/services"
	"app/accounts/domain/factories"
	"app/accounts/domain/models"
	"app/accounts/persistence/fake/repositories"
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

func getTestContext() (*AuthenticationController, func(), func() (*models.Account, error)) {
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

	createAccount := func() (*models.Account, error) {
		return authenticationController.AuthenticationService.AccountsService.Create(accountPayload.CreateAccountPayload{
			FirstName: "Dummy first name",
			LastName: "Dummy last name",
			Email: "dummy@dummy.com",
			Password: "1234",
			PlanId: "dummyPlanId",
		})
	}

	return authenticationController, repositories.ResetFakeAccountsRepository, createAccount
}

func TestLoginController(t *testing.T) {
	authenticationController, reset, createAccount := getTestContext()
	defer reset()

	newAccount, _ := createAccount()

	// Creates request body
	requestBody, err := json.Marshal(dtos.LoginDto{
		Email: newAccount.Email,
		Password: "1234",
	})
	if err != nil {
		t.Errorf("Expected to create request body but got %v", err)
		return
	}

	// Creates request
	request, err := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(requestBody))
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