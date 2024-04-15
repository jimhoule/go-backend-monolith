package movies

import (
	"app/database/postgres"
	"app/files"
	"app/movies/presenters/http/controllers"
	"app/router"
)

func Init(mainRouter *router.MainRouter, db *postgres.Db) {
	filesController := &controllers.MoviesController{
		FilesService: files.GetService(),
	}

	mainRouter.Post("/movies/uploads", filesController.Upload)
}