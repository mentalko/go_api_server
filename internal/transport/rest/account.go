package rest

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mentalko/go_api_server/internal/core"
)

func (h *Handler) createAccount(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)

	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	var account core.Account
	if err = json.Unmarshal(reqBytes, &account); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.accountsService.Create(context.TODO(), account)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, "created")
}

func (h *Handler) getAccountByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	account, err := h.accountsService.GetByID(context.TODO(), id)
	if err != nil {
		if errors.Is(err, core.ErrAccountNotFound) {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, account)
}

func (h *Handler) getAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.accountsService.GetAll(context.TODO())
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, accounts)
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
