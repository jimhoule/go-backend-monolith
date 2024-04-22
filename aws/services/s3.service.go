package services

import (
	"app/aws/configs"
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Service struct{
	Client *s3.Client
	BucketName string
}

func CreateS3Client() *s3.Client {
	return s3.NewFromConfig(configs.Get())
}

// Downloads file from bucket
func (ss *S3Service) GetObject(filePath string) (*s3.GetObjectOutput, error) {
	output, err := ss.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(ss.BucketName),
		Key:    aws.String(filePath),
	})
	if err != nil {
		return nil, err
	}

	return output, err
}

// Uploads file into Bucket
func (ss *S3Service) PutObject(filePath string, file []byte) (*s3.PutObjectOutput, error) {
	output, err := ss.Client.PutObject(
		context.TODO(),
		&s3.PutObjectInput{
			Bucket: aws.String(ss.BucketName),
			Key:    aws.String(filePath),
			Body:  	bytes.NewReader(file),
		},
	)
	if err != nil {
		return nil, err
	}

	return output, err
}

// Uses an upload manager to upload file into bucket
// NOTE: The upload manager breaks large data into parts and uploads the parts concurrently
func (ss *S3Service) PutLargeObject(fileName string, largeFile []byte) (*manager.UploadOutput, error) {
	var partMiBs int64 = 10
	uploader := manager.NewUploader(
		ss.Client,
		func(u *manager.Uploader) {
			u.PartSize = partMiBs * 1024 * 1024
			u.Concurrency = 8
		},
	)

	output, err := uploader.Upload(
		context.TODO(),
		&s3.PutObjectInput{
			Bucket: aws.String(ss.BucketName),
			Key:    aws.String(fileName),
			Body:   bytes.NewReader(largeFile),
		},
	)
	if err != nil {
		return nil, err
	}

	return output, err
}