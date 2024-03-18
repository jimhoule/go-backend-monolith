package main

import (
	"app/accounts"
	"app/router"
	"fmt"
	"log"
	"net/http"
)

const httpPort = 3000

func main() {
	// Gets router
	router := router.Get()

	// Inits modules
	accounts.Init(router)

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