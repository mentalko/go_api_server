package rest

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mentalko/go_api_server/internal/core"
)

type Accounts interface {
	Create(ctx context.Context, account core.Account) error
	GetByID(ctx context.Context, id int64) (core.Account, error)
	GetAll(ctx context.Context) ([]core.Account, error)
	Update(ctx context.Context, id int64, inp core.Account) error
	GetIdByUsername(ctx context.Context, username string) (core.Account, error)
}

type Transactions interface {
	Create(ctx context.Context, transaction core.Transaction) error
	GetAll(ctx context.Context) ([]core.Transaction, error)
}

type Handler struct {
	accountsService     Accounts
	transactionsService Transactions
}

func NewHandler(accounts Accounts, transactions Transactions) *Handler {
	return &Handler{
		accountsService:     accounts,
		transactionsService: transactions,
	}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddlewhare)

	accounts := r.PathPrefix("/accounts").Subrouter()
	{
		accounts.HandleFunc("", h.createAccount).Methods(http.MethodPost)
		accounts.HandleFunc("", h.getAllAccounts).Methods(http.MethodGet)
		accounts.HandleFunc("/{id:[0-9]+}", h.getAccountByID).Methods(http.MethodGet)

	}

	r.HandleFunc("/topup", h.topupBalance).Methods(http.MethodGet)
	r.HandleFunc("/exchange", h.exchangeBalance).Methods(http.MethodGet)

	transactions := r.PathPrefix("/transactions").Subrouter()
	{
		transactions.HandleFunc("", h.createTransaction).Methods(http.MethodPost)
		transactions.HandleFunc("", h.getAllTransactions).Methods(http.MethodGet)
	}
	return r
}
