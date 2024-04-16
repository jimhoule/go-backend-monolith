package controllers

import (
	"app/movies/application/payloads"
	"app/movies/application/services"
	"app/movies/presenters/http/dtos"
	"app/utils/json"
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type MoviesController struct{
    MoviesService *services.MoviesService
}

func (mc *MoviesController) FindAll(writer http.ResponseWriter, request *http.Request) {
    movies, err := mc.MoviesService.FindAll()
    if err != nil {
        json.WriteHttpError(writer, http.StatusInternalServerError, err)
        return
    }

    json.WriteHttpResponse(writer, http.StatusOK, movies)
}

func (mc *MoviesController) FindById(writer http.ResponseWriter, request *http.Request) {
    id := chi.URLParam(request, "id")
    movie, err := mc.MoviesService.FindById(id)
    if err != nil {
        json.WriteHttpError(writer, http.StatusNotFound, err)
        return
    }

    json.WriteHttpResponse(writer, http.StatusOK, movie)
}

func (mc *MoviesController) Update(writer http.ResponseWriter, request *http.Request) {
    var updateMovieDto dtos.UpdateMovieDto
    err := json.ReadHttpRequestBody(writer, request, &updateMovieDto)
    if err != nil {
        json.WriteHttpError(writer, http.StatusBadRequest, err)
        return
    }

    id := chi.URLParam(request, "id")
    movie, err := mc.MoviesService.Update(id, &payloads.UpdateMoviePayload{
        GenreId: updateMovieDto.GenreId,
        UpdateTitleTranslationPayloads: updateMovieDto.UpdateTitleTranslationPayloads,
        UpdateDescriptionTranslationPayloads: updateMovieDto.UpdateDescriptionTranslationPayloads,
    })
    if err != nil {
        json.WriteHttpError(writer, http.StatusNotFound, err)
        return
    }

    json.WriteHttpResponse(writer, http.StatusOK, movie)
}

func (mc *MoviesController) Delete(writer http.ResponseWriter, request *http.Request) {
    id := chi.URLParam(request, "id")
    mc.MoviesService.Delete(id)
    
    json.WriteHttpResponse(writer, http.StatusNoContent, nil)
}

func (mc *MoviesController) Create(writer http.ResponseWriter, request *http.Request) {
    var createMovieDto dtos.CreateMovieDto
    err := json.ReadHttpRequestBody(writer, request, &createMovieDto)
    if err != nil {
        json.WriteHttpError(writer, http.StatusBadRequest, err)
        return
    }

    movie, err := mc.MoviesService.Create(&payloads.CreateMoviePayload{
        GenreId: createMovieDto.GenreId,
        CreateTitleTranslationPayloads: createMovieDto.CreateTitleTranslationPayloads,
        CreateDescriptionTranslationPayloads: createMovieDto.CreateDescriptionTranslationPayloads,
    })
    if err != nil {
        json.WriteHttpError(writer, http.StatusNotFound, err)
        return
    }

    json.WriteHttpResponse(writer, http.StatusCreated, movie)
}

func (mc *MoviesController) Upload(writer http.ResponseWriter, request *http.Request) {
	// limit your max input length!
	//request.ParseMultipartForm(32 << 20)

    // Gets request file
    file, header, err := request.FormFile("file")
    if err != nil {
        json.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
    }
    defer file.Close()

    // Gets request file name
    splittedFileName := strings.Split(header.Filename, ".")

    // Copies file into buffer
    var buffer bytes.Buffer
    _, err = io.Copy(&buffer, file)
    if err != nil {
        json.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
    }

    // Uploads file
    isUpload, err := mc.MoviesService.Upload(splittedFileName[0], buffer.Bytes())
    if err != nil {
        json.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
    }

    // Resets the buffer in case I want to use it again
    // NOTE: Reduces memory allocations in more intense projects
    buffer.Reset()

    json.WriteHttpResponse(writer, http.StatusOK, isUpload)
}