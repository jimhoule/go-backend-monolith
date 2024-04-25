package services

type TranscoderService interface {
	TranscodeDash(
		dirPath string,
		fileName string,
		fileExtension string,
		onTranscodingProgressSent func(text string),
	) error
}