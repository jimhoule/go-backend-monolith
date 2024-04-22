package payloads

type TranscodeDashAndUploadMoviePayload struct {
	File          []byte
	FileName      string
	FileExtension string
}