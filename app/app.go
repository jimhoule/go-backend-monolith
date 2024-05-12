package app

import (
	"app/app/application/services"
	"app/app/infrastructures/persitence/postgres/repositories"
	"app/app/presenters/http/controllers"
	"app/database"
	"app/router"
)

func GetService(db *database.Db) *services.AppService {
	return &services.AppService{
		AppRepository: &repositories.PostgresAppRepository{
			Db: db,
		},
	}
}

func Init(mainRouter *router.MainRouter, db *database.Db) {
	appController := &controllers.AppController{
		AppService: GetService(db),
	}

	mainRouter.Get("/app/health", appController.Diagnose)
}