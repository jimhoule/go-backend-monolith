package services

import (
	"app/transactions/application/ports"
	"context"
)

type TransactionsService struct {
	TransactionsRepository ports.TransactionsRepositoryPort
}

func (ts *TransactionsService) Execute(
	ctx context.Context,
	executeQuery func(ctx context.Context) (any, error),
) (any, error) {
	return ts.TransactionsRepository.Execute(ctx, executeQuery)
}