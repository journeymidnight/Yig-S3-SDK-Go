package main

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/sample"
)

func main() {
	sample.MakeBucketSample()
	sample.NewObjectSample()
	sample.ListBucketsSample()
	sample.BucketACLSample()
	//sample.BucketLifecycleSample()
	//sample.BucketRefererSample()
	//sample.BucketLoggingSample()
	//sample.BucketCORSSample()
	sample.AppendObjectSample()
	sample.ObjectACLSample()
	//sample.ObjectMetaSample()
	//sample.ListObjectsSample()
	//sample.DeleteObjectSample()
	//sample.AppendObjectSample()
	//sample.MySample()
	//sample.S2()
	//sample.CopyObjectSample()
	//sample.PutObjectSample()
	//sample.GetObjectSample()
	//
	//sample.CnameSample()
	//sample.SignURLSample()
	//sample.ArchiveSample()
	sample.MySample()

	fmt.Println("All samples completed !")
}
