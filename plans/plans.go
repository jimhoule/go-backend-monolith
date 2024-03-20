package plans

import (
	"app/authentication/middlewares"
	"app/database/postgres"
	"app/plans/controllers"
	"app/plans/domain/factories"
	"app/plans/persistence/mappers"
	"app/plans/persistence/repositories"
	"app/plans/services"
	"app/router"

	uuidService "app/uuid/services"
)

func GetService(db *postgres.Db) *services.PlansService {
	return &services.PlansService{
		PlansFactory: factories.PlansFactory{
			UuidService: &uuidService.NativeUuidService{},
		},
		PlansRepository: &repositories.PostgresPlansRepository{
			PlansMapper: mappers.PlansMapper{},
			Db: db,
		},
	}
}

func Init(mainRouter *router.MainRouter, db *postgres.Db) {
	plansController := controllers.PlansController{
		PlansService: *GetService(db),
	}

	// NOTE: In a mux, all middleware must be defined before routes so we have to wrap the routes around a Group
	mainRouter.Group(func(groupRouter router.GroupRouter) {
		groupRouter.Use(middlewares.VerifyAccessToken)

		groupRouter.Get("/plans", plansController.FindAll)
		groupRouter.Get("/plans/{id}", plansController.FindById)

		groupRouter.Post("/plans", plansController.Create)
	})
}