package main

import (
	"fmt"
)

func main() {
	info := S3{}

	accessKey := ""
	secretKey := ""
	region := ""

	info.Init(accessKey, secretKey, region)

	err := info.Set_s3_config()
	if err != nil {
		fmt.Println(err)
	}

	err = info.Get_s3_bucket_list()
	if err != nil {
		fmt.Println(err)
	}

	err = info.Get_s3_bucket_item_list("s3-ann-test000")
	if err != nil {
		fmt.Println(err)
	}
}
