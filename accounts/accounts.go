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

func Init(router *router.Router, db *postgres.Db) {
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

	router.Get("/accounts", accountsController.FindAll)
	router.Get("/accounts/{id}", accountsController.FindById)

	router.Post("/accounts", accountsController.Create)
}