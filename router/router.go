package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type MainRouter = chi.Mux
type GroupRouter = chi.Router

var mainRouter *MainRouter

func Get() *MainRouter {
	if mainRouter == nil {
		mainRouter := chi.NewRouter()

		mainRouter.Use(cors.Handler(cors.Options{
			AllowedOrigins: []string{
				"https://*",
				"http://*",
			},
			AllowedMethods: []string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"OPTIONS",
			},
			AllowedHeaders: []string{
				"Accept",
				"Authorization",
				"Content-Type",
				"X-CRSF-TOKEN",
			},
			ExposedHeaders: []string{
				"Link",
			},
			AllowCredentials: true,
			MaxAge: 300,
		}))

		return mainRouter
	}

	return mainRouter;
}