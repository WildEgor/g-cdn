// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package pkg

import (
	"github.com/WildEgor/g-cdn/internal/adapters/storage"
	"github.com/WildEgor/g-cdn/internal/config"
	"github.com/WildEgor/g-cdn/internal/db"
	"github.com/WildEgor/g-cdn/internal/handlers/delete"
	"github.com/WildEgor/g-cdn/internal/handlers/download"
	"github.com/WildEgor/g-cdn/internal/handlers/get-files"
	"github.com/WildEgor/g-cdn/internal/handlers/health-check"
	"github.com/WildEgor/g-cdn/internal/handlers/metadata"
	"github.com/WildEgor/g-cdn/internal/handlers/upload"
	"github.com/WildEgor/g-cdn/internal/repositories"
	"github.com/WildEgor/g-cdn/internal/router"
	"github.com/google/wire"
)

// Injectors from server.go:

func NewServer() (*Server, error) {
	configurator := config.NewConfigurator()
	appConfig := config.NewAppConfig(configurator)
	mongoConfig := config.NewMongoConfig(configurator)
	mongoDBConnection := db.NewMongoDBConnection(mongoConfig)
	fileRepository := repositories.NewFileRepository(mongoDBConnection)
	minioConfig := config.NewMinioConfig(configurator)
	storageConfig := adapters.NewStorageConfig(minioConfig)
	storageProvider := adapters.NewStorage(storageConfig)
	storageAdapter := adapters.NewStorageAdapter(storageProvider)
	uploadHandler := upload_handler.NewUploadHandler(fileRepository, storageAdapter)
	downloadHandler := download_handler.NewDownloadHandler(storageAdapter)
	deleteHandler := delete_handler.NewDeleteHandler(fileRepository, storageAdapter)
	getFilesHandler := get_files_handler.NewGetFilesHandler(fileRepository, appConfig)
	metadataHandler := metadata_handler.NewMetadataHandler(storageAdapter)
	filesRouter := router.NewFilesRouter(uploadHandler, downloadHandler, deleteHandler, getFilesHandler, metadataHandler)
	healthCheckHandler := health_check_handler.NewHealthCheckHandler(storageAdapter, appConfig)
	healthRouter := router.NewHealthRouter(healthCheckHandler)
	server := NewApp(appConfig, filesRouter, healthRouter, mongoDBConnection)
	return server, nil
}

// server.go:

var ServerSet = wire.NewSet(AppSet)
