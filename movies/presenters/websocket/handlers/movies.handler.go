package handlers

import (
	"app/movies/application/payloads"
	"app/movies/application/services"
	"app/movies/presenters/websocket/dtos"
	"app/movies/presenters/websocket/events"
	"app/websocket"
	"encoding/base64"
	"encoding/json"
)

type MoviesHandler struct{
	MoviesService *services.MoviesService
}

func (mh *MoviesHandler) HandleTranscodeDashAndUploadVideo(
	client *websocket.Client,
	payload json.RawMessage,
) {
	// Gets payload
	var transcodeDashAndUploadMovieDto dtos.TranscodeDashAndUploadVideoDto
	err := json.Unmarshal(payload, &transcodeDashAndUploadMovieDto); 
	if err != nil {
		client.Emit(events.TranscodeDashAndUploadVideoError, err)
		return
	}

	// Convert file base64 string to bytes ([]byte)
	file, err := base64.StdEncoding.DecodeString(transcodeDashAndUploadMovieDto.FileBase64String)
	if err != nil {
		client.Emit(events.TranscodeDashAndUploadVideoError, err)
		return
	}

	// Transcode to hash and uploads video
	isFinished, err := mh.MoviesService.TranscodeDashAndUploadVideo(&payloads.TranscodeDashAndUploadVideoPayload{
		File: file,
		FileName: transcodeDashAndUploadMovieDto.FileName,
		FileExtension: transcodeDashAndUploadMovieDto.FileExtension,
		OnTranscodingProgressSent: func(text string) {
			client.Emit(events.TranscodingProgressSent, text)
		},
		OnUploadStarted: func() {
			client.Emit(events.UploadStarted, nil)
		},
	})
	if err != nil {
		client.Emit(events.TranscodeDashAndUploadVideoError, err)
		return
	}

	client.Emit(events.TranscodeDashAndUploadVideoFinished, isFinished)
}