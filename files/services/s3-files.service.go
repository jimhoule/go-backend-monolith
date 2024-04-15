package services

import (
	"app/aws/services"
	"io"
)

type S3FilesService struct {
	S3Service *services.S3Service
}

func (sfs *S3FilesService) Upload(dirName string, fileName string, file []byte) (bool, error) {
	_, err := sfs.S3Service.PutObject(dirName, fileName, file)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (sfs *S3FilesService) Download(dirName string, fileName string) ([]byte, error) {
	output, err := sfs.S3Service.GetObject(dirName, fileName)
	if err != nil {
		return nil, err
	}
	defer output.Body.Close()

	file, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, err
	}

	return file, nil
}