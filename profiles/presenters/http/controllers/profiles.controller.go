package controllers

import (
	"app/profiles/application/payloads"
	"app/profiles/application/services"
	"app/profiles/presenters/http/dtos"
	"app/utils/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProfilesController struct {
	ProfilesService *services.ProfilesService
}

func (pc *ProfilesController) FindAllByAccountId(writer http.ResponseWriter, request *http.Request) {
	accountId := chi.URLParam(request, "accountId");
	profiles, err := pc.ProfilesService.FindAllByAccountId(accountId);
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, profiles)
}

func (pc *ProfilesController) FindById(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id");
	profile, err := pc.ProfilesService.FindById(id);
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, profile)
}

func (pc *ProfilesController) Update(writer http.ResponseWriter, request *http.Request) {
	// Gets request body
	var updateProfileDto dtos.UpdateProfileDto
	err := json.ReadHttpRequestBody(writer, request, &updateProfileDto)
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	// Updates profile
	id := chi.URLParam(request, "id");
	profile, err := pc.ProfilesService.Update(
		id,
		&payloads.UpdateProfilePayload{
			Name: updateProfileDto.Name,
		},
	)
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, profile)
}

func (pc *ProfilesController) Delete(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id");
	_, err := pc.ProfilesService.Delete(id);
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusNoContent, nil)
}

func (pc *ProfilesController) Create(writer http.ResponseWriter, request *http.Request) {
	// Gets request body
	var createProfileDto dtos.CreateProfileDto
	err := json.ReadHttpRequestBody(writer, request, &createProfileDto)
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	// Creates profile
	profile, err := pc.ProfilesService.Create(&payloads.CreateProfilePayload{
		Name: createProfileDto.Name,
		AccountId: createProfileDto.AccountId,
	})
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusCreated, profile)
}