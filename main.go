package main

import (
	"app/accounts"
	"app/authentication"
	"app/database"
	"app/genres"
	"app/gql"
	"app/languages"
	"app/movies"
	"app/plans"
	"app/profiles"
	"app/router"
	"app/websocket"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Loads .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err);
	}

	// Gets database connection
	db := database.Get()

	// Gets websocket
	websocketServer := websocket.Get()

	// Gets gql
	gqlServer := gql.Get()

	// Gets router
	mainRouter := router.Get()

	// Inits modules
	authentication.Init(mainRouter, db)
	accounts.Init(mainRouter, db)
	genres.Init(mainRouter, db)
	languages.Init(mainRouter, db)
	movies.Init(mainRouter, websocketServer, gqlServer, db)
	plans.Init(mainRouter, db)
	profiles.Init(mainRouter, db)

	// Mounts websocket server and gql server to router
	mainRouter.Handle("/ws", websocketServer.ServeWS())
	mainRouter.Handle("/graphql", gqlServer.ServeGQL())
	mainRouter.Handle("/graphql/sandbox", gqlServer.ServeSandbox())

	// Creates server
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler: mainRouter,
	}

	// Starts server
	err = server.ListenAndServe();
	if err != nil {
		log.Panic(err);
	}
}