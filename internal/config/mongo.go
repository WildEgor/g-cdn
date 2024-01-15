package config

import (
	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
)

type MongoConfig struct {
	URI string `env:"MONGODB_URI"`
}

func NewMongoConfig(c *Configurator) *MongoConfig {
	cfg := MongoConfig{}

	if err := env.Parse(&cfg); err != nil {
		log.Printf("[MinioConfig] %+v\n", err)
	}

	return &cfg
}
