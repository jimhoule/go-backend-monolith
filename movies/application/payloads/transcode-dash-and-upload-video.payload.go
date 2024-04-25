package payloads

type TranscodeDashAndUploadVideoPayload struct {
	File                      []byte
	FileName                  string
	FileExtension             string
	OnTranscodingProgressSent func(text string)
	OnUploadStarted           func()
}