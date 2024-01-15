package config

import (
	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
)

type MinioConfig struct {
	Endpoint string `env:"MINIO_ENDPOINT"`
	AccessId string `env:"MINIO_ACCESS_KEY"`
	Secret   string `env:"MINIO_SECRET_KEY"`
	Bucket   string `env:"MINIO_BUCKET"`
	Region   string `env:"MINIO_REGION"`
	UseSSL   bool   `env:"MINIO_SECURE"`
}

func NewMinioConfig(c *Configurator) *MinioConfig {
	cfg := MinioConfig{}

	if err := env.Parse(&cfg); err != nil {
		log.Printf("[MinioConfig] %+v\n", err)
	}

	return &cfg
}
