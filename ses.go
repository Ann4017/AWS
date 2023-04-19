package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/aws-sdk-go/aws"
)

type SES_info struct {
	s_region     string
	s_access_Key string
	s_secret_key string
	pc_client    *ses.Client
	s_sender     string
	s_recipient  string
	s_subject    string
	s_body       string
}

func (c *SES_info) ses_init(access_key string, secret_key string, region string) {
	c.s_region = region
	c.s_access_Key = access_key
	c.s_secret_key = secret_key
}

func (c *SES_info) write_msg(sender, recipient, subject, body string) {
	c.s_sender = sender
	c.s_recipient = recipient
	c.s_subject = subject
	c.s_body = body
}

func (c *SES_info) set_cfg() error {
	cred := credentials.NewStaticCredentialsProvider(c.s_access_Key, c.s_secret_key, "")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(cred), config.WithRegion(c.s_region))
	if err != nil {
		return err
	}

	c.pc_client = ses.NewFromConfig(cfg)

	return nil
}

func (c *SES_info) send_email(client *ses.Client, sender, recipient, subject, body string) error {
	input := ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{recipient},
		},
		Message: &types.Message{
			Subject: &types.Content{
				Data: aws.String(subject),
			},
			Body: &types.Body{
				Text: &types.Content{
					Data: aws.String(body),
				},
			},
		},
		Source: aws.String(sender),
	}

	_, err := c.pc_client.SendEmail(context.Background(), &input)
	if err != nil {
		return err
	}

	return nil
}
