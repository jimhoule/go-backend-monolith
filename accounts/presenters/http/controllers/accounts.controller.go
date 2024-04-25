package controllers

import (
	"app/accounts/application/services"
	"app/router"
	"app/utils/json"
	"net/http"
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
	id := router.GetUrlParam(request, "id");
	account, err := ac.AccountsService.FindById(id);
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, account)
}