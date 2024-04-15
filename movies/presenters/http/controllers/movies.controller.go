package controllers

import (
	"app/files/services"
	"app/utils/json"
	"bytes"
	"io"
	"net/http"
	"strings"
)

type MoviesController struct{
    FilesService services.FilesService
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
    isUpload, err := mc.FilesService.Upload("showtime-transcoded-video-dev", splittedFileName[0], buffer.Bytes())
    if err != nil {
        json.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
    }

    // Resets the buffer in case I want to use it again
    // NOTE: Reduces memory allocations in more intense projects
    buffer.Reset()

    json.WriteHttpResponse(writer, http.StatusOK, isUpload)
}