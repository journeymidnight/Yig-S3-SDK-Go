package s3lib

import (
	"github.com/journeymidnight/aws-sdk-go/aws"
	"github.com/journeymidnight/aws-sdk-go/service/s3"
)

func (s3client *S3Client) MakeBucket(bucketName string) (err error) {
	params := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}
	if _, err = s3client.Client.CreateBucket(params); err != nil {
		return err
	}
	return
}

func (s3client *S3Client) DeleteBucket(bucketName string) (err error) {
	params := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}
	if _, err = s3client.Client.DeleteBucket(params); err != nil {
		return err
	}
	return
}

func (s3client *S3Client) HeadBucket(bucketName string) (err error) {
	params := &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	}
	if _, err = s3client.Client.HeadBucket(params); err != nil {
		return err
	}
	return
}

func (s3client *S3Client) ListBuckets() (buckets []*s3.Bucket, err error) {
	params := &s3.ListBucketsInput{}
	out, err := s3client.Client.ListBuckets(params)
	if err != nil {
		return nil, err
	}
	return out.Buckets, err
}
