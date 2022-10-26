package service

import (
	"context"

	"github.com/mentalko/go_api_server/internal/core"
)

type AccountRepository interface {
	Create(ctx context.Context, account core.Account) error
	GetByID(ctx context.Context, id int64) (core.Account, error)
	GetAll(ctx context.Context) ([]core.Account, error)
	Update(ctx context.Context, id int64, inp core.Account) error
	GetIdByUsername(ctx context.Context, username string) (core.Account, error)
}

type Account struct {
	repo AccountRepository
}

func NewAccount(repo AccountRepository) *Account {
	return &Account{
		repo: repo,
	}
}

func (a *Account) Create(ctx context.Context, account core.Account) error {
	return a.repo.Create(ctx, account)
}

func (a *Account) GetByID(ctx context.Context, id int64) (core.Account, error) {
	return a.repo.GetByID(ctx, id)
}

func (a *Account) GetIdByUsername(ctx context.Context, username string) (core.Account, error) {
	return a.repo.GetIdByUsername(ctx, username)
}

func (a *Account) GetAll(ctx context.Context) ([]core.Account, error) {
	return a.repo.GetAll(ctx)
}

func (a *Account) Update(ctx context.Context, id int64, inp core.Account) error {
	return a.repo.Update(ctx, id, inp)
}
