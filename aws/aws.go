package aws

import (
	"app/aws/services"
)

// NOTE: Allows creation of multiple s3 services instead of always getting a pointer to the same one
func CreateS3Service(bucketName string) services.S3Service {
	return services.S3Service{
		Client: services.CreateS3Client(),
		BucketName: bucketName,
	}
}