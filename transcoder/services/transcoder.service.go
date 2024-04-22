package services

type TranscoderService interface {
	TranscodeDash(dirPath string, fileName string, fileExtension string) error
}