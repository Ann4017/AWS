package main

import (
	"AWS/s3"
	"fmt"
)

func main() {
	info := s3.Info{}

	accessKey := ""
	secretKey := ""
	region := ""

	info.Set_session(accessKey, secretKey, region)

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
