package config

import (
	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
	"time"
)

type MongoConfig struct {
	URI               string        `env:"MONGODB_URI,required"`
	DbName            string        `env:"MONGODB_NAME,required"`
	ct                int64         `env:"MONGODB_TIMEOUT"`
	ConnectionTimeout time.Duration `env:"-"`
}

func NewMongoConfig(c *Configurator) *MongoConfig {
	cfg := MongoConfig{}

	if err := env.Parse(&cfg); err != nil {
		log.Printf("[MinioConfig] %+v\n", err)
	}

	cfg.ConnectionTimeout = time.Duration(cfg.ct) * time.Millisecond

	return &cfg
}
