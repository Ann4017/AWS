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
	s_region      string
	s_access_Key  string
	s_secret_key  string
	s_bucket_name string
	pc_clint      *s3.Client
}

func (c *S3) Init(access_key string, secret_key string, region string) {
	c.s_access_Key = access_key
	c.s_secret_key = secret_key
	c.s_region = region
}

func (c *S3) Set_s3_config() error {
	cred := credentials.NewStaticCredentialsProvider(c.s_access_Key, c.s_secret_key, "")
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithCredentialsProvider(cred),
		config.WithRegion(c.s_region))
	if err != nil {
		return err
	}
	c.pc_clint = s3.NewFromConfig(cfg)

	return nil
}

func (c *S3) Insert_s3_bucket(new_bucket_name string, region types.BucketLocationConstraint) (bool, error) {
	_, err := c.pc_clint.CreateBucket(context.Background(), &s3.CreateBucketInput{
		Bucket: aws.String(new_bucket_name),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: region,
		},
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *S3) Delete_s3_bucket(bucket_name string) (bool, error) {
	_, err := c.pc_clint.DeleteBucket(context.Background(), &s3.DeleteBucketInput{
		Bucket: aws.String(bucket_name),
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *S3) Get_s3_bucket_list() (bucket_list []string, err error) {
	output, err := c.pc_clint.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	item := make([]string, len(output.Buckets))
	for i, bucket := range output.Buckets {
		item[i] = *bucket.Name
	}

	return item, nil
}

func (c *S3) Get_s3_bucket_item_list(bucket_name string) (bucket_item_list []string, err error) {
	resp, err := c.pc_clint.ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket_name),
	})
	if err != nil {
		return nil, err
	}

	item := make([]string, len(resp.Contents))
	for i, v := range resp.Contents {
		fmt.Printf("num: %d, file: %s, size: %v\n", i, *v.Key, v.Size)
		item[i] = *v.Key
	}

	return item, nil
}

func (c *S3) Upload_file(file_name string, folder string, bucket_name string) (*manager.UploadOutput, error) {
	file, err := os.Open(file_name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	uploader := manager.NewUploader(c.pc_clint)
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

	downloader := manager.NewDownloader(c.pc_clint)
	_, err = downloader.Download(context.TODO(), fd, &s3.GetObjectInput{
		Bucket: &c.s_bucket_name,
		Key:    &key,
	})

	return err
}
