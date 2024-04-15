package aws

import "app/aws/services"

func GetS3Service() *services.S3Service {
	return &services.S3Service{
		Client: services.CreateS3Client(),
	}
}