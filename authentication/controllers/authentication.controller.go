package controllers

import (
	"app/authentication/dtos"
	"app/authentication/services"
	"app/utils/json"
	"errors"
	"net/http"
)

type AuthenticationController struct {
	AuthenticationService services.AuthenticationService
}

func (ac *AuthenticationController) Login(writer http.ResponseWriter, request *http.Request) {
	loginError := errors.New("Invalid email or password")
	var loginDto dtos.LoginDto

	err := json.ReadHttpRequestBody(writer, request, &loginDto)
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, loginError)
		return
	}

	tokens, err := ac.AuthenticationService.Login(loginDto.Email, loginDto.Password)
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, loginError)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, tokens)
}