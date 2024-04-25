package controllers

import (
	genrePayloads "app/genres/application/payloads"
	"app/genres/application/services"
	"app/genres/presenters/http/dtos"
	"app/router"
	"app/utils/json"
	"net/http"
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
	id := router.GetUrlParam(request, "id")
	genre, err := gc.GenresService.FindById(id)
	if (err != nil) {
		json.WriteHttpError(writer, http.StatusNotFound, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, genre)
}

func (gc *GenresController) Update(writer http.ResponseWriter, request *http.Request) {
	// Gets request body
	var updateGenreDto dtos.UpdateGenreDto
	err := json.ReadHttpRequestBody(writer, request, &updateGenreDto)
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	// Updates genre
	id := router.GetUrlParam(request, "id")
	genre, err := gc.GenresService.Update(id, &genrePayloads.UpdateGenrePayload{
		UpdateLabelTranslationPayloads: updateGenreDto.UpdateLabelTranslationPayloads,
	})
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, genre)
}

func (gc *GenresController) Delete(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")
	gc.GenresService.Delete(id)

	json.WriteHttpResponse(writer, http.StatusNoContent, id)
}

func (gc *GenresController) Create(writer http.ResponseWriter, request *http.Request) {
	// Gets request body
	var createGenreDto dtos.CreateGenreDto
	err := json.ReadHttpRequestBody(writer, request, &createGenreDto)
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	// Creates genre
	genre, err := gc.GenresService.Create(&genrePayloads.CreateGenrePayload{
		CreateLabelTranslationPayloads: createGenreDto.CreateLabelTranslationPayloads,
	})
	if (err != nil) {
		json.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusCreated, genre)
}