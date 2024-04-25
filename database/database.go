package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type NamedArgs = pgx.NamedArgs
type Batch = pgx.Batch
type Identifier = pgx.Identifier

type Db struct{
	Connection *pgx.Conn
	CopyFromSlice func(length int, next func(int) ([]any, error)) pgx.CopyFromSource
	CopyFromRows func(rows [][]any) pgx.CopyFromSource
}

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

		db = &Db{
			Connection: connection,
			CopyFromSlice: pgx.CopyFromSlice,
			CopyFromRows: pgx.CopyFromRows,
		}
	}

	return db
}