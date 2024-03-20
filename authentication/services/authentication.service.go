package services

import (
	"app/authentication/dtos"

	accountsService "app/accounts/services"
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

func (as *AuthenticationService) Login(loginDto dtos.LoginDto) (*Tokens, error) {
	tokens := &Tokens{}

	account, err := as.AccountsService.FindByEmail(loginDto.Email)
	if err != nil {
		return tokens, err
	}

	isValid := as.CryptoService.ComparePassword(account.Password, loginDto.Password)
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