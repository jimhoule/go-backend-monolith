package transcoder

import "app/transcoder/services"

func GetService() services.TranscoderService {
	return &services.FfmpegTranscoderService{}
}