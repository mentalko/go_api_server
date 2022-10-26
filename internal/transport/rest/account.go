package rest

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mentalko/go_api_server/internal/core"
)

func (h *Handler) createAccount(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println("createAccount() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var account core.Account
	if err = json.Unmarshal(reqBytes, &account); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("createAccount() error:", err)
		return
	}

	err = h.accountsService.Create(context.TODO(), account)
	if err != nil {
		log.Println("createAccount() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getAccountByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("getAccountByID() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	account, err := h.accountsService.GetByID(context.TODO(), id)
	if err != nil {
		if errors.Is(err, core.ErrAccountNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("getAccountByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(account)
	if err != nil {
		log.Println("getAccountByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) getAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.accountsService.GetAll(context.TODO())
	if err != nil {
		log.Println("getAllAccounts() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responce, err := json.Marshal(accounts)
	if err != nil {
		log.Println("getAllAccounts() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(responce)
}

func getIdFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("id can't be 0")
	}

	return id, nil

}
