package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Db = pgx.Conn
type NamedArgs = pgx.NamedArgs

var db *Db

func Get() *Db {
	if db == nil {
		//Creates a new Postgresql database connection
		dsn := "host=localhost user=postgres password=password dbname=go-backend-monolith port=5432"
		connection, err := pgx.Connect(context.Background(), dsn)
		if err != nil {
			fmt.Printf("Unable to connect to database: %v\n", err)
			os.Exit(1)
		}

		db = connection

		return db
	}

	return db
}