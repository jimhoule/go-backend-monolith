package services

import (
	"app/accounts/application/services"
	"app/accounts/domain/factories"
	"app/accounts/persistence/fake/repositories"
	"app/authentication/application/payloads"
	"app/crypto"
	"app/tokens"
	"app/uuid"
	"testing"
)

func getTestContext() (*AuthenticationService, func(), func(email string) (*Tokens, error)) {
	authenticationService := &AuthenticationService{
		AccountsService: &services.AccountsService{
			AccountsFactory: &factories.AccountsFactory{
				UuidService:   uuid.GetService(),
				CryptoService: crypto.GetService(),
			},
			AccountsRepository: &repositories.FakeAccountsRepository{},
		},
		TokensService: tokens.GetService(),
		CryptoService: crypto.GetService(),
	}

	register := func(email string) (*Tokens, error) {
		return authenticationService.Register(&payloads.RegisterPayload{
			FirstName: "Dummy first name",
			LastName: "Dummy last name",
			Email: email,
			Password: "1234",
			PlanId: "dummyPlanId",
		})
	}

	return authenticationService, repositories.ResetFakeAccountsRepository, register
}

func TestRegisterService(t *testing.T) {
	authenticationService, reset, register := getTestContext()
	defer reset()

	email := "dummy@dummy.com"
	tokens, err := register(email)
	if err != nil {
		t.Errorf("Expected Tokens but got %v", err)
		return
	}

	accessTokenPayload, err := authenticationService.TokensService.Decode(tokens.AccessToken)
	if err != nil {
		t.Errorf("Expected Access Token payload but got %v", err)
		return
	}

	if accessTokenPayload.Email != email {
		t.Errorf("Expected Access Token payload email to equal registration email but got %s", email)
		return
	}

	refreshTokenPayload, err := authenticationService.TokensService.Decode(tokens.RefreshToken)
	if err != nil {
		t.Errorf("Expected Refresh Token payload but got %v", err)
		return
	}

	if refreshTokenPayload.Email != email {
		t.Errorf("Expected Refresh Token payload email to equal registration email but got %s", email)
	}
}

func TestLoginService(t *testing.T) {
	authenticationService, reset, register := getTestContext()
	defer reset()

	email := "dummy@dummy.com"
	register(email)

	tokens, err := authenticationService.Login(&payloads.LoginPayload{
		Email: "dummy@dummy.com",
		Password: "1234",
	})
	if err != nil {
		t.Errorf("Expected Tokens but got %v", err)
	}

	accessTokenPayload, err := authenticationService.TokensService.Decode(tokens.AccessToken)
	if err != nil {
		t.Errorf("Expected Access Token payload but got %v", err)
		return
	}

	if accessTokenPayload.Email != email {
		t.Errorf("Expected Access Token payload email to equal registration email but got %s", email)
		return
	}

	refreshTokenPayload, err := authenticationService.TokensService.Decode(tokens.RefreshToken)
	if err != nil {
		t.Errorf("Expected Refresh Token payload but got %v", err)
		return
	}

	if refreshTokenPayload.Email != email {
		t.Errorf("Expected Refresh Token payload email to equal registration email but got %s", email)
	}
}