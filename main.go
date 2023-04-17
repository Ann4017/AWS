package main

import (
	"fmt"
)

func main() {
	info := S3{}

	accessKey := ""
	secretKey := ""
	region := ""

	info.init(accessKey, secretKey, region)

	err := info.set_s3_config()
	if err != nil {
		fmt.Println(err)
	}

	list, err := info.get_s3_bucket_list()
	if err != nil {
		fmt.Println(err)
	}
	for _, name := range list {
		fmt.Println(name)
	}
}
