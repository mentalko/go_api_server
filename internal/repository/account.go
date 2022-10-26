package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/mentalko/go_api_server/internal/core"
)

type Account struct {
	db *sql.DB
}

func NewAccount(db *sql.DB) *Account {
	return &Account{db}
}

func (a *Account) Create(ctx context.Context, account core.Account) error {
	_, err := a.db.Exec("INSERT INTO accounts (id, username, phone_number, balance_usd, balance_rub ) values ($1, $2, $3, $4, $5)",
		account.ID, account.Username, account.Phone_number, account.Balance_USD, account.Balance_RUB)
	return err
}

func (a *Account) GetByID(ctx context.Context, id int64) (core.Account, error) {
	var account core.Account
	err := a.db.QueryRow("SELECT id, username, phone_number, balance_usd, balance_rub FROM accounts WHERE id = $1", id).
		Scan(&account.ID, &account.Username, &account.Phone_number, &account.Balance_USD, &account.Balance_RUB)
	if err == sql.ErrNoRows {
		return account, core.ErrAccountNotFound
	}

	return account, err
}

func (a *Account) GetIdByUsername(ctx context.Context, username string) (core.Account, error) {
	var account core.Account
	err := a.db.QueryRow("SELECT id FROM accounts WHERE username = $1", username).
		Scan(&account.ID)
	if err == sql.ErrNoRows {
		return account, core.ErrAccountNotFound
	}

	return account, err
}

func (a *Account) GetAll(ctx context.Context) ([]core.Account, error) {
	rows, err := a.db.Query("SELECT id, username, phone_number, balance_usd, balance_rub FROM accounts")
	if err != nil {
		return nil, err
	}

	accounts := make([]core.Account, 0)
	for rows.Next() {
		var account core.Account
		if err := rows.Scan(&account.ID, &account.Username, &account.Phone_number, &account.Balance_USD, &account.Balance_RUB); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, rows.Err()
}

func (a *Account) Delete() {
	//TODO:
}

func (a *Account) Update(ctx context.Context, id int64, inp core.Account) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1

	if inp.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argsId))
		args = append(args, *inp.Username)
		argsId++
	}

	if inp.Phone_number != nil {
		setValues = append(setValues, fmt.Sprintf("phone_number=$%d", argsId))
		args = append(args, *inp.Phone_number)
		argsId++
	}

	if inp.Balance_USD != nil {
		setValues = append(setValues, fmt.Sprintf("balance_usd=$%d", argsId))
		args = append(args, *inp.Balance_USD)
		argsId++
	}

	if inp.Balance_RUB != nil {
		setValues = append(setValues, fmt.Sprintf("balance_rub=$%d", argsId))
		args = append(args, *inp.Balance_RUB)
		argsId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE accounts SET %s WHERE id=$%d", setQuery, argsId)
	args = append(args, id)

	log.Println(query)
	log.Println(args)

	_, err := a.db.Exec(query, args...)
	return err
}
