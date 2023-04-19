package main

import (
	"fmt"
)

func main() {
	accessKey := "ses-smtp-user.20230418-200907"
	secretKey := "AKIAVOZYFWFTBRFZPU6H,BF2GwJlO8sQtmgOvtpuKfUwnm8g0dXiDvxxIRMOhmR2F"
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

	// sns_info := SNS{}

	// sns_info.sns_init(accessKey, secretKey, region)

	// svc := sns_info.get_sess()
	// err := sns_info.send_sms("안녕하세용", "+821045196551", svc)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	ses := SES_info{}

	ses.ses_init(accessKey, secretKey, region)
	ses.write_msg("abh4017@naver.com", "qudgusyou012@gmail.com", "test", "testEmail")

	err := ses.set_cfg()
	if err != nil {
		fmt.Println(err)
	}

	err = ses.send_email(ses.pc_client, ses.s_sender,
		ses.s_recipient, ses.s_subject, ses.s_body)
	if err != nil {
		fmt.Println(err)
	}

}
