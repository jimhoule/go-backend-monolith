package repositories

import (
	"app/database"
	"context"
)

type PostgresAppRepository struct {
	Db *database.Db
}

func (par *PostgresAppRepository) Diagnose() error {
	err := par.Db.Connection.Ping(context.Background())
	if err != nil {
		return err
	}

	return nil
}