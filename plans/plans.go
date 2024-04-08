package plans

import (
	"app/authentication/presenters/http/middlewares"
	"app/database/postgres"
	"app/plans/application/services"
	"app/plans/domain/factories"
	"app/plans/persistence/postgres/repositories"
	"app/plans/presenters/http/controllers"
	"app/router"
	"app/transactions"
	"app/translations"
	"app/uuid"
)

func GetService(db *postgres.Db) *services.PlansService {
	return &services.PlansService{
		PlansFactory: &factories.PlansFactory{
			UuidService: uuid.GetService(),
		},
		PlansRepository: &repositories.PostgresPlansRepository{
			Db: db,
		},
		TranslationsService: translations.GetService(db),
		TransactionsService: transactions.GetService(db),
	}
}

func Init(mainRouter *router.MainRouter, db *postgres.Db) {
	plansController := &controllers.PlansController{
		PlansService: GetService(db),
	}

	// NOTE: In a mux, all middleware must be defined before routes so we have to wrap the routes around a Group
	mainRouter.Group(func(groupRouter router.GroupRouter) {
		groupRouter.Use(middlewares.VerifyAccessToken)

		groupRouter.Get("/plans", plansController.FindAll)
		groupRouter.Get("/plans/{id}", plansController.FindById)
		groupRouter.Post("/plans", plansController.Create)
	})
}