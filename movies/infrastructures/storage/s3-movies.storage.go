package storage

import (
	"app/aws/services"
	"io"
)

type S3Storge struct {
	S3Service services.S3Service
}

func (ss *S3Storge) Upload(filePath string, file []byte) (bool, error) {
	_, err := ss.S3Service.PutObject(filePath, file)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (ss *S3Storge) Download(filePath string) ([]byte, error) {
	output, err := ss.S3Service.GetObject(filePath)
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