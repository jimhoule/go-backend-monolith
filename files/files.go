package files

import (
	"app/aws"
	"app/files/services"
	"os"
)

func GetService() services.FilesService {
	return &services.S3FilesService{
		S3Service: aws.CreateS3Service(os.Getenv("AWS_VIDEO_UPLOADS_BUCKET_NAME")),
	}
}