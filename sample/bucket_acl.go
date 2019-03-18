package sample

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
)

func BucketACLSample()  {
	DeleteTestBucketAndObject()
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Printf("BucketACLSample Run Success!\n\n")
}
