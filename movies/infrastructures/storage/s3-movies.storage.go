package storage

import (
	"app/aws/services"
)

type S3Storge struct {
	S3Service services.S3Service
}

func (ss *S3Storge) Upload(filePath string, file []byte) ([]byte, error) {
	_, err := ss.S3Service.PutObject(filePath, file)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (ss *S3Storge) UploadLarge(filePath string, file []byte) ([]byte, error) {
	_, err := ss.S3Service.PutLargeObject(filePath, file)
	if err != nil {
		return nil, err
	}

	return file, nil
}