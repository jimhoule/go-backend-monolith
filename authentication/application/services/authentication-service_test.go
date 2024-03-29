package services

import (
	accountPayload "app/accounts/application/payloads"
	"app/accounts/application/services"
	"app/accounts/domain/factories"
	"app/accounts/domain/models"
	"app/accounts/persistence/fake/repositories"
	"app/crypto"
	"app/tokens"
	"app/uuid"
	"testing"
)

func getTestContext() (*AuthenticationService, func(), func() (*models.Account, error)) {
	authenticationService := &AuthenticationService{
		AccountsService: services.AccountsService{
			AccountsFactory: factories.AccountsFactory{
				UuidService:   uuid.GetService(),
				CryptoService: crypto.GetService(),
			},
			AccountsRepository: &repositories.FakeAccountsRepository{},
		},
		TokensService: tokens.GetService(),
		CryptoService: crypto.GetService(),
	}

	createAccount := func() (*models.Account, error) {
		return authenticationService.AccountsService.Create(accountPayload.CreateAccountPayload{
			FirstName: "Dummy first name",
			LastName: "Dummy last name",
			Email: "dummy@dummy.com",
			Password: "1234",
			PlanId: "dummyPlanId",
		})
	}

	return authenticationService, repositories.ResetFakeAccountsRepository, createAccount
}

func TestLoginService(t *testing.T) {
	authenticationService, reset, createAccount := getTestContext()
	defer reset()

	account, _ := createAccount()

	_, err := authenticationService.Login(account.Email, "1234")
	if err != nil {
		t.Errorf("Expected Tokens but got %v", err)
	}
}