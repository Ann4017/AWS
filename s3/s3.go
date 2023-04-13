package s3

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3_info struct {
	S_AwsS3Region    string
	S_AwsAccessKey   string
	S_AwsSecretkey   string
	S_AwsProfileName string
	S_BucketName     string
	S3Client         *s3.Client
}

func (s *S3_info) Get_AccessKey(acccess_key string, secret_key string) {
	s.S_AwsAccessKey = acccess_key
	s.S_AwsSecretkey = secret_key
}

func (s *S3_info) Set_S3ConfigDefault() error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return errors.New(err.Error())
	}
	s.S3Client = s3.NewFromConfig(cfg)

	return nil
}

func (s *S3_info) Set_S3ConfigByKey() error {
	creds := credentials.NewStaticCredentialsProvider(s.S_AwsAccessKey, s.S_AwsSecretkey, "")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(s.S_AwsS3Region), config.WithCredentialsProvider(creds))
	if err != nil {
		return err
	}

	s.S3Client = s3.NewFromConfig(cfg)

	return nil
}

func (s *S3_info) Upload_S3_file(filename, preFix string) *manager.UploadOutput {
	file, err := ioutil.ReadFile(filename)
	uploader := manager.NewUploader(s.S3Client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.S_BucketName),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(file),
	})
	if err != nil {
		panic(err)
	}
	return result
}

func (s *S3_info) Insert_bucket(bucket_name string, region types.BucketLocationConstraint) {
	output, err := s.S3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: &bucket_name,
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: region,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(output.Location)
}
