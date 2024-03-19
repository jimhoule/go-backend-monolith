package plans

import (
	"app/database/postgres"
	"app/plans/controllers"
	"app/plans/domain/factories"
	"app/plans/persistence/mappers"
	"app/plans/persistence/repositories"
	"app/plans/services"
	"app/router"

	uuidService "app/uuid/services"
)

func Init(router *router.Router, db *postgres.Db) {
	plansController := controllers.PlansController{
		PlansService: services.PlansService{
			PlansFactory: factories.PlansFactory{
				UuidService: &uuidService.NativeUuidService{},
			},
			PlansRepository: &repositories.PostgresPlansRepository{
				PlansMapper: mappers.PlansMapper{},
				Db: db,
			},
		},
	}

	router.Get("/plans", plansController.FindAll)
	router.Get("/plans/{id}", plansController.FindById)

	router.Post("/plans", plansController.Create)
}