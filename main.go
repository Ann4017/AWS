package main

import (
	"fmt"
)

func main() {
	accessKey := ""
	secretKey := ""
	region := ""

	// info := S3{}

	// info.Init(accessKey, secretKey, region)

	// err := info.Set_s3_config()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// list, err := info.Get_s3_bucket_list()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, name := range list {
	// 	fmt.Println(name)
	// }

	// sns_info := SNS{}

	// sns_info.Init(accessKey, secretKey, region)

	// svc := sns_info.Get_sess()
	// err := sns_info.Send_sms("안녕하세용", "+821045196551", svc)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	ses := SES_info{}

	ses.Init(accessKey, secretKey, region)
	ses.Write_msg("abh4017@naver.com", "qudgusyou012@gmail.com", "test", "testEmail")

	err := ses.Set_cfg()
	if err != nil {
		fmt.Println(err)
	}

	err = ses.Send_email(ses.pc_client, ses.s_sender,
		ses.s_recipient, ses.s_subject, ses.s_body)
	if err != nil {
		fmt.Println(err)
	}

}
