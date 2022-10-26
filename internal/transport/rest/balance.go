package rest

import (
	"context"
	"log"
	"net/http"
	"strconv"
	// "github.com/mentalko/go_api_server/internal/core"
)

func (h *Handler) topupBalance(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, curency, value := query.Get("id"), query.Get("curency"), query.Get("value")

	log.Printf(" id=%s, curency=%s, value=%s", id, curency, value)

	account_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("topupBalance() error:", err)
	}

	account, err := h.accountsService.GetByID(context.TODO(), account_id)
	if err != nil {
		log.Println("topupBalance() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Шиткод
	if curency == "balance_usd" {
		new_value, _ := strconv.ParseFloat(value, 64)
		*account.Balance_USD += new_value
		log.Printf("new_balance: %.2f", *account.Balance_USD)

	} else if curency == "balance_rub" {
		new_value, _ := strconv.ParseFloat(value, 64)
		*account.Balance_RUB += new_value
		log.Printf("new_balance: %.2f", *account.Balance_RUB)
	}

	err = h.accountsService.Update(context.TODO(), account_id, account)
	if err != nil {
		log.Println("topupBalance() error: cant update data", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (h *Handler) exchangeBalance(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, value := query.Get("id"), query.Get("value")
	curency_from, curency_to := query.Get("curency_from"), query.Get("curency_to")

	log.Printf(" id=%s, curency_from=%s, value=%s", id, curency_from, value)

	account_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("exchangeBalance() error:", err)
	}

	account, err := h.accountsService.GetByID(context.TODO(), account_id)
	if err != nil {
		log.Println("exchangeBalance() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Шиткод
	if curency_from == "balance_usd" && curency_to == "balance_rub" {
		new_value, _ := strconv.ParseFloat(value, 64)
		*account.Balance_USD -= new_value
		*account.Balance_RUB += new_value * 55
		log.Printf("new_balance: %.2f", *account.Balance_USD)

	} else if curency_from == "balance_rub" && curency_to == "balance_usd" {
		new_value, _ := strconv.ParseFloat(value, 64)
		*account.Balance_RUB -= new_value
		*account.Balance_USD += new_value / 55
		log.Printf("new_balance: %.2f", *account.Balance_RUB)
	}

	err = h.accountsService.Update(context.TODO(), account_id, account)
	if err != nil {
		log.Println("exchangeBalance() error: cant update data", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
