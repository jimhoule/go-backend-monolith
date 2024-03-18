package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Router = chi.Mux

var router *Router

func Get() *Router {
	if router == nil {
		router := chi.NewRouter()

		router.Use(cors.Handler(cors.Options{
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

		return router
	}

	return router;
}