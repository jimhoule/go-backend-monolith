package controllers

import (
	"app/accounts/application/payloads"
	"app/accounts/application/services"
	"app/accounts/domain/factories"
	"app/accounts/domain/models"
	"app/accounts/infrastructures/persistence/fake/repositories"
	"app/crypto"
	"app/router/mock"
	"app/uuid"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestContext() (*AccountsController, func(), func() (*models.Account, error)) {
	accountsController := &AccountsController{
		AccountsService: &services.AccountsService{
			AccountsFactory: &factories.AccountsFactory{
				UuidService: uuid.GetService(),
				CryptoService: crypto.GetService(),
			},
			AccountsRepository: &repositories.FakeAccountsRepository{},
		},
	}

	createAccount := func() (*models.Account, error) {
		return accountsController.AccountsService.Create(&payloads.CreateAccountPayload{
			FirstName: "Dummy first name",
			LastName:  "Dummy last name",
			Email:     "dummy@dummy.com",
			Password:  "1234",
			PlanId:    "dummyPlanId",
		})
	}

	return accountsController, repositories.ResetFakeAccountsRepository, createAccount
}

func TestFindAllAccountsController(t *testing.T) {
	accountsController, reset, createAccount := getTestContext()
	defer reset()

	newAccount, _ := createAccount()

	// Creates request
	request, err := http.NewRequest(http.MethodGet, "/accounts", nil)
	if err != nil {
		t.Errorf("Expected to create a new request but got %v", err)
		return
	}

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(accountsController.FindAll)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %v", responseRecorder.Code)
		return
	}

	// Validates response body
	var accounts []*models.Account
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &accounts)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	// NOTE: Dereferences pointers to compares the values and not the memory addresses (memory addresses are different but values are the same)
	if *accounts[0] != *newAccount {
		t.Errorf("Expected first element of Accounts slice to equal New Account but got %v", *accounts[0])
		return
	}
}

func TestFindAccountByIdController(t *testing.T) {
	accountsController, reset, createAccount := getTestContext()
	defer reset()

	newAccount, _ := createAccount()

	// Creates request
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/accounts/%s", newAccount.Id), nil)
	if err != nil {
		t.Errorf("Expected to create a new request but got %v", err)
		return
	}

	// NOTE: Adds chi URL params context to request
	urlParams := map[string]string{
		"id": newAccount.Id,
	}
	request = mock.GetRequestWithUrlParams(request, urlParams)

	// Creates response recorder
	responseRecorder := httptest.NewRecorder()
	// Creates handler
	handler := http.HandlerFunc(accountsController.FindById)
	// Executes request
	handler.ServeHTTP(responseRecorder, request)

	// Validates status code
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected http.StatusOK but got %d", responseRecorder.Code)
		return
	}

	// Validates response body
	var account *models.Account
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &account)
	if err != nil {
		t.Errorf("Expected to unmarshal response body but got %v", err)
		return
	}

	// NOTE: Dereferencing pointers to compare their values and not their memory addresses
	if *account != *newAccount {
		t.Errorf("Expected Account to equal New Account but got %v", *account)
		return
	}
}