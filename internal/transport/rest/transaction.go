package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mentalko/go_api_server/internal/core"
)

func (h *Handler) createTransaction(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	var transaction core.Transaction
	if err = json.Unmarshal(reqBytes, &transaction); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Println(transaction)

	id_account_to, err := h.accountsService.GetIdByUsername(context.TODO(), transaction.Account_to)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("account_to ID:", id_account_to)

	transaction.Account_to = fmt.Sprintf("%v", (id_account_to.ID))

	err = h.transactionsService.Create(context.TODO(), transaction)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	account_from, err := h.accountsService.GetByID(context.TODO(), transaction.Account_from)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	account_to, err := h.accountsService.GetByID(context.TODO(), id_account_to.ID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Шиткод
	if transaction.Curency == "USD" {
		*account_from.Balance_USD -= transaction.Value
		*account_to.Balance_USD += transaction.Value

	} else if transaction.Curency == "RUB" {
		*account_from.Balance_RUB -= transaction.Value
		*account_to.Balance_RUB += transaction.Value

	}

	err = h.accountsService.Update(context.TODO(), account_from.ID, account_from)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.accountsService.Update(context.TODO(), account_to.ID, account_to)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, "created")

}

func (h *Handler) getAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.transactionsService.GetAll(context.TODO())
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, transactions)
}

func (h *Handler) transferBalance(w http.ResponseWriter, r *http.Request) {

}
