package main

import "fmt"

func main() {
	accessKey := ""
	secretKey := ""
	region := "ap-northeast-1"

	// info := S3{}

	// info.init(accessKey, secretKey, region)

	// err := info.set_s3_config()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// list, err := info.get_s3_bucket_list()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, name := range list {
	// 	fmt.Println(name)
	// }

	sns_info := SNS{}

	sns_info.sns_init(accessKey, secretKey, region)

	svc := sns_info.get_sess()
	err := sns_info.send_sms("안녕하세용", "+821045196551", svc)
	if err != nil {
		fmt.Println(err)
	}

}
