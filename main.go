package main

import (
	"app/accounts"
	"app/authentication"
	"app/database/postgres"
	"app/genres"
	"app/languages"
	"app/movies"
	"app/plans"
	"app/profiles"
	"app/router"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

const httpPort = 3000

func main() {
	// Loads .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err);
	}

	// Gets database connection
	db := postgres.Get()

	// Gets router
	mainRouter := router.Get()

	// Inits modules
	authentication.Init(mainRouter, db)
	accounts.Init(mainRouter, db)
	genres.Init(mainRouter, db)
	languages.Init(mainRouter, db)
	movies.Init(mainRouter, db)
	plans.Init(mainRouter, db)
	profiles.Init(mainRouter, db)

	// Creates server
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", httpPort),
		Handler: mainRouter,
	}

	// Starts server
	err = server.ListenAndServe();
	if err != nil {
		log.Panic(err);
	}
}