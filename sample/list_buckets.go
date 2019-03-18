package sample

import (
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
	"fmt"
)

func ListBucketsSample()  {
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	DeleteTestBucketAndObject()

	out, err := sc.ListBuckets()
	if err != nil {
		HandleError(err)
	}
	for _, b := range out {
		fmt.Println("bucket:", *b.Name)
	}

	DeleteTestBucketAndObject()
	fmt.Printf("ListBucketsSample Run Success!\n\n")
}