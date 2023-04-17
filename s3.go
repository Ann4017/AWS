package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3 struct {
	S_Region      string
	S_Access_Key  string
	S_Secret_key  string
	S_Bucket_name string
	P_S3_clint    *s3.Client
}

func (c *S3) Init(access_key string, secret_key string, region string) {
	c.S_Access_Key = access_key
	c.S_Secret_key = secret_key
	c.S_Region = region
}

func (c *S3) Set_s3_config() error {
	cred := credentials.NewStaticCredentialsProvider(c.S_Access_Key, c.S_Secret_key, "")
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithCredentialsProvider(cred),
		config.WithRegion(c.S_Region))
	if err != nil {
		return err
	}
	c.P_S3_clint = s3.NewFromConfig(cfg)

	return nil
}

func (c *S3) Insert_s3_bucket(new_bucket_name string, region types.BucketLocationConstraint) error {
	_, err := c.P_S3_clint.CreateBucket(context.Background(), &s3.CreateBucketInput{
		Bucket: aws.String(new_bucket_name),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: region,
		},
	})
	if err != nil {
		return err
	}

	fmt.Printf("Insert bucket: %s\n", new_bucket_name)

	return nil
}

func (c *S3) Delete_s3_bucket(bucket_name string) (bool, error) {
	_, err := c.P_S3_clint.DeleteBucket(context.Background(), &s3.DeleteBucketInput{
		Bucket: aws.String(bucket_name),
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *S3) Get_s3_bucket_list() error {
	output, err := c.P_S3_clint.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return err
	}

	for i, bucket := range output.Buckets {
		fmt.Printf("%d: %s\n", i, *bucket.Name)
	}

	return nil
}

func (c *S3) Get_s3_bucket_item_list(bucket_name string) error {
	resp, err := c.P_S3_clint.ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket_name),
	})
	if err != nil {
		return err
	}

	for i, v := range resp.Contents {
		fmt.Printf("num: %d, file: %s, size: %v\n", i, *v.Key, v.Size)
	}

	return nil
}

func (c *S3) Upload_file(file_name string, folder string, bucket_name string) (*manager.UploadOutput, error) {
	file, err := os.Open(file_name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	uploader := manager.NewUploader(c.P_S3_clint)
	path_key := filepath.Join(folder, file_name)

	result, err := uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(bucket_name),
		Key:    aws.String(path_key),
		Body:   file,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *S3) Download_file(directory string, key string) error {
	file := filepath.Join(directory, key)
	if err := os.MkdirAll(filepath.Dir(file), 7750); err != nil {
		return err
	}

	fd, err := os.Create(file)
	if err != nil {
		return err
	}

	defer fd.Close()

	downloader := manager.NewDownloader(c.P_S3_clint)
	_, err = downloader.Download(context.TODO(), fd, &s3.GetObjectInput{
		Bucket: &c.S_Bucket_name,
		Key:    &key,
	})

	return err
}
