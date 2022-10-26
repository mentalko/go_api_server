package service

import (
	"context"
	"time"

	"github.com/mentalko/go_api_server/internal/core"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction core.Transaction) error
	GetAll(ctx context.Context) ([]core.Transaction, error)
}

type Transaction struct {
	repo TransactionRepository
}

func NewTransaction(repo TransactionRepository) *Transaction {
	return &Transaction{
		repo: repo,
	}
}

func (a *Transaction) Create(ctx context.Context, transaction core.Transaction) error {
	if transaction.Date.IsZero() {
		transaction.Date = time.Now()
	}
	return a.repo.Create(ctx, transaction)
}

func (a *Transaction) GetAll(ctx context.Context) ([]core.Transaction, error) {
	return a.repo.GetAll(ctx)
}
