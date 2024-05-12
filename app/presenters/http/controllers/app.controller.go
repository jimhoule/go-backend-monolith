package controllers

import (
	"app/app/application/services"
	"app/utils/json"
	"net/http"
)

type AppController struct {
	AppService *services.AppService
}

func (ac *AppController) Diagnose(writer http.ResponseWriter, request *http.Request) {
	err := ac.AppService.Diagnose()
	if err != nil {
		json.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, true)
}