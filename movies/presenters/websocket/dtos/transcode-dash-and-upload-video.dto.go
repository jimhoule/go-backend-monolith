package dtos

type TranscodeDashAndUploadVideoDto struct {
	FileBase64String string `json:"fileBase64String"`
	FileName         string `json:"fileName"`
	FileExtension    string `json:"fileExtension"`
}