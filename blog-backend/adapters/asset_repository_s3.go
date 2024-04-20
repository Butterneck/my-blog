package adapters

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AssetRepositoryS3 struct {
	s3         *s3.Client
	presignS3  *s3.PresignClient
	bucketName string
}

type AssetRepositoryS3Config struct {
	BucketName string
}

func NewAssetRepositoryS3(s3Client *s3.Client, s3PresignClient *s3.PresignClient, config AssetRepositoryS3Config) *AssetRepositoryS3 {
	return &AssetRepositoryS3{
		s3:         s3Client,
		bucketName: config.BucketName,
	}
}

func (s *AssetRepositoryS3) MoveAsset(ctx context.Context, oldName, newName string) error {
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

func (s *AssetRepositoryS3) PublishAsset(ctx context.Context, name string) (string, error) {
	// TODO: Move object from draft to published

	return name, nil
}

func (s *AssetRepositoryS3) DeleteAsset(ctx context.Context, name string) error {
	_, err := s.s3.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &s.bucketName,
		Key:    &name,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *AssetRepositoryS3) GeneratePresignedURL(ctx context.Context, name string) (string, error) {
	url, err := s.presignS3.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.bucketName,
		Key:    &name,
	})
	if err != nil {
		return "", err
	}

	return url.URL, nil
}
