package controllers

import (
	"app/genres/application/payloads"
	"app/genres/application/services"
	"app/utils/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GenresController struct {
	GenresService *services.GenresService
}

func (gc *GenresController) FindAll(writer http.ResponseWriter, request *http.Request) {
	genres, err := gc.GenresService.FindAll()
	if (err != nil) {
		json.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, genres)
}

func (gc *GenresController) FindById(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	genre, err := gc.GenresService.FindById(id)
	if (err != nil) {
		json.WriteHttpError(writer, http.StatusNotFound, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, genre)
}

func (gc *GenresController) Create(writer http.ResponseWriter, request *http.Request) {
	genre, err := gc.GenresService.Create(&payloads.CreateGenrePayload{})
	if (err != nil) {
		json.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, genre)
}