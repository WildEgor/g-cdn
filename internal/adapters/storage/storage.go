package adapters

import (
	"context"
	"github.com/WildEgor/g-cdn/internal/config"
	domains "github.com/WildEgor/g-cdn/internal/domain"
	"io"
)

type StorageProvider interface {
	Upload(ctx context.Context, objectName string, reader io.Reader) error
	Download(ctx context.Context, objectName string) (io.Reader, error)
	Delete(objectName string) error
	Exists(ctx context.Context, objectName string) (bool, error)
	Metadata(objectName string) (*domains.FileMetadata, error)
}

type StoragePing interface {
	Ping() error
}

type StorageConfig struct {
	Provider string
	s3       *S3StorageConfig
}

func NewStorageConfig(s3 *config.MinioConfig) *StorageConfig {
	return &StorageConfig{
		Provider: "s3",
		s3: &S3StorageConfig{
			Endpoint: s3.Endpoint,
			Region:   s3.Region,
			Bucket:   s3.Bucket,
			AccessId: s3.AccessId,
			Secret:   s3.Secret,
		},
	}
}

func NewStorage(config *StorageConfig) StorageProvider {
	switch config.Provider {
	case "s3":
		return NewS3Storage(config.s3)
	default:
		return nil
	}
}

type StorageAdapter struct {
	provider StorageProvider
}

func NewStorageAdapter(provider StorageProvider) *StorageAdapter {
	return &StorageAdapter{
		provider: provider,
	}
}

func (s *StorageAdapter) Metadata(objectName string) (*domains.FileMetadata, error) {
	return s.provider.Metadata(objectName)
}

func (s *StorageAdapter) Upload(ctx context.Context, objectName string, reader io.Reader) error {
	return s.provider.Upload(ctx, objectName, reader)
}

func (s *StorageAdapter) Download(ctx context.Context, objectName string) (io.Reader, error) {
	return s.provider.Download(ctx, objectName)
}

func (s *StorageAdapter) Delete(objectName string) error {
	return s.provider.Delete(objectName)
}

func (s *StorageAdapter) Exists(ctx context.Context, objectName string) (bool, error) {
	return s.provider.Exists(ctx, objectName)
}

func (s *StorageAdapter) GetProvider() StorageProvider {
	return s.provider
}

func (s *StorageAdapter) SetProvider(provider StorageProvider) {
	s.provider = provider
}

func (s *StorageAdapter) Ping() error {
	if ping, ok := s.provider.(StoragePing); ok {
		return ping.Ping()
	}
	return nil
}
