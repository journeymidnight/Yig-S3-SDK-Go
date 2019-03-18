package sample

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
	"io/ioutil"
	"strings"
)

func AppendObjectSample() {
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	DeleteTestBucketAndObject()
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	var nextPos int64

	// 1. Append strings to an object
	strs := []string{"yig1", "yig2", "yig3"}
	for _, s := range strs {
		fmt.Println("Append String:", s)
		nextPos, err = sc.AppendObject(bucketName, objectKey, strings.NewReader(s), nextPos)
		if err != nil {
			HandleError(err)
		}
	}
	out, err := sc.GetObject(bucketName, objectKey)
	b, _ := ioutil.ReadAll(out)
	fmt.Println("Get appended string:", string(b))

	// TODO 2. Append files to an object

	DeleteTestBucketAndObject()
	fmt.Printf("AppendObjectSample Run Success !\n\n")
}
