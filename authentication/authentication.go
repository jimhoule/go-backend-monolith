package authentication

import (
	"app/accounts"
	"app/authentication/application/services"
	"app/authentication/presenters/http/controllers"
	"app/crypto"
	"app/database"
	"app/router"
	"app/tokens"
)

func GetService(db *database.Db) *services.AuthenticationService {
	return &services.AuthenticationService{
		AccountsService: accounts.GetService(db),
		TokensService: tokens.GetService(),
		CryptoService: crypto.GetService(),
	}
}

func Init(mainRouter *router.MainRouter, db *database.Db) {
	authenticationController := &controllers.AuthenticationController{
		AuthenticationService: GetService(db),
	}

	mainRouter.Post("/authentication/login", authenticationController.Login)
	mainRouter.Post("/authentication/register", authenticationController.Register)
}