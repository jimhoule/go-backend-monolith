package controllers

import (
	"app/accounts/application/payloads"
	"app/accounts/application/services"
	"app/accounts/presenters/http/dtos"
	"app/utils/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AccountsController struct {
	AccountsService services.AccountsService
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

func (ac *AccountsController) Create(writer http.ResponseWriter, request *http.Request) {
	var createAccountDto dtos.CreateAccountDto

	// Gets request body
	err := json.ReadHttpRequestBody(writer, request, &createAccountDto)
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	// Creates account
	account, err := ac.AccountsService.Create(payloads.CreateAccountPayload{
		FirstName: createAccountDto.FirstName,
		LastName: createAccountDto.LastName,
		Email: createAccountDto.Email,
		Password: createAccountDto.Password,
		PlanId: createAccountDto.PlanId,
	})
	if err != nil {
		json.WriteHttpError(writer, http.StatusNotFound, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusCreated, account)
}