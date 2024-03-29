package services

import (
	accountsService "app/accounts/application/services"
	"app/authentication/application/payloads"
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

func (as *AuthenticationService) Login(loginPayload payloads.LoginPayload) (*Tokens, error) {
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