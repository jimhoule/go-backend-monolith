package controllers

import (
	"app/accounts/application/services"
	"app/utils/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AccountsController struct {
	AccountsService *services.AccountsService
}

func (ac *AccountsController) FindAll(writer http.ResponseWriter, request *http.Request) {
	accounts, err := ac.AccountsService.FindAll();
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, accounts)
}

func (ac *AccountsController) FindById(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id");
	account, err := ac.AccountsService.FindById(id);
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, account)
}