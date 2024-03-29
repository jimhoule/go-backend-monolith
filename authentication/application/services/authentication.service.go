package services

import (
	accountPayloads "app/accounts/application/payloads"
	accountsService "app/accounts/application/services"
	"app/accounts/domain/models"
	authenticationPayloads "app/authentication/application/payloads"
	cryptoService "app/crypto/services"
	tokensService "app/tokens/services"
)

type AuthenticationService struct{
	AccountsService accountsService.AccountsService
	TokensService   tokensService.TokensService
	CryptoService   cryptoService.CryptoService
}

type Tokens struct{
	AccessToken  string
	RefreshToken string
}

func (as *AuthenticationService) Login(loginPayload authenticationPayloads.LoginPayload) (*Tokens, error) {
	tokens := &Tokens{}

	account, err := as.AccountsService.FindByEmail(loginPayload.Email)
	if err != nil {
		return tokens, err
	}

	isValid, err := as.CryptoService.ComparePassword(account.Password, loginPayload.Password)
	if !isValid {
		return tokens, err
	}

	accessToken, err := as.TokensService.GenerateAccessToken(account.Id, account.Email)
	if err != nil {
		return tokens, err
	}

	refreshToken, err := as.TokensService.GenerateRefreshToken(account.Id, account.Email)
	if err != nil {
		return tokens, err
	}

	tokens.AccessToken = accessToken
	tokens.RefreshToken = refreshToken

	return tokens, nil
}

func (as *AuthenticationService) Register(registerPayload authenticationPayloads.RegisterPayload) (*models.Account, error) {
	return as.AccountsService.Create(accountPayloads.CreateAccountPayload{
		FirstName: registerPayload.FirstName,
		LastName: registerPayload.LastName,
		Email: registerPayload.Email,
		Password: registerPayload.Password,
		PlanId: registerPayload.PlanId,
	})
}