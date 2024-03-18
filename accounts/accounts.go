package accounts

import (
	"app/accounts/controllers"
	"app/accounts/persistence/repositories"
	"app/accounts/services"
	"app/database/postgres"
	"app/router"
)

// NOTE: Not working
func Init(router *router.Router, db *postgres.Db) {
	accountsController := controllers.AccountsController{
		AccountsService: services.AccountsService{
			AccountsRepository: &repositories.PostgresAccountsRepository{
				Db: db,
			},
		},
	}

	router.Get("/accounts", accountsController.FindAll)
	router.Get("/accounts/{id}", accountsController.FindById)

	router.Post("/accounts", accountsController.Save)
}