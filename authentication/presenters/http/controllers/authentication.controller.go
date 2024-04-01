package controllers

import (
	"app/authentication/application/payloads"
	"app/authentication/application/services"
	"app/authentication/presenters/http/dtos"
	"app/utils/json"
	"errors"
	"net/http"
)

type AuthenticationController struct {
	AuthenticationService services.AuthenticationService
}

func (ac *AuthenticationController) Login(writer http.ResponseWriter, request *http.Request) {
	loginError := errors.New("invalid email or password")

	// Gets request body
	var loginDto dtos.LoginDto

	err := json.ReadHttpRequestBody(writer, request, &loginDto)
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, loginError)
		return
	}

	// Logs in
	tokens, err := ac.AuthenticationService.Login(&payloads.LoginPayload{
		Email: loginDto.Email,
		Password: loginDto.Password,
	})
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, loginError)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, tokens)
}

func (ac *AuthenticationController) Register(writer http.ResponseWriter, request *http.Request) {
	// Gets request body
	var registerDto dtos.RegisterDto

	err := json.ReadHttpRequestBody(writer, request, &registerDto)
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	// Registers
	tokens, err := ac.AuthenticationService.Register(&payloads.RegisterPayload{
		FirstName: registerDto.FirstName,
		LastName: registerDto.LastName,
		Email: registerDto.Email,
		Password: registerDto.Password,
		PlanId: registerDto.PlanId,
	})
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusCreated, tokens)
}