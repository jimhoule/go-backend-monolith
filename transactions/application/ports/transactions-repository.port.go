package ports

import "context"

type TransactionsRepositoryPort interface {
	Execute(ctx context.Context, executeQuery func(ctx context.Context) (any, error)) (any, error)
}