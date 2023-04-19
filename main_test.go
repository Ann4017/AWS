package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	accessKey := ""
	secretKey := ""
	region := ""

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
