package controllers

import (
	"app/plans/application/payloads"
	"app/plans/application/services"
	"app/plans/presenters/http/dtos"
	"app/utils/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PlansController struct {
	PlansService services.PlansService
}

func (pc *PlansController) FindAll(writer http.ResponseWriter, request *http.Request) {
	plans, err := pc.PlansService.FindAll()
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, plans)
}

func (pc *PlansController) FindById(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	plan, err := pc.PlansService.FindById(id)
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, plan)
}

func (pc *PlansController) Create(writer http.ResponseWriter, request *http.Request) {
	// Gets request body
	var createPlanDto dtos.CreatePlanDto
	err := json.ReadHttpRequestBody(writer, request, &createPlanDto)
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	// Creates plan
	plan, err := pc.PlansService.Create(payloads.CreatePlanPayload{
		Name: createPlanDto.Name,
		Description: createPlanDto.Description,
		Price: createPlanDto.Price,
	})
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, plan)
}