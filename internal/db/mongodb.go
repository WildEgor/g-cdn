package db

import (
	"context"
	"github.com/WildEgor/g-cdn/internal/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConnection struct {
	Client *mongo.Client
	mc     *config.MongoConfig
}

func NewMongoDBConnection(
	mc *config.MongoConfig,
) *MongoDBConnection {
	conn := &MongoDBConnection{
		nil,
		mc,
	}

	defer conn.Disconnect()

	return conn
}

func (mc *MongoDBConnection) Connect() {
	opts := options.Client().ApplyURI(mc.mc.URI)

	opts.ConnectTimeout = &mc.mc.ConnectionTimeout

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Panic("Fail connect to Mongo", err)
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Panic("Fail connect to Mongo", err)
	}

	log.Info("Success connect to Mongo")

	mc.Client = client
}

func (mc *MongoDBConnection) Disconnect() {
	if mc.Client == nil {
		return
	}

	err := mc.Client.Disconnect(context.TODO())
	if err != nil {
		log.Panic("Fail disconnect Mongo", err)
		panic(err)
	}

	log.Info("Connection to Mongo closed.")
}

func (mc *MongoDBConnection) Instance() *mongo.Client {
	return mc.Client
}

func (mc *MongoDBConnection) DbName() string {
	return mc.mc.DbName
}
