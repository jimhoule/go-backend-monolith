package services

import (
	accountsService "app/accounts/application/services"
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

func (as *AuthenticationService) Login(email string, password string) (*Tokens, error) {
	tokens := &Tokens{}

	account, err := as.AccountsService.FindByEmail(email)
	if err != nil {
		return tokens, err
	}

	isValid, err := as.CryptoService.ComparePassword(account.Password, password)
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