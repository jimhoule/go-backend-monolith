package services

import (
	"app/transactions/infrastructures/persistence/fake/repositories"
	"context"
	"fmt"
	"testing"
)

func getTestContext() (*TransactionsService) {
	transactionsService := &TransactionsService{
		TransactionsRepository: &repositories.FakeTransactionsRepository{},
	}

	return transactionsService
}

func TestTransactionsExecuteServiceSuccess(t *testing.T) {
	transactionsService := getTestContext()

	_, err := transactionsService.Execute(
		context.Background(),
		func(ctx context.Context) (any, error) {
			return true, nil
		},
	)
	if err != nil {
		t.Errorf("Expected Transaction to be successfull but got %v", err)
	}
}

func TestTransactionsExecuteServiceFailure(t *testing.T) {
	transactionsService := getTestContext()

	_, err := transactionsService.Execute(
		context.Background(),
		func(ctx context.Context) (any, error) {
			return nil, fmt.Errorf("Failed to execute Transaction")
		},
	)
	if err == nil {
		t.Errorf("Expected Transaction to fail but got %v", err)
	}
}