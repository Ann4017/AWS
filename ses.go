package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
)

type SES struct {
	s_region     string
	s_access_Key string
	s_secret_key string
	ses_cfg
}

func (c *SES) set_cfg() error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

}
