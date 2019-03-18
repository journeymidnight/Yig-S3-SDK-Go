package sample

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
	"strings"
	"os"
	"io"
)

func NewObjectSample() {
	DeleteTestBucketAndObject()
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// 1. Put a string object
	err = sc.PutObject(bucketName, objectKey, strings.NewReader("NewBucketAndObjectSample"))
	if err != nil {
		HandleError(err)
	}

	err = sc.DeleteObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}

	// 2. Put a file
	f, err := os.Open(localFilePath)
	if err != nil {
		HandleError(err)
	}
	err = sc.PutObject(bucketName, objectKey, f)
	if err != nil {
		HandleError(err)
	}
	out, err := sc.GetObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}
	// Download file
	f2, err := os.OpenFile("sample/Download.jpeg", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		HandleError(err)
	}
	io.Copy(f2, out)

	DeleteTestBucketAndObject()
	fmt.Printf("NewObjectSample Run Success !\n\n")
}
