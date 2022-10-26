package core

import (
	"errors"
	"time"
)

var (
	ErrTransactionNotFound = errors.New("Transaction not found")
)

type Transaction struct {
	ID           int64     `json:"id"`
	Account_from int64     `json:"account_from"`
	Account_to   string    `json:"account_to"`
	Value        float64   `json:"value"`
	Curency      string    `json:"curency"`
	Date         time.Time `json:"date"`
}

// type UpdateTransactionInput struct {
// 	Account_from *int64     `json:"account_from"`
// 	Account_to   *int64     `json:"account_to"`
// 	Value        *float64   `json:"value"`
// 	Curency      *string    `json:"curency"`
// 	Date         *time.Time `json:"date"`
// }
