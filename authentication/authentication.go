package authentication

import (
	"app/accounts"
	"app/authentication/controllers"
	authenticationService "app/authentication/services"
	cryptoService "app/crypto/services"
	"app/database/postgres"
	"app/router"
	tokensService "app/tokens/services"
)

func GetService(db *postgres.Db) *authenticationService.AuthenticationService {
	return &authenticationService.AuthenticationService{
		AccountsService: *accounts.GetService(db),
		TokensService: &tokensService.JwtTokensService{},
		CryptoService: &cryptoService.BcryptCryptoService{},
	}
}

func Init(mainRouter *router.MainRouter, db *postgres.Db) {
	authenticationController := controllers.AuthenticationController{
		AuthenticationService: *GetService(db),
	}

	mainRouter.Post("/authentication/login", authenticationController.Login)
}