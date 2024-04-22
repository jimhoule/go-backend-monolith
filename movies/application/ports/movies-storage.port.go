package ports

type MoviesStoragePort interface {
	Upload(filePath string, file []byte) ([]byte, error)
	UploadLarge(filePath string, file []byte) ([]byte, error)
}