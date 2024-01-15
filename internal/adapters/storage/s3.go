package adapters

import (
	"bytes"
	"context"
	s3 "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
)

type S3StorageConfig struct {
	Endpoint string
	Region   string
	Bucket   string
	AccessId string
	Secret   string
	Secure   bool
}

type S3Storage struct {
	client *s3.Client `wire:"-"`
	config *S3StorageConfig
}

func NewS3Storage(config *S3StorageConfig) *S3Storage {
	client, err := s3.New(config.Endpoint, &s3.Options{
		Creds:  credentials.NewStaticV4(config.AccessId, config.Secret, ""),
		Secure: config.Secure,
	})
	if err != nil {
		return nil
	}

	return &S3Storage{
		client: client,
		config: config,
	}
}

func (s *S3Storage) Upload(ctx context.Context, objectName string, reader io.Reader) error {
	_, err := s.client.PutObject(ctx, s.config.Bucket, objectName, reader, -1, s3.PutObjectOptions{})

	if err != nil {
		return err
	}
	return nil
}

func (s *S3Storage) Download(ctx context.Context, objectName string) (io.Reader, error) {

	data, err := s.client.GetObject(ctx, s.config.Bucket, objectName, s3.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, data)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *S3Storage) Delete(objectName string) error {
	err := s.client.RemoveObject(
		context.Background(),
		s.config.Bucket, objectName,
		s3.RemoveObjectOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *S3Storage) Exists(ctx context.Context, objectName string) (bool, error) {
	_, err := s.client.StatObject(ctx, s.config.Bucket, objectName, s3.StatObjectOptions{})
	if err != nil {
		// If the error is due to the object not found, return false with no error
		if errResp := s3.ToErrorResponse(err); errResp.Code == "NoSuchKey" {
			return false, nil
		}
		// For other errors, return them
		return false, err
	}
	// If there is no error, the object exists
	return true, nil
}

func (s *S3Storage) Ping() error {
	_, err := s.client.BucketExists(context.Background(), s.config.Bucket)
	if err != nil {
		return err
	}
	return nil
}
