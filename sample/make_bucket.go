package sample

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
)

func MakeBucketSample() {
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	DeleteTestBucketAndObject()
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// Delete a bucket
	err = sc.DeleteBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	fmt.Printf("CreateBucketSample Run Success!\n\n")
}
