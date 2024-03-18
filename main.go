package main

import (
	"app/accounts"
	"app/database/postgres"
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
	router := router.Get()

	// Inits modules
	accounts.Init(router, db)

	// Creates server
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", httpPort),
		Handler: router,
	}

	// Starts server
	err := server.ListenAndServe();
	if err != nil {
		log.Panic(err);
	}
}