package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func main() {
	// Specify AWS region
	region := "ap-northeast-2"

	// Load AWS SDK config with region
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region))
	if err != nil {
		fmt.Println("Error loading AWS SDK config:", err)
		return
	}

	// Create IAM client
	client := iam.NewFromConfig(cfg)

	// Set IAM user name
	userName := "ann"

	// Create new access key for specified IAM user
	resp, err := client.CreateAccessKey(context.TODO(), &iam.CreateAccessKeyInput{
		UserName: aws.String(userName),
	})
	if err != nil {
		fmt.Println("Error creating access key for IAM user:", err)
		return
	}

	// Print access key ID and secret access key
	fmt.Println("Access key ID:", *resp.AccessKey.AccessKeyId)
	fmt.Println("Secret access key:", *resp.AccessKey.SecretAccessKey)
}
