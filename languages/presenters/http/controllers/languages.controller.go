package controllers

import (
	"app/languages/application/payloads"
	"app/languages/application/services"
	"app/languages/presenters/http/dtos"
	"app/router"
	"app/utils/json"
	"net/http"
)

type LanguagesController struct {
	LanguagesService *services.LanguagesService
}

func (lc *LanguagesController) FindAll(writer http.ResponseWriter, request *http.Request) {
	languages, err := lc.LanguagesService.FindAll()
	if err != nil {
		json.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, languages)
}

func (lc *LanguagesController) FindById(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")
	language, err := lc.LanguagesService.FindById(id)
	if err != nil {
		json.WriteHttpError(writer, http.StatusNotFound, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, language)
}

func (lc *LanguagesController) Update(writer http.ResponseWriter, request *http.Request) {
	var updateLanguageDto dtos.UpdateLanguageDto
	err := json.ReadHttpRequestBody(writer, request, &updateLanguageDto)
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	id := router.GetUrlParam(request, "id")
	language, err := lc.LanguagesService.Update(id, &payloads.UpdateLanguagePayload{
		Code: updateLanguageDto.Code,
		UpdateLabelTranslationPayloads: updateLanguageDto.UpdateLabelTranslationPayloads,
	})
	if err != nil {
		json.WriteHttpError(writer, http.StatusNotFound, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusOK, language)
}

func (lc *LanguagesController) Delete(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")
	lc.LanguagesService.Delete(id)

	json.WriteHttpResponse(writer, http.StatusNoContent, nil)
}

func (lc *LanguagesController) Create(writer http.ResponseWriter, request *http.Request) {
	var createLanguageDto dtos.CreateLanguageDto
	err := json.ReadHttpRequestBody(writer, request, &createLanguageDto)
	if err != nil {
		json.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	language, err := lc.LanguagesService.Create(&payloads.CreateLanguagePayload{
		Code: createLanguageDto.Code,
		CreateLabelTranslationPayloads: createLanguageDto.CreateLabelTranslationPayloads,
	})
	if err != nil {
		json.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	json.WriteHttpResponse(writer, http.StatusCreated, language)
}