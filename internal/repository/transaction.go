package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/mentalko/go_api_server/internal/core"
)

type Transaction struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) *Transaction {
	return &Transaction{db}
}

func (a *Transaction) Create(ctx context.Context, transaction core.Transaction) error {
	log.Println(transaction.Account_from, transaction.Account_to, transaction.Value,
		transaction.Curency, transaction.Date)
	_, err := a.db.Exec("INSERT INTO transactions (account_from, account_to, value, curency, date) values ($1, $2, $3, $4, $5)",
		transaction.Account_from, transaction.Account_to, transaction.Value,
		transaction.Curency, transaction.Date)
	return err
}

func (a *Transaction) GetByAccountID() {

}

func (a *Transaction) GetAll(ctx context.Context) ([]core.Transaction, error) {
	rows, err := a.db.Query("SELECT id, account_from, account_to, value, curency, date FROM transactions")
	if err != nil {
		return nil, err
	}

	transactions := make([]core.Transaction, 0)
	for rows.Next() {
		var transaction core.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.Account_from, &transaction.Account_to, &transaction.Value,
			&transaction.Curency, &transaction.Date); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, rows.Err()
}

func (a *Transaction) Delete() {

}

func (a *Transaction) Update() {

}
