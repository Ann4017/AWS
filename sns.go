package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SNS struct {
	s_region     string
	s_access_Key string
	s_secret_key string
	pc_session   *session.Session
}

func (c *SNS) Init(access_key string, secret_key string, region string) {
	c.s_access_Key = access_key
	c.s_secret_key = secret_key
	c.s_region = region
}

func (c *SNS) Get_sess() *sns.SNS {
	cred := credentials.NewStaticCredentials(c.s_access_Key, c.s_secret_key, "")
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(c.s_region),
		Credentials: cred,
	}))

	svc := sns.New(sess)

	return svc
}

func (c *SNS) Send_sms(msg string, phone_num string, svc *sns.SNS) error {
	params := &sns.PublishInput{
		Message:     aws.String(msg),
		PhoneNumber: aws.String(phone_num),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"AWS.SNS.SMS.SMSType": {
				DataType:    aws.String("String"),
				StringValue: aws.String("Transactional"),
			},
		},
	}

	resp, err := svc.Publish(params)
	if err != nil {
		return err
	}

	fmt.Printf("SMS message sent successfully id:%v", resp.MessageId)

	return nil
}
