package adapters

import (
	"bytes"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3AssetStore struct {
	s3         *s3.Client
	bucketName string
}

type S3AssetStoreConfig struct {
	BucketName string
}

func NewS3AssetStore(s3Client *s3.Client, config S3AssetStoreConfig) *S3AssetStore {
	return &S3AssetStore{
		s3:         s3Client,
		bucketName: config.BucketName,
	}
}

func (s *S3AssetStore) UploadAsset(ctx context.Context, asset []byte, name string) error {
	var partMiBs int64 = 10
	uploader := manager.NewUploader(s.s3, func(u *manager.Uploader) {
		u.PartSize = partMiBs * 1024 * 1024
		u.Concurrency = 10
	})

	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: &s.bucketName,
		Key:    &name,
		Body:   bytes.NewReader(asset),
	})
	if err != nil {
		log.Printf("Couldn't upload large object to %v:%v. Here's why: %v\n",
			s.bucketName, name, err)
	}

	return nil
}

func (s *S3AssetStore) MoveAsset(ctx context.Context, oldName, newName string) error {
	_, err := s.s3.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     &s.bucketName,
		CopySource: &oldName,
		Key:        &newName,
	})
	if err != nil {
		return err
	}

	err = s.DeleteAsset(ctx, oldName)
	if err != nil {
		return err
	}

	return nil
}

func (s *S3AssetStore) DeleteAsset(ctx context.Context, name string) error {
	_, err := s.s3.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &s.bucketName,
		Key:    &name,
	})
	return err
}
