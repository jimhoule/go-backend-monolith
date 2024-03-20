package main

import (
	"app/accounts"
	"app/authentication"
	"app/database/postgres"
	"app/plans"
	"app/router"
	"fmt"
	"log"
	"net/http"
)

const httpPort = 3000

func main() {
	// Gets database connection
	db := postgres.Get()

	// Gets router
	mainRouter := router.Get()

	// Inits modules
	authentication.Init(mainRouter, db)
	accounts.Init(mainRouter, db)
	plans.Init(mainRouter, db)

	// Creates server
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", httpPort),
		Handler: mainRouter,
	}

	// Starts server
	err := server.ListenAndServe();
	if err != nil {
		log.Panic(err);
	}
}