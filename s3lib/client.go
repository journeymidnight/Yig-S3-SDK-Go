package s3lib

import (
	"github.com/journeymidnight/aws-sdk-go/aws"
	"github.com/journeymidnight/aws-sdk-go/aws/credentials"
	"github.com/journeymidnight/aws-sdk-go/aws/session"
	"github.com/journeymidnight/aws-sdk-go/service/s3"
)

type S3Client struct {
	Client *s3.S3
}

func NewS3(endpoint, accessKey, secretKey string) *S3Client {
	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")

	// By default make sure a region is specified
	s3client := s3.New(session.Must(session.NewSession(
		&aws.Config{
			Credentials: creds,
			DisableSSL:  aws.Bool(true),
			Endpoint:    aws.String(endpoint),
			Region:      aws.String("r"),
		},
	),
	),
	)
	return &S3Client{s3client}
}
