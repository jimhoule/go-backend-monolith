package repositories

import (
	"context"
)

type FakeTransactionsRepository struct {}

func (ftr *FakeTransactionsRepository) Execute(
	ctx context.Context,
	executeQuery func(ctx context.Context) (any, error),
) (any, error) {
	return executeQuery(ctx)
}