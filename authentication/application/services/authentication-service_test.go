package services

import (
	"app/accounts/application/services"
	"app/accounts/domain/factories"
	"app/accounts/domain/models"
	"app/accounts/persistence/fake/repositories"
	"app/authentication/application/payloads"
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

	register := func() (*models.Account, error) {
		return authenticationService.Register(payloads.RegisterPayload{
			FirstName: "Dummy first name",
			LastName: "Dummy last name",
			Email: "dummy@dummy.com",
			Password: "1234",
			PlanId: "dummyPlanId",
		})
	}

	return authenticationService, repositories.ResetFakeAccountsRepository, register
}

func TestRegisterService(t *testing.T) {
	_, reset, register := getTestContext()
	defer reset()

	_, err := register()
	if err != nil {
		t.Errorf("Expected Tokens but got %v", err)
	}
}

func TestLoginService(t *testing.T) {
	authenticationService, reset, register := getTestContext()
	defer reset()

	account, _ := register()

	_, err := authenticationService.Login(payloads.LoginPayload{
		Email: account.Email,
		Password: "1234",
	})
	if err != nil {
		t.Errorf("Expected Tokens but got %v", err)
	}
}