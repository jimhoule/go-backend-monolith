package payloads

type UploadMoviePayload struct {
	File     []byte
	FilePath string
}