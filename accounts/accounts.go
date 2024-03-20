package accounts

import (
	"app/accounts/controllers"
	"app/accounts/domain/factories"
	"app/accounts/persistence/mappers"
	"app/accounts/persistence/repositories"
	"app/accounts/services"
	"app/database/postgres"
	"app/router"

	cryptoService "app/crypto/services"
	uuidService "app/uuid/services"
)

func Init(mainRouter *router.MainRouter, db *postgres.Db) {
	accountsController := controllers.AccountsController{
		AccountsService: services.AccountsService{
			AccountsFactory: factories.AccountsFactory{
				UuidService: &uuidService.NativeUuidService{},
				CryptoService: &cryptoService.BcryptCryptoService{},
			},
			AccountsRepository: &repositories.PostgresAccountsRepository{
				AccountsMapper: mappers.AccountsMapper{},
				Db: db,
			},
		},
	}

	mainRouter.Get("/accounts", accountsController.FindAll)
	mainRouter.Get("/accounts/{id}", accountsController.FindById)

	mainRouter.Post("/accounts", accountsController.Create)
}