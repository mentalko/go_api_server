package core

import (
	"errors"
)

var (
	ErrAccountNotFound = errors.New("Account not found")
)

type Account struct {
	ID           int64   `json:"id"`
	Username     *string  `json:"username"`
	Phone_number *string  `json:"phone_number"`
	Balance_USD  *float64 `json:"balance_usd"`
	Balance_RUB  *float64 `json:"balance_rub"`
}

// type UpdateAccount struct {
// 	Username     string  `json:"username"`
// 	Phone_number string  `json:"phone_number"`
// 	Balance_USD  float64 `json:"balance_usd"`
// }
