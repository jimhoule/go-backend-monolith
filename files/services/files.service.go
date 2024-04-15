package services

type FilesService interface {
	Upload(dirName string, fileName string, file []byte) (bool, error)
	Download(dirName string, fileName string) ([]byte, error)
}