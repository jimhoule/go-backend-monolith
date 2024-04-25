package repositories

import (
	"app/database"
	"context"
)

type PostgresTransactionsRepository struct {
	Db *database.Db
}

func (ptr *PostgresTransactionsRepository) Execute(
	ctx context.Context,
	executeQuery func(ctx context.Context) (any, error),
) (any, error) {
	transaction, err := ptr.Db.Connection.Begin(ctx)
	if err != nil {
		return nil, err
	}

	result, err := executeQuery(ctx)
	if err != nil {
		transaction.Rollback(ctx)
		return nil, err
	}

	transaction.Commit(ctx)

	return result, nil
}