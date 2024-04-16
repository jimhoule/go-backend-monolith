package ports

type MoviesStoragePort interface {
	Upload(filePath string, file []byte) (bool, error)
	Download(filePath string) ([]byte, error)
}