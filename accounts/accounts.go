package accounts

import (
	"app/accounts/controllers"
	"app/accounts/persistence/repositories"
	"app/accounts/services"
	"app/router"
)

// NOTE: Not working
func Init(router *router.Router) {
	accountsController := controllers.AccountsController{
		AccountsService: services.AccountsService{
			AccountsRepository: &repositories.FakeAccountsRepository{},
		},
	}

	router.Get("/accounts", accountsController.FindAll)
	router.Get("/accounts/{id}", accountsController.FindById)

	router.Post("/accounts", accountsController.Save)
}