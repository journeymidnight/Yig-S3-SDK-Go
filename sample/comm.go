package sample

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
	"os"
)

func HandleError(err error) {
	fmt.Println("panic err:", err)
	err = DeleteTestBucketAndObject()
	if err != nil {
		fmt.Println("DeleteTestBucketAndObject err:", err)
	}
	os.Exit(-1)
}

func DeleteTestBucketAndObject() error {
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	err := sc.DeleteObject(bucketName, objectKey)
	if err != nil {
		return err
	}

	err = sc.DeleteBucket(bucketName)
	if err != nil {
		return err
	}
	return nil
}
